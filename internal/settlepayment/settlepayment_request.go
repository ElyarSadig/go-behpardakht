package settlepayment

import (
	"encoding/xml"
	"errors"

	"github.com/elyarsadig/behpardakht/internal/soap"
)

type settlePaymentRequest struct {
	XMLName         xml.Name `xml:"ns1:bpSettleRequest"`
	TerminalId      string   `xml:"terminalId"`
	UserName        string   `xml:"userName"`
	Password        string   `xml:"userPassword"`
	OrderId         string   `xml:"orderId"`
	SaleOrderId     string   `xml:"saleOrderId"`
	SaleReferenceId string   `xml:"saleReferenceId"`
}

func NewSettlePaymentRequest(orderId, saleOrderId, saleReferenceId string) (*settlePaymentRequest, error) {
	if len(orderId) == 0 {
		return nil, errors.New("orderId cannot be empty")
	}
	if len(saleOrderId) == 0 {
		return nil, errors.New("saleOrderId cannot be empty")
	}
	if len(saleReferenceId) == 0 {
		return nil, errors.New("saleReferenceId cannot be empty")
	}
	return &settlePaymentRequest{
		OrderId:         orderId,
		SaleOrderId:     saleOrderId,
		SaleReferenceId: saleReferenceId,
	}, nil
}

func (s *settlePaymentRequest) PrepareSOAPRequest(userId string, password string) ([]byte, error) {
	s.TerminalId = userId
	s.UserName = userId
	s.Password = password
	root := soap.NewRoot()
	root.Body.Request = s
	return root.Marshal()
}
