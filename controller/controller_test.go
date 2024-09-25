package controller_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/herux/indegooweather/config"
	"github.com/herux/indegooweather/controller"
	"github.com/herux/indegooweather/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) FetchAndStoreIndegoData() error {
	args := m.Called()
	return args.Error(0)
}

func TestHandleFetchIndego_Success(t *testing.T) {
	app := fiber.New()

	_ = config.Load("../config.yaml")
	db.Init(true)

	mockService := new(MockService)
	mockService.On("FetchAndStoreIndegoData").Return(nil)

	app.Post("/api/v1/indego-data-fetch-and-store-it-db", func(c *fiber.Ctx) error {
		return controller.HandleFetchIndego(c)
	})

	req := httptest.NewRequest("POST", "/api/v1/indego-data-fetch-and-store-it-db", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	assert.Contains(t, buf.String(), "Data fetched and stored succesfully")
}

func TestHandleStationSnapshot_Success(t *testing.T) {
	app := fiber.New()

	_ = config.Load("../config.yaml")
	db.Init(true)

	app.Get("/api/v1/stations", func(c *fiber.Ctx) error {
		return controller.HandleStationSnapshot(c)
	})

	req := httptest.NewRequest("GET", "/api/v1/stations?at=2024-09-01T10:00:00Z", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	assert.Contains(t, buf.String(), "stations")
	assert.Contains(t, buf.String(), "weather")
}

func TestHandleKiosk_Success(t *testing.T) {
	app := fiber.New()

	_ = config.Load("../config.yaml")
	db.Init(true)

	app.Get("/api/v1/stations/:kioskId", func(c *fiber.Ctx) error {
		return controller.HandleKiosk(c)
	})

	req := httptest.NewRequest("GET", "/api/v1/stations/123?at=2019-09-01T10:00:00Z", nil)
	req.Header.Set("Authorization", "Bearer valid-token")
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	assert.Contains(t, buf.String(), "station")
	assert.Contains(t, buf.String(), "weather")
}
