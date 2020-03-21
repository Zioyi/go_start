package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// LogWriter - 声明日志写入起接口
type LogWriter interface {
	Write(data interface{}) error
}

// Logger - 日志器
type Logger struct {
	writeList []LogWriter
}

// RegisterWriter - 注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writeList = append(l.writeList, writer)

}

// Log - 将一个data类型的数据写入日志
func (l *Logger) Log(data interface{}) {
	// 遍历所有注册的写入器
	for _, writer := range l.writeList {
		// 将日志写入到每一个写入器中
		writer.Write(data)
	}
}

// NewLogger - 创建日志器的实例
func NewLogger() *Logger {
	return &Logger{}
}

// 声明文件写入器
type fileWriter struct {
	file *os.File
}

// SetFile - 设置文件写入器的文件名
func (f *fileWriter) SetFile(filename string) (err error) {
	// 如果文件已经打开，关闭前一个文件
	if f.file != nil {
		f.file.Close()
	}

	// 创建一个文件并保存文件句柄
	// f.file, err = os.Create(filename)
	// 追加写入
	f.file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	// 如果创建的过程出现错误，则返回错误
	return err
}

func getLocalTime() string {
	timeStr := time.Now().String()
	fmt.Println(timeStr)
	newTimeStr := strings.Split(timeStr, ".")[0]
	return newTimeStr
}

func (f *fileWriter) Write(data interface{}) error {
	if f.file == nil {
		return errors.New("file not created")
	}

	str := fmt.Sprintf("file::%v::%v\n", getLocalTime(), data)

	_, err := f.file.Write([]byte(str))

	return err
}

func newFileWriter() *fileWriter {
	return &fileWriter{}
}

// 命令行写入器
type consoleWriter struct {
}

func (f *consoleWriter) Write(data interface{}) error {

	str := fmt.Sprintf("consele::%v::%v\n", getLocalTime(), data)

	_, err := os.Stdout.Write([]byte(str))

	return err
}

func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}

func main() {

	l := NewLogger()

	fw := newFileWriter()
	fw.SetFile("file.log")

	cw := newConsoleWriter()

	l.RegisterWriter(fw)
	l.RegisterWriter(cw)

	l.Log("hello")

}
