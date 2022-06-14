package hubspot

// Contacts client
type Contacts struct {
	Client
}

// Contacts constructor (from Client)
func (c Client) Contacts() Contacts {
	return Contacts{
		Client: c,
	}
}

// ContactsRequest object
type ContactsRequest struct {
	Properties ContactsRequestProperty `json:"properties"`
}

type ContactsRequestProperty struct {
	Company   string `json:"company" example:"Biglytics"`
	Email     string `json:"email" example:"bcooper@biglytics.net"`
	Firstname string `json:"firstname" example:"Bryan"`
	Lastname  string `json:"lastname" example:"Cooper"`
	Phone     string `json:"phone" example:"(877) 929-0687"`
	Website   string `json:"website" example:"biglytics.net"`
}

// ContactsResponse object
type ContactsResponse struct {
	ErrorResponse
	Id         string                   `json:"id"`
	Properties ContactsResponseProperty `json:"properties"`
	CreatedAt  string                   `json:"createdAt"`
	UpdatedAt  string                   `json:"updatedAt"`
	Archived   bool                     `json:"archived"`
}

type ContactsResponseProperty struct {
	Company          string `json:"company" example:"Biglytics"`
	Createdate       string `json:"createdate" example:"2019-10-30T03:30:17.883"`
	Email            string `json:"email" example:"bcooper@biglytics.net"`
	Firstname        string `json:"firstname" example:"Bryan"`
	Lastname         string `json:"lastname" example:"Cooper"`
	Lastmodifieddate string `json:"lastmodifieddate" example:"2019-12-07T16:50:06.678Z"`
	Phone            string `json:"phone" example:"(877) 929-0687"`
	Website          string `json:"website" example:"biglytics.net"`
	ClassinAccountId string `json:"classin_account_id"`
}

// AssociatedCompany object
type AssociatedCompany struct {
	CompanyID  int                       `json:"company-id"`
	PortalID   int                       `json:"portal-id"`
	Properties AssociatedCompanyProperty `json:"properties"`
}
type ObjectValue struct {
	Value string `json:"value"`
}
type AssociatedCompanyProperty struct {
	HsNumOpenDeals               ObjectValue `json:"hs_num_open_deals"`
	FirstContactCreatedate       ObjectValue `json:"first_contact_createdate"`
	Website                      ObjectValue `json:"website"`
	HsLastmodifieddate           ObjectValue `json:"hs_lastmodifieddate"`
	HsNumDecisionMakers          ObjectValue `json:"hs_num_decision_makers"`
	NumAssociatedContacts        ObjectValue `json:"num_associated_contacts"`
	NumConversionEvents          ObjectValue `json:"num_conversion_events"`
	Domain                       ObjectValue `json:"domain"`
	HsNumChildCompanies          ObjectValue `json:"hs_num_child_companies"`
	HsNumContactsWithBuyingRoles ObjectValue `json:"hs_num_contacts_with_buying_roles"`
	HsObjectId                   ObjectValue `json:"hs_object_id"`
	Createdate                   ObjectValue `json:"createdate"`
	HsNumBlockers                ObjectValue `json:"hs_num_blockers"`
}

// CreateOrUpdateContactResponse object
type CreateOrUpdateContactResponse struct {
	ErrorResponse
	Vid   int  `json:"vid"`
	IsNew bool `json:"isNew"`
}

// DeleteContactResponse object
type DeleteContactResponse struct {
	ErrorResponse
	Vid     int    `json:"vid"`
	Deleted bool   `json:"deleted"`
	Reason  string `json:"reason"`
}

// IdentityProfile response object
type IdentityProfile struct {
	Identities []struct {
		Timestamp int64  `json:"timestamp"`
		Type      string `json:"type"`
		Value     string `json:"value"`
	} `json:"identities"`
	Vid int `json:"vid"`
}

// ContactProperty response object
type ContactProperty struct {
	Value    string `json:"value"`
	Versions []struct {
		Value       string      `json:"value"`
		Timestamp   int64       `json:"timestamp"`
		SourceType  string      `json:"source-type"`
		SourceID    interface{} `json:"source-id"`
		SourceLabel interface{} `json:"source-label"`
		Selected    bool        `json:"selected"`
	} `json:"versions"`
}

type SearchContactRequest struct {
	FilterGroups []SearchContactFilterGroups `json:"filterGroups,omitempty"`
	Sorts        string                      `json:"sorts,omitempty"`
	Query        string                      `json:"query,omitempty"`
	Properties   []string                    `json:"properties,omitempty"`
	Limit        int                         `json:"limit,omitempty"`
	After        int                         `json:"after,omitempty"`
}

type SearchContactFilterGroups struct {
	Filters []SearchContactFilter `json:"filters,omitempty"`
}

type SearchContactFilter struct {
	Value        string   `json:"value,omitempty"`
	Values       []string `json:"values,omitempty"`
	PropertyName string   `json:"propertyName,omitempty"`
	Operator     string   `json:"operator,omitempty"`
}

type SearchContactResponse struct {
	Total   int                `json:"total"`
	Results []ContactsResponse `json:"results"`
	Paging  ContactPagination  `json:"paging"`
}

type ContactPagination struct {
	Next ContactPaginationNext `json:"next"`
}
type ContactPaginationNext struct {
	After string `json:"after"`
	Link  string `json:"link"`
}

// Get Contact
func (c Contacts) Get(contactID string) (ContactsResponse, error) {
	r := ContactsResponse{}
	err := c.Client.Request("GET", "/crm/v3/objects/contacts/"+contactID, nil, &r)
	return r, err
}

// Search a Contact
func (c Contacts) Search(body SearchContactRequest) (SearchContactResponse, error) {
	r := SearchContactResponse{}
	err := c.Client.Request("POST", "/crm/v3/objects/contacts/search", body, &r)
	return r, err
}

// GetByEmail a Contact [deprecated]
func (c Contacts) GetByEmail(email string) (ContactsResponse, error) {
	r := ContactsResponse{}
	err := c.Client.Request("GET", "/contacts/v1/contact/email/"+email+"/profile", nil, &r)
	return r, err
}

// Create new Contact [deprecated]
func (c Contacts) Create(data ContactsRequest) (ContactsResponse, error) {
	r := ContactsResponse{}
	err := c.Client.Request("POST", "/crm/v3/objects/contacts", data, &r)
	return r, err
}

// Update Contact
func (c Contacts) Update(contactID string, data ContactsRequest) error {
	return c.Client.Request("PATCH", "/crm/v3/objects/contacts/"+contactID, data, nil)
}

// UpdateByEmail a Contact [deprecated]
func (c Contacts) UpdateByEmail(email string, data ContactsRequest) error {
	return c.Client.Request("POST", "/contacts/v1/contact/email/"+email+"/profile", data, nil)
}

// CreateOrUpdate a Contact [deprecated]
func (c Contacts) CreateOrUpdate(email string, data ContactsRequest) (CreateOrUpdateContactResponse, error) {
	r := CreateOrUpdateContactResponse{}
	err := c.Client.Request("POST", "/contacts/v1/contact/createOrUpdate/email/"+email, data, &r)
	return r, err
}

// Delete Contact
func (c Contacts) Delete(contactID string) error {
	return c.Client.Request("DELETE", "/crm/v3/objects/contacts/"+contactID, nil, nil)
}
