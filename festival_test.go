package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/constants"
	"github.com/kindlewit/go-festility/models"
	"github.com/kindlewit/go-festility/router"
	"github.com/kindlewit/go-festility/services"
	"github.com/stretchr/testify/assert"
)

var FEST_DATA_CORRECT = models.Fest{
	Id:   "Fest101",
	Name: "Test Fest 101",
	From: 1672294677,
	To:   1672467473,
}

func startupFestival() func() {
	fmt.Println("\n\n\n====\tStartup\t====")
	services.Clear("festival")

	return func() {
		fmt.Println("\n\n\n====\tTeardown\t====")
		services.Clear("festival")
	}
}

func Test_WhenEmptyCreateFest_ShouldRetError(t *testing.T) {
	app := gin.Default()
	router.SetupRouter(app)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/fest", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Received different response status:")
	assert.Equal(t, w.Body.String(), constants.MsgMissingFestParams)
}

func Test_WhenCreateFest_ShouldRetSuccess(t *testing.T) {
	app := gin.Default()
	router.SetupRouter(app)

	w := httptest.NewRecorder()

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(FEST_DATA_CORRECT)
	if err != nil {
		t.FailNow()
	}

	req, _ := http.NewRequest(http.MethodPost, "/fest", &body)
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Received different response status:")
	assert.Contains(t, w.Body.String(), FEST_DATA_CORRECT.Id)
}

func Test_WhenDuplicateCreateFest_ShouldRetError(t *testing.T) {
	app := gin.Default()
	router.SetupRouter(app)

	w := httptest.NewRecorder()

	var body bytes.Buffer
	err := json.NewEncoder(&body).Encode(FEST_DATA_CORRECT)
	if err != nil {
		t.FailNow()
	}

	req, _ := http.NewRequest(http.MethodPost, "/fest", &body)
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code, "Received different response status:")
	assert.Equal(t, w.Body.String(), constants.MsgDuplicateRecord)
}

func Test_WhenGetFest_ShouldRetData(t *testing.T) {
	app := gin.Default()
	router.SetupRouter(app)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/fest/%s", FEST_DATA_CORRECT.Id), nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Received different response status!")
	assert.Contains(t, w.Body.String(), FEST_DATA_CORRECT.Id)
	assert.Contains(t, w.Body.String(), FEST_DATA_CORRECT.Name)
}

func Test_WhenGetInvalidFest_ShouldRetError(t *testing.T) {
	app := gin.Default()
	router.SetupRouter(app)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/fest/Nope", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code, "Received different response status!")
	assert.Equal(t, constants.MsgNoSuchRecord, w.Body.String())
}

func TestMain(m *testing.M) {
	teardown := startupFestival()
	code := m.Run()
	teardown()
	os.Exit(code)
}
