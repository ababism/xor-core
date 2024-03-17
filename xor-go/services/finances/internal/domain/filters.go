package domain

import "github.com/google/uuid"

func CreateBankAccountFilter(
	uuid *uuid.UUID,
	accountId *uuid.UUID,
	login *string,
	funds *float64,
	status *string,
) BankAccountFilter {
	return BankAccountFilter{
		UUID:        uuid,
		AccountUUID: accountId,
		Login:       login,
		Funds:       funds,
		Status:      status,
	}
}

func CreateBankAccountFilterLogin(login *string) BankAccountFilter {
	return CreateBankAccountFilter(nil, nil, login, nil, nil)
}

func CreatePaymentFilter(
	uuid *uuid.UUID,
	senderUUID *uuid.UUID,
	receiverUUID *uuid.UUID,
	url *string,

	status *string,
) PaymentFilter {
	return PaymentFilter{
		UUID:     uuid,
		Sender:   senderUUID,
		Receiver: receiverUUID,
		URL:      url,
		Status:   status,
	}
}

func CreatePaymentFilterId(uuid *uuid.UUID) PaymentFilter {
	return CreatePaymentFilter(uuid, nil, nil, nil, nil)
}
