package controllers


import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/sblausten/time-tracker/mocks"
	"github.com/sblausten/time-tracker/models"
	"github.com/sblausten/time-tracker/server"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSessionDao := mocks.NewMockSessionDaoInterface(ctrl)

	userId := "1234"
	url := "/v1/users/" + userId + "/session"
	input := models.Session{SessionName: "test", Start: "2021-07-12T22:00:03.394Z", End: "2021-07-12T22:01:03.394Z", Duration: 1000}
	expected := models.Session{UserId: userId, SessionName: "test", Start: "2021-07-12T22:00:03.394Z", End: "2021-07-12T22:01:03.394Z", Duration: 1000}
	requestByte, err := json.Marshal(input)
	requestReader := bytes.NewReader(requestByte)

	r := server.NewRouter(mockSessionDao)
	mockServer := httptest.NewServer(r)
	resp, err := http.Post(mockServer.URL + "/hello", "application/json", requestReader)

	// Handle any unexpected error
	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// In the next few lines, the response body is read, and converted to a string
	defer resp.Body.Close()
	// read the body into a bunch of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// convert the bytes to a string
	respString := string(b)

	// We want our response to match the one defined in our handler.
	// If it does happen to be "Hello world!", then it confirms, that the
	// route is correct
	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

	//req, err := http.NewRequest("POST", url, requestReader)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//recorder := httptest.NewRecorder()
	//hf := http.HandlerFunc(SaveSession(mockSessionDao))
	//hf.ServeHTTP(recorder, req)

	mockSessionDao.EXPECT().InsertSession(expected)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}
