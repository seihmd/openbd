# openbd

--

    import "github.com/seihmd/openbd"

Package openbd is API wrapper library for openBD: (openbd.jp)[https://openbd.jp/]

## Usage

```go
o := openbd.New()

// 一冊の情報を取得
book, err := o.Get("477418392X")
if err != nil {
    ...
}
book.GetISBN() // 477418392X

// 複数冊の情報を取得 1000冊まで可
isbns := []string{"9780061234002", "9780071123419"}
books, err := o.GetBooks(isbns)
if err != nil {
    ...
}
for _, book := range *books {
    if !book.IsValidData() {
        // fail to get data
    }
}
```

#### type OpenBD

```go
type OpenBD struct {
    Client *http.Client
}
```

OpenBD has api request funcs

#### func  New

```go
func New() OpenBD
```
New returns OpenBD

#### func (*OpenBD) Get

```go
func (o *OpenBD) Get(isbn string) (*Book, error)
```
Get requests single Book data

#### func (*OpenBD) GetBooks

```go
func (o *OpenBD) GetBooks(isbns []string) (*[]Book, error)
```
GetBooks requests multiple Book data


#### type Book

```go
type Book struct {
	Onix    `json:"onix"`    // JPRO-onix準拠項目
	Hanmoto `json:"hanmoto"` // 版元ドットコム独自項目
	Summary `json:"summary"` // 基本事項
}
```

Book :struct to map whole response

#### func (*Book) GetAuthor

```go
func (b *Book) GetAuthor() string
```
GetAuthor returns Author

#### func (*Book) GetCover

```go
func (b *Book) GetCover() string
```
GetCover returns Cover

#### func (*Book) GetDescription

```go
func (b *Book) GetDescription() string
```
GetDescription returns Description

#### func (*Book) GetISBN

```go
func (b *Book) GetISBN() string
```
GetISBN returns ISBN

#### func (*Book) GetImageLink

```go
func (b *Book) GetImageLink() (l string)
```
GetImageLink returns ImageLink

#### func (*Book) GetPages

```go
func (b *Book) GetPages() (int, error)
```
GetPages returns Pages

#### func (*Book) GetPubdate

```go
func (b *Book) GetPubdate() (d time.Time, err error)
```
GetPubdate returns Pubdate

#### func (*Book) GetPubdateStr

```go
func (b *Book) GetPubdateStr() string
```
GetPubdateStr returns PubdateStr

#### func (*Book) GetPublisher

```go
func (b *Book) GetPublisher() string
```
GetPublisher returns Publisher

#### func (*Book) GetSeries

```go
func (b *Book) GetSeries() string
```
GetSeries returns Series

#### func (*Book) GetTableOfContents

```go
func (b *Book) GetTableOfContents() string
```
GetTableOfContents returns TableOfContents

#### func (*Book) GetTitle

```go
func (b *Book) GetTitle() string
```
GetTitle returns Title

#### func (*Book) GetVolume

```go
func (b *Book) GetVolume() string
```
GetVolume returns Volume

#### func (*Book) IsValidData

```go
func (b *Book) IsValidData() bool
```
IsValidData validates book has valid data

#### type Hanmoto

```go
type Hanmoto struct {
	Dateshuppan    string // 出版年月日
	Datemodified   string // 情報更新日時
	Datecreated    string // 情報作成日時
	Lanove         bool   // ライトノベルフラグ
	Hasshohyo      string // 書評データありフラグ
	Hastameshiyomi bool   // ためしよみありフラグ
	Reviews        []struct {
		Han        string // 版数
		Appearance string // 掲載日
		PostUser   string `json:"post_user"` // 入力者
		KubunID    int    `json:"kubun_id"`  // 掲載元区分ID
		Source     string // 掲載元
		Choyukan   string // 朝刊/夕刊
		SourceID   int    `json:"source_id"` // 掲載元ID
		Reviewer   string // 書評者
		Link       string // リンク
	}
	// 版元ドットコム会員社独自項目 ->
	Genshomei                       string // 原書名
	Han                             string // 版"` // 新版、改訂などの版次(エディショ
	Datejuuhanyotei                 string // 重版予定日
	Datezeppan                      string // 絶版日
	Toji                            string // 製本"`  // 上製="上製",並製="並
	Zaiko                           int    // 在庫"` // 在庫あり=11／在庫僅少=21／重版中=／品切れ・重版未定=33／絶版=34
	Maegakinado                     string // まえがきなど
	Hanmotokarahitokoto             string // 版元からひとこと
	Kaisetsu105W                    string // 内容紹介TRC(105字)
	Kanrensho                       string // 関連書
	KanrenLink                      string // 関連リンク	htmlで記入
	Tsuiki                          string // サイト追記
	Genrecodetrc                    int    // ジャンルコード
	Genrecodetrcjidou               string // ジャンルコード(TRC児童書)
	Rubynoumu                       string // ルビの有無
	Ndccode                         string // NDCコード
	Kankoukeitai                    string // 刊行形態
	Sonotatokkijikou                string // その他特記事項
	Jushoujouhou                    string // 受賞情報
	Furoku                          string // 付録
	Furokusonota                    string // 付録その他
	Dokushakakikomi                 string // 読者書き込み
	Dokushakakikomipagesuu          int    // 読者書き込みページ数
	Fuzokushiryounokangaikashidashi int    // 付属資料館外貸出 可=1,不可=2,館内のみ=3
	Obinaiyou                       string // 帯内容
	Ruishokyougousho                string // 類書・競合書
	Bessoushiryou                   string // 別送資料
	Zasshicode                      string // 雑誌コード
	Bikoujpo                        string // 備考(JPO)
	Bikoutrc                        string // 備考(TRC)
	Kanrenshoisbn                   string // 関連書ISBN
	Author                          []struct {
		Dokujikubun string // 独自著者区分
		Listseq     int    // リスト順
	}
	Jyuhan []struct {
		Date    string // 重版日
		Ctime   string // レコード作成日時
		Suri    int    // 刷り数
		Comment string // 重版コメント
	}
}
```

Hanmoto :版元ドットコム独自項目

#### type Onix

```go
type Onix struct {
	RecordReference   string // ISBNコード
	ProductIdentifier struct {
		IDValue       string // ISBNコード
		ProductIDType string // ISBNコード
	}
	DescriptiveDetail struct {
		ProductComposition string // セット商品分売可否
		ProductFormDetail  string // 判型
		Measure            []struct {
			MeasureType     string
			MeasureUnitCode string
			Measurement     string
		}
		ProductPart []struct {
			NumberOfItemsOfThisForm json.Number
			ProductForm             string
			ProductFormDescription  string
		}
		Collection struct {
			TitleDetail struct {
				TitleType    string
				TitleElement []struct {
					TitleElementLevel string
					TitleText         struct {
						Content      string `json:"content"`
						CollationKey string `json:"collationkey"`
					}
				}
			}
			CollectionType string
		}
		TitleDetail struct {
			TitleType    string
			TitleElement struct {
				TitleElementLevel string
				TitleText         struct {
					Content      string
					Collationkey string
				}
			}
		}
		Contributor []struct {
			ContributorRole []string // 著者区分
			PersonName      struct {
				Content      string // 氏名
				Collationkey string // 照合キー
			}
			BiographicalNote string // 著者略歴
			SequenceNumber   json.Number
		}
		Language []struct {
			LanguageCode string
			LanguageRole string
			CountryCode  string
		}
		Extent []struct {
			ExtentValue json.Number
			ExtentUnit  string
			ExtentType  string
		}
		ProductForm string
		Subject     []struct {
			SubjectSchemeIdentifier string
			SubjectCode             string
			SubjectHeadingText      string
		}
		Audience []struct {
			AudienceCodeType  string
			AudienceCodeValue string
		}
	}
	Measure []struct {
		MeasureType     string
		MeasureUnitCode string
		Measurement     string
	}
	CollateralDetail struct {
		TextContent []struct {
			Text            string
			TextType        string
			ContentAudience string
		}
		SupportingResource []struct {
			ResourceContentType string // ISBNコード
			ResourceMode        string
			ContentAudience     string
			ResourceVersion     []struct {
				ResourceLink           string // 画像パス
				ResourceForm           string
				ResourceVersionFeature []struct {
					ResourceVersionFeatureType string
					FeatureValue               string
				}
			}
		}
	}
	PublishingDetail struct {
		Imprint struct {
			ImprintName       string
			ImprintIdentifier []struct {
				IDValue       string
				ImprintIDType string
			}
		}
		Publisher struct {
			PublisherIdentifier []struct {
				PublisherIDType string
				IDValue         string
			}
			PublisherName  string
			PublishingRole string
		}
		PublishingDate []struct {
			Date               string
			PublishingDateRole string
		}
	}
	NotificationType string
	ProductSupply    struct {
		SupplyDetail struct {
			ReturnsConditions struct {
				ReturnsCode     string
				ReturnsCodeType string
			}
			ProductAvailability string
			Price               []struct {
				PriceAmount  json.Number
				CurrencyCode string
				PriceType    string
			}
		}
	}
}
```

Onix :JPRO-onix準拠項目

#### type Summary

```go
type Summary struct {
	Publisher string
	ISBN      string
	Pubdate   string
	Title     string
	Series    string
	Author    string
	Cover     string
	Volume    string
}
```

Summary :使いやすい基本事項
