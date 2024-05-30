package verifypayment

import (
	"encoding/xml"
	"errors"

	"github.com/elyarsadig/behpardakht/internal/soap"
)

type verifyPaymentRequest struct {
	XMLName         xml.Name `xml:"ns1:bpVerifyRequest"`
	TerminalID      string   `xml:"terminalId"`
	UserName        string   `xml:"userName"`
	Password        string   `xml:"userPassword"`
	OrderID         string   `xml:"orderId"`
	SaleOrderID     string   `xml:"saleOrderId"`
	SaleReferenceID string   `xml:"saleReferenceId"`
}

func NewVerifyPaymentRequest(orderID, saleOrderID, saleReferenceID string) (*verifyPaymentRequest, error) {
	if len(orderID) == 0 {
		return nil, errors.New("orderID cannot be empty")
	}
	if len(saleOrderID) == 0 {
		return nil, errors.New("saleOrderID cannot be empty")
	}
	if len(saleReferenceID) == 0 {
		return nil, errors.New("saleReferenceID cannot be empty")
	}
	return &verifyPaymentRequest{
		OrderID:         orderID,
		SaleOrderID:     saleOrderID,
		SaleReferenceID: saleReferenceID,
	}, nil
}

func (v *verifyPaymentRequest) PrepareSOAPRequest(userId string, password string) ([]byte, error) {
	v.TerminalID = userId
	v.UserName = userId
	v.Password = password
	root := soap.NewRoot()
	root.Body.Request = v
	return root.Marshal()
}
