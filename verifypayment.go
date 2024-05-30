package behpardakht

import (
	"net/http"

	"github.com/elyarsadig/behpardakht/internal/verifypayment"
)

func (b *behPardakht) VerifyPayment(orderID, saleOrderID, saleReferenceID string) error {
	verifyPaymentRequest, err := verifypayment.NewVerifyPaymentRequest(orderID, saleOrderID, saleReferenceID)
	if err != nil {
		return err
	}
	response, err := b.sendRequest(http.MethodPost, VERIFY_TRANSACTION_URL, verifyPaymentRequest)
	if err != nil {
		return err
	}
	err = verifypayment.ProcessVerifyPaymentResponse(response)
	if err != nil {
		return err
	}
	return nil
}
