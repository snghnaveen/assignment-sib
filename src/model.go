package src

// customer
type customer struct {
	FirstName string
	LastName  string
	UserName  string
}

// getFullName returns full name of customer
func (c *customer) getFullName() string {
	return c.FirstName + " " + c.LastName
}

// getUsername returns getUsername of customer
func (c *customer) getUsername() string {
	return c.UserName
}
