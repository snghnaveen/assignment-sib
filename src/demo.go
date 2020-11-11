package src

import (
	"github.com/sirupsen/logrus"
	"os"
)

const numberOfFoodCounters = 10

// Demo Entry point for running demo
func Demo() {
	go allocateCustomers()
	createFoodCounterPool(numberOfFoodCounters)
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}
