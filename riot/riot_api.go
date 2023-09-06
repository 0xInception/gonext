package riot

import (
	"time"

	"github.com/0xInception/gonext/ratelimiter"
	"github.com/0xInception/gonext/riot/endpoints"
)

type RiotApi struct {
	apiKey      string
	rateLimiter *ratelimiter.RegionRatelimiter
	Summoner    *endpoints.SummonerEndpoint
	Spectator   *endpoints.SpectatorEndpoint
	League      *endpoints.LeagueEndpoint
}

func NewRiotApi(apiKey string, rateLimits map[time.Duration]int, retries int) *RiotApi {
	limiter := ratelimiter.NewRegionRatelimiter(rateLimits)
	return &RiotApi{apiKey: apiKey, rateLimiter: limiter,
		Summoner:  endpoints.NewSummonerEndpoint(apiKey, limiter, retries),
		Spectator: endpoints.NewSpectatorEndpoint(apiKey, limiter, retries),
		League:    endpoints.NewLeagueEndpoint(apiKey, limiter, retries),
	}
}
