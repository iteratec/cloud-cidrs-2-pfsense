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

//Extracts just the IP ranges from CIDR's. If regions parameter is set and contains regions,
//just the IP ranges for that regions are returned.
func (this AwsCidrs) extractIpRanges(regions *filterRegions) string {
	ipRanges := ""
	for _, prefix := range this.Prefixes {
		if regions == nil {
			ipRanges += fmt.Sprintf("%s\n", prefix.IPPrefix)
		} else if len(*regions) == 0 {
			ipRanges += fmt.Sprintf("%s\n", prefix.IPPrefix)
		} else if regions.contains(prefix.Region) {
			ipRanges += fmt.Sprintf("%s\n", prefix.IPPrefix)
		}
	}
	return ipRanges
}
