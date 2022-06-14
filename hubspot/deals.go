package hubspot

import "fmt"

// Deals client
type Deals struct {
	Client
}

// Deals constructor (from Client)
func (c Client) Deals() Deals {
	return Deals{
		Client: c,
	}
}

// DealsRequest object
type DealsRequest struct {
	Properties DealsProperty `json:"properties"`
}

// DealsProperty object
type DealsProperty struct {
	Amount         string `json:"amount"`
	Closedate      string `json:"closedate"`
	Dealname       string `json:"dealname"`
	Dealstage      string `json:"dealstage"`
	HubspotOwnerId string `json:"hubspot_owner_id"`
	Pipeline       string `json:"pipeline"`
}

// DealsResponse object
type DealsResponse struct {
	ErrorResponse
	Id           string         `json:"id"`
	Properties   DealProperties `json:"properties"`
	CreatedAt    string         `json:"createdAt"`
	UpdatedAt    string         `json:"updatedAt"`
	Archived     bool           `json:"archived"`
	Associations Associations   `json:"associations"`
}

// DealProperties object
type DealProperties struct {
	Amount             string `json:"amount"`
	Closedate          string `json:"closedate"`
	Createdate         string `json:"createdate"`
	Dealname           string `json:"dealname"`
	Dealstage          string `json:"dealstage"`
	HsLastmodifieddate string `json:"hs_lastmodifieddate"`
	HubspotOwnerID     string `json:"hubspot_owner_id"`
	Pipeline           string `json:"pipeline"`
}

// Get Deal
func (d Deals) Get(dealID string) (DealsResponse, error) {
	r := DealsResponse{}
	err := d.Client.Request("GET", fmt.Sprintf(
		"/crm/v3/objects/deals/%s", dealID), nil, &r, []string{
		"associations=contacts",
		"associations=line_items",
		"archived=false",
	})
	return r, err
}

// Create new Deal
func (d Deals) Create(data DealsRequest) (DealsResponse, error) {
	r := DealsResponse{}
	err := d.Client.Request("POST", "/crm/v3/objects/deals", data, &r, nil)
	return r, err
}

// Update Deal
func (d Deals) Update(dealID string, data DealsRequest) (DealsResponse, error) {
	r := DealsResponse{}
	err := d.Client.Request("PATCH", "/crm/v3/objects/deals/"+dealID, data, &r, nil)
	return r, err
}

// Delete Deal
func (d Deals) Delete(dealID string) error {
	err := d.Client.Request("DELETE", "/crm/v3/objects/deals/"+dealID, nil, nil, nil)
	return err
}
