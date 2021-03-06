package servicenowclient

type BaseTicket struct {
	ID         string       `json:"sys_id"`
	ClassName  string       `json:"sys_class_name"`
	Tags       string       `json:"sys_tags"`
	Domain     TicketDomain `json:"sys_domain"`
	DomainPath string       `json:"sys_domain_path"`
	ModCount   string       `json:"sys_mod_count"`
	UpdatedBy  string       `json:"sys_updated_by"`
	CreatedBy  string       `json:"sys_created_by"`
	// UpdatedOn  time.Time    `json:"sys_updated_on"`
	// CreatedOn  time.Time    `json:"sys_created_on"`
}

type TicketDomain struct {
	Link  string `json:"link"`
	Value string `json:"value"`
}

type CreateTicketResponse struct {
	Result BaseTicket `json:"result"`
}
