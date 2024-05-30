package behpardakht

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/elyarsadig/behpardakht/internal/soap"
)

type behPardakht struct {
	client   *http.Client
	username string
	password string
}

func New(client *http.Client, username string, password string) *behPardakht {
	if client == nil {
		client = &http.Client{}
	}
	return &behPardakht{
		client:   client,
		username: username,
		password: password,
	}
}

func (b *behPardakht) sendRequest(method, url string, soapRequest soap.SOAPRequestPreparer) ([]byte, error) {
	body, err := soapRequest.PrepareSOAPRequest(b.username, b.password)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "text/xml")
	request.Header.Set("charset", "utf-8")
	response, err := b.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to pay with status %d", response.StatusCode)
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
