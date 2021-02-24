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

//Fetches all current, puplic IP ranges in cidr notation from AWS and transforms them to a plain text format.
//
//The format is very simple and matches pfSense requirements for a URL table Alias:
//One IP range in cidr notation per line.
func FetchAwsCidrs() (string, error) {
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
		return awsCidrs.extractIpRanges(), nil
	}
	return "", errors.Errorf("aws ip-ranges request returns status code %s", resp.Status)
}
