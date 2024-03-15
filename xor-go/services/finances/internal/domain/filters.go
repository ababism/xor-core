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

func CreateBankAccountFilterByLogin(login string) BankAccountFilter {
	return CreateBankAccountFilter(nil, nil, &login, nil, nil)
}
