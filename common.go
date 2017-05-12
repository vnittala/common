package common

import (
	"net/http"
	"os"

	"database/sql"

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
	LOG_NAME           = "/tmp/hcare.log"
	DATE_LAYOUT        = "02012006"

	//Database exception codes
	DB_OPEN = 1000
	DB_PING = 2000
)

var (
	LOG_FILE_NAME string
)

//getLogFileName retrieves log file name based on
//OUT_LOG_FILE environment variable.
//If not exported defaults to tmp folder
func getLogFileName() string {
	LOG_FILE_NAME = os.Getenv("OUT_LOG_FILE")

	if LOG_FILE_NAME == "" {
		LOG_FILE_NAME = LOG_NAME
	}
	return LOG_FILE_NAME
}

//LogOutput logs the records to log file
func LogOutput(level string, msg string) {
	logger, _ := NewFileOutLog()

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

}

//CheckErr function check if there is error and logs it
func CheckErr(err error, msg string) {
	if err != nil {
		LogOutput("error", msg+"-"+err.Error())
	}
}

//CheckFatalErr function shuts the application in case of fatal error
func CheckFatalErr(err error, errCode int, msg string) {
	if err != nil {
		LogOutput("error", msg+"-"+err.Error())
		os.Exit(errCode)
	}
}

//NewUUID function generates a unique identifier
func NewUUID() string {

	return uuid.NewV4().String()
}

/*

Below functions include middlewares for gin-gonic framework

*/

//RequestIdMiddleware is s GIN middleware function to add request ID to the incoming rest request
func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Header.Set("X-Request-Id", NewUUID())
		c.Next()
	}
}

func CheckRowsCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		CheckErr(err, err.Error())
	}
	return count
}

func NewFileOutLog() (*zap.Logger, error) {

	cfg := zap.NewDevelopmentConfig()
	cfg.OutputPaths = []string{
		getLogFileName(),
	}
	return cfg.Build()

}
