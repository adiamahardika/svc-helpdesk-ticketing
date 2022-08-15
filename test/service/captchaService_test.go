package service_test

import (
	"image/color"
	"svc-myg-ticketing/model"
	"svc-myg-ticketing/service"
	"testing"

	"github.com/mojocn/base64Captcha"
	"github.com/stretchr/testify/require"
)

var captchaService = service.CapthcaService()

func TestGenerateCaptcha(t *testing.T) {

	tests := []struct {
		name          string
		request       *model.ConfigJsonBody
		expectedError error
	}{
		{
			name: "Success Generate Captcha",
			request: &model.ConfigJsonBody{
				CaptchaId:   "",
				VerifyValue: "",
				DriverString: &base64Captcha.DriverString{
					Height:          60,
					Width:           240,
					ShowLineOptions: 0,
					NoiseCount:      40,
					Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
					Length:          6,
					Fonts:           []string{"wqy-microhei.ttc"},
					BgColor: &color.RGBA{
						R: 0,
						G: 0,
						B: 0,
						A: 0,
					},
				},
			},
			expectedError: nil,
		},
	}
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			_, _, error := captchaService.GenerateCaptcha(test.request)
			require.Equal(t, test.expectedError, error)
		})
	}
}

func BenchmarkGenerateCaptcha(b *testing.B) {
	request := model.ConfigJsonBody{
		CaptchaId:   "",
		VerifyValue: "",
		DriverString: &base64Captcha.DriverString{
			Height:          60,
			Width:           240,
			ShowLineOptions: 0,
			NoiseCount:      40,
			Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
			Length:          6,
			Fonts:           []string{"wqy-microhei.ttc"},
			BgColor: &color.RGBA{

				R: 0,
				G: 0,
				B: 0,
				A: 0,
			},
		},
	}
	for index := 0; index < b.N; index++ {
		captchaService.GenerateCaptcha(&request)
	}
}