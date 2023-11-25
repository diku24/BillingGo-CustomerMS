build:
	@go build -o bin/billingGo

run: build
	@./bin/billingGo

test:
	@go test ./... -coverprofile=coverage

testHTMLCoverage: test
	@go tool cover -html=coverage

testFuncCoverage: test
	@go tool cover -func=coverage

mody:
	@go mod tidy

mockGenerateRepo:
#Make sure The Path for the Source and Destination should be full path for the file
	@mockgen -source=C:/Users/Dinesh/go/src/BillingGo/CustomerMS/repository/bill-repo.go -destination=C:/Users/Dinesh/go/src/BillingGo/CustomerMS/mocks/repositoryMocks.go -package=mocks

mockGenerateService:
#Make sure The Path for the Source and Destination should be full path for the file
	@mockgen -source=C:/Users/Dinesh/go/src/BillingGo/CustomerMS/services/service.go -destination=C:/Users/Dinesh/go/src/BillingGo/CustomerMS/mocks/serviceMocks.go -package=mocks

mockGenerateHandler:
#Make sure The Path for the Source and Destination should be full path for the file
	@mockgen -source=C:/Users/Dinesh/go/src/BillingGo/CustomerMS/handler/handler.go -destination=C:/Users/Dinesh/go/src/BillingGo/CustomerMS/mocks/handlerMocks.go -package=mocks

mockGenerateAPI:
#Make sure The Path for the Source and Destination should be full path for the file
	@mockgen -source=C:/Users/Dinesh/go/src/BillingGo/CustomerMS/api/router.go -destination=C:/Users/Dinesh/go/src/BillingGo/CustomerMS/mocks/apiMocks.go -package=mocks

generateAllMocks: mockGenerateRepo mockGenerateService  mockGenerateHandler mockGenerateAPI


