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

type SummonerEndpoint struct {
	apiKey      string
	rateLimiter *ratelimiter.RegionRatelimiter
	retries     int
}

func NewSummonerEndpoint(apiKey string, rateLimits *ratelimiter.RegionRatelimiter, retries int) *SummonerEndpoint {
	return &SummonerEndpoint{
		apiKey:      apiKey,
		rateLimiter: rateLimits,
		retries:     retries,
	}
}

func (s *SummonerEndpoint) GetSummonerByName(region constants.Region, name string) (GetSummonerResponse, error) {
	if !s.rateLimiter.WaitRetries(region, s.retries) {
		return GetSummonerResponse{}, errors.New("rate limit")
	}
	resp, err := http.Get(fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s?api_key=%s", region, name, s.apiKey))
	if err != nil {
		return GetSummonerResponse{}, nil
	}
	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)
	if resp.StatusCode == http.StatusOK {
		var p GetSummonerResponse
		err = json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {
			return p, err
		}
		return p, nil
	}
	return GetSummonerResponse{}, errors.New(resp.Status)
}

func (s *SummonerEndpoint) GetSummonerBySummonerId(region constants.Region, summonerId string) (GetSummonerResponse, error) {
	if !s.rateLimiter.WaitRetries(region, s.retries) {
		return GetSummonerResponse{}, errors.New("rate limit")
	}
	resp, err := http.Get(fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/v4/summoners/%s?api_key=%s", region, summonerId, s.apiKey))
	if err != nil {
		return GetSummonerResponse{}, nil
	}
	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)
	if resp.StatusCode == http.StatusOK {
		var p GetSummonerResponse
		err = json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {
			return p, err
		}
		return p, nil
	}
	return GetSummonerResponse{}, errors.New(resp.Status)
}

func (s *SummonerEndpoint) GetSummonerByPuuid(region constants.Region, puuid string) (GetSummonerResponse, error) {
	if !s.rateLimiter.WaitRetries(region, s.retries) {
		return GetSummonerResponse{}, errors.New("rate limit")
	}
	resp, err := http.Get(fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/by-puuid/%s?api_key=%s", region, puuid, s.apiKey))
	if err != nil {
		return GetSummonerResponse{}, nil
	}
	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)
	if resp.StatusCode == http.StatusOK {
		var p GetSummonerResponse
		err = json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {
			return p, err
		}
		return p, nil
	}
	return GetSummonerResponse{}, errors.New(resp.Status)
}

func (s *SummonerEndpoint) GetSummonerByAccountId(region constants.Region, accountId string) (GetSummonerResponse, error) {
	if !s.rateLimiter.WaitRetries(region, s.retries) {
		return GetSummonerResponse{}, errors.New("rate limit")
	}
	resp, err := http.Get(fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/by-account/%s?api_key=%s", region, accountId, s.apiKey))
	if err != nil {
		return GetSummonerResponse{}, nil
	}
	defer resp.Body.Close()
	defer io.Copy(io.Discard, resp.Body)
	if resp.StatusCode == http.StatusOK {
		var p GetSummonerResponse
		err = json.NewDecoder(resp.Body).Decode(&p)
		if err != nil {
			return p, err
		}
		return p, nil
	}
	return GetSummonerResponse{}, errors.New(resp.Status)
}
