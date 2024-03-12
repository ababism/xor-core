package scraper

import (
	"context"
	"github.com/google/uuid"
	"github.com/juju/zaputil/zapctx"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/domain"
	"gitlab.com/ArtemFed/mts-final-taxi/projects/template/internal/service/adapters"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"time"
)

type Scraper struct {
	stop          chan bool
	driverService adapters.DriverService
	logger        *zap.Logger
}

func NewScraper(logger *zap.Logger, driverService adapters.DriverService) *Scraper {
	return &Scraper{
		logger:        logger,
		driverService: driverService,
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
	_, span := tr.Start(ctxLogger, "driver.daemon.scraper: Scrape", trace.WithNewRoot())
	defer span.End()

	time.Sleep(scrapeInterval)
}
