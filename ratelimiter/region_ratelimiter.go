package ratelimiter

import (
	"time"

	"github.com/0xInception/gonext/riot/constants"
)

type RegionRatelimiter struct {
	rateLimiter map[constants.Region]*RateLimiter
}

func NewRegionRatelimiter(rateLimits map[time.Duration]int) *RegionRatelimiter {
	regions := []constants.Region{
		constants.BR, constants.EUNE, constants.EUW, constants.JP, constants.KR, constants.LAN,
		constants.LAS, constants.NA, constants.OCE, constants.PBE, constants.PH, constants.RU,
		constants.SG, constants.TH, constants.TR, constants.TW, constants.VN,
	}

	res := RegionRatelimiter{
		rateLimiter: make(map[constants.Region]*RateLimiter),
	}
	for _, x := range regions {
		res.rateLimiter[x] = NewRateLimiter(rateLimits)
	}
	return &res
}

func (x *RegionRatelimiter) WaitRetries(region constants.Region, retries int) bool {
	rl, contains := x.rateLimiter[region]
	if !contains {
		return false
	}
	return rl.WaitRetries(retries)
}
