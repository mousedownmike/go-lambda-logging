package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main(){
	metrics := make(map[string]interface{}, 3)
	hasDiff, additions, subtractions := randomDiff()
	metrics["hasDiff"] = hasDiff
	if hasDiff {
		metrics["additions"] = additions
		metrics["subtractions"] = subtractions
	}
	metricsJson, _ := json.Marshal(metrics)
	log.Print(string(metricsJson))
}

func randomDiff() (bool, int, int) {
	var diff bool
	var adds, subs int
	if r.Intn(2) == 0 {
		diff = false
	} else {
		diff = true
	}
	if diff {
		return diff, r.Intn(100), r.Intn(50)
	} else {
		return diff, adds, subs
	}
}