package service

import (
	"svc-myg-ticketing/model"

	"github.com/mojocn/base64Captcha"
)

type CaptchaServiceInterface interface {
	GenerateCaptcha(param model.ConfigJsonBody) (string, string, error)
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
	// body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}

	// if err != nil {
	// 	body = map[string]interface{}{"code": 0, "msg": err.Error()}
	// }

}

// base64Captcha verify http handler
// func captchaVerifyHandle(w http.ResponseWriter, r *http.Request) {

// 	//parse request json body
// 	decoder := json.NewDecoder(r.Body)
// 	var param model.ConfigJsonBody
// 	err := decoder.Decode(&param)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer r.Body.Close()
// 	//verify the captcha
// 	body := map[string]interface{}{"code": 0, "msg": "failed"}
// 	if store.Verify(param.Id, param.VerifyValue, true) {
// 		body = map[string]interface{}{"code": 1, "msg": "ok"}
// 	}

// 	//set json response
// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")

// 	json.NewEncoder(w).Encode(body)
// }
