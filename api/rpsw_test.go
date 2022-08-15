package api

import (
	"bytes"
	"encoding/json"
	"github.com/kinbiko/jsonassert"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testTemplateForTestPlayRPSW(t *testing.T, api *API, testName string, buf *bytes.Buffer, code int, result string) {
	t.Run(testName, func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/rpsw/play", buf)
		res := httptest.NewRecorder()
		api.router.ServeHTTP(res, req)
		assert.Equal(t, code, res.Code)
		jsonassert.New(t).Assertf(res.Body.String(), result)
	})
}

func TestPlayRPSW(t *testing.T) {
	api := NewAPI()
	testCases := []struct {
		name    string
		reqBody playRockPaperScissorsWellReqBody
		code    int
		resp    string
	}{
		{
			name:    "successfully play considering no errors occurs",
			reqBody: playRockPaperScissorsWellReqBody{"rock"},
			code:    http.StatusOK,
			resp:    `{"data": {"user_move": "rock","computer_move": "<<PRESENCE>>","user_result":"<<PRESENCE>>"}}`,
		},
		{
			name:    "invalid user move",
			reqBody: playRockPaperScissorsWellReqBody{"rocks"},
			code:    http.StatusUnprocessableEntity,
			resp:    `{"error_message": "invalid move"}`,
		},
	}

	for _, tc := range testCases {
		var buf bytes.Buffer
		_ = json.NewEncoder(&buf).Encode(tc.reqBody)
		testTemplateForTestPlayRPSW(t, api, tc.name, &buf, tc.code, tc.resp)
	}
}

// testTemplate for TestValidateGameCharacter
func testTemplate(t *testing.T, testName, input string, result bool) {
	t.Run(testName, func(t *testing.T) {
		assert.Equal(t, result, validateGameCharacter(input))
	})
}

func TestValidateGameCharacter(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		result bool
	}{
		{"for input 'rock' it should return true", "rock", true},
		{"for input 'Rock' it should return true", "Rock", true},
		{"for input 'scissors' it should return true", "scissors", true},
		{"for input 'paper' it should return true", "paper", true},
		{"for input 'well' it should return true", "well", true},
		{"for input 'welcome' it should return false", "welcome", false},
		{"for input 'PaPeR' it should return true", "PaPeR", true},
		{"for input 'papers' it should return false", "papers", false},
		{"for input as 'empty string' it should return false", " ", false},
		{"for input '12345' it should return false", "12345", false},
	}

	for _, tc := range testCases {
		testTemplate(t, tc.name, tc.input, tc.result)
	}
}
