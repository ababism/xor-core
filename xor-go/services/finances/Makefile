include .env

codegen-all:
	oapi-codegen --config .codegen/bank-account-codegen-config.yaml .codegen/bank-account-codegen.yaml
	oapi-codegen --config .codegen/discount-codegen-config.yaml .codegen/discount-codegen.yaml
	oapi-codegen --config .codegen/payment-codegen-config.yaml .codegen/payment-codegen.yaml
	oapi-codegen --config .codegen/payout-request-codegen-config.yaml .codegen/payout-request-codegen.yaml
	oapi-codegen --config .codegen/product-codegen-config.yaml .codegen/product-codegen.yaml
	oapi-codegen --config .codegen/purchase-request-codegen-config.yaml .codegen/purchase-request-codegen.yaml

codegen:
	oapi-codegen --config .codegen/finance-codegen-config.yaml .codegen/finance-codegen.yaml

migrate-up:
	migrate -path ./migrations -database 'postgres://$(FINANCES_POSTGRES_USER):$(FINANCES_POSTGRES_PASSWORD)@$(FINANCES_POSTGRES_HOST):$(FINANCES_POSTGRES_PORT)/$(FINANCES_POSTGRES_NAME)?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://$(FINANCES_POSTGRES_USER):$(FINANCES_POSTGRES_PASSWORD)@$(FINANCES_POSTGRES_HOST):$(FINANCES_POSTGRES_PORT)/$(FINANCES_POSTGRES_NAME)?sslmode=disable' down
