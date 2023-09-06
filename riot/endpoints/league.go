package endpoints

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/0xInception/gonext/ratelimiter"
	"github.com/0xInception/gonext/riot/constants"
)

type LeagueEndpoint struct {
	apiKey      string
	rateLimiter *ratelimiter.RegionRatelimiter
	retries     int
}

func NewLeagueEndpoint(apiKey string, rateLimits *ratelimiter.RegionRatelimiter, retries int) *LeagueEndpoint {
	return &LeagueEndpoint{
		apiKey:      apiKey,
		rateLimiter: rateLimits,
		retries:     retries,
	}
}

func (s *LeagueEndpoint) GetBySummonerId(region constants.Region, summonerId string) (GetLeague, error) {
	if !s.rateLimiter.WaitRetries(region, s.retries) {
		return GetLeague{}, errors.New("rate limit")
	}
	resp, err := http.Get(fmt.Sprintf("https://%s.api.riotgames.com/lol/league/v4/entries/by-summoner/%s?api_key=%s", region, summonerId, s.apiKey))
	if err != nil {
		return GetLeague{}, nil
	}
	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)
	if resp.StatusCode == http.StatusOK {
		var p GetLeague
		err = json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {
			return p, err
		}
		return p, nil
	}
	return GetLeague{}, errors.New(resp.Status)
}
