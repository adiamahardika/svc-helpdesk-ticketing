package model

import "github.com/mojocn/base64Captcha"

type ConfigJsonBody struct {
	CaptchaId    string                      `json:"captchaId"`
	VerifyValue  string                      `json:"verifyValue"`
	DriverString *base64Captcha.DriverString `json:"driverString"`
}
