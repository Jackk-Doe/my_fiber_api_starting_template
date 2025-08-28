package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jackk-Doe/my_fiber_api_starting_template/api"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestAPIHealthzEndpoint(t *testing.T) {
	app := api.SetUpAPI()
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestAPIRootEndpoint(t *testing.T) {
	app := api.SetUpAPI()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestAPITestSampleEndpoint(t *testing.T) {
	app := api.SetUpAPI()
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
}

func TestAPIUnknownPathEndpoint(t *testing.T) {
	app := api.SetUpAPI()
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	resp, _ := app.Test(req)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}
