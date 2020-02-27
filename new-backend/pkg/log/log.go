package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Log 封包 LOG
func Log(title string, msg interface{}) {
	logTitle := `[` + title + `]`
	logMsg := msg
	time := time.Now().Format(`2006-01-02 15:04:05`)
	fmt.Fprintln(gin.DefaultWriter, time, logTitle, `：`, logMsg)
}

// InitLogFile 初始化 Log file
func InitLogFile(logFilePath string) *os.File {
	// logger file
	src, err := os.Create(logFilePath)
	if err != nil {
		fmt.Println(`err`, err)
	}

	gin.DefaultWriter = io.MultiWriter(src, os.Stdout)

	return src
}
