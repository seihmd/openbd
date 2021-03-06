package test

// BooksResponse is response of https://api.openbd.jp/v1/get?isbn=4798142417,4873117526,4798031801&pretty
const BooksResponse = `
[
    {
        "onix":{
            "CollateralDetail":{
                "SupportingResource":[
                    {
                        "ResourceContentType":"01",
                        "ResourceMode":"03",
                        "ContentAudience":"01",
                        "ResourceVersion":[
                            {
                                "ResourceLink":"https://cover.openbd.jp/9784798142418.jpg",
                                "ResourceForm":"02",
                                "ResourceVersionFeature":[
                                    {
                                        "ResourceVersionFeatureType":"01",
                                        "FeatureValue":"D502"
                                    },
                                    {
                                        "ResourceVersionFeatureType":"04",
                                        "FeatureValue":"9784798142418.jpg"
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "TextContent":[
                    {
                        "Text":"本書はGo言語の入門書です。Go言語の基本的な構文から、特徴的な機能、実際にプログラムを書いてみたい方に必要な知識を解説。",
                        "TextType":"02",
                        "ContentAudience":"00"
                    },
                    {
                        "Text":"Go 1.6に対応！\n構文や動作を「なぜ・どうして」から解説。\n読みやすさ、理解の深さを追求したGo入門書！\n\n本書は、Goプログラミングをこれから始める方のための学習書です。Go言語の基本的な構文から、特徴的な機能、開発ツールや使用頻度の高いパッケージの使い方まで、実際にプログラムを書いてみたい方に必要な知識を解説しています。\n\nGo言語の最新バージョン1.6に対応。データ型からチャネルとゴルーチンなどの特徴的な機能、各種パッケージまで、コード例をふんだんに使って「なぜそうなるのか」から説明していますので、C/C++、Java、C#、PythonやPHP、Rubyなど何かしらのプログラミング言語を学んだ方であれば「なるほど、なるほど」と腹落ちしながら読み進められます。\n\n使用頻度の高いパッケージの使い方は集中的に解説するほか、巻末には標準ライブラリのパッケージカタログ付き。座右において便利に使える一冊です。\n\n■Go言語とは\nGo言語(Golangとも呼ばれます)は、2009年に米Googleが公開したプログラミング言語です。構文は簡潔で可読性が高く、コンパイル言語ゆえの安全性とインタプリタ言語のような開発の軽快さを両立し、マルチコアや並行処理などに対応していることから、次世代のソフトウェア開発を支えるプログラミング言語の1つと期待されています。\n\n",
                        "TextType":"03",
                        "ContentAudience":"00"
                    },
                    {
                        "Text":"はじめに\n  Goとは\n  Goの特徴\n  本書の想定読者\n  本書の構成\n  対象とするGoのバージョン\n\nChapter 1：開発環境\n  ～ Windows・OS X・LinuxへのGoのインストール\n  1.1 はじめに\n  1.2 公式ページ\n  1.3 Goのダウンロード\n  1.4 Windows環境への導入\n  1.5 OS X環境への導入\n  1.6 Linux環境への導入\n  1.7 SCMとの連携\n  1.8 開発環境について\n\nChapter 2：プログラムの構成と実行\n  ～ Goプログラムの書き方とビルド・実行・パッケージ作成\n  2.1 はじめに\n  2.2 テキストエディターとエンコーディング\n  2.3 go run\n  2.4 プログラムのビルド\n  2.5 パッケージと構成\n\nChapter 3：言語の基本\n  ～変数・型・演算子・関数・定数・スコープ・制御構文・ゴルーチン\n  3.1 コメント\n  3.2 文\n  3.3 定義済み識別子\n  3.4 コード例の表記について\n  3.5 fmtパッケージ\n  3.6 変数\n  3.7 基本型\n  3.8 配列型\n  3.9 interface{}とnil\n  3.10 演算子\n  3.11 関数\n  3.12 定数\n  3.13 スコープ\n  3.14 制御構文\n\nChapter 4：参照型\n  ～スライス・マップ・チャネル\n  4.1 参照型とは\n  4.2 組み込み関数make\n  4.3 スライス\n  4.4 マップ\n  4.5 チャネル\n\nChapter 5：構造体とインターフェース\n  ～ポインタ・構造体・メソッド・タグ・インターフェース\n  5.1 はじめに\n  5.2 ポインタ\n  5.3 構造体\n  5.4 インターフェース\n\nChapter 6：Goのツール\n  ～さまざまなGoコマンド\n  6.1 Goのツール群について\n  6.2 goコマンド\n  6.3 go version\n  6.4 go env\n  6.5 go fmt\n  6.6 go doc\n  6.7 go build\n  6.8 go install\n  6.9 go get\n  6.10 go test\n  6.11 ベンダリング\n\nChapter 7：Goのパッケージ\n  ～よく使われるパッケージとコーディング例\n  7.1 はじめに\n  7.2 os\n  7.3 time\n  7.4 math\n  7.5 math/rand\n  7.6 flag\n  7.7 fmt\n  7.8 log\n  7.9 strconv\n  7.10 unicode\n  7.11 strings\n  7.12 io\n  7.13 bufio\n  7.14 ioutil\n  7.15 regexp\n  7.16 json\n  7.17 net/url\n  7.18 net/http\n  7.19 sync\n  7.20 crypto/*\n\n巻末付録：Go標準ライブラリオーバービュー\n",
                        "TextType":"04",
                        "ContentAudience":"00"
                    }
                ]
            },
            "PublishingDetail":{
                "Imprint":{
                    "ImprintName":"翔泳社",
                    "ImprintIdentifier":[
                        {
                            "IDValue":"7981",
                            "ImprintIDType":"19"
                        },
                        {
                            "IDValue":"3602",
                            "ImprintIDType":"24"
                        }
                    ]
                },
                "Publisher":{
                    "PublisherIdentifier":[
                        {
                            "PublisherIDType":"19",
                            "IDValue":"7981"
                        },
                        {
                            "PublisherIDType":"24",
                            "IDValue":"3602"
                        }
                    ],
                    "PublisherName":"翔泳社",
                    "PublishingRole":"01"
                },
                "PublishingDate":[
                    {
                        "Date":"20160414",
                        "PublishingDateRole":"01"
                    },
                    {
                        "Date":"20160311",
                        "PublishingDateRole":"09"
                    },
                    {
                        "Date":"20160322",
                        "PublishingDateRole":"25"
                    }
                ]
            },
            "RecordReference":"9784798142418",
            "NotificationType":"03",
            "ProductIdentifier":{
                "IDValue":"9784798142418",
                "ProductIDType":"15"
            },
            "ProductSupply":{
                "SupplyDetail":{
                    "ReturnsConditions":{
                        "ReturnsCode":"03",
                        "ReturnsCodeType":"04"
                    },
                    "ProductAvailability":"99",
                    "Price":[
                        {
                            "PriceAmount":"2980",
                            "CurrencyCode":"JPY",
                            "PriceType":"03"
                        }
                    ]
                }
            },
            "DescriptiveDetail":{
                "Language":[
                    {
                        "LanguageCode":"jpn",
                        "LanguageRole":"01"
                    }
                ],
                "TitleDetail":{
                    "TitleType":"01",
                    "TitleElement":{
                        "TitleElementLevel":"01",
                        "TitleText":{
                            "content":"スターティングGo言語",
                            "collationkey":"スターティングゴーゲンゴ"
                        }
                    }
                },
                "ProductForm":"BA",
                "Collection":{
                    "CollectionType":"10",
                    "TitleDetail":{
                        "TitleType":"01",
                        "TitleElement":[
                            {
                                "TitleElementLevel":"03",
                                "TitleText":{
                                    "content":"CodeZine BOOKS",
                                    "collationkey":"コードジンブックス"
                                }
                            }
                        ]
                    }
                },
                "Audience":[
                    {
                        "AudienceCodeType":"22",
                        "AudienceCodeValue":"00"
                    }
                ],
                "Extent":[
                    {
                        "ExtentValue":"432",
                        "ExtentUnit":"03",
                        "ExtentType":"11"
                    }
                ],
                "Contributor":[
                    {
                        "ContributorRole":[
                            "A01"
                        ],
                        "PersonName":{
                            "content":"松尾 愛賀",
                            "collationkey":"マツオ アイガ"
                        },
                        "SequenceNumber":"1"
                    }
                ],
                "ProductComposition":"00",
                "ProductFormDetail":"B108",
                "Subject":[
                    {
                        "SubjectSchemeIdentifier":"78",
                        "SubjectCode":"3055"
                    },
                    {
                        "SubjectSchemeIdentifier":"20",
                        "SubjectHeadingText":"Go言語;Golang;Go Lang;Go;プログラミング言語Go;Goプログラミング;Go入門;並列;並行;ゴルーチン;Goパッケージ;標準ライブラリ;はじめてのGo;やさしいGo;たのしいGo;モダンプログラミング;プログラミング;入門;初心者"
                    }
                ]
            }
        },
        "hanmoto":{
            "datemodified":"2016-03-17 10:02:12",
            "datecreated":"2016-03-12 16:02:11"
        },
        "summary":{
            "publisher":"翔泳社",
            "isbn":"9784798142418",
            "pubdate":"20160414",
            "title":"スターティングGo言語",
            "series":"CodeZine BOOKS",
            "author":"松尾愛賀／著",
            "cover":"https://cover.openbd.jp/9784798142418.jpg",
            "volume":""
        }
    },
    {
        "onix":{
            "CollateralDetail":{
                "TextContent":[
                    {
                        "Text":"Goプログラミングによる開発手法および関連技術の使い方をマスターし、ツールやプログラムの開発スキルの向上が可能。",
                        "TextType":"02",
                        "ContentAudience":"00"
                    },
                    {
                        "Text":"Go言語の実践テクニックを身につけたいGoプログラミング経験者は必読！\nGoプログラミングについて一歩踏み込んだプロユースの解説書。読者はシンプルなコードを書きながら、実運用アプリケーションの開発で使うスキルとテクニックを学ぶことができます。本書のサンプルプログラムはどれもシンプルですがとても実践的です。拡張性、並行処理、高可用性といったエンタープライズアプリケーションの開発で直面する現実的な問題に対するソリューションが含まれています。本書を読めば、実際の業務に必要な技能――Goによる開発手法および関連技術の使い方――をマスターし、ツールやプログラムの開発スキルを迅速かつ簡単に向上できます。日本語版では、監訳者の鵜飼文敏氏による巻末付録「Goらしいコードの書き方」を収録しました。",
                        "TextType":"03",
                        "ContentAudience":"00"
                    }
                ]
            },
            "PublishingDetail":{
                "Imprint":{
                    "ImprintName":"オライリー・ジャパン",
                    "ImprintIdentifier":[
                        {
                            "IDValue":"87311",
                            "ImprintIDType":"19"
                        },
                        {
                            "IDValue":"0742",
                            "ImprintIDType":"24"
                        }
                    ]
                },
                "Publisher":{
                    "PublisherIdentifier":[
                        {
                            "PublisherIDType":"19",
                            "IDValue":"274"
                        },
                        {
                            "PublisherIDType":"24",
                            "IDValue":"0742"
                        }
                    ],
                    "PublisherName":"株式会社オーム社",
                    "PublishingRole":"01"
                },
                "PublishingDate":[
                    {
                        "Date":"20160121",
                        "PublishingDateRole":"01"
                    }
                ]
            },
            "RecordReference":"9784873117522",
            "NotificationType":"03",
            "ProductIdentifier":{
                "IDValue":"9784873117522",
                "ProductIDType":"15"
            },
            "ProductSupply":{
                "SupplyDetail":{
                    "ReturnsConditions":{
                        "ReturnsCode":"03",
                        "ReturnsCodeType":"04"
                    },
                    "ProductAvailability":"99",
                    "Price":[
                        {
                            "PriceAmount":"3200",
                            "CurrencyCode":"JPY",
                            "PriceType":"03"
                        }
                    ]
                }
            },
            "DescriptiveDetail":{
                "TitleDetail":{
                    "TitleType":"01",
                    "TitleElement":{
                        "TitleElementLevel":"01",
                        "TitleText":{
                            "content":"Go言語によるWebアプリケーション開発",
                            "collationkey":"ゴーゲンゴニヨル　ウェブアプリケーションカイハツ"
                        }
                    }
                },
                "Language":[
                    {
                        "LanguageCode":"jpn",
                        "LanguageRole":"01",
                        "CountryCode":"JP"
                    }
                ],
                "ProductForm":"BA",
                "Audience":[
                    {
                        "AudienceCodeType":"22",
                        "AudienceCodeValue":"00"
                    }
                ],
                "Extent":[
                    {
                        "ExtentValue":"280",
                        "ExtentUnit":"03",
                        "ExtentType":"11"
                    }
                ],
                "Contributor":[
                    {
                        "ContributorRole":[
                            "A01"
                        ],
                        "PersonName":{
                            "content":"Mat Ryer",
                            "collationkey":"マット ライヤー"
                        },
                        "SequenceNumber":"1"
                    },
                    {
                        "ContributorRole":[
                            "B20"
                        ],
                        "PersonName":{
                            "content":"鵜飼 文敏",
                            "collationkey":"ウカイ フミトシ"
                        },
                        "SequenceNumber":"2"
                    },
                    {
                        "ContributorRole":[
                            "B06"
                        ],
                        "PersonName":{
                            "content":"牧野 聡",
                            "collationkey":"マキノ サトシ"
                        },
                        "SequenceNumber":"3"
                    }
                ],
                "ProductComposition":"00",
                "ProductFormDetail":"B124",
                "Subject":[
                    {
                        "SubjectSchemeIdentifier":"78",
                        "SubjectCode":"3055"
                    },
                    {
                        "SubjectSchemeIdentifier":"79",
                        "SubjectCode":"20"
                    }
                ]
            }
        },
        "hanmoto":{
            "datemodified":"2016-01-06 16:04:16",
            "datecreated":"2015-12-24 16:04:45"
        },
        "summary":{
            "publisher":"オライリー・ジャパン",
            "isbn":"9784873117522",
            "pubdate":"20160121",
            "title":"Go言語によるWebアプリケーション開発",
            "series":"",
            "author":"MatRyer／著 鵜飼文敏／監修 牧野聡／翻訳",
            "cover":"",
            "volume":""
        }
    },
    {
        "onix":{
            "CollateralDetail":{
                "SupportingResource":[
                    {
                        "ResourceContentType":"01",
                        "ResourceMode":"03",
                        "ContentAudience":"01",
                        "ResourceVersion":[
                            {
                                "ResourceLink":"https://cover.openbd.jp/9784798031804.jpg",
                                "ResourceForm":"02",
                                "ResourceVersionFeature":[
                                    {
                                        "ResourceVersionFeatureType":"01",
                                        "FeatureValue":"D502"
                                    },
                                    {
                                        "ResourceVersionFeatureType":"04",
                                        "FeatureValue":"9784798031804.jpg"
                                    }
                                ]
                            }
                        ]
                    }
                ],
                "TextContent":[
                    {
                        "Text":"Go言語とGoogle App EngineでWebアプリを作る。クラウド環境でWebアプリ製作。",
                        "TextType":"03",
                        "ContentAudience":"00"
                    },
                    {
                        "Text":"1 はじめに\n2 開発環境の構築\n3 Go言語基礎\n4 Google App Engine API\n5 Google App Engine for Goを使ったWebアプリケーション作成\n補章 今後のSDKの変更と、APIの追加に備えて",
                        "TextType":"04",
                        "ContentAudience":"00"
                    }
                ]
            },
            "PublishingDetail":{
                "Imprint":{
                    "ImprintName":"秀和システム",
                    "ImprintIdentifier":[
                        {
                            "IDValue":"7980",
                            "ImprintIDType":"19"
                        }
                    ]
                },
                "PublishingDate":[
                    {
                        "Date":"",
                        "PublishingDateRole":"01"
                    }
                ]
            },
            "RecordReference":"9784798031804",
            "NotificationType":"03",
            "ProductIdentifier":{
                "IDValue":"9784798031804",
                "ProductIDType":"15"
            },
            "ProductSupply":{
                "SupplyDetail":{
                    "ReturnsConditions":{
                        "ReturnsCode":"02",
                        "ReturnsCodeType":"04"
                    },
                    "ProductAvailability":"99"
                }
            },
            "DescriptiveDetail":{
                "TitleDetail":{
                    "TitleType":"01",
                    "TitleElement":{
                        "TitleElementLevel":"01",
                        "TitleText":{
                            "content":"Go言語プログラミング入門 : on Google App Engine",
                            "collationkey":"Go ゲンゴ プログラミング ニュウモン : on Google App Engine"
                        }
                    }
                },
                "Language":[
                    {
                        "LanguageCode":"jpn",
                        "LanguageRole":"01",
                        "CountryCode":"JP"
                    }
                ],
                "ProductForm":"BZ",
                "Extent":[
                    {
                        "ExtentValue":"364",
                        "ExtentUnit":"03",
                        "ExtentType":"11"
                    }
                ],
                "Measure":[
                    {
                        "MeasureType":"01",
                        "MeasureUnitCode":"mm",
                        "Measurement":"240"
                    },
                    {
                        "MeasureType":"02",
                        "MeasureUnitCode":"mm",
                        "Measurement":"0"
                    }
                ],
                "Contributor":[
                    {
                        "ContributorRole":[
                            "A01"
                        ],
                        "PersonName":{
                            "content":"横山 隆司"
                        },
                        "SequenceNumber":"1"
                    }
                ],
                "ProductComposition":"00"
            }
        },
        "hanmoto":{
            "dateshuppan":"2011-12",
            "datemodified":"2016-06-18 12:01:02",
            "datecreated":"2016-06-18 12:01:02"
        },
        "summary":{
            "publisher":"秀和システム",
            "isbn":"9784798031804",
            "pubdate":"2011-12",
            "title":"Go言語プログラミング入門 : on Google App Engine",
            "series":"",
            "author":"横山隆司／著",
            "cover":"https://cover.openbd.jp/9784798031804.jpg",
            "volume":""
        }
    }
]
`
