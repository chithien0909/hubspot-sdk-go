package hubspot

import "fmt"

// Products client
type Products struct {
	Client
}

// Products constructor (from Client)
func (c Client) Products() Products {
	return Products{
		Client: c,
	}
}

// ProductsRequest object
type ProductsRequest struct {
	Properties ProductsProperty `json:"properties"`
}

// ProductsProperty object
type ProductsProperty struct {
	Description              string `json:"description"`
	HsCostOfGoodsSold        string `json:"hs_cost_of_goods_sold"`
	HsRecurringBillingPeriod string `json:"hs_recurring_billing_period"`
	HsSku                    string `json:"hs_sku"`
	Name                     string `json:"name"`
	Price                    string `json:"price"`
}

// ProductsResponse object
type ProductsResponse struct {
	ErrorResponse
	Id         string           `json:"id"`
	Properties ProductsProperty `json:"properties"`
	CreatedAt  string           `json:"createdAt"`
	UpdatedAt  string           `json:"updatedAt"`
	Archived   bool             `json:"archived"`
}

// ProductResponse object
type ProductResponse struct {
	Properties ProductsProperty `json:"properties"`
}

// ProductsPropertyResponse object
type ProductsPropertyResponse struct {
	ProductsProperty
	Createdate         string `json:"createdate"`
	HsLastmodifieddate string `json:"hs_lastmodifieddate"`
}

// Get products
func (l Products) Get(productId string) (ProductsResponse, error) {
	r := ProductsResponse{}
	err := l.Client.Request("GET", fmt.Sprintf("/crm/v3/objects/products/%v", productId), nil, &r)
	return r, err
}

// Create new products
func (l Products) Create(data ProductsRequest) (ProductsResponse, error) {
	r := ProductsResponse{}
	err := l.Client.Request("POST", "/crm/v3/objects/products", data, &r)
	return r, err
}

// Update product
func (l Products) Update(productId string, data ProductsRequest) (ProductsResponse, error) {
	r := ProductsResponse{}
	err := l.Client.Request("PATCH", "/crm/v3/objects/products/"+productId, data, &r)
	return r, err
}

// Delete product
func (l Products) Delete(productId string) error {
	err := l.Client.Request("DELETE", "/crm/v3/objects/products/"+productId, nil, nil)
	return err
}
