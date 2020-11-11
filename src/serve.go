package src

import (
	"errors"
	"github.com/sirupsen/logrus"
	"strings"
	"sync"
	"time"
)

// customerChan keeps a customer for serving from a unoccupied food counter
var customerChan = make(chan customer)

// serveToCustomer business logic for serving to a customer
func serveToCustomer(c *customer, foodCounterID int) (string, error) {

	if c.UserName == "" {
		return "", errors.New("unable to get username")
	}

	// Food needs some time to process
	time.Sleep(time.Millisecond * 80)

	logrus.Info("#FoodCounterID : ", foodCounterID,
		" | customer : ", c.getFullName(), " / ", c.getUsername())

	return c.getUsername(), nil
}

// foodCounterProcessor pulls customer from customerChan for food serving
func foodCounterProcessor(foodCounterID int, wg *sync.WaitGroup) {
	for job := range customerChan {
		processedUsername, err := serveToCustomer(&job, foodCounterID)
		if err == nil {
			logrus.Debug("processed - ", processedUsername)
		}
	}
	wg.Done()
}

// createFoodCounterPool prepares food counter for serving to customer
func createFoodCounterPool(numberOfCounters int) {
	var wg sync.WaitGroup
	for w := 1; w <= numberOfCounters; w++ {
		wg.Add(1)
		go foodCounterProcessor(w, &wg)
	}
	wg.Wait()
}

// allocateCustomers will push each customer to customerChan
func allocateCustomers() {
	for _, eachCustomer := range customerList {
		customerDetails := strings.Split(eachCustomer, " ")

		// customer details are present - firstname, lastname & getUsername
		if len(customerDetails) == 3 {
			customerChan <- customer{
				FirstName: customerDetails[0],
				LastName:  customerDetails[1],
				UserName:  customerDetails[2],
			}
		}
	}
	close(customerChan)
}
