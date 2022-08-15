package service_test

import (
	"errors"
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
			name: "Success",
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
		{
			name: "Error Empty Request 1",
			request: &model.ConfigJsonBody{
				CaptchaId:    "",
				VerifyValue:  "",
				DriverString: &base64Captcha.DriverString{},
			},
			expectedError: errors.New("text must not be empty, there is nothing to draw"),
		},
		{
			name: "Error Empty Request 2",
			request: &model.ConfigJsonBody{
				CaptchaId:   "",
				VerifyValue: "",
				DriverString: &base64Captcha.DriverString{
					Height:          0,
					Width:           0,
					ShowLineOptions: 0,
					NoiseCount:      0,
					Source:          "",
					Length:          0,
					Fonts:           []string{"wqy-microhei.ttc"},
					BgColor: &color.RGBA{
						R: 0,
						G: 0,
						B: 0,
						A: 0,
					},
				},
			},
			expectedError: errors.New("text must not be empty, there is nothing to draw"),
		},
	}
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			_, _, error := captchaService.GenerateCaptcha(test.request)
			require.Equal(t, test.expectedError, error)
		})
	}
}

func TestCaptchaVerify(t *testing.T) {

	tests := []struct {
		name          string
		request       *model.ConfigJsonBody
		expectedError error
	}{
		{
			name: "Error With Filled Req",
			request: &model.ConfigJsonBody{
				CaptchaId:   "2RHFkONE8GODe913MAS9",
				VerifyValue: "oa2tbe",
			},
			expectedError: errors.New("Captcha Not Match"),
		},
		{
			name: "Error With Empty Req 1",
			request: &model.ConfigJsonBody{
				CaptchaId:   "",
				VerifyValue: "",
			},
			expectedError: errors.New("Captcha Not Match"),
		},
		{
			name:          "Error With Empty Req 2",
			request:       &model.ConfigJsonBody{},
			expectedError: errors.New("Captcha Not Match"),
		},
	}
	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			_, error := captchaService.CaptchaVerify(test.request)
			require.Equal(t, test.expectedError, error)
		})
	}
}

func BenchmarkGenerateCaptcha(b *testing.B) {
	benchmarks := []struct {
		name    string
		request model.ConfigJsonBody
	}{{
		name: "Generate Captcha",
		request: model.ConfigJsonBody{
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
	}}

	for _, benchmark := range benchmarks {
		for index := 0; index < b.N; index++ {
			b.Run(benchmark.name, func(b *testing.B) {
				captchaService.GenerateCaptcha(&benchmark.request)
			})
		}
	}
}

func BenchmarkCaptchaVerify(b *testing.B) {
	benchmarks := []struct {
		name    string
		request model.ConfigJsonBody
	}{{
		name: "Captcha Verify",
		request: model.ConfigJsonBody{
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
	}}

	for _, benchmark := range benchmarks {
		for index := 0; index < b.N; index++ {
			b.Run(benchmark.name, func(b *testing.B) {
				captchaService.CaptchaVerify(&benchmark.request)
			})
		}
	}
}
