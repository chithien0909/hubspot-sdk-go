package hubspot

import (
	"testing"
)

func TestDeals(t *testing.T) {

	data := DealsRequest{
		Properties: DealsProperty{
			Amount:    "1500.00",
			Closedate: "2019-12-07T16:50:06.678Z",
			Dealname:  "Custom data integrations",
			Pipeline:  "default",
		},
	}
	c := NewClient(NewClientConfig("", ""))
	// Create new deal
	r, err := c.Deals().Create(data)

	if err != nil {
		t.Error(err)
	}
	if r.ErrorResponse.Status == "error" {
		t.Error(r.ErrorResponse.Message)
	}

	if r.Id != "" {
		// Get Deal by ID
		r, err = c.Deals().Get(r.Id)
		if err != nil {
			t.Error(err)
		}
		if r.ErrorResponse.Status == "error" {
			t.Error(r.ErrorResponse.Message)
		}
	}

	data.Properties.Dealname = data.Properties.Dealname + " updated"

	if r.Id != "" {
		// Update deal by id
		r, err = c.Deals().Update(r.Id, data)
		if err != nil {
			t.Error(err)
		}
		if r.ErrorResponse.Status == "error" {
			t.Error(r.ErrorResponse.Message)
		}
	}

	if r.Id != "" {
		// Delete deal By Id
		err = c.Deals().Delete(r.Id)
		if err != nil {
			t.Error(err)
		}
	}

	t.Logf("%+v", r)
}
