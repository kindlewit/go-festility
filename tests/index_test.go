package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/router"
	"github.com/stretchr/testify/assert"
)

// func Test_WhenDbDisconnected_ShouldRetSuccessFalse(t *testing.T) {
// 	app := gin.Default()
// 	router.SetupRouter(app)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/", nil)
// 	app.ServeHTTP(w, req)

//		assert.Equal(t, http.StatusOK, w.Code, "Received invalid response status")
//		assert.Contains(t, w.Body.String(), "timestamp", "Invalid response structure: 'timestamp' missing")
//		assert.Contains(t, w.Body.String(), "success", "Invalid response structure: 'success' missing")
//		assert.Contains(t, w.Body.String(), "success\":false", "Invalid response structure: 'success' is truthy")
//		assert.Contains(t, w.Body.String(), "message", "Invalid response structure: 'message' missing")
//	}
func init() {
	app = gin.Default()
	router.SetupRouter(app)
}

func Test_WhenDbConnected_ShouldRetSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	app.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Received invalid response status")
	assert.Contains(t, w.Body.String(), "timestamp", "Invalid response structure: 'timestamp' missing")
	assert.Contains(t, w.Body.String(), "success", "Invalid response structure: 'success' missing")
	assert.Contains(t, w.Body.String(), "success\":true", "Invalid response structure: 'success' is falsy")
	assert.Contains(t, w.Body.String(), "message", "Invalid response structure: 'message' missing")
}
