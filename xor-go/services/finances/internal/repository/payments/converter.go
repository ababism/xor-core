package payments

import "xor-go/services/finances/internal/domain"

func convertToCreatePurchase(model domain.PaymentsCreatePurchase) CreatePurchase {
	return CreatePurchase{
		Currency:    &model.Currency,
		Email:       model.Email,
		FullName:    model.FullName,
		Money:       model.Money,
		PaymentName: model.PaymentName,
		PaymentUuid: model.PaymentUUID.String(),
		Phone:       model.Phone,
		Products:    convertToCreatePurchaseProducts(model.Products),
	}
}

func convertToCreatePurchaseProducts(model []domain.PaymentsCreatePurchaseProduct) []CreatePurchaseProduct {
	createPurchaseProducts := make([]CreatePurchaseProduct, len(model))
	for i, product := range model {
		createPurchaseProducts[i] = CreatePurchaseProduct{
			Currency:    &product.Currency,
			Description: product.Description,
			Money:       product.Money,
			PaymentMode: &product.PaymentMode,
			Quantity:    product.Quantity,
		}
	}
	return createPurchaseProducts
}

func convertToCreatePayout(model domain.PaymentsCreatePayout) CreatePayout {
	return CreatePayout{
		CardInfo: CardInfo{
			&model.CardInfo.CardType,
			model.CardInfo.First6,
			&model.CardInfo.IssuerCountry,
			model.CardInfo.IssuerName,
			model.CardInfo.Last4,
		},
		Currency:    &model.Currency,
		Email:       &model.Email,
		FullName:    model.FullName,
		IsTest:      &model.IsTest,
		Money:       model.Money,
		PaymentName: model.PaymentName,
		PaymentUuid: model.PaymentUUID,
		Phone:       &model.Phone,
	}
}
