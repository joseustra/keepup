package cfgo

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cloudflare/cloudflare-go"
)

// Client the client that will handle the Cloudflare api
type Client struct {
	API *cloudflare.API
}

// NewClient ...
func NewClient(cfKey, cfEmail string) (*Client, error) {
	api, err := cloudflare.New(cfKey, cfEmail)
	if err != nil {
		return nil, err
	}
	client := &Client{API: api}

	return client, nil
}

// GetIPV4IP gets the current external ipv4 ip
func GetIPV4IP() (string, error) {
	return getExternalIP("ipv4")
}

// GetIPV6IP gets the current external ipv6 ip
func GetIPV6IP() (string, error) {
	return getExternalIP("ipv6")
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
func (c *Client) GetDNSRecord(zone, dns string) (cloudflare.DNSRecord, error) {
	zoneID, err := c.API.ZoneIDByName(zone)
	if err != nil {
		return cloudflare.DNSRecord{}, err
	}

	rr := cloudflare.DNSRecord{Name: dns}
	rrs, err := c.API.DNSRecords(zoneID, rr)
	if err != nil {
		return cloudflare.DNSRecord{}, err
	}

	return rrs[0], nil
}

// UpdateDNSRecord updates the Cloudflare DNSRecord
func (c *Client) UpdateDNSRecord(record cloudflare.DNSRecord) error {
	return c.API.UpdateDNSRecord(record.ZoneID, record.ID, record)
}

func getExternalIP(p string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s.myexternalip.com/raw", p))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}
