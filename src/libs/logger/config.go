package logger

import (
	"be-blog/src/libs/errors"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

type Config struct {
	Folder string
	Ext    string
}

type Logger struct {
	log *log.Logger
}

func (l *Logger) SetError(err error) *Logger {
	isDev := os.Getenv("ENV_MODE") == "dev"
	// format: Type, Status, Message, Time, File, Line
	_, file, line, _ := runtime.Caller(2)
	time := time.Now()
	switch e := err.(type) {
	case *errors.Error:
		if isDev {
			log.Println(time.Format("2006-01-02 15:04:05"), e.Type, e.Status, e.Message, file, line)
		} else {
			l.log.Printf("%s\t%s\t%d\t%s\t%s:%d", time.Format("2006-01-02 15:04:05"), e.Type, e.Status, e.Message, file, line)
		}
	default:
		if isDev {
			log.Println(time.Format("2006-01-02 15:04:05"), "error", http.StatusInternalServerError, err.Error(), file, line)
		} else {
			l.log.Printf("%s\t%s\t%d\t%s\t%s:%d", time.Format("2006-01-02 15:04:05"), "error", http.StatusInternalServerError, err.Error(), file, line)
		}
	}
	return l
}

var logger Logger

func nameFileLog(folder, ext string) string {
	time := time.Now()
	return folder + time.Format("2006-01-02") + ext
}

func InitLog(config Config) {
	if config.Ext == "" {
		config.Ext = ".txt"
	}
	fullPath := nameFileLog(config.Folder, config.Ext)
	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger = Logger{
		log: log.New(file, "", 0),
	}
	logger.log.SetOutput(file)
}
