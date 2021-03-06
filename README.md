# openbd-api -- APIs for openBD by Golang 

[![check vulns](https://github.com/spiegel-im-spiegel/openbd-api/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/openbd-api/actions)
[![lint status](https://github.com/spiegel-im-spiegel/openbd-api/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/openbd-api/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/openbd-api/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/openbd-api.svg)](https://github.com/spiegel-im-spiegel/openbd-api/releases/latest)

This package is required Go 1.16 or later.

## Usage of package

### Import Package

```
import "github.com/spiegel-im-spiegel/openbd-api"
```

### Lookup openBD Book Data

```go
books, err := openbd.DefaultClient().LookupBook([]string{"9784797369915", "9784274069321"})
```

## Entities for openBD

### Book type

```go
//Book is entity class of book info.
type Book struct {
    Onix    Onix    `json:"onix"`
    Hanmoto Hanmoto `json:"hanmoto"`
    Summary Summary `json:"summary"`
}

//Onix is entity class of JPRO-onix items
type Onix struct {
    RecordReference   string //ISBN code (key code)
    NotificationType  string
    ProductIdentifier struct {
        ProductIDType string
        IDValue       string //ISBN ?
    }
    DescriptiveDetail struct {
        ProductComposition string
        ProductForm        string
        Measure            []struct {
            MeasureType     string
            Measurement     string
            MeasureUnitCode string
        } `json:",omitempty"`
        Collection struct {
            CollectionType     string
            CollectionSequence *struct {
                CollectionSequenceType     string `json:",omitempty"`
                CollectionSequenceTypeName string `json:",omitempty"`
                CollectionSequenceNumber   string `json:",omitempty"`
            } `json:",omitempty"`
            TitleDetail *struct {
                TitleType    string `json:",omitempty"`
                TitleElement []struct {
                    TitleElementLevel string
                    TitleText         struct {
                        Content      string `json:"content"`
                        CollationKey string `json:"collationkey,omitempty"`
                    }
                } `json:",omitempty"`
            } `json:",omitempty"`
        }
        TitleDetail struct {
            TitleType    string
            TitleElement struct {
                TitleElementLevel string
                TitleText         struct {
                    Content      string `json:"content"`
                    Collationkey string `json:"collationkey,omitempty"`
                }
            }
        }
        Contributor []struct {
            SequenceNumber  string
            ContributorRole []string `json:",omitempty"`
            PersonName      struct {
                Content      string `json:"content"`
                Collationkey string `json:"collationkey,omitempty"`
            }
            BiographicalNote string `json:",omitempty"`
        } `json:",omitempty"`
        Language []struct {
            LanguageRole string
            LanguageCode string
            CountryCode  string
        } `json:",omitempty"`
        Extent []struct {
            ExtentType  string
            ExtentValue string
            ExtentUnit  string
        } `json:",omitempty"`
        Subject []struct {
            SubjectSchemeIdentifier string
            SubjectCode             string
            SubjectHeadingText      string `json:",omitempty"`
        } `json:",omitempty"`
        Audience []struct {
            AudienceCodeType  string
            AudienceCodeValue string
        } `json:",omitempty"`
    }
    CollateralDetail *struct {
        TextContent []struct {
            TextType        string
            ContentAudience string
            Text            string
        } `json:",omitempty"`
        SupportingResource []struct {
            ResourceContentType string
            ContentAudience     string
            ResourceMode        string
            ResourceVersion     []struct {
                ResourceForm           string
                ResourceVersionFeature []struct {
                    ResourceVersionFeatureType string
                    FeatureValue               string
                } `json:",omitempty"`
                ResourceLink string
            } `json:",omitempty"`
        } `json:",omitempty"`
    } `json:",omitempty"`
    PublishingDetail struct {
        Imprint struct {
            ImprintIdentifier []struct {
                ImprintIDType string
                IDValue       string
            } `json:",omitempty"`
            ImprintName string
        }
        Publisher struct {
            PublisherIdentifier []struct {
                PublisherIDType string
                IDValue         string
            } `json:",omitempty"`
            PublishingRole string
            PublisherName  string
        }
        PublishingDate []struct {
            Date               Date
            PublishingDateRole string
        } `json:",omitempty"`
    }
    ProductSupply struct {
        SupplyDetail struct {
            ReturnsConditions struct {
                ReturnsCodeType string
                ReturnsCode     string
            }
            ProductAvailability string
            Price               []struct {
                PriceType    string
                CurrencyCode string
                PriceAmount  string
            } `json:",omitempty"`
        }
    }
}

//Hanmoto is entity class of Hanmoto dot com items
type Hanmoto struct {
    DatePublished Date `json:"dateshuppan"`
    DateModified  Date `json:"datemodified"`
    DateCreated   Date `json:"datecreated"`
    DateReleased  Date `json:"datekoukai"`
    IsLightNovel  bool `json:"lanove,omitempty"`
    HasReview     bool `json:"hasshohyo,omitempty"`
    Reviews       []struct {
        Reviewer       string `json:"reviewer"`
        Link           string `json:"link"`
        DateAppearance Date   `json:"appearance"`
        SourceKindID   int    `json:"kubun_id"`
        SourceID       int    `json:"source_id"`
        Source         string `json:"source"`
        PaperType      string `json:"choyukan"`
        PostUser       string `json:"post_user"`
        Han            string `json:"han"`
        Gou            string `json:"gou"`
    } `json:"reviews,omitempty"`
    HasSample bool `json:"hastameshiyomi,omitempty"`
}

//Summary is entity class of summary data
type Summary struct {
    ISBN      string `json:"isbn"`
    Title     string `json:"title"`
    Volume    string `json:"volume"`
    Series    string `json:"series"`
    Publisher string `json:"publisher"`
    PubDate   Date   `json:"pubdate"`
    Author    string `json:"author"`
    Cover     string `json:"cover"`
}
```

## Command Line Interface (Sample Code)

### Download and Build

```
$ go get github.com/spiegel-im-spiegel/openbd-api/cli/openbd
```

### Lookup openBD Books Data

```
$ openbd lookup 9784797369915 9784274069321
```

[openbd-api]: https://github.com/spiegel-im-spiegel/openbd-api "spiegel-im-spiegel/openbd-api: APIs for openBD by Golang"
