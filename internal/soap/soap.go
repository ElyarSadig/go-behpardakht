package soap

import "encoding/xml"

type soapRoot struct {
	XMLName xml.Name `xml:"x:Envelope"`
	X       string   `xml:"xmlns:x,attr"`
	Ns1     string   `xml:"xmlns:ns1,attr"`
	Body    soapBody
}

type soapBody struct {
	XMLName xml.Name `xml:"x:Body"`
	Request any
}

func (r *soapRoot) Marshal() ([]byte, error) {
	return xml.MarshalIndent(r, "", "  ")
}

func NewRoot() *soapRoot {
	return &soapRoot{
		X:   "http://schemas.xmlsoap.org/soap/envelope/",
		Ns1: "http://interfaces.core.sw.bps.com/",
	}
}
