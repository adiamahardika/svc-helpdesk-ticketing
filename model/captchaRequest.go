package model

import "github.com/mojocn/base64Captcha"

type ConfigJsonBody struct {
	VerifyValue  string                      `json:"verifyValue"`
	DriverString *base64Captcha.DriverString `json:"driverString"`
}
