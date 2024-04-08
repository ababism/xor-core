package app

import (
	"context"
	"fmt"
	"github.com/juju/zaputil/zapctx"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"xor-go/pkg/apperror"
	"xor-go/pkg/metrics"
	"xor-go/pkg/mylogger"
	"xor-go/pkg/xshutdown"
	"xor-go/pkg/xtracer"
	"xor-go/services/courses/internal/config"
	kafkaConsumer "xor-go/services/courses/internal/daemons/kafkaConsumer"
	"xor-go/services/courses/internal/daemons/scraper"
	financesClient "xor-go/services/courses/internal/repository/financesclient"
	financesClientGen "xor-go/services/courses/internal/repository/financesclient/generated"
	"xor-go/services/courses/internal/repository/mongo"
	"xor-go/services/courses/internal/repository/mongo/collections"
	"xor-go/services/courses/internal/service"
	"xor-go/services/courses/internal/service/adapters"
)

type App struct {
	cfg            *config.Config
	address        string
	logger         *zap.Logger
	tracerProvider *trace.TracerProvider
	service        adapters.CoursesService
}

func NewApp(cfg *config.Config) (*App, error) {
	startCtx := context.Background()

	// INFRASTRUCTURE ----------------------------------------------------------------------

	// Инициализируем logger
	logger, err := mylogger.InitLogger(cfg.Logger, cfg.App.Name)
	if err != nil {
		return nil, err
	}
	// Чистим кэш logger при shutdown
	xshutdown.AddCallback(
		&xshutdown.Callback{
			Name: "ZapLoggerCacheWipe",
			FnCtx: func(ctx context.Context) error {
				return logger.Sync()
			},
		})
	logger.Info("Init Logger – success")

	// сохраняем logger в контекст
	ctx := zapctx.WithLogger(startCtx, logger)

	// Инициализируем обработку ошибок
	err = apperror.InitAppError(cfg.App)
	if err != nil {
		logger.Fatal("while initializing App Error handling package", zap.Error(err))
	}

	//logger.Info("Importing constants from driver openApi – success")

	// Инициализируем трассировку
	tp, err := xtracer.Init(cfg.Tracer, cfg.App)
	if err != nil {
		return nil, err
	}
	xshutdown.AddCallback(
		&xshutdown.Callback{
			Name: "OpenTelemetryShutdown",
			FnCtx: func(ctx context.Context) error {
				if err := tp.Shutdown(context.Background()); err != nil {
					logger.Error("Error shutting down tracer provider: %v", zap.Error(err))
					return err
				}
				return nil
			},
		})
	logger.Info("Init Tracer – success")

	// Инициализируем Prometheus
	metrics.InitOnce(cfg.Metrics, logger, metrics.AppInfo{
		Name:        cfg.App.Name,
		Environment: string(cfg.App.Environment),
		Version:     cfg.App.Version,
	})
	logger.Info("Init Metrics – success")

	// REPOSITORY ----------------------------------------------------------------------

	// Инициализируем MongoDB
	mongoDatabase := mongo.NewDatabase(logger)

	mongoDisconnect, err := mongoDatabase.Connect(ctx, cfg.Mongo, cfg.MigrationsMongo)
	if err != nil {
		logger.Fatal("error while connecting to Mongo DB:", zap.Error(err))
	}
	xshutdown.AddCallback(
		&xshutdown.Callback{
			Name: "MongoClientDisconnect",
			FnCtx: func(ctx context.Context) error {
				return mongoDisconnect(ctx)
			},
		})
	logger.Info("Mongo connect – success")

	// All repositories
	courseRepo := mongo.NewCourseRepository(mongoDatabase, collections.CourseCollectionName)
	lessonRepo := mongo.NewLessonRepository(mongoDatabase, collections.LessonCollectionName)
	courseEditRepo := mongo.NewCourseRepository(mongoDatabase, collections.CourseEditorCollectionName)
	lessonEditRepo := mongo.NewLessonRepository(mongoDatabase, collections.LessonEditorCollectionName)
	teacherRepo := mongo.NewTeacherRepository(mongoDatabase)
	studentRepo := mongo.NewStudentRepository(mongoDatabase)
	publicationRepo := mongo.NewPublicationRepository(mongoDatabase)

	// Инициализируем Location microservice client repository layer
	client, err := financesClientGen.NewClientWithResponses(cfg.FinancesClient.Uri)
	if err != nil {
		logger.Fatal("cannot initialize generated location client:", zap.Error(err))
		return nil, err
	}
	financesCLi := financesClient.NewClient(client)
	logger.Info("Init Location client – success")

	// Kafka Producer in repository layer
	//_ = kafkaproducer.NewKafkaProducer(cfg.KafkaWriter)
	//logger.Info("Init Kafka Producer – success")

	// SERVICE LAYER ----------------------------------------------------------------------

	// Service layer
	coursesService := service.NewCoursesService(courseRepo, courseEditRepo,
		lessonRepo, lessonEditRepo, teacherRepo, studentRepo, publicationRepo, financesCLi)

	logger.Info(fmt.Sprintf("Init %s – success", cfg.App.Name))

	// Scraper for event calling
	scr := scraper.NewScraper(logger, coursesService)
	xshutdown.AddCallback(
		&xshutdown.Callback{
			Name:  "Data scraper stop",
			FnCtx: scr.StopFunc(),
		})

	scrapeInterval, err := cfg.Scraper.GetScrapeInterval()
	if err != nil {
		logger.Fatal("can't parse time from scraper LongPollTimeout config string:", zap.Error(err))
	}
	logger.Info("Scraper interval – ", zap.Duration("interval", scrapeInterval))
	// TODO Move to Start() ?
	//scr.Start(scrapeInterval)

	logger.Info("Init Scraper – success")

	// TRANSPORT LAYER ----------------------------------------------------------------------

	// Kafka consumer in transport layer
	consumer := kafkaConsumer.NewKafkaConsumer(cfg.KafkaReader, coursesService)
	kafkaConsumerClose := consumer.Start(ctx)
	xshutdown.AddCallback(&xshutdown.Callback{
		Name: "kafkaConsumer Close",
		FnCtx: func(ctx context.Context) error {
			return kafkaConsumerClose(consumer)
		},
	})
	logger.Info("Init Kafka Consumer – success")

	// инициализируем адрес сервера
	address := fmt.Sprintf(":%d", cfg.Http.Port)

	return &App{
		cfg:            cfg,
		logger:         logger,
		service:        coursesService,
		address:        address,
		tracerProvider: tp,
	}, nil
}
