package controllers

import (
	"api/controllers"
	"api/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateOrderInvalidCreds(t *testing.T) {
	order := models.Order{
		Id:       "js",
		Products: []models.Product{},
	}
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/createorder", controllers.MakeOrder)

	marshalled, err := json.Marshal(order)
	if err != nil {
		t.Errorf("marshalling error")
	}
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/createorder", bytes.NewReader(marshalled))
	router.ServeHTTP(rec, req)
	if rec.Code != 400 {
		t.Errorf("error occured")
	}

}
func TestCreateOrderInvalidDataType(t *testing.T) {
	order := "ez"
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/createorder", controllers.MakeOrder)

	marshalled, err := json.Marshal(order)
	if err != nil {
		t.Errorf("marshalling error")
	}
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/createorder", bytes.NewReader(marshalled))
	router.ServeHTTP(rec, req)
	if rec.Code != 400 {
		t.Errorf("error occured")
	}

}

func TestQueryOrderInvalidCreds(t *testing.T) {
	order := models.Order{
		Id:       "js",
		Products: []models.Product{},
	}
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/queryorder", controllers.MakeOrder)

	marshalled, err := json.Marshal(order)
	if err != nil {
		t.Errorf("marshalling error")
	}
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/queryorder", bytes.NewReader(marshalled))
	router.ServeHTTP(rec, req)
	if rec.Code != 400 {
		t.Errorf("error occured")
	}

}
func TestQueryOrderInvalidDataType(t *testing.T) {

	order := "ez"
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/queryorder", controllers.QueryOrder)

	marshalled, err := json.Marshal(order)
	if err != nil {
		t.Errorf("marshalling error")
	}
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/queryorder", bytes.NewReader(marshalled))
	router.ServeHTTP(rec, req)
	if rec.Code != 400 {
		t.Errorf("error occured")
	}

}

func TestDeleteOrderInvalidCreds(t *testing.T) {
	order := models.Order{
		Id:       "js",
		Products: []models.Product{},
	}
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/deleteorder", controllers.DeleteOrder)

	marshalled, err := json.Marshal(order)
	if err != nil {
		t.Errorf("marshalling error")
	}
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/deleteorder", bytes.NewReader(marshalled))
	router.ServeHTTP(rec, req)
	if rec.Code != 400 {
		t.Errorf("error occured")
	}

}
func TestDeleteOrderInvalidDataType(t *testing.T) {
	order := "alkfklaafkl"
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/deleteorder", controllers.DeleteOrder)

	marshalled, err := json.Marshal(order)
	if err != nil {
		t.Errorf("marshalling error")
	}
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodPost, "/deleteorder", bytes.NewReader(marshalled))
	router.ServeHTTP(rec, req)
	if rec.Code != 400 {
		t.Errorf("error occured")
	}

}
