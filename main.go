package main

import (
	"fmt"
	"os"
	"time"

	"github.com/0xInception/gonext/riot"
	"github.com/0xInception/gonext/riot/constants"
)

func main() {
	api := riot.NewRiotApi(os.Getenv("riotkey"), map[time.Duration]int{
		time.Second:     2,
		time.Minute * 2: 95,
	}, 8)
	g := time.Now()
	for i := 0; i < 10000; i++ {
		_, err := api.Summoner.GetSummonerByName(constants.EUW, "uciu1")
		if err != nil {
			fmt.Println(err, i)
			i -= 1
			continue
		}
		println(fmt.Sprintf("%d | %s", i, time.Since(g)))
	}
}
