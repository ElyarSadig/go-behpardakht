package behpardakht

import (
	"fmt"
	"net/http"

	"github.com/elyarsadig/behpardakht/internal/payment"
)

func (b *behPardakht) Pay(orderID, userID, callBackUrl string, amount uint64) (string, error) {
	paymentRequest, err := payment.NewPaymentRequest(orderID, userID, callBackUrl, amount)
	if err != nil {
		return "", err
	}
	response, err := b.sendRequest(http.MethodPost, CREATE_TRANSACTION_URL, paymentRequest)
	if err != nil {
		return "", err
	}
	refID, err := payment.ProcessPaymentResponse(response)
	if err != nil {
		return "", fmt.Errorf("failed to process payment response: %v", err)
	}
	return refID, nil
}
