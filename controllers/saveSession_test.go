package controllers


import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/sblausten/time-tracker/mocks"
	"github.com/sblausten/time-tracker/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveSession(t *testing.T) {
	//ctrl := gomock.NewController(t)
	//defer ctrl.Finish()
	//
	//mockSessionDao := mocks.NewMockSessionDaoInterface(ctrl)
	//
	//userId := "1234"
	//url := "/v1/users/" + userId + "/session"
	//input := models.Session{SessionName: "test", Start: "2021-07-12T22:00:03.394Z", End: "2021-07-12T22:01:03.394Z", Duration: 1000}
	//expected := models.Session{UserId: userId, SessionName: "test", Start: "2021-07-12T22:00:03.394Z", End: "2021-07-12T22:01:03.394Z", Duration: 1000}
	//requestByte, err := json.Marshal(input)
	//requestReader := bytes.NewReader(requestByte)
	//
	//r := server.NewRouter(mockSessionDao)
	//httptest.NewServer(r)
	//resp, err := http.Post(url, "application/json", requestReader)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//mockSessionDao.EXPECT().InsertSession(expected)
	//
	//if status := resp.StatusCode; status != http.StatusCreated {
	//	t.Errorf("Handler returned wrong status code: got %v, wanted %v",
	//		status, http.StatusCreated)
	//}

	//req, err := http.NewRequest("POST", url, requestReader)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//recorder := httptest.NewRecorder()
	//hf := http.HandlerFunc(SaveSession(mockSessionDao))
	//hf.ServeHTTP(recorder, req)


}
