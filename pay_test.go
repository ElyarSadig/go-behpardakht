package behpardakht

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPay(t *testing.T) {
	tests := []struct {
		name             string
		orderId          string
		callBackUrl      string
		userId           string
		amount           uint64
		serverStatusCode int
		serverResponse   string
		expectedRefID    string
		expectedError    bool
	}{
		{
			name:             "Successful Payment",
			orderId:          "123",
			callBackUrl:      "http://example.com/callback",
			userId:           "user123",
			amount:           1000,
			serverStatusCode: 200,
			serverResponse:   `<x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/"><x:Body><bpPayRequestResponse><return>0,ref123</return></bpPayRequestResponse></x:Body></x:Envelope>`,
			expectedRefID:    "ref123",
			expectedError:    false,
		},
		{
			name:             "Invalid Response Format",
			orderId:          "123",
			callBackUrl:      "http://example.com/callback",
			userId:           "user123",
			amount:           1000,
			serverStatusCode: 200,
			serverResponse:   `<x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/"><x:Body><bpPayRequestResponse><return>54,error</return></bpPayRequestResponse></x:Body></x:Envelope>`,
			expectedRefID:    "",
			expectedError:    true,
		},
		{
			name:             "Server Error",
			orderId:          "123",
			callBackUrl:      "http://example.com/callback",
			userId:           "user123",
			amount:           1000,
			serverStatusCode: 200,
			serverResponse:   `Invalid XML Payload`,
			expectedRefID:    "",
			expectedError:    true,
		},
		{
			name:             "Server Error",
			orderId:          "123",
			callBackUrl:      "http://example.com/callback",
			userId:           "user123",
			amount:           1000,
			serverResponse:   `Internal Server Error`,
			serverStatusCode: 500,
			expectedRefID:    "",
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

			CREATE_TRANSACTION_URL = server.URL

			refID, err := b.Pay(tt.orderId, tt.callBackUrl, tt.userId, tt.amount)

			if (err != nil) != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}

			if refID != tt.expectedRefID {
				t.Errorf("Expected refID: %v, got: %v", tt.expectedRefID, refID)
			}
		})
	}
}
