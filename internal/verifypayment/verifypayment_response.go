package verifypayment

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/elyarsadig/behpardakht/internal/bankerrors"
)

type verifyPaymentResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"Body"`
		BpPay   struct {
			XMLName xml.Name `xml:"bpVerifyRequestResponse"`
			Return  string   `xml:"return"`
		}
	}
	responseCode int    `xml:"-"`
	rawResponse  []byte `xml:"-"`
}

func (v *verifyPaymentResponse) parseAndSetResponseFields() error {
	returnCode := v.Body.BpPay.Return
	code, err := strconv.Atoi(returnCode)
	if err != nil {
		return err
	}
	v.responseCode = code
	return nil
}

func ProcessVerifyPaymentResponse(response []byte) error {
	paymentResponse := new(verifyPaymentResponse)
	if err := xml.Unmarshal(response, paymentResponse); err != nil {
		return fmt.Errorf("error in unmarshal response: %v", err)
	}
	if err := paymentResponse.parseAndSetResponseFields(); err != nil {
		return fmt.Errorf("error in parseAndSetResponseFields: %v", err)
	}
	if err := bankerrors.GetBankErrorMessage(paymentResponse.responseCode); err != nil {
		return err
	}
	return nil
}
