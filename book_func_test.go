package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/seihmd/openbd/test"
)

func TestGetter(t *testing.T) {
	b, err := mapToBook([]byte(test.BookResponse))

	if err != nil {
		t.Fatal(err.Error())
	}

	if !b.IsValidData() {
		t.Fatalf("failed: Book.IsValidData")
	}
	expectISBN := "9784774183923"
	if b.GetISBN() != expectISBN {
		t.Errorf(fmt.Sprintf("GetISBN expect:%s actual:%s", expectISBN, b.GetISBN()))
	}
	expectPublisher := "技術評論社"
	if b.GetPublisher() != expectPublisher {
		t.Errorf(fmt.Sprintf("GetPublisher expect:%s actual:%s", expectPublisher, b.GetPublisher()))
	}
	if b.GetPubdateStr() != "2016-10" {
		t.Errorf(fmt.Sprintf("GetISBN expect:2016-10 actual:%s", b.GetPubdateStr()))
	}
	actualPubdate, _ := time.Parse("2006-01", "2016-10")
	expectPubdate, _ := b.GetPubdate()
	if actualPubdate != expectPubdate {
		t.Errorf(fmt.Sprintf("GetISBN expect:%d actual:%d", expectPubdate, actualPubdate))
	}

	expectTitle := "みんなのGo言語 : 現場で使える実践テクニック"
	if b.GetTitle() != expectTitle {
		t.Errorf(fmt.Sprintf("GetTitle expect:%s actual:%s", expectTitle, b.GetTitle()))
	}

	expectSeries := "dummy_series"
	if b.GetSeries() != expectSeries {
		t.Errorf(fmt.Sprintf("GetSeries expect:%s actual:%s", expectSeries, b.GetSeries()))
	}
	expectAuthor := "mattn／著 中島大一／著 牧大輔／著 藤原俊一郎／著 鈴木健太／著 松木雅幸／著"
	if b.GetAuthor() != expectAuthor {
		t.Errorf(fmt.Sprintf("GetAuthor expect:%s actual:%s", expectAuthor, b.GetAuthor()))

	}
	expectCover := "http://example.com/xxx.jpg"
	if b.GetCover() != expectCover {
		t.Errorf(fmt.Sprintf("GetCover expect:%s actual:%s", expectCover, b.GetCover()))
	}
	expectVolume := "dummy_volume"
	if b.GetVolume() != expectVolume {
		t.Errorf(fmt.Sprintf("GetVolume expect:%s actual:%s", expectVolume, b.GetVolume()))

	}

	expectImageLink := "http://example.com/xxx.jpg"
	if b.GetImageLink() != expectImageLink {
		t.Errorf(fmt.Sprintf("GetImageLink expect:%s actual:%s", expectImageLink, b.GetImageLink()))
	}

	expectDescription := "dummy_description"
	if b.GetDescription() != expectDescription {
		t.Errorf(fmt.Sprintf("GetDescription expect:%s actual:%s", expectDescription, b.GetDescription()))
	}

	expectTOC := "dummy_toc"
	if b.GetTableOfContents() != expectTOC {
		t.Errorf(fmt.Sprintf("GetTableOfContents expect:%s actual:%s", expectTOC, b.GetTableOfContents()))
	}

	expectPages := 143
	pages, err := b.GetPages()
	if err != nil {
		t.Errorf(err.Error())
	}
	if pages != expectPages {
		t.Errorf(fmt.Sprintf("GetPages expect: %d actual: %d", expectPages, pages))
	}
}

func TestPubdate(t *testing.T) {
	expect, _ := time.Parse("20060102", "20171001")

	jsons := []string{
		test.PubdateYMDHyphen,
		test.PubdateYMHyphen,
		test.PubdateYMD,
		test.PubdateYM,
	}
	for _, j := range jsons {
		b, _ := mapToBook([]byte(j))
		actual, _ := b.GetPubdate()
		if actual != expect {
			t.Errorf(fmt.Sprintf("expect:%s actual:%s", expect, actual))
		}
	}
}
