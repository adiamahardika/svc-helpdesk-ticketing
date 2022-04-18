package service

import (
	"fmt"
	"svc-myg-ticketing/model"

	"github.com/mojocn/base64Captcha"
)

type CaptchaServiceInterface interface {
	GenerateCaptcha(param model.ConfigJsonBody) (string, string, error)
	CaptchaVerify(request model.ConfigJsonBody) (bool, error)
}

type captchaService struct{}

func CapthcaService() *captchaService {
	return &captchaService{}
}

var store = base64Captcha.DefaultMemStore

func (captchaService *captchaService) GenerateCaptcha(request model.ConfigJsonBody) (string, string, error) {

	var driver base64Captcha.Driver

	driver = request.DriverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, store)

	id, b64s, error := c.Generate()

	return id, b64s, error
}

func (captchaService *captchaService) CaptchaVerify(request model.ConfigJsonBody) (bool, error) {

	is_valid := false
	error := fmt.Errorf("Captcha Not Match")
	if store.Verify(request.CaptchaId, request.VerifyValue, true) && request.CaptchaId != "" && request.VerifyValue != "" {
		is_valid = true
		error = nil
	}

	return is_valid, error
}
