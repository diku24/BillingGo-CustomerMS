//swagger:meta
package main

import (
	"CustomerMS/api"
	dbInit "CustomerMS/db"
	_ "CustomerMS/docs"
	"CustomerMS/handler"
	"CustomerMS/logger"
	"CustomerMS/repository"
	"CustomerMS/services"
	"CustomerMS/utils"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

var (
	httpRouter  api.Router                 = api.NewMuxRouter()
	billRepo    repository.BillRespository = repository.NewMySQLReopsitory()
	billService services.BillService       = services.NewCustomerService(billRepo)
	billHandler handler.BillHandler        = handler.NewCustomerController(billService)
)

const (
	uriCustomer = "/customer"
)

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func init() {

	utils.InitUtils()

	logger.InitLogger()

	// Open a file for appending logs
	file, err := os.OpenFile("logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Failed to open log file: %v", err)
	}
	//defer file.Close()

	hook := &logger.FileHook{File: file}
	logrus.AddHook(hook)

	//Log as JSON Instead of default ASCII formatter
	//logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: false,
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableColors:    false,
		//ForceColors:            true,
		QuoteEmptyFields:       true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// this function is required when you want to introduce your custom format.
			// In my case I wanted file and line to look like this `file="engine.go:141`
			// but f.File provides a full path along with the file name.
			// So in `formatFilePath()` function I just trimmet everything before the file name
			// and added a line number in the end
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
		PadLevelText:  true,
		FullTimestamp: false,
		// Customizing delimiters
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFile:  "file",
			//logrus.FieldKeyFunc:  "caller",
		},
	})

	//Output tos stdout instead of default stderr
	//can be any io.writter
	multi := io.MultiWriter(file, os.Stdout)
	logrus.SetOutput(multi)

	//Only log the warning serverity or above
	//logrus.SetLevel(logrus.WarnLevel)

	//add the calling method
	logrus.SetReportCaller(true)
}

//	@Title			Customer Service API
//	@Version		1.0
//	@Description	this is the Commerce Billing application for the Server.
//	@termsOfService	http://billingapplication.io/terms/
//	@contact.name	Dinesh Thakur
//	@contact.email	dineshthakur.24@outlook.com
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:8382
//	@accept			json
//	@produce		json

func main() {

	//Create Table in DB Database Setup
	dbInit.TableCreation()

	//Handle the Graceful Shutdown if service is interrupted
	idelConnectionClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)
		logrus.Infoln("Signal Recvied:", sigint)
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

	httpRouter.GET(uriCustomer, api.MakeHTTPHandlerFunction(billHandler.GET))
	httpRouter.DELETE(uriCustomer+"/{customer_id}", api.MakeHTTPHandlerFunction(billHandler.DELETE))
	httpRouter.POST(uriCustomer, api.MakeHTTPHandlerFunction(billHandler.POST))
	httpRouter.UPDATE(uriCustomer, api.MakeHTTPHandlerFunction(billHandler.PUT))

	// httpRouter.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	httpRouter.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8382/docs/swagger.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
		httpSwagger.Layout(httpSwagger.StandaloneLayout),
		httpSwagger.InstanceName("swagger"),
	)).Methods(http.MethodGet)

	httpRouter.SERVE(port)

	<-idelConnectionClosed
	logrus.Infoln("Service is Stopped Successfully!")
}
