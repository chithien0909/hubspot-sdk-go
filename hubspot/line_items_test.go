package hubspot

import (
	"testing"
)

func TestLineItems(t *testing.T) {
	// Create product for testing
	dataProduct := ProductsRequest{
		Properties: ProductsProperty{
			Description:       "Onboarding service for data product",
			HsCostOfGoodsSold: "600.00",
			HsSku:             "191902",
			Name:              "Implementation Service",
			Price:             "6000.00",
		},
	}
	c := NewClient(NewClientConfig("", ""))
	// Create new line items
	rProduct, err := c.Products().Create(dataProduct)

	data := LineItemsRequest{
		Properties: LineItemsProperty{
			Name:        "1 year implementation consultation",
			HsProductId: rProduct.Id,
			//HsRecurringBillingPeriod: "24",
			Recurringbillingfrequency: "monthly",
			Quantity:                  "2",
			Price:                     "6000.00",
		},
	}
	// Create new line items
	r, err := c.LineItems().Create(data)

	if err != nil {
		t.Error(err)
	}
	if r.ErrorResponse.Status == "error" {
		t.Error(r.ErrorResponse.Message)
	}

	if r.Id != "" {
		// Get Deal by ID
		lineItem, err := c.LineItems().Get(r.Id)
		if err != nil {
			t.Error(err)
		}
		if lineItem.ErrorResponse.Status == "error" {
			t.Error(r.ErrorResponse.Message)
		}
	}

	data.Properties.Name = data.Properties.Name + " updated"

	if r.Id != "" {
		// Update line item by id
		r, err = c.LineItems().Update(r.Id, data)
		if err != nil {
			t.Error(err)
		}
		if r.ErrorResponse.Status == "error" {
			t.Error(r.ErrorResponse.Message)
		}
	}

	if r.Id != "" {
		// Delete line items by Id
		err = c.LineItems().Delete(r.Id)
		if err != nil {
			t.Error(err)
		}
	}

	t.Logf("%+v", r)

	// Clear product after testing
	if rProduct.Id != "" {
		// Delete product by Id
		err = c.Products().Delete(rProduct.Id)
		if err != nil {
			t.Error(err)
		}
	}

}
