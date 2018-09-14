package helper

import (
	"log/syslog"
	"runtime"
	"path"
	"strconv"

	"github.com/astaxie/beego"
)

func LogDebug(msg string) error {
	return logPaperTrail(msg, "debug")
}

func LogInfo(msg string) error {
	return logPaperTrail(msg, "info")
}

func LogWarning(msg string) error {
	return logPaperTrail(msg, "warning")
}

func LogError (msg string) error {
	return logPaperTrail(msg, "error")
}

func logPaperTrail (msg string, tags string) (error) {
	w, err := syslog.Dial("udp", beego.AppConfig.String("papertrail_url"), syslog.LOG_EMERG | syslog.LOG_KERN, tags)
	if err != nil {
		beego.Error("failed to dial syslog")
		return err
	}

	// Skip 2 because there are bridge function
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	msg = "[" + filename + ":" + strconv.Itoa(line) + "] " + msg

	if tags == "info" {
		beego.Info(msg)
	} else if tags == "debug" {
		beego.Debug(msg)
	} else if tags == "warning" {
		beego.Warn(msg)
	} else if tags == "error" {
		beego.Error(msg)
	}

	// Only log in production
	if beego.BConfig.RunMode == "prod" {
		w.Info(msg)
	}

	return nil
}
