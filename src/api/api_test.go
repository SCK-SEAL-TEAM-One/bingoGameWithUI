package api_test

import (
	. "api"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_PlayHandle_Should_Be_Json_number_9_winner_Empty(t *testing.T) {
	expected := Play{
		number: 9,
		winner: "",
	}

	rwq := httptest.NewRequest("GET", "localhost:3000/bigo/play")
	w := httptest.NewRecorder()

	Play(w, rwq)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	var actaul Play
	json.Unmarshal(body, actaul)

	if expected != actaul {
		t.Errorf("expected %v but got %v", expected, actaul)
	}
}
