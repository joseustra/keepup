package cfgo

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	// ErrSameIP returns when the current ip is alread set on the local storage
	ErrSameIP = errors.New("current and old ip are the same, no need to update")
)

// UpdateRecord updates the DNS record
func UpdateRecord(client Client, zone, dns, ip string) error {
	re := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	if !re.MatchString(ip) {
		return fmt.Errorf("invalid ip address, too short: %s", ip)
	}

	if !strings.Contains(dns, zone) {
		dns = fmt.Sprintf("%s.%s", dns, zone)
	}

	domain, err := client.GetDNSRecord(zone, dns)
	if err != nil {
		return err
	}

	if domain.IP != ip {
		domain.IP = strings.TrimSpace(ip)
		err = client.UpdateDNSRecord(domain)
		if err != nil {
			return err
		}
	}

	fmt.Printf("DNS record updated: %s %s\n", domain.DNS, domain.IP)

	return nil
}
