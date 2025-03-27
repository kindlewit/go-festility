package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/router"
	"github.com/stretchr/testify/assert"
)

var app *gin.Engine

func init() {
	app = gin.Default()
	router.SetupRouter(app)
}

func Test_WhenCreateCinema_ShouldRetValidRes(t *testing.T) {
	validData := map[string]interface{}{
		"name":             "SPI Sathyam Cinemas",
		"city":             "Chennai",
		"google_plus_code": "3745+45 Chennai, Tamil Nadu, India",
	}
	body, _ := json.Marshal(validData)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/cinema", bytes.NewReader(body))
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code, "Received invalid response status")
}
