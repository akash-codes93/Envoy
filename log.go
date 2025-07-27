package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLog() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}
