package banker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"net/http"
	"runtime"
	"time"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/log"
	"xor-go/services/finances/internal/service/adapters"
)

type Banker struct {
	stop               chan bool
	payoutsService     adapters.PayoutRequestService
	purchaseService    adapters.PurchaseRequestService
	productService     adapters.ProductService
	bankAccountService adapters.BankAccountService
	paymentsClient     adapters.PaymentsClient
}

func NewBanker(
	payoutsService adapters.PayoutRequestService,
	purchaseService adapters.PurchaseRequestService,
	productService adapters.ProductService,
	bankAccountService adapters.BankAccountService,
	paymentsClient adapters.PaymentsClient,
) *Banker {
	return &Banker{
		stop:               make(chan bool),
		payoutsService:     payoutsService,
		purchaseService:    purchaseService,
		productService:     productService,
		bankAccountService: bankAccountService,
		paymentsClient:     paymentsClient,
	}
}

func (b *Banker) stopCallback(_ context.Context) error {
	b.stop <- true
	return nil
}

func (b *Banker) StopFunc() func(context.Context) error {
	return b.stopCallback
}

func (b *Banker) Start(scrapeInterval time.Duration) {
	go func() {
		stop := b.stop
		go func() {
			for {
				b.checkForPayments(scrapeInterval)
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

func (b *Banker) checkForPayments(scrapeInterval time.Duration) {
	ctx := context.Background()

	requestIdCtx := WithRequestID(ctx)
	tr := global.Tracer("Banker")
	ctxNew, span := tr.Start(requestIdCtx, "finances/daemon/banker.checkForPayments", trace.WithNewRoot())
	defer span.End()

	purchaseRequests, err := b.purchaseService.List(ctxNew, nil)
	if err != nil {
		return
	}

	for _, req := range purchaseRequests {
		err = b.handlePurchaseRequest(tr, requestIdCtx, req)
		if err != nil {
			log.Logger.Error("Error while handling purchase request", zap.Error(err))
		}
	}

	time.Sleep(scrapeInterval)
}

func (b *Banker) handlePurchaseRequest(tr trace.Tracer, ctx context.Context, req domain.PurchaseRequestGet) error {
	ctxPurchaseReq, spanPurchaseReq := tr.Start(
		ctx,
		fmt.Sprintf("finances/daemons/banker.handlePurchaseRequest.id=%s", req.UUID),
		trace.WithNewRoot(),
	)
	defer spanPurchaseReq.End()

	log.Logger.Info(fmt.Sprintf("Found '%s' Purchase request='%s'", req.Status, req.UUID))
	if req.Status == domain.PaymentsStatusPending {
		status, err := b.paymentsClient.GetStatus(ctxPurchaseReq, req.UUID)
		if err != nil {
			return err
		}

		req.Status = status.Status
	}

	if req.Status == domain.PaymentsStatusSucceeded {
		products := make([]domain.ProductGet, 0)
		for _, productId := range req.Products {
			product, err := b.productService.Get(ctx, productId)
			if err != nil {
				log.Logger.Error(fmt.Sprintf(
					"Purchase request=%s: Error while finding a product with id=%s: %v", req.UUID, productId, zap.Error(err),
				))
				return err
			}
			products = append(products, *product)
		}

		webhook := domain.PurchaseRequestWebhook{
			Sender:   *req.Sender,
			Receiver: *req.Receiver,
			Products: domain.ConvertProductsToSmall(products),
		}
		jsonBody, err := json.Marshal(webhook)
		if err != nil {
			return err
		}
		bodyReader := bytes.NewReader(jsonBody)
		_, err = http.NewRequest(http.MethodPost, req.WebhookURL, bodyReader)
		if err != nil {
			return err
		}
		log.Logger.Info(fmt.Sprintf("Purchase request result sended to %s", req.WebhookURL))

		err = b.purchaseService.Archive(ctxPurchaseReq, req.UUID)
		if err != nil {
			return err
		}
	} else if req.Status == domain.PaymentsStatusCanceled {
		err := b.purchaseService.Archive(ctxPurchaseReq, req.UUID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Banker) handlePayoutRequest(tr trace.Tracer, ctx context.Context, req domain.PayoutRequestGet) error {
	ctxPayoutReq, spanPayoutReq := tr.Start(
		ctx,
		fmt.Sprintf("finances/daemons/banker.handlePayoutRequest.id=%s", req.UUID),
		trace.WithNewRoot(),
	)
	defer spanPayoutReq.End()

	log.Logger.Info(fmt.Sprintf("Found '%s' Payout request='%s'", req.Status, req.UUID))
	if req.Status == domain.PaymentsStatusPending {
		status, err := b.paymentsClient.GetStatus(ctxPayoutReq, req.UUID)
		if err != nil {
			return err
		}
		req.Status = status.Status
	}

	if req.Status == domain.PaymentsStatusSucceeded {
		err := b.bankAccountService.ChangeFunds(ctx, req.Receiver, -1*req.Amount)
		if err != nil {
			return err
		}

		err = b.payoutsService.Archive(ctxPayoutReq, req.UUID)
		if err != nil {
			return err
		}
	} else if req.Status == domain.PaymentsStatusCanceled {
		err := b.payoutsService.Archive(ctxPayoutReq, req.UUID)
		if err != nil {
			return err
		}
	}

	return nil
}
