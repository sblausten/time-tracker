package controllers


import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveSession(t *testing.T) {
	session := Session{UserId: "1234", SessionName: "testSession", Start: "2021-07-12T22:00:03.394Z", End: "2021-07-12T22:01:03.394Z"}
	requestByte, err := json.Marshal(session)
	requestReader := bytes.NewReader(requestByte)

	req, err := http.NewRequest("POST", "", requestReader)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(SaveSession)


	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}
