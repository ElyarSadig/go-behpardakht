package verifypayment

import (
	"testing"
)

func TestProcessVerifyPaymentResponse(t *testing.T) {
	tests := []struct {
		name          string
		xmlResponse   string
		expectedError bool
	}{
		{
			name:          "Valid Response",
			xmlResponse:   `<Envelope><Body><bpVerifyRequestResponse><return>0</return></bpVerifyRequestResponse></Body></Envelope>`,
			expectedError: false,
		},
		{
			name:          "Invalid Code",
			xmlResponse:   `<Envelope><Body><bpVerifyRequestResponse><return>abc</return></bpVerifyRequestResponse></Body></Envelope>`,
			expectedError: true,
		},
		{
			name:          "Non-zero Code",
			xmlResponse:   `<Envelope><Body><bpVerifyRequestResponse><return>123</return></bpVerifyRequestResponse></Body></Envelope>`,
			expectedError: true,
		},
		{
			name: "Invalid Response",
			xmlResponse: `Invalid XML Response`,
			expectedError: true,
		},
		{
			name: "Empty Response",
			xmlResponse: ``,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ProcessVerifyPaymentResponse([]byte(tt.xmlResponse))
			if (err != nil) != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}
