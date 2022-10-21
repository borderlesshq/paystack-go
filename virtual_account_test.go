package paystack

import "testing"

func TestVirtualAccountCRUD(t *testing.T) {
	cust := &Customer{
		FirstName: "User333",
		LastName:  "Admin-User",
		Email:     "user123-vAccount@gmail.com",
		Phone:     "+23400000000000001",
	}
	// create the customer
	customer, err := c.Customer.Create(cust)
	if err != nil {
		t.Errorf("CREATE Subscription Customer returned error: %v", err)
	}

	req := &VirtualAccountRequest{
		Customer: customer.CustomerCode,
	}

	vc, err := c.VirtualAccount.Create(req)
	if err != nil {
		t.Errorf("CREATE VirtualAccount returned error: %v", err)
	}

	_, err = c.VirtualAccount.Get(vc.ID)
	if err != nil {
		t.Errorf("GET VirtualAccount returned error: %v", err)
	}

	_, err = c.VirtualAccount.List(true, "NGN")
	if err != nil {
		t.Errorf("LIST VirtualAccount returned error: %v", err)
	}

	_, err = c.VirtualAccount.Deactivate(vc.ID)
	if err != nil {
		t.Errorf("DEACTIVATE VirtualAccount returned error: %v", err)
	}
}
