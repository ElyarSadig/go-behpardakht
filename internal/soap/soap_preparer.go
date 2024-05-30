package soap

type SOAPRequestPreparer interface {
	PrepareSOAPRequest(userId string, password string) ([]byte, error)
}
