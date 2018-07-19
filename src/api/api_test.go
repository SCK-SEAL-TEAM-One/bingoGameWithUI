package api_test

import (
	"testing"
	"bytes"
	"net/http/httptest"
	"encoding/json"
	."api"
)

func Test_StartGameHandler(t *testing.T){
	request := StartGameRequest{PlayerOne:"A",PlayerTwo:"B",}
	requestByte, _ := json.Marshal(request)

	req := httptest.NewRequest("POST", "/bingo/start", bytes.NewBuffer(requestByte))
	w := httptest.NewRecorder()
	api := Api{}
	api.StartGameHandler(w, req)
	expectedRespondstatus := 200
	response := w.Result()
	if expectedRespondstatus != response.StatusCode{
		t.Errorf("expect %d but got %d",expectedRespondstatus, response.StatusCode)
	}
}