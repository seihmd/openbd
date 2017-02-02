package main

import (
	"errors"
	"strings"
	"time"
)

// IsValidData returns book has valid data
func (b Book) IsValidData() bool {
	return b.GetISBN() != ""
}

func (b Book) GetPublisher() string {
	return b.Summary.Publisher
}

func (b Book) GetISBN() string {
	return b.Summary.ISBN
}

func (b Book) GetPubdateStr() string {
	return b.Summary.Pubdate
}

func (b Book) GetPubdate() (d time.Time, err error) {
	p := b.Summary.Pubdate
	if p == "" {
		err = errors.New("no pubdate")
		return
	}
	layoutYMD := "20060102"
	layoutYM := "200601"
	if strings.Contains(p, "-") {
		layoutYMD = "2006-01-02"
		layoutYM = "2006-01"
	}
	d, err = time.Parse(layoutYMD, p)
	if err != nil {
		d, err = time.Parse(layoutYM, p)
	}
	return
}

func (b Book) GetTitle() string {
	return b.Summary.Title
}

func (b Book) GetSeries() string {
	return b.Summary.Series
}

func (b Book) GetAuthor() string {
	return b.Summary.Author
}

func (b Book) GetCover() string {
	return b.Summary.Cover
}

func (b Book) GetVolume() string {
	return b.Summary.Volume
}

func (b Book) GetImageLink() (l string) {
	srs := b.Onix.CollateralDetail.SupportingResource
	if len(srs) == 0 {
		return
	}
	for _, sr := range srs {
		if len(sr.ResourceVersion) > 0 {
			l = sr.ResourceVersion[0].ResourceLink
			break
		}
	}
	return
}

func (b Book) GetDescription() string {
	tcs := b.Onix.CollateralDetail.TextContent
	d := ""
	for _, tc := range tcs {
		if tc.TextType == textTypeBrief && d == "" {
			d = tc.Text
		}
		if tc.TextType == textTypeDescription {
			d = tc.Text
		}
	}
	return d
}

func (b Book) GetTableOfContents() string {
	tcs := b.Onix.CollateralDetail.TextContent
	toc := ""
	for _, tc := range tcs {
		if tc.TextType == textTypeToc {
			toc = tc.Text
		}
	}
	return toc
}
