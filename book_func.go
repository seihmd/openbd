package openbd

import (
	"errors"
	"strings"
	"time"
)

// IsValidData validates book has valid data
func (b *Book) IsValidData() bool {
	return b.GetISBN() != ""
}

// GetPublisher returns Publisher
func (b *Book) GetPublisher() string {
	return b.Summary.Publisher
}

// GetISBN returns ISBN
func (b *Book) GetISBN() string {
	return b.Summary.ISBN
}

// GetPubdateStr returns PubdateStr
func (b *Book) GetPubdateStr() string {
	return b.Summary.Pubdate
}

// GetPubdate returns Pubdate
func (b *Book) GetPubdate() (d time.Time, err error) {
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

// GetTitle returns Title
func (b *Book) GetTitle() string {
	return b.Summary.Title
}

// GetSeries returns Series
func (b *Book) GetSeries() string {
	return b.Summary.Series
}

// GetAuthor returns Author
func (b *Book) GetAuthor() string {
	return b.Summary.Author
}

// GetCover returns Cover
func (b *Book) GetCover() string {
	return b.Summary.Cover
}

// GetVolume returns Volume
func (b *Book) GetVolume() string {
	return b.Summary.Volume
}

// GetImageLink returns ImageLink
func (b *Book) GetImageLink() (l string) {
	if b.Summary.Cover != "" {
		l = b.Summary.Cover
		return
	}
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

// GetDescription returns Description
func (b *Book) GetDescription() string {
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

// GetTableOfContents returns TableOfContents
func (b *Book) GetTableOfContents() string {
	tcs := b.Onix.CollateralDetail.TextContent
	toc := ""
	for _, tc := range tcs {
		if tc.TextType == textTypeToc {
			toc = tc.Text
		}
	}
	return toc
}

// GetPages returns Pages
func (b *Book) GetPages() (int, error) {
	extents := b.Onix.DescriptiveDetail.Extent
	if len(extents) > 0 {
		pages, err := extents[0].ExtentValue.Int64()
		return int(pages), err
	}
	return 0, errors.New("no page data")
}
