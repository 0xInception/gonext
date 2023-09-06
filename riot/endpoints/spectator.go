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

type SpectatorEndpoint struct {
	apiKey      string
	rateLimiter *ratelimiter.RegionRatelimiter
	retries     int
}

func NewSpectatorEndpoint(apiKey string, rateLimits *ratelimiter.RegionRatelimiter, retries int) *SpectatorEndpoint {
	return &SpectatorEndpoint{
		apiKey:      apiKey,
		rateLimiter: rateLimits,
		retries:     retries,
	}
}

func (s *SpectatorEndpoint) GetBySummonerId(region constants.Region, summonerId string) (GetSpectatorGame, error) {
	if !s.rateLimiter.WaitRetries(region, s.retries) {
		return GetSpectatorGame{}, errors.New("rate limit")
	}
	resp, err := http.Get(fmt.Sprintf("https://%s.api.riotgames.com/lol/spectator/v4/active-games/by-summoner/%s?api_key=%s", region, summonerId, s.apiKey))
	if err != nil {
		return GetSpectatorGame{}, nil
	}
	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)
	if resp.StatusCode == http.StatusOK {
		var p GetSpectatorGame
		err = json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {
			return p, err
		}
		return p, nil
	}
	return GetSpectatorGame{}, errors.New(resp.Status)
}
