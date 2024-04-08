package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"net/http"
	"time"
	"xor-go/services/finances/internal/domain"
	"xor-go/services/finances/internal/log"
	"xor-go/services/finances/internal/service/adapters"
)

const (
	spanDefaultPurchaseRequest = "purchase-request/service."
)

var _ adapters.PurchaseRequestService = &purchaseRequestService{}

type purchaseRequestService struct {
	rPurchase adapters.PurchaseRequestRepository
	rProduct  adapters.ProductRepository
	rPayment  adapters.PaymentRepository
	rDiscount adapters.DiscountRepository
	cPayment  adapters.PaymentsClient
}

func NewPurchaseRequestService(
	purchaseRequestRepository adapters.PurchaseRequestRepository,
	paymentsClient adapters.PaymentsClient,
	productRepo adapters.ProductRepository,
	paymentRepo adapters.PaymentRepository,
	discountRepo adapters.DiscountRepository,
) adapters.PurchaseRequestService {
	return &purchaseRequestService{
		rPurchase: purchaseRequestRepository,
		rProduct:  productRepo,
		rPayment:  paymentRepo,
		rDiscount: discountRepo,
		cPayment:  paymentsClient,
	}
}

func getPurchaseRequestTracerSpan(ctx context.Context, name string) (trace.Tracer, context.Context, trace.Span) {
	tr := global.Tracer(adapters.ServiceNamePurchaseRequest)
	newCtx, span := tr.Start(ctx, spanDefaultPurchaseRequest+name)
	return tr, newCtx, span
}

func (s *purchaseRequestService) Get(ctx context.Context, id uuid.UUID) (*domain.PurchaseRequestGet, error) {
	_, newCtx, span := getPurchaseRequestTracerSpan(ctx, ".Get")
	defer span.End()

	request, err := s.rPurchase.Get(newCtx, id)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (s *purchaseRequestService) GetPrice(ctx context.Context, id uuid.UUID) (*float32, error) {
	_, newCtx, span := getPurchaseRequestTracerSpan(ctx, ".GetPrice")
	defer span.End()

	request, err := s.rPurchase.GetPrice(newCtx, id)
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (s *purchaseRequestService) List(
	ctx context.Context,
	filter *domain.PurchaseRequestFilter,
) ([]domain.PurchaseRequestGet, error) {
	_, newCtx, span := getPurchaseRequestTracerSpan(ctx, ".List")
	defer span.End()

	purchaseRequests, err := s.rPurchase.List(newCtx, filter)
	if err != nil {
		return nil, err
	}

	return purchaseRequests, nil
}

func (s *purchaseRequestService) Create(ctx context.Context, purchase *domain.PurchaseRequestCreate) error {
	_, newCtx, span := getPurchaseRequestTracerSpan(ctx, ".Create")
	defer span.End()

	id, err := s.rPurchase.Create(newCtx, purchase, 0)
	if err != nil {
		return err
	}

	products := make([]domain.ProductGet, 0)
	productsNames := ""
	for _, productId := range purchase.Products {
		product, err := s.rProduct.Get(ctx, productId)
		if err != nil {
			log.Logger.Error(
				fmt.Sprintf("Error while finding a product with id=%s: %v", productId, zap.Error(err)),
			)
			return nil
		}
		products = append(products, *product)
		productsNames += " " + product.Name
	}

	err := s.rPurchase.Create(newCtx, purchase, 0)

	createPurchase, err := s.cPayment.CreatePurchase(ctx, &domain.PaymentsCreatePurchase{
		PaymentUUID: *id,
		PaymentName: fmt.Sprintf("Payment for:%s", productsNames),
		Money:       sum,
		Currency:    "RUB",
		FullName:    purchase.Sender.String(),
		Phone:       "",
		Email:       "",
		Products:    []domain.PaymentsCreatePurchaseProduct{},
	})
	if err != nil {
		return err
	}

	log.Logger.Info(fmt.Sprintf("Purchase created: %v", createPurchase))

	jsonBody := []byte(`{"client_message": "Hello, Courses!"}`)
	bodyReader := bytes.NewReader(jsonBody)

	log.Logger.Info(fmt.Sprintf("Sending Purchase request result to %s", purchase.WebhookURL))

	_, err = http.NewRequest(http.MethodPost, purchase.WebhookURL, bodyReader)
	if err != nil {
		return err
	}

	return nil
}

func (s *purchaseRequestService) Archive(ctx context.Context, id uuid.UUID) error {
	_, newCtx, span := getPurchaseRequestTracerSpan(ctx, ".Archive")
	defer span.End()

	purchase, err := s.rPurchase.Get(ctx, id)
	if err != nil {
		return err
	}

	err = s.rPayment.Create(ctx, &domain.PaymentCreate{
		Sender:   purchase.Sender,
		Receiver: purchase.Receiver,
		Data:     domain.PaymentData{},
		URL:      purchase.WebhookURL,
		Status:   domain.STATUS_COMPLETED,
		EndedAt:  time.Now(),
	})
	if err != nil {
		return err
	}

	err = s.rPurchase.Delete(newCtx, id)
	if err != nil {
		return err
	}

	return nil
}
