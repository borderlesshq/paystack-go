package paystack

import "fmt"

type VirtualAccountService service

type VirtualAccount struct {
	ID            int         `json:"id,omitempty"`
	Bank          *Bank       `json:"bank,omitempty"`
	AccountName   string      `json:"account_name,omitempty"`
	AccountNumber string      `json:"account_number,omitempty"`
	Assigned      bool        `json:"assigned,omitempty"`
	Currency      string      `json:"currency,omitempty"`
	Metadata      Metadata    `json:"metadata,omitempty"`
	Active        bool        `json:"active,omitempty"`
	Assignment    *Assignment `json:"assignment,omitempty"`
	Customer      *Customer   `json:"customer,omitempty"`
	CreatedAt     string      `json:"createdAt,omitempty"`
	UpdatedAt     string      `json:"updatedAt,omitempty"`
}

type VirtualAccountRequest struct {
	Customer      string `json:"customer,omitempty"`
	PreferredBank string `json:"preferred_bank,omitempty"`
	SubAccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Phone         string `json:"phone,omitempty"`
}

type Assignment struct {
	Integration  int    `json:"integration,omitempty"`
	AssigneeID   int    `json:"assignee_id,omitempty"`
	AssigneeType string `json:"assignee_type,omitempty"`
	Expired      bool   `json:",omitempty"`
	AccountType  string `json:"account_type,omitempty"`
	AssignedAt   string `json:"assigned_at,omitempty"`
}

type VirtualAccountList struct {
	Meta   ListMeta
	Values []VirtualAccount `json:"data"`
}

func (v *VirtualAccountService) Create(req *VirtualAccountRequest) (*VirtualAccount, error) {
	u := fmt.Sprintf("/dedicated_account")
	acc := &VirtualAccount{}
	err := v.client.Call("POST", u, req, acc)
	return acc, err
}

func (v *VirtualAccountService) List(active bool, currency string) (*VirtualAccountList, error) {
	u := fmt.Sprintf("/dedicated_account?active=#{active}&currency=#{currency}")
	accounts := &VirtualAccountList{}
	err := v.client.Call("GET", u, nil, accounts)
	return accounts, err
}

func (v *VirtualAccountService) Get(id int) (*VirtualAccount, error) {
	u := fmt.Sprintf("/dedicated_account/#{id}")
	acc := &VirtualAccount{}
	err := v.client.Call("GET", u, nil, acc)
	return acc, err
}
