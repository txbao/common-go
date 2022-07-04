/**
 * 日志操作
 * Create by whimp(whimp@189.cn)
 * Date: 2020\4\29 18:43
 */
package logs

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"

	log "github.com/sirupsen/logrus"
)

func Init(level string) error {
	//d, err := NewDailyLogWriter(prefix, path)
	//if err != nil {
	//	return err
	//}
	//log.SetOutput(d)

	lv, err := log.ParseLevel(level)
	if err != nil {
		return err
	}

	log.SetLevel(lv)
	return nil
}

func entry(tag ...interface{}) *log.Entry {
	_, f, l, ok := runtime.Caller(2)
	if !ok {
		f, l = "", 0
	}
	//设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return log.WithFields(log.Fields{
		"file": f,
		"line": l,
		"tag":  tag,
	})
}

func Debug(data ...interface{}) {
	entry().Debug(data)
}

func Info(data ...interface{}) {
	entry().Info(data)
}

func Warn(data ...interface{}) {
	entry().Warn(data)
}

func Error(data ...interface{}) {
	entry().Error(data)
}

func DebugTag(tag string, data ...interface{}) {
	entry(tag).Debug(data)
}

func InfoTag(tag string, data ...interface{}) {
	entry(tag).Info(data)
}

func WarnTag(tag string, data ...interface{}) {
	entry(tag).Warn(data)
}

func ErrorTag(tag string, data ...interface{}) {
	entry(tag).Error(data)
}

func Fatal(data ...interface{}) {
	entry().Fatal(data)
}

type DailyLogWriter struct {
	prefix string
	path   string
	file   *os.File
}

func (d *DailyLogWriter) Write(p []byte) (int, error) {
	return d.file.Write(p)
}

func (d *DailyLogWriter) init() error {
	if d.file != nil {
		_ = d.file.Close()
	}
	if len(d.path) == 0 {
		return errors.New("invalid log path")
	}
	s, err := os.Stat(d.path)
	if err != nil {
		err := os.MkdirAll(d.path, 0777)
		if err != nil {
			return err
		}
	} else {
		if !s.IsDir() {
			return errors.New("conflict log path ")
		}
	}

	t := time.Now().Format("2006-01-02")
	if d.prefix != "" {
		t = "-" + t
	}
	f := fmt.Sprintf("%s/%s%s.log", d.path, d.prefix, t)
	file, err := os.OpenFile(f, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	d.file = file
	return nil
}

func NewDailyLogWriter(prefix string, path string) (*DailyLogWriter, error) {
	d := &DailyLogWriter{
		prefix: prefix,
		path:   path,
	}
	if err := d.init(); err != nil {
		return nil, err
	}
	return d, nil
}
