package handlers

import (
	"bytes"
	"github.com/stenio-lima/gologin/internal/app/helpers"
	"github.com/stenio-lima/gologin/internal/app/http/router"
	"net/http/httptest"
	"testing"
	"time"

	utils "github.com/gofiber/utils"
)

func Test_SuccessLogin(t *testing.T) {
	t.Parallel()

	app := router.InitRoutes()

	body := `{"login":"teste@test.com","password":"teste"}`

	resp, err := app.Test(httptest.NewRequest(helpers.MethodPost, "/api/v1/login", bytes.NewReader([]byte(body))), time.Second)

	utils.AssertEqual(t, nil, err, "No errors should be returned")
	utils.AssertEqual(t, 200, resp.StatusCode, "Nothing shoud be returned")
}

func Test_EmptyBodyLogin(t *testing.T) {
	t.Parallel()

	app := router.InitRoutes()

	body := `{"login":"","password":""}`

	resp, err := app.Test(httptest.NewRequest(helpers.MethodPost, "/api/v1/login", bytes.NewReader([]byte(body))), time.Second)

	utils.AssertEqual(t, nil, err, "No errors should be returned")
	utils.AssertEqual(t, 200, resp.StatusCode, "Nothing shoud be returned")
}
