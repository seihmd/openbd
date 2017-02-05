package main

import (
	"testing"

	"github.com/seihmd/openbd/test"
)

func TestMapToBooks(t *testing.T) {
	books, err := mapToBooks([]byte(test.BooksResponse))
	if err != nil {
		t.Fatal(err.Error())
	}
	l := len(books)
	if l != 3 {
		t.Errorf("failed books mapping expect: 3, actual: %d", l)
	}

	isbn0 := books[0].GetISBN()
	if isbn0 != "9784798142418" {
		t.Errorf("books[0] isbn expect: 9784798142418, actual: %s", isbn0)
	}

	isbn1 := books[1].GetISBN()
	if isbn1 != "9784873117522" {
		t.Errorf("books[1] isbn expect: 9784873117522, actual: %s", isbn1)
	}

	isbn2 := books[2].GetISBN()
	if isbn2 != "9784798031804" {
		t.Errorf("books[2] isbn expect: 9784798031804 , actual: %s", isbn2)
	}
}

func TestGet(t *testing.T) {
	o := New()
	for _, isbn := range test.ISBNs {
		b, err := o.Get(isbn)
		if err != nil {
			t.Errorf("isbn: %s, msg: %s", isbn, err.Error())
		} else if b.ISBN != isbn {
			t.Errorf("different isbn code Book.ISBN: %s, expect: %s", b.ISBN, isbn)
		}
		// assert Getter make no panic
		b.GetISBN()
		b.GetPublisher()
		b.GetPubdateStr()
		b.GetPubdate()
		b.GetTitle()
		b.GetSeries()
		b.GetAuthor()
		b.GetCover()
		b.GetVolume()
		b.GetImageLink()
		b.GetDescription()
		b.GetTableOfContents()
		b.GetPages()
	}
}
