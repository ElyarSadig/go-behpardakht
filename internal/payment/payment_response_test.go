package payment

import (
	"strings"
	"testing"
)

func TestProcessPaymentResponse(t *testing.T) {
	tests := []struct {
		name          string
		response      []byte
		expectedRefID string
		expectedError string
	}{
		{
			name: "Successful processing",
			response: []byte(`
				<Envelope>
					<Body>
						<bpPayRequestResponse>
							<return>0,123456789</return>
						</bpPayRequestResponse>
					</Body>
				</Envelope>
			`),
			expectedRefID: "123456789",
			expectedError: "",
		},
		{
			name:          "XML unmarshal error",
			response:      []byte(`<InvalidXML`),
			expectedRefID: "",
			expectedError: "error in unmarshal response",
		},
		{
			name: "Parse and set response fields error",
			response: []byte(`
				<Envelope>
					<Body>
						<bpPayRequestResponse>
							<return>invalid_code</return>
						</bpPayRequestResponse>
					</Body>
				</Envelope>
			`),
			expectedRefID: "",
			expectedError: "error in parseAndSetResponseFields",
		},
		{
			name: "Bank error message",
			response: []byte(`
				<Envelope>
					<Body>
						<bpPayRequestResponse>
							<return>12,123456789</return>
						</bpPayRequestResponse>
					</Body>
				</Envelope>
			`),
			expectedRefID: "",
			expectedError: "موجودی کافی نیست",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			refID, err := ProcessPaymentResponse(tt.response)
			if err != nil {
				if tt.expectedError == "" || !errorContains(err, tt.expectedError) {
					t.Errorf("expected error: %v, got: %v", tt.expectedError, err)
				}
			} else {
				if tt.expectedError != "" {
					t.Errorf("expected error: %v, got none", tt.expectedError)
				}
			}
			if refID != tt.expectedRefID {
				t.Errorf("expected refID: %v, got: %v", tt.expectedRefID, refID)
			}
		})
	}
}

func errorContains(out error, want string) bool {
	return strings.Contains(out.Error(), want)
}
