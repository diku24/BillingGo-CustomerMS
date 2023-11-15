//swagger:meta
package main

import (
	"BillingGo/api"
	dbInit "BillingGo/db"
	_ "BillingGo/docs"
	"BillingGo/handler"
	"BillingGo/repository"

	"BillingGo/services"
	"BillingGo/utils"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	httpRouter  api.Router                 = api.NewMuxRouter()
	billRepo    repository.BillRespository = repository.NewMySQLReopsitory()
	billService services.BillService       = services.NewCustomerService(billRepo)
	billHandler handler.BillHandler        = handler.NewCustomerController(billService)
)

func init() {
	//Log as JSON Instead of default ASCII formatter
	//logrus.SetFormatter(&logrus.JSONFormatter{})

	//Output tos stdout instead of default stderr
	//can be any io.writter
	//logrus.SetOutput(os.Stdout)

	//Only log the warning serverity or above
	//logrus.SetLevel(logrus.WarnLevel)
}

//	@Title			Tag Service API
//	@Description:	this is the Commerce Billing application for the Server.
//
// @termsOfService: http://billingapplication.io/terms/
// @contact:
// @email: dineshthakur.24@outlook.com
// @license:
// @name: Apache 2.0
// @url: http://www.apache.org/licenses/LICENSE-2.0.html
// @version: 1.0.11
//
//go:generate swagger generate spec
func main() {

	//Create Table in DB Database Setup
	dbInit.TableCreation()

	//Handle the Graceful Shutdown if service is interrupted
	idelConnectionClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		logrus.Infoln("Service interrupt received !!")

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		err := httpRouter.GraceFulShutDown(ctx)
		if err != nil {
			logrus.Infof("Unable to Perform shutdown Please wait for time: %v", err)
		}
		logrus.Println("Shutdown Completed !!")

		close(idelConnectionClosed)

	}()

	//Load Configuration.env file
	//port is the constant for port number
	port := utils.EnvVarRead("SERVERPORT")

	//http GET endpoint for checking Server is Running Status
	httpRouter.GET("/", func(responce http.ResponseWriter, request *http.Request) {
		logrus.Println("SERVER IS RUNNNING !!")
	})

	httpRouter.GET("/dbtest", func(responce http.ResponseWriter, request *http.Request) {
		dbInit.PingServer()
	})

	httpRouter.GET("/customer", api.MakeHTTPHandlerFunction(billHandler.GET))
	httpRouter.GET("/customer/{customer_id}", api.MakeHTTPHandlerFunction(billHandler.GET))
	httpRouter.DELETE("/customer/{customer_id}", api.MakeHTTPHandlerFunction(billHandler.DELETE))
	httpRouter.POST("/customer", api.MakeHTTPHandlerFunction(billHandler.POST))

	httpRouter.UPDATE("/customer", api.MakeHTTPHandlerFunction(billHandler.PUT))

	httpRouter.SERVE(port)

	<-idelConnectionClosed
	logrus.Infoln("Service is Stopped Successfully!")
}
