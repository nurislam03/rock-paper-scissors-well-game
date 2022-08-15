package api

import (
	"net/http"
	"rpsw/game"
	"rpsw/utils"
	"strings"
)

type playRockPaperScissorsWellReqBody struct {
	Move string `json:"move"`
}

type playRockPaperScissorsWellResponse struct {
	UserMove     string `json:"user_move"`
	ComputerMove string `json:"computer_move"`
	UserResult   string `json:"user_result"`
}

// playRPSW accepts user's choice/move as a request body and response back the result state of that game for that user
func (api *API) playRPSW(w http.ResponseWriter, r *http.Request) {
	var usersChoice playRockPaperScissorsWellReqBody

	// parse user's choice from request body
	err := utils.ParseBody(r, &usersChoice)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "error in parsing request body")
		return
	}

	// considering case-insensitive
	userInput := strings.ToLower(usersChoice.Move)

	// validate user's choice
	isValidChoice := validateGameCharacter(userInput)
	if !isValidChoice {
		utils.HandleError(w, http.StatusUnprocessableEntity, "invalid move")
		return
	}

	// construct players for the game
	humanPlayer := game.NewHuman(userInput)
	botPlayer := game.NewBot(game.NewSimpleRandGenerator())

	// calculate result state of the match between two players
	gameResult := game.NewGame(humanPlayer, botPlayer).Play()

	response := &utils.Response{
		Code: http.StatusOK,
		Data: playRockPaperScissorsWellResponse{
			UserMove:     gameResult.Choice1.Name(),
			ComputerMove: gameResult.Choice2.Name(),
			UserResult:   gameResult.Result,
		},
	}

	response.ServeJSON(w)
}

// validateGameCharacter validates weather a specific character selection is valid or not
func validateGameCharacter(usersChoice string) bool {
	choices := map[string]bool{"rock": true, "paper": true, "scissors": true, "well": true} // valid characters

	uChoice := strings.ToLower(usersChoice)
	return choices[uChoice]
}
