package hubspot

import "fmt"

// LineItems client
type LineItems struct {
	Client
}

// LineItems constructor (from Client)
func (c Client) LineItems() LineItems {
	return LineItems{
		Client: c,
	}
}

// LineItemsRequest object
type LineItemsRequest struct {
	Properties LineItemsProperty `json:"properties"`
}

// LineItemsProperty object
type LineItemsProperty struct {
	Name                      string `json:"name"`
	HsProductId               string `json:"hs_product_id"`
	HsRecurringBillingPeriod  string `json:"hs_recurring_billing_period"`
	Recurringbillingfrequency string `json:"recurringbillingfrequency"`
	Quantity                  string `json:"quantity"`
	Price                     string `json:"price"`
}

// LineItemsResponse object
type LineItemsResponse struct {
	ErrorResponse
	Id         string            `json:"id"`
	Properties LineItemsProperty `json:"properties"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Archived   bool              `json:"archived"`
}

// LineItemResponse object
type LineItemResponse struct {
	ErrorResponse
	Properties LineItemsProperty `json:"properties"`
}

// Get Line Items
func (l LineItems) Get(lineItemId string) (LineItemResponse, error) {
	r := LineItemResponse{}
	err := l.Client.Request("GET", fmt.Sprintf("/crm/v3/objects/line_items/%v", lineItemId), nil, &r)
	return r, err
}

// Create new Line Items
func (l LineItems) Create(data LineItemsRequest) (LineItemsResponse, error) {
	r := LineItemsResponse{}
	err := l.Client.Request("POST", "/crm/v3/objects/line_items", data, &r)
	return r, err
}

// Update Line Items
func (l LineItems) Update(lineItem string, data LineItemsRequest) (LineItemsResponse, error) {
	r := LineItemsResponse{}
	err := l.Client.Request("PATCH", "/crm/v3/objects/line_items/"+lineItem, data, &r)
	return r, err
}

// Delete Deal
func (l LineItems) Delete(lineItemId string) error {
	err := l.Client.Request("DELETE", "/crm/v3/objects/line_items/"+lineItemId, nil, nil)
	return err
}
