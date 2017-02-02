package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/seihmd/goisbn"
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
	errUnmartial   = errors.New("ERROR: failed unmarshal json to book")
	errInvalidISBN = errors.New("ERROR: invalid isbn code")
	errRequest     = errors.New("ERROR: request error")
	errOverGetMax  = errors.New("ERROR: isbns over 1000")
	errOverPostMax = errors.New("ERROR: isbns over 10000")
)

func Get(isbn string) (Book, error) {
	if !goisbn.IsISBN(isbn) {
		return Book{}, errInvalidISBN
	}
	u := createGetURL(isbn)
	body, err := requestGet(u)
	if err != nil {
		return Book{}, errRequest
	}
	b, err := mapToBook(body)
	if err != nil {
		return Book{}, err
	}
	return b, nil
}

func GetBooks(isbns []string) ([]Book, error) {
	if len(isbns) > maxForGetRequest {
		return nil, errOverGetMax
	}
	u := createISBNsURL(isbns)
	body, err := requestGet(u)
	if err != nil {
		return nil, errRequest
	}
	b, err := mapToBooks(body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func createGetURL(isbn string) string {
	return getAPI + "?isbn=" + isbn
}

func createISBNsURL(isbns []string) string {
	param := strings.Join(isbns, ",")
	return getAPI + "?isbn=" + param
}

func requestGet(url string) ([]byte, error) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func mapToBook(response []byte) (Book, error) {
	var books []Book
	err := json.Unmarshal(response, &books)
	if err != nil {
		return Book{}, err
	}
	b := books[0]
	if !b.IsValidData() {
		return b, errNoData
	}
	return b, nil
}

func mapToBooks(response []byte) ([]Book, error) {
	var b []Book
	err := json.Unmarshal(response, &b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
