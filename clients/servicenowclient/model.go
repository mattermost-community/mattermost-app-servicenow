package servicenowclient

import "time"

type BaseTicket struct {
	ID         string       `json:"sys_id"`
	ClassName  string       `json:"sys_class_name"`
	Tags       string       `json:"sys_tags"`
	Domain     TicketDomain `json:"sys_domain"`
	DomainPath string       `json:"sys_domain_path"`
	ModCount   int          `json:"sys_mod_count"`
	UpdatedOn  time.Time    `json:"sys_updated_on"`
	UpdatedBy  string       `json:"sys_updated_by"`
	CreatedOn  time.Time    `json:"sys_created_on"`
	CreatedBy  string       `json:"sys_created_by"`
}

type TicketDomain struct {
	Link  string `json:"link"`
	Value string `json:"value"`
}

type CreateTicketResponse struct {
	Result BaseTicket `json:"result"`
}
