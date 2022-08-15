package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testTemplate(t *testing.T, moc *MockRandGenerator, testName string, rand int, humanChoice, botChoice, result string) {
	t.Run(testName, func(t *testing.T) {
		moc.On("Rand", 4).Return(rand).Once()
		bp := NewBot(moc)
		hp := NewHuman(humanChoice)
		res := NewGame(hp, bp).Play()

		assert.Equal(t, humanChoice, res.Choice1.Name())
		assert.Equal(t, botChoice, res.Choice2.Name())
		assert.Equal(t, result, res.Result)
		moc.AssertExpectations(t)
	})
}

func TestGamePlay(t *testing.T) {

	mockRandGen := &MockRandGenerator{}
	testCases := []struct {
		name                           string
		rand                           int
		humanChoice, botChoice, result string
	}{
		{"Rock Vs Paper - rock loses", 1, "rock", "paper", "lose"},
		{"Rock Vs Scissors - rock wins", 2, "rock", "scissors", "win"},
	}

	for _, tc := range testCases {
		testTemplate(t, mockRandGen, tc.name, tc.rand, tc.humanChoice, tc.botChoice, tc.result)
	}
}

//func TestGamePlay_RockVsScissors(t *testing.T) {
//	mockRandGen := MockRandGen{2}
//	bp := NewBot(mockRandGen)
//	hp := NewHuman("rock")
//	res := NewGame(hp, bp).Play()
//
//	assert.Equal(t, "rock", res.Choice1.Name())
//	assert.Equal(t, "scissors", res.Choice2.Name())
//	assert.Equal(t, "win", res.Result)
//}
//
//func TestGamePlay_RockVsWell(t *testing.T) {
//	mockRandGen := MockRandGen{3}
//	bp := NewBot(mockRandGen)
//	hp := NewHuman("rock")
//	res := NewGame(hp, bp).Play()
//
//	assert.Equal(t, "rock", res.Choice1.Name())
//	assert.Equal(t, "well", res.Choice2.Name())
//	assert.Equal(t, "lose", res.Result)
//}

//
//func TestGamePlay(t *testing.T) {
//	testdata := []struct {
//		randomNumber int
//		bot          string
//		human        string
//		resp         string
//	}{
//		{
//			des:          "successfully get tasks with operator",
//			randomNumber: 1,
//			bot:          "1",
//			code:         http.StatusOK,
//			resp:         `{"data":[{"id":"000000000000000000000000","task_id":"1","name":"Assessment","start_date":"<<PRESENCE>>","category":"Test","status":"processing","customer_id":"1","created_at":"<<PRESENCE>>","updated_at":"<<PRESENCE>>"},{"id":"000000000000000000000000","task_id":"2","name":"Dentist Appointment","start_date":"<<PRESENCE>>","category":"Health","status":"processing","customer_id":"2","created_at":"<<PRESENCE>>","updated_at":"<<PRESENCE>>"}],"meta":{"current_page":1,"per_page":2,"total_page":0,"total":0}}`,
//		},
//	}
//
//	for _, td := range testdata {
//		t.Run(td.des, func(t *testing.T) {
//			req.Header.Add("Role", td.role)
//			req.Header.Add("UserID", td.userID)
//			res := httptest.NewRecorder()
//			api.router.ServeHTTP(res, req)
//			assert.Equal(t, td.code, res.Code)
//			jsonassert.New(t).Assertf(res.Body.String(), td.resp)
//		})
//	}
//}
