package behpardakht

import (
	"net/http"

	"github.com/elyarsadig/behpardakht/internal/settlepayment"
)

func (b *behPardakht) SettlePayment(orderId, saleOrderId, saleReferenceId string) error {
	settlePayment, err := settlepayment.NewSettlePaymentRequest(orderId, saleOrderId, saleReferenceId)
	if err != nil {
		return err
	}
	response, err := b.sendRequest(http.MethodPost, SETTLE_TRANSACTION_URL, settlePayment)
	if err != nil {
		return err
	}
	err = settlepayment.ProcessSettlePaymentResponse(response)
	if err != nil {
		return err
	}
	return nil
}
