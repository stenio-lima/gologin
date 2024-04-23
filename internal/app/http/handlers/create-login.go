package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v3"
)

type LoginBody struct {
	Login string `json:"login"`
	Senha string `json:"password"`
}

func CreateLoginHandler(c fiber.Ctx) error {
	loginBody := new(LoginBody)

	if err := parseValidate(c.Body(), loginBody); err != nil {
		return err
	}

	//Temporary validade fake infos
	if loginBody.Login != "teste@test.com" || loginBody.Senha != "teste" {
		return errors.New("Wrong login or senha")
	}
	return nil
}

func parseValidate(body []byte, loginBody interface{}) error {
	if len(body) == 0 {
		return errors.New("Body is empty")
	}

	if err := json.Unmarshal(body, &loginBody); err != nil {
		return err
	}

	if loginBody == nil {
		return errors.New("LoginBody is nil")
	}

	return nil
}
