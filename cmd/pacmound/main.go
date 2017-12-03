package main

import (
	"log"
	"net/http"

	"github.com/crhntr/pacmound"
	"github.com/crhntr/pacmound/agents"
)

func getGopher() pacmound.Agent {
	return &agents.Naive{}
}

func getPython() pacmound.Agent {
	return &agents.Ghost{}
}

func main() {
	mux := pacmound.NewGameMux(getGopher, getPython)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
