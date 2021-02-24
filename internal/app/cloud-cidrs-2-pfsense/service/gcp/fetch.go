package gcp

import (
	"fmt"
	"github.com/pkg/errors"
	"net"
	"strings"
)

var (
	googleNetblocksDomains = []string{
		"_netblocks.google.com",
		"_netblocks2.google.com",
		"_netblocks3.google.com",
	}
)

//Fetches all current, puplic IP ranges in cidr notation from GCP and transforms them to a plain text format.
//
//The format is very simple and matches pfSense requirements for a URL table Alias:
//One IP range in cidr notation per line.
func FetchGcpCidrs() (string, error) {
	cidrs := ""
	for _, netblocksDomain := range googleNetblocksDomains {
		cidrsOfDomain, err := fetchSingle(netblocksDomain)
		if err != nil {
			return "", err
		}
		cidrs += cidrsOfDomain
	}
	return cidrs, nil
}

func fetchSingle(domain string) (string, error) {
	txts, err := net.LookupTXT(domain)
	if err != nil {
		return "", err
	}
	if len(txts) == 0 {
		return "", errors.Errorf("no txt lookup data for domain: ", domain)
	}
	cidrsOfDomain := ""
	for _, txt := range txts {
		for _, token := range strings.Split(txt, " ") {
			ip4Tokens := strings.Split(token, "ip4:")
			if len(ip4Tokens) == 2 {
				cidrsOfDomain += fmt.Sprintf("%s\n", ip4Tokens[1])
			}
		}
	}
	return cidrsOfDomain, nil
}
