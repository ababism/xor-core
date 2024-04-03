package kafkaconsumer

import (
	"context"
	"encoding/json"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"time"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/service/adapters"
)

type KafkaConsumer struct {
	reader        *kafka.Reader
	driverService adapters.CourseService
}

func NewKafkaConsumer(cfg *Config, driverService adapters.CourseService) *KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        cfg.Brokers,
		Topic:          cfg.Topic,
		GroupID:        cfg.IdGroup,
		MinBytes:       cfg.MinBytes,
		MaxBytes:       cfg.MaxBytes,
		SessionTimeout: 100 * time.Second,
	})

	return &KafkaConsumer{
		reader:        reader,
		driverService: driverService,
	}
}

func (kc *KafkaConsumer) Start(ctx context.Context) func(kc *KafkaConsumer) error {
	go kc.consumeMessages(ctx)

	return closeKafka
}

func closeKafka(kc *KafkaConsumer) error {
	return kc.reader.Close()
}

func (kc *KafkaConsumer) consumeMessages(mainCtx context.Context) {
	logger := zapctx.Logger(mainCtx)
	for {
		// Block (wait)
		// Создаю новый контекст и логгер для отслеживания trace
		ctx := context.Background()
		message, err := kc.reader.ReadMessage(ctx)
		if err != nil {
			logger.Error("error while reading message from kafka", zap.Error(err))
			continue
		}

		tr := global.Tracer(domain.ServiceName)
		ctxTrace, span := tr.Start(ctx, "driver.daemon.kafkaConsumer: ConsumeMessage", trace.WithNewRoot())
		ctxLog := zapctx.WithLogger(ctxTrace, logger)

		var event Event
		errCreate := json.Unmarshal(message.Value, &event)
		if errCreate != nil {
			logger.Debug("error unmarshalling message from kafka (unsupported schema)", zap.Error(err), zap.ByteString("json:", message.Value))
			continue
		}

		// TODO Delete
		ctxLog.Done()

		// close span and to next message in cycle
		span.End()
	}
}
