package main

import (
	"fmt"

	log "github.com/jeffizhungry/logrus"
)

// Test case to see if we can track caller info properly.
func callerDepth3() {
	log.Debug("Test message")
	log.Info("Test message")
	log.Warn("Test message")
	log.Error("Test message")
	fmt.Println("")
}

// Test case to see if we can track caller info properly.
func callerDepth2() {
	log.Debug("Test message")
	log.Info("Test message")
	log.Warn("Test message")
	log.Error("Test message")
	fmt.Println("")
	callerDepth3()
}

// Test case to see if we can track caller info properly.
func callerDepth1() {
	log.Debug("Test message")
	log.Info("Test message")
	log.Warn("Test message")
	log.Error("Test message")
	fmt.Println("")
	callerDepth2()
}

func panicker() {
	defer func() {
		if err := recover(); err != nil {
			log.Info("PANIC CAUGHT!!")
		}
	}()
	log.Panic("Panicking!!")
}

func fatality() {
	defer func() {
		if err := recover(); err != nil {
			log.Info("PANIC CAUGHT!!")
		}
	}()
	log.Fatal("Fatality!!")
}

func main() {
	// DO THIS ONCE AT THE VERY BEGINNING OF THE APPLICATIONS
	log.Setup(false, log.DebugLevel)

	// Show basic debugging
	log.Debug("Test message")
	log.Info("Test message")
	log.Warn("Test message")
	log.Error("Test message")
	fmt.Println("")

	log.Debugf("Test message: %v", "FORMATTED")
	log.Infof("Test message: %v", "FORMATTED")
	log.Warnf("Test message: %v", "FORMATTED")
	log.Errorf("Test message: %v", "FORMATTED")
	fmt.Println("")

	// Show caller info tracking
	callerDepth1()

	// Show field tagging
	err := fmt.Errorf("NULL ptr dereference")
	log.WithError(err).Info("Segmentation Fault")

	log.WithField("keyA", "valA").Info("Test message with ONE Fields")
	testStruct := struct {
		name string
		id   int
	}{
		name: "Jeff",
		id:   8989,
	}
	log.WithFields(log.Fields{
		"str":           "helloworld",
		"strWithSpaces": "hello world",
		"int":           1992,
		"bool":          true,
		"struct":        testStruct,
	}).Info("Test message with Fields")
	fmt.Println("")

	// Show Panic and Fatal logging
	log.Info("Continuting after panic")
	panicker()
	fatality()
	log.Error("INVALID STATE, EXAMPLE SHOULD HAVE EXITED AFTER FATAL")
}
