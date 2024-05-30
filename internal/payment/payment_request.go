package payment

import (
	"encoding/xml"
	"errors"
	"time"

	"github.com/elyarsadig/behpardakht/internal/soap"
)

type paymentRequest struct {
	XMLName      xml.Name `xml:"ns1:bpPayRequest"`
	TerminalId   string   `xml:"terminalId"`
	UserName     string   `xml:"userName"`
	UserPassword string   `xml:"userPassword"`
	OrderId      string   `xml:"orderId"`
	Amount       uint64   `xml:"amount"`
	LocalDate    string   `xml:"localDate"`
	LocalTime    string   `xml:"localTime"`
	CallBackUrl  string   `xml:"callBackUrl"`
	PayerId      string   `xml:"payerId"`
}

func NewPaymentRequest(orderID, payerID, callBackUrl string, amount uint64) (*paymentRequest, error) {
	if len(orderID) == 0 {
		return nil, errors.New("orderId cannot be empty")
	}
	if len(callBackUrl) == 0 {
		return nil, errors.New("callBackUrl cannot be empty")
	}
	if len(payerID) == 0 {
		return nil, errors.New("payerId cannot be empty")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	return &paymentRequest{
		OrderId:     orderID,
		Amount:      amount,
		LocalDate:   time.Now().Format("20060402"),
		LocalTime:   time.Now().Format("150405"),
		CallBackUrl: callBackUrl,
		PayerId:     payerID,
	}, nil
}

func (pr *paymentRequest) PrepareSOAPRequest(userId string, password string) ([]byte, error) {
	pr.TerminalId = userId
	pr.UserName = userId
	pr.UserPassword = password
	root := soap.NewRoot()
	root.Body.Request = pr
	return root.Marshal()
}
