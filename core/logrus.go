package core

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

// 配置格式化，实现打印文件路径，日志时间，日志等级，日志提示颜色
type MyLog struct {
}

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

func (MyLog) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据不同的level展示颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	// 定义日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	if entry.HasCaller() {
		// 定义文件路径
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		// 定义输出格式
		fmt.Fprintf(b, "[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message)
	}
	return b.Bytes(), nil
}

func InitLogger() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(MyLog{})
	logrus.AddHook(&MyHook{
		logDir: "logs",
	})
}

// 配置hook，实现日志产生后写入到日志文件之中，并按天进行分片存储日志，另外错误的日志需要单独存放
type MyHook struct {
	file    *os.File   // 当前打开的日志文件
	errFile *os.File   // 错误日志的日志文件
	date    string     // 当前日志的时间
	logDir  string     // 日志路径
	mu      sync.Mutex // 文件互斥锁
}

func (hook *MyHook) Fire(entry *logrus.Entry) error {
	hook.mu.Lock()
	defer hook.mu.Unlock()
	msg, _ := entry.String()
	date := entry.Time.Format("2006-01-02")
	if hook.date != date {
		// 如果时间不同即需要更换时间，打开新的日志文件
		hook.rotateFiles(date)
		hook.date = date
	}
	// 判断是否为错误及以上级别的日志等级，如果是则单独存放在错误日志的日志文件中
	if entry.Level <= logrus.ErrorLevel {
		hook.errFile.Write([]byte(msg))
	}

	hook.file.Write([]byte(msg))
	//fmt.Println(entry)
	return nil

}

func (hook *MyHook) rotateFiles(timer string) error {
	if hook.file != nil {
		hook.file.Close()
	}
	if hook.file == nil {
		// 判断文件是否存在并创建写入
		logDir := fmt.Sprintf("%s/%s", hook.logDir, timer)
		os.MkdirAll(logDir, 0666)
		// 日常日志文件的存放
		logPath := fmt.Sprintf("%s/info.log", logDir)
		file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		hook.file = file
		// 错误日志文件的存放
		errlogPath := fmt.Sprintf("%s/err.log", logDir)
		errFile, _ := os.OpenFile(errlogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		hook.errFile = errFile
	}
	return nil
}

func (*MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
