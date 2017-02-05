// Package openbd is API wrapper library for openBD: https://openbd.jp/
package openbd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	getAPI      = "https://api.openbd.jp/v1/get"
	postAPI     = "https://api.openbd.jp/v1/post"
	coverageAPI = "https://api.openbd.jp/v1/coverage"
)

const (
	maxForGetRequest  = 1000
	maxForPostRequest = 10000
)

var (
	errNoData      = errors.New("ERROR: no data")
	errInvalidISBN = errors.New("ERROR: invalid isbn code")
	errRequest     = errors.New("ERROR: request error")
	errOverGetMax  = errors.New("ERROR: isbns over 1000")
)

// OpenBD has api request funcs
type OpenBD struct {
	Client *http.Client
}

// New returns OpenBD
func New() OpenBD {
	return OpenBD{
		Client: http.DefaultClient,
	}
}

// Get requests single Book data
func (o *OpenBD) Get(isbn string) (*Book, error) {
	u := createGetURL(isbn)
	body, err := o.requestGet(u)
	if err != nil {
		return nil, errRequest
	}
	b, err := mapToBook(body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GetBooks requests multiple Book data
func (o *OpenBD) GetBooks(isbns []string) (*[]Book, error) {
	if len(isbns) > maxForGetRequest {
		return nil, errOverGetMax
	}
	u := createISBNsURL(isbns)
	body, err := o.requestGet(u)
	if err != nil {
		return nil, errRequest
	}
	b, err := mapToBooks(body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (o *OpenBD) requestGet(url string) ([]byte, error) {
	resp, err := o.Client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func createGetURL(isbn string) string {
	return getAPI + "?isbn=" + isbn
}

func createISBNsURL(isbns []string) string {
	return getAPI + "?isbn=" + strings.Join(isbns, ",")
}

func mapToBook(response []byte) (b *Book, err error) {
	var books []Book
	if err = json.Unmarshal(response, &books); err != nil {
		return
	}
	b = &books[0]
	if !b.IsValidData() {
		err = errNoData
	}
	return
}

func mapToBooks(response []byte) (*[]Book, error) {
	var b []Book
	err := json.Unmarshal(response, &b)
	return &b, err
}
