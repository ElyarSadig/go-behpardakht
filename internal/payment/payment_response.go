package payment

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"

	"github.com/elyarsadig/behpardakht/internal/bankerrors"
)

type paymentResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"Body"`
		BpPay   struct {
			XMLName xml.Name `xml:"bpPayRequestResponse"`
			Return  string   `xml:"return"`
		}
	}
	responseCode int    `xml:"-"`
	refId        string `xml:"-"`
}

func (pr *paymentResponse) parseAndSetResponseFields() error {
	params := strings.Split(pr.Body.BpPay.Return, ",")
	pr.responseCode = -1
	if params[0] == "0" {
		pr.responseCode = 0
	} else if params[0] != "" {
		code, err := strconv.Atoi(params[0])
		if err != nil {
			return fmt.Errorf("invalid response code: %v", err)
		}
		pr.responseCode = code
	}
	if len(params) > 1 {
		pr.refId = params[1]
	}
	return nil
}

func ProcessPaymentResponse(response []byte) (string, error) {
	paymentResponse := new(paymentResponse)
	if err := xml.Unmarshal(response, paymentResponse); err != nil {
		return "", fmt.Errorf("error in unmarshal response: %v", err)
	}
	if err := paymentResponse.parseAndSetResponseFields(); err != nil {
		return "", fmt.Errorf("error in parseAndSetResponseFields: %v", err)
	}
	if err := bankerrors.GetBankErrorMessage(paymentResponse.responseCode); err != nil {
		return "", err
	}
	return paymentResponse.refId, nil
}
