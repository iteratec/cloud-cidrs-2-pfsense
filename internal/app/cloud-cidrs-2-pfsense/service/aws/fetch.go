package aws

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
)

var (
	AwsIpRangesUrl = "https://ip-ranges.amazonaws.com/ip-ranges.json"
)

type filterRegions []string

func (this filterRegions) contains(region string) bool {
	for _, r := range this {
		if r == region {
			return true
		}
	}
	return false
}

//Fetches all current, puplic IP ranges in cidr notation from AWS and transforms them to a plain text format.
//
//The format is very simple and matches pfSense requirements for a URL table Alias:
//One IP range in cidr notation per line.
//If regions parameter is set and contains regions, just the IP ranges for that regions are returned.
func FetchAwsCidrs(regions *[]string) (string, error) {
	logger := log.With().Str("url", AwsIpRangesUrl).Logger()
	resp, err := http.Get(AwsIpRangesUrl)
	if err != nil {
		logger.Error().Err(err).Msg("couldn't get aws cidrs")
		return "", err
	}
	if resp.StatusCode == http.StatusOK {
		var awsCidrs AwsCidrs
		err = json.NewDecoder(resp.Body).Decode(&awsCidrs)
		if err != nil {
			logger.Error().Err(err).Msg("couldn't parse response")
		}
		return awsCidrs.extractIpRanges((*filterRegions)(regions)), nil
	}
	return "", errors.Errorf("aws ip-ranges request returns status code %s", resp.Status)
}
