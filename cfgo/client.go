package cfgo

// Client defines how to client will update records on DNS
type Client interface {
	GetDNSRecord(zone, dns string) (*Domain, error)
	UpdateDNSRecord(domain *Domain) error
}
