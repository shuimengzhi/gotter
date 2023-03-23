package test

import (
	"encoding/json"
	"fmt"
	"gotter/request_struct"
	router2 "gotter/router"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewRouter(t *testing.T) {
	router := router2.NewRouter()
	w := httptest.NewRecorder()

	request := request_struct.DemoRequest{Account: "okokjj"}
	infoJson, _ := json.Marshal(request)
	req, _ := http.NewRequest("POST", "/demo", strings.NewReader(string(infoJson)))
	router.ServeHTTP(w, req)
	fmt.Println(fmt.Sprintf("%#v", w.Body.String()))
	//assert.Equal(t, 200, w.Code)
	//assert.Equal(t, "get pong", w.Body.String())
}
