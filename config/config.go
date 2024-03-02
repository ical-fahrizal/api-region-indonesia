package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	// co "api-gateway-dc/controllers"

	"github.com/ipsusila/opt"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"

	// u "itemcode-dc/utils"

	// _ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
)

// type DataWsUser struct {
// 	Token   string      `json:"token"`
// 	Data    interface{} `json:"data"`
// 	Request string      `json:request`
// 	Uuid    uuid.UUID   `json:"-"`
// }

//variable global validation with regexp
// var re = regexp.MustCompile("'")

type Error string

var TestMsg string

// var DataAuth2 []DataWsUser
var configTokenExpiration int
var configAppName string

// var secretKey []byte
var configPort string

// var rh *rejson.Handler
// var cli *goredis.Client

// var db *sql.DB

// var DataRequest []interface{}
var DataRequest map[string]interface{}

var LogInfo *log.Logger
var LogError *log.Logger
var LogStatus *log.Logger

func init() {
	// with loggger lumberjack
	ex, err := os.Executable()
	if err != nil {
		LogInfo.Println("Error Executable: ", err)
	}
	exPath := filepath.Dir(ex) + "/"
	confPath := exPath + "config.hjson"
	// re-open file
	file, err := os.Open(confPath)
	if err != nil {
		confPath = "config.hjson"
	}
	defer file.Close()

	//parse configurationf file
	cfgFile := flag.String("conf", confPath, "Configuration file")
	flag.Parse()
	// LogInfo.Print("masuk config")

	//load options
	config, err := opt.FromFile(*cfgFile, opt.FormatAuto)
	if err != nil {
		LogError.Printf("Error while loading configuration file %v -> %v\n", *cfgFile, err)
		return
	}

	// DataAuth2 := make([]DataAuth, 0)
	// LogInfo.Printf("%v", DataAuth2)

	// DataRequest := make(map[string]interface{}, 0)
	// LogInfo.Printf("DataRequest : %v", DataRequest)

	// Port Hostname
	configPort = config.Get("server").GetString("port", "3008")

	//logging INFO ----------------
	logName := fmt.Sprintf(`./%v`, config.Get("log").Get("info").GetString("fileName", "info.log"))
	fileLogInfo, err := openLogFile(logName)
	if err != nil {
		LogError.Fatal(err)
	}
	LogInfo = log.New(fileLogInfo, "[INFO]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	cfgLog := config.Get("log")
	cfgLogInfo := cfgLog.Get("info")
	logMaxSize := cfgLogInfo.GetInt("maxSize", 10)
	logMaxBackups := cfgLogInfo.GetInt("maxBackups", 3)
	logMaxAge := cfgLogInfo.GetInt("maxAge", 3)
	logCompress := cfgLogInfo.GetBool("compress", false)

	LogInfo.SetOutput(&lumberjack.Logger{
		Filename:   logName,
		MaxSize:    logMaxSize, // megabytes
		MaxBackups: logMaxBackups,
		MaxAge:     logMaxAge,   //days
		Compress:   logCompress, // disabled by default
	})
	//End logging INFO ----------------

	//logging ERROR ----------------
	logName = fmt.Sprintf(`./%v`, config.Get("log").Get("error").GetString("fileName", "error.log"))
	fileLogError, err := openLogFile(logName)
	if err != nil {
		LogError.Fatal(err)
	}
	LogError = log.New(fileLogError, "[ERROR]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	cfgLogError := cfgLog.Get("error")
	logMaxSize = cfgLogError.GetInt("maxSize", 10)
	logMaxBackups = cfgLogError.GetInt("maxBackups", 3)
	logMaxAge = cfgLogError.GetInt("maxAge", 3)
	logCompress = cfgLogError.GetBool("compress", false)
	LogError.SetOutput(&lumberjack.Logger{
		Filename:   logName,
		MaxSize:    logMaxSize, // megabytes
		MaxBackups: logMaxBackups,
		MaxAge:     logMaxAge,   //days
		Compress:   logCompress, // disabled by default
	})
	//End logging ERROR ----------------

	//logging SERVICE STATUS ----------------
	logName = fmt.Sprintf(`./%v`, config.Get("log").Get("status").GetString("fileName", "status.log"))
	fileLogStatus, err := openLogFile(logName)
	if err != nil {
		LogError.Fatal(err)
	}
	LogStatus = log.New(fileLogStatus, "[STATUS]", log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	cfgLogStatus := cfgLog.Get("status")
	logMaxSize = cfgLogStatus.GetInt("maxSize", 10)
	logMaxBackups = cfgLogStatus.GetInt("maxBackups", 3)
	logMaxAge = cfgLogStatus.GetInt("maxAge", 3)
	logCompress = cfgLogStatus.GetBool("compress", false)
	LogStatus.SetOutput(&lumberjack.Logger{
		Filename:   logName,
		MaxSize:    logMaxSize, // megabytes
		MaxBackups: logMaxBackups,
		MaxAge:     logMaxAge,   //days
		Compress:   logCompress, // disabled by default
	})
	//End logging SERVICE STATUS ----------------

	LogStatus.Println("RUN")

	//app_name config
	configAppName = config.Get("server").GetString("appName", "api-gateway-dc")
	LogInfo.Println("init() -> configAppName: ", configAppName)

	//token_expiration config
	// wib := 7 * 60
	// configTokenExpiration = config.Get("server").GetInt("tokenExpiration", 15) + wib
	// LogInfo.Println("init() -> configTokenExpiration: ", configTokenExpiration)

}

//database
// func GetDB() *sql.DB {
// 	return db
// }

// Config func to get env value
func Config(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetTokenExpiration() int {
	return configTokenExpiration
}

func GetAppName() string {
	return configAppName
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func GetPort() string {
	return configPort
}
