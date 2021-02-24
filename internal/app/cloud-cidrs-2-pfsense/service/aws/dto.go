package aws

import "fmt"

//Response DTO for AWS endpoint (https://ip-ranges.amazonaws.com/ip-ranges.json)
type AwsCidrs struct {
	SyncToken  string `json:"syncToken"`
	CreateDate string `json:"createDate"`
	Prefixes   []struct {
		IPPrefix           string `json:"ip_prefix"`
		Region             string `json:"region"`
		Service            string `json:"service"`
		NetworkBorderGroup string `json:"network_border_group"`
	} `json:"prefixes"`
}

func (this AwsCidrs) extractIpRanges() string {
	ipRanges := ""
	for _, prefix := range this.Prefixes {
		ipRanges += fmt.Sprintf("%s\n", prefix.IPPrefix)
	}
	return ipRanges
}