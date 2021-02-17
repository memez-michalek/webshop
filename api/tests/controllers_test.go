package controllers

import (
	"api/controllers"
	"api/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
)

var apikey = models.APILOGIN{
	Username: "cypis",
	Email:    "cypis@cypis.com",
	Key:      "1ntoZdCfBUb6u1RTUfkXKioqvxc",
}

func TestApiMainPageBadCreds(t *testing.T) {
	data := models.APIUSER{
		Token: faker.Email(),
	}
	marshalled, err := json.Marshal(&data)
	if err != nil {
		t.Errorf(err.Error())
	}
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/main", controllers.MainPage)

	req, err := http.NewRequest(http.MethodPost, "/api/main", bytes.NewReader(marshalled))
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)
	log.Print(rec.Code)
	if rec.Code != 404 {
		t.Errorf("check bad cred handling")
	}

}

func TestAPIManipage(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/main", controllers.MainPage)
	router.POST("/api/login", controllers.LogInToApi)

	marshalled, err := json.Marshal(&apikey)
	if err != nil {
		t.Errorf("internal marshal error")
	}
	req := httptest.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(marshalled))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	data := models.APIUSER{
		Token: w.Body.String(),
	}

	marshalled, err = json.Marshal(&data)
	if err != nil {
		t.Errorf("internal marshal error")
	}

	req = httptest.NewRequest(http.MethodPost, "/api/main", bytes.NewReader(marshalled))

	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("check token verification ")
	}
}

func Test_LoginToApiAgain(t *testing.T) {

	marshalled, err := json.Marshal(apikey)
	if err != nil {
		t.Fatalf("internal json marshalling error")
	}
	log.Print(apikey)
	d := bytes.NewReader(marshalled)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/login", controllers.LogInToApi)

	request, err := http.NewRequest(http.MethodPost, "/api/login", d)
	if err != nil {
		t.Fatalf("failed to send request")
	}
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, request)

	if rec.Code != 400 {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, rec.Code)
	}

}

func TestLoginWrongCreds(t *testing.T) {
	data := models.APILOGIN{
		Username: faker.Username(),
		Email:    faker.Email(),
		Key:      ksuid.New().String(),
	}
	marshalled, err := json.Marshal(&data)
	if err != nil {
		t.Errorf("marshll error")
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/login", controllers.LogInToApi)

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(marshalled))

	router.ServeHTTP(recorder, request)

	if recorder.Code != 400 {
		t.Errorf("wrong status code check your code{}" + fmt.Sprint(recorder.Code))
	}

}
func TestLoginToApiWrongBinding(t *testing.T) {
	data := struct {
		name string
		key  string
	}{
		name: faker.Name(),
		key:  ksuid.New().String(),
	}

	marshalled, err := json.Marshal(&data)
	if err != nil {
		t.Errorf("marshall error")
	}
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/api/login", controllers.LogInToApi)

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, "/api/login", bytes.NewReader(marshalled))
	if err != nil {
		t.Errorf("error when creating request")
	}
	router.ServeHTTP(recorder, request)
	if recorder.Code != 400 {
		t.Errorf("binded to the wrong interface check binding")
	}

}

func TestAPIManipageFakeJWT(t *testing.T) {

	data := models.APIUSER{
		Token: faker.JWT,
	}

	marshalled, err := json.Marshal(&data)
	if err != nil {
		t.Errorf("internal marshal error")
	}
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/api/main", controllers.MainPage)

	req := httptest.NewRequest(http.MethodPost, "/api/main", bytes.NewReader(marshalled))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	log.Print(w.Code)
	if w.Code != 404 {
		t.Errorf("check token verification ")
	}
}

func TestRenewWebtoken(t *testing.T) {
	data := models.APIUSER{
		Token: faker.JWT,
	}
	marshalled, err := json.Marshal(data)
	if err != nil {
		t.Errorf("internal marshall error")
	}

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/api/renew", controllers.RenewApiKey)

	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/api/renew", bytes.NewReader(marshalled))
	if err != nil {
		t.Errorf("cannot create new request")
	}
	router.ServeHTTP(rec, req)

	if rec.Code != 400 {
		t.Errorf("renew function returned wrong http status code")
	}
}
