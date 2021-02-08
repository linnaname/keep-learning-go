package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {

	log.SetFormatter(&log.TextFormatter{})

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	//log.WithFields(log.Fields{
	//	"omg":    true,
	//	"number": 100,
	//}).Fatal("The ice breaks!")

	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")

	req, _ := http.NewRequest("Get", "127.0.0.1", nil)
	log.WithFields(log.Fields{
		"req": req,
	}).Error("new request failed")
}
