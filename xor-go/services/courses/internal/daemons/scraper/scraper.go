package scraper

import (
	"context"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"runtime"
	"time"
	"xor-go/services/courses/internal/domain"
	"xor-go/services/courses/internal/service/adapters"
)

type Scraper struct {
	stop          chan bool
	courseService adapters.CourseService
	logger        *zap.Logger
}

func NewScraper(logger *zap.Logger, courseService adapters.CourseService) *Scraper {
	return &Scraper{
		logger:        logger,
		courseService: courseService,
		stop:          make(chan bool)}
}

func (s *Scraper) stopCallback(ctx context.Context) error {
	s.stop <- true
	return nil
}

func (s Scraper) StopFunc() func(context.Context) error {
	return s.stopCallback
}

func (s *Scraper) Start(scrapeInterval time.Duration) {
	go func() {
		stop := s.stop
		go func() {
			for {
				s.scrape(scrapeInterval)
				runtime.Gosched()
			}
		}()
		<-stop
	}()
}

func generateRequestID() string {
	id := uuid.New()
	return id.String()
}

func WithRequestID(ctx context.Context) context.Context {
	requestID := generateRequestID()
	return context.WithValue(ctx, domain.KeyRequestID, requestID)
}

func (s *Scraper) scrape(scrapeInterval time.Duration) {
	ctx := context.Background()

	requestIdCtx := WithRequestID(ctx)
	ctxLogger := zapctx.WithLogger(requestIdCtx, s.logger)

	tr := global.Tracer(domain.ServiceName)
	_, span := tr.Start(ctxLogger, "courses/daemon/scraper.Scrape", trace.WithNewRoot())
	defer span.End()

	// Your functionality

	time.Sleep(scrapeInterval)
}
