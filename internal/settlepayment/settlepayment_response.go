package settlepayment

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/elyarsadig/behpardakht/internal/bankerrors"
)

type settlePaymentResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		XMLName xml.Name `xml:"Body"`
		BpPay   struct {
			XMLName xml.Name `xml:"bpSettleRequestResponse"`
			Return  string   `xml:"return"`
		}
	}
	responseCode int    `xml:"-"`
	rawResponse  []byte `xml:"-"`
}

func (s *settlePaymentResponse) parseAndSetResponseFields() error {
	returnCode := s.Body.BpPay.Return
	code, err := strconv.Atoi(returnCode)
	if err != nil {
		return err
	}
	s.responseCode = code
	return nil
}

func ProcessSettlePaymentResponse(response []byte) error {
	settlePaymentResponse := new(settlePaymentResponse)
	if err := xml.Unmarshal(response, settlePaymentResponse); err != nil {
		return fmt.Errorf("error in unmarshal response: %v", err)
	}
	if err := settlePaymentResponse.parseAndSetResponseFields(); err != nil {
		return fmt.Errorf("error in parseAndSetResponseFields: %v", err)
	}
	if err := bankerrors.GetBankErrorMessage(settlePaymentResponse.responseCode); err != nil {
		return err
	}
	return nil
}
