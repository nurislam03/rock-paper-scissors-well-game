package api

import (
	"encoding/json"
	"log"
	"net/http"
	"rpsw/utils"
	"strings"
)

type playRockPaperScissorsWellReqBody struct {
	Move string `json:"move"`
}

// playRPSW accepts a user request with a game move and response back the result of the game
func (api *API) playRPSW(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered into playRPSW function")

	var usersChoice playRockPaperScissorsWellReqBody

	// parse user's choice from request body
	err := utils.ParseBody(r, &usersChoice)
	if err != nil {
		log.Println("error in parsing request body")
	}

	// validate user's choice
	isValidChoice := validateGameCharacter(usersChoice.Move)
	if !isValidChoice {
		log.Println("isValidChoice: ", isValidChoice)
		json.NewEncoder(w).Encode("invalid move")
		return
	}

	//w.Write([]byte("Getting POST index route!"))
}

// validateGameCharacter validates weather a specific character selection is valid or not
func validateGameCharacter(usersChoice string) bool {
	// game choice
	choices := map[string]bool{"rock": true, "paper": true, "scissors": true, "well": true}
	uChoice := strings.ToLower(usersChoice)
	return choices[uChoice]
}
