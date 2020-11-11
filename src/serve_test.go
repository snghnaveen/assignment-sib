package src

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServeToCustomer(t *testing.T) {
	is := assert.New(t)

	// Positive case testing
	processed1, err1 := serveToCustomer(&customer{
		FirstName: "testfirst",
		LastName:  "testlast",
		UserName:  "testusername",
	}, 1)

	is.NoError(err1)
	is.Equal(processed1, "testusername")

	// Negative case testing
	processed2, err2 := serveToCustomer(&customer{
		FirstName: "testfirst",
		LastName:  "testlast",
		UserName:  "",
	}, 1)
	is.Error(err2)
	is.Equal(processed2, "")

}
