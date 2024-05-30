package behpardakht

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVerifyPayment(t *testing.T) {
	tests := []struct {
		name             string
		serverStatusCode int
		serverResponse   string
		expectedError    bool
	}{
		{
			name:             "Valid Response",
			serverStatusCode: http.StatusOK,
			serverResponse:   `<Envelope><Body><bpVerifyRequestResponse><return>0</return></bpVerifyRequestResponse></Body></Envelope>`,
			expectedError:    false,
		},
		{
			name:             "Invalid XML Response",
			serverStatusCode: http.StatusOK,
			serverResponse:   `<Envelope><Body><bpVerifyRequestResponse><return>invalid</return></bpVerifyRequestResponse></Body></Envelope>`,
			expectedError:    true,
		},
		{
			name:             "Non-zero Response Code",
			serverStatusCode: http.StatusOK,
			serverResponse:   `<Envelope><Body><bpVerifyRequestResponse><return>123</return></bpVerifyRequestResponse></Body></Envelope>`,
			expectedError:    true,
		},
		{
			name:             "HTTP Error",
			serverStatusCode: http.StatusInternalServerError,
			serverResponse:   ``,
			expectedError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &behPardakht{
				client:   &http.Client{},
				username: "testUser",
				password: "testPassword",
			}

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.serverStatusCode)
				w.Write([]byte(tt.serverResponse))
			}))
			defer server.Close()

			VERIFY_TRANSACTION_URL = server.URL

			err := b.VerifyPayment("dummydata", "dummydata", "dummydata")

			if (err != nil) != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
		})
	}
}
