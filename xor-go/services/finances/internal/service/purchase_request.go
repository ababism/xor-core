package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	global "go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
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
	_, newCtx, span := getPurchaseRequestTracerSpan(ctx, ".GetByLogin")
	defer span.End()

	request, err := s.rPurchase.Get(newCtx, id)
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

func (s *purchaseRequestService) Create(
	ctx context.Context,
	purchase *domain.PurchaseRequestCreate,
) (*uuid.UUID, error) {
	_, newCtx, span := getPurchaseRequestTracerSpan(ctx, ".Create")
	defer span.End()

	_, newCtxSpanCreate, spanCreate := getPurchaseRequestTracerSpan(newCtx, ".Create.PurchaseRequest")
	defer spanCreate.End()
	price, err := s.rProduct.GetPrice(ctx, purchase.Products)
	if err != nil {
		return nil, err
	}

	id, err := s.rPurchase.Create(newCtx, purchase, *price)
	if err != nil {
		return nil, err
	}
	spanCreate.End()

	_, _, spanPayments := getPurchaseRequestTracerSpan(newCtxSpanCreate, ".Create.SendPayment")
	defer spanPayments.End()
	products := make([]domain.ProductGet, 0)
	productsNames := ""
	for _, productId := range purchase.Products {
		product, err := s.rProduct.Get(ctx, productId)
		if err != nil {
			log.Logger.Error(
				fmt.Sprintf("Error while finding a product with id=%s: %v", productId, zap.Error(err)),
			)
			return nil, err
		}
		products = append(products, *product)
		productsNames += " " + product.Name
	}

	createPurchase, err := s.cPayment.CreatePurchase(newCtx, &domain.PaymentsCreatePurchase{
		PaymentUUID: *id,
		PaymentName: fmt.Sprintf("Payment for:%s", productsNames),
		Money:       *price,
		Currency:    "RUB",
		FullName:    fmt.Sprintf("%s,", purchase.Sender),
		Phone:       "",
		Email:       "",
		Products:    []domain.PaymentsCreatePurchaseProduct{},
	})
	if err != nil {
		return nil, err
	}
	spanPayments.End()

	log.Logger.Info(fmt.Sprintf("Purchase created status='%s': %v", createPurchase.Status, createPurchase))
	return id, nil
}

func (s *purchaseRequestService) Archive(ctx context.Context, id uuid.UUID) error {
	_, newCtx, span := getPurchaseRequestTracerSpan(ctx, ".Archive")
	defer span.End()

	purchase, err := s.rPurchase.Get(ctx, id)
	if err != nil {
		return err
	}

	_, err = s.rPayment.Create(ctx, &domain.PaymentCreate{
		Sender:   *purchase.Sender,
		Receiver: *purchase.Receiver,
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
