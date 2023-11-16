build:
	@go build -o bin/billingGo

run: build
	@./bin/billingGo

test:
	@go test ./... -coverprofile=coverage

mockGenerateRepo:
#Make sure The Path for the Source and Destination should be full path for the file
	@mockgen -source=C:/Users/Dinesh/go/src/BillingGo/repository/bill-repo.go -destination=C:/Users/Dinesh/go/src/BillingGo/mocks/repositoryMocks.go -package=mocks

mockGenerateService:
#Make sure The Path for the Source and Destination should be full path for the file
	@mockgen -source=C:/Users/Dinesh/go/src/BillingGo/services/service.go -destination=C:/Users/Dinesh/go/src/BillingGo/mocks/serviceMocks.go -package=mocks

mockGenerateHandler:
#Make sure The Path for the Source and Destination should be full path for the file
	@mockgen -source=C:/Users/Dinesh/go/src/BillingGo/handler/handler.go -destination=C:/Users/Dinesh/go/src/BillingGo/mocks/handlerMocks.go -package=mocks

mockGenerateAPI:
#Make sure The Path for the Source and Destination should be full path for the file
	@mockgen -source=C:/Users/Dinesh/go/src/BillingGo/api/router.go -destination=C:/Users/Dinesh/go/src/BillingGo/mocks/apiMocks.go -package=mocks

generateAllMocks: mockGenerateRepo mockGenerateAPI mockGenerateHandler mockGenerateService


