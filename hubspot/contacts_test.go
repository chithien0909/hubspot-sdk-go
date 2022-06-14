package hubspot

import (
	"github.com/leonelquinteros/gorand"
	"strings"
	"testing"
)

func TestContacts(t *testing.T) {
	tEmailUser, err := gorand.GetAlphaNumString(8)
	if err != nil {
		t.Fatal(err)
	}
	tCompanyName, err := gorand.GetAlphaNumString(9)
	if err != nil {
		t.Fatal(err)
	}
	tEmailHost, err := gorand.GetAlphaNumString(6)
	if err != nil {
		t.Fatal(err)
	}
	tPhone, err := gorand.GetNumString(10)
	if err != nil {
		t.Fatal(err)
	}
	tEmail := tEmailUser + "@" + tEmailHost + ".com"
	if err != nil {
		t.Fatal(err)
	}
	data := ContactsRequest{
		Properties: ContactsRequestProperty{
			Company:   tCompanyName,
			Email:     tEmail,
			Firstname: tEmailUser,
			Lastname:  tEmailUser,
			Phone:     tPhone,
			Website:   strings.ToLower(tCompanyName) + ".net",
		},
	}
	// Create
	c := NewClient(NewClientConfig("", ""))
	r, err := c.Contacts().Create(data)

	if err != nil {
		t.Fatal(err)
	}
	if r.ErrorResponse.Status == "error" {
		t.Fatal(r.ErrorResponse.Message)
	}

	//Update
	data.Properties.Company = data.Properties.Company + "1"
	if r.Id != "" {
		err = c.Contacts().Update(r.Id, data)
		if err != nil {
			t.Fatal(err)
		}
		if r.ErrorResponse.Status == "error" {
			t.Fatal(r.ErrorResponse.Message)
		}
	}
	// Get contact
	contact, err := c.Contacts().Get(r.Id)
	if err != nil {
		t.Fatal(err)
	}
	if r.ErrorResponse.Status == "error" {
		t.Fatal(r.ErrorResponse.Message)
	}

	// Delete
	err = c.Contacts().Delete(contact.Id)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", r)
	t.Logf("%+v", contact)
}
