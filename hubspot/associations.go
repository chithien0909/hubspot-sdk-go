package hubspot

type Associations struct {
	LineItems AssociationsLineItems `json:"line items"`
	Contacts  AssociationsLineItems `json:"contacts"`
}

type AssociationsLineItems struct {
	Results []AssociationsLineItemsResult `json:"results"`
}

type AssociationsLineItemsResult struct {
	Id   string `json:"id"`
	Type string `json:"deal_to_line_item"`
}
type AssociationsContact struct {
	Results []AssociationsContactResult `json:"results"`
}
type AssociationsContactResult struct {
	Id   string `json:"id"`
	Type string `json:"deal_to_line_item"`
}
