package cfgo

import (
	"github.com/cloudflare/cloudflare-go"
)

// CloudflareClient the Client that will handle the Cloudflare api
type CloudflareClient struct {
	API *cloudflare.API
}

// NewCloudflareClient creates a new client for the cloudflare api
func NewCloudflareClient(cfKey, cfEmail string) (*CloudflareClient, error) {
	api, err := cloudflare.New(cfKey, cfEmail)
	if err != nil {
		return nil, err
	}

	return &CloudflareClient{API: api}, nil
}

// GetDNSRecord gets the Cloudflare DNS record for the given value
//
// zone is the cloudflare domain you are working on
// eg: domain.com
//
// dns is the address you want get the DNSRecord
// eg: my.domain.com
//
// The first matching DNSRecord will be returned
func (c *CloudflareClient) GetDNSRecord(zone, dns string) (*Domain, error) {
	zoneID, err := c.API.ZoneIDByName(zone)
	if err != nil {
		return &Domain{}, err
	}

	rr := cloudflare.DNSRecord{Name: dns}
	rrs, err := c.API.DNSRecords(zoneID, rr)
	if err != nil {
		return &Domain{}, err
	}
	record := rrs[0]

	return &Domain{
		RecordID: record.ID,
		ZoneID:   record.ZoneID,
		Type:     record.Type,
		Zone:     record.ZoneName,
		DNS:      record.Name,
		IP:       record.Content,
	}, nil
}

// UpdateDNSRecord updates the Cloudflare DNSRecord
func (c *CloudflareClient) UpdateDNSRecord(domain *Domain) error {
	record := cloudflare.DNSRecord{
		ID:      domain.RecordID,
		ZoneID:  domain.ZoneID,
		Content: domain.IP,
	}

	return c.API.UpdateDNSRecord(domain.ZoneID, domain.RecordID, record)
}
