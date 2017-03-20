package cfgo

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrSameIP returns when the current ip is alread set on the local storage
	ErrSameIP = errors.New("current and old ip are the same, no need to update")
)

// UpdateRecord updates the DNS record
func UpdateRecord(db Storage, client Client, zone, dns, ip string, force bool) error {
	var err error
	var currentIP string
	domain := &Domain{}

	if len(ip) >= 7 {
		currentIP = ip
	} else {
		currentIP, err = GetIPV4IP()
		if err != nil {
			return err
		}
	}

	if !force {
		domain, err = db.Find(fmt.Sprintf("%s-%s", zone, dns))
		if err != nil {
			return err
		}

		if len(domain.IP) >= 7 {
			if domain.IP == currentIP {
				return ErrSameIP
			}
		}
	}

	domain, err = client.GetDNSRecord(zone, dns)
	if err != nil {
		return err
	}

	if domain.IP != currentIP {
		domain.IP = strings.TrimSpace(currentIP)
		err = client.UpdateDNSRecord(domain)
		if err != nil {
			return err
		}
	}

	err = db.Save(domain)
	if err != nil {
		return err
	}

	fmt.Printf("DNS record updated: %s %s\n", domain.DNS, domain.IP)

	return nil
}
