package api

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"rpsw/model"
	"rpsw/utils"
	"strings"
	"time"
)

// playRPSW accepts a user request with a game move and response back the result of the game
func (api *API) playRPSW(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered into playRPSW function")

	var usersChoice model.PlayRockPaperScissorsWellReqBody

	// parse user's choice from request body
	err := utils.ParseBody(r, &usersChoice)
	if err != nil {
		log.Println("error in parsing request body")
	}

	// validate user's choice
	isValidChoice := validateGameChoice(usersChoice.Move)
	if !isValidChoice {
		log.Println("isValidChoice: ", isValidChoice)
		json.NewEncoder(w).Encode("invalid move")
		return
	}

	// computer's choice for the game
	computersChoice := computersMove()

	// todo: result

	//w.Write([]byte("Getting POST index route!"))
}

func computersMove() string {
	// game choice
	choices := []string{"rock", "paper", "scissors", "well"}

	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(4)

	return choices[randNumber]
}

// validateGameChoice validates weather a specific move is valid or not
func validateGameChoice(usersChoice string) bool {
	// game choice
	choices := map[string]bool{"rock": true, "paper": true, "scissors": true, "well": true}
	uChoice := strings.ToLower(usersChoice)

	return choices[uChoice]
}


