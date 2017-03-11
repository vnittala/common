package common

import (
	"net/http"
	"os"
	//	"time"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type ERROR struct {
	ErrorCode    string
	ErrorMessage string
}

const (
	DATA_EXCEPTION     = http.StatusUnprocessableEntity
	AUTH_EXCEPTION     = http.StatusUnauthorized
	GENERAL_EXCEPTION  = http.StatusUnprocessableEntity
	BAD_REQUEST        = http.StatusBadRequest
	INTERNAL_ERROR     = http.StatusInternalServerError
	NOT_IMPLEMENTED    = http.StatusNotImplemented
	RESOURCE_NOT_FOUND = http.StatusNotFound
	STATUS_OK          = http.StatusOK
	LOG_NAME           = "E:\\logFiles\\"
	DATE_LAYOUT        = "02012006"

	// DB Exceptions
	DB_OPEN    = 1000
	CRATE_PING = 2000
)

var (
	LOG_FILE_NAME string
)

func getLogFileName() string {
	LOG_FILE_NAME = os.Getenv("LOG_FOLDER")

	if LOG_FILE_NAME == "" {
		LOG_FILE_NAME = LOG_NAME
	}
	return LOG_FILE_NAME
}

func LogOutput(level string, msg string) {

	/*currentTime := time.Now()
	logName := getLogFileName() + "log-" + currentTime.Format(DATE_LAYOUT) + ".log"
	f, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil {
		panic("failed to create temporary file")
	}
	defer f.Close()
	*/

	//logger := zap.New(zapcore.NewJSONEncoder(), zap.Output(f))
	logger, _ := zap.NewDevelopment()

	switch level {

	case "info":
		logger.Info(msg)
	case "error":
		logger.Error(msg)
	case "fatal":
		logger.Fatal(msg)
	default:
		logger.Debug(msg)
	}
	//f.Sync()

}

func CheckErr(err error, msg string) {
	if err != nil {
		LogOutput("error", msg+"-"+err.Error())
	}
}

func CheckFatalErr(err error, errCode int, msg string) {
	if err != nil {
		LogOutput("error", msg+"-"+err.Error())
		os.Exit(errCode)
	}
}

func NewUUID() string {

	return uuid.NewV4().String()
}

/*

Below functions include middlewares for gin-gonic framework

*/

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("X-Request-Id", NewUUID())
		//c.Next()
	}
}

func FileLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		LogOutput("info", c.Request.Method+"|"+c.Request.Header.Get("X-User-Id")+"|"+c.Request.URL.RequestURI())
	}
}
