package cfgo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Domain a single domain on the dns
type Domain struct {
	ZoneID   string // the identifier of the zone on cloudflare
	RecordID string // the identifer of the record on cloudflare
	Type     string // valid values: A, AAAA, CNAME, TXT, SRV, LOC, MX, NS, SPF
	Zone     string // cloudflare zone
	DNS      string // cloudflare dns (the domain that will be handled)
	IP       string // the ip of the domain
}

// GetIPV4IP gets the current external ipv4 ip
func GetIPV4IP() (string, error) {
	return getExternalIP("ipv4")
}

// GetIPV6IP gets the current external ipv6 ip
func GetIPV6IP() (string, error) {
	return getExternalIP("ipv6")
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

	return strings.TrimSpace(string(ip)), nil
}
