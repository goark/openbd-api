package openbd

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/spiegel-im-spiegel/errs"
)

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

func (book *Book) String() string {
	if b, err := EncodeBook(book); err == nil {
		return string(b)
	}
	return ""
}

//Id returns id code (= Book.Onix.RecordReference)
func (book *Book) Id() string {
	if book == nil {
		return ""
	}
	return book.Onix.RecordReference
}

//IsValid returns true if Book is valid data
func (book *Book) IsValid() bool {
	return len(book.Id()) > 0
}

//ISBN returns ISBN code
func (book *Book) ISBN() string {
	if !book.IsValid() {
		return ""
	}
	if len(book.Onix.ProductIdentifier.IDValue) > 0 {
		return book.Onix.ProductIdentifier.IDValue
	}
	if len(book.Summary.ISBN) > 0 {
		return book.Summary.ISBN
	}
	return book.Onix.RecordReference
}

//Title returns string of Book Title
func (book *Book) Title() string {
	if !book.IsValid() {
		return ""
	}
	title := ""
	if len(book.Onix.DescriptiveDetail.TitleDetail.TitleType) > 0 {
		title = book.Onix.DescriptiveDetail.TitleDetail.TitleElement.TitleText.Content
	}
	if len(title) == 0 {
		title = book.Summary.Title
	}
	return title
}

//SubTitle returns string of sub-title
func (book *Book) SubTitle() string {
	if !book.IsValid() {
		return ""
	}
	//TODO: Subtitle
	return ""
}

//SeriesTitle returns string of series title
func (book *Book) SeriesTitle() string {
	if !book.IsValid() {
		return ""
	}
	if book.Onix.DescriptiveDetail.Collection.TitleDetail != nil {
		for _, elm := range book.Onix.DescriptiveDetail.Collection.TitleDetail.TitleElement {
			if elm.TitleElementLevel == "01" && len(elm.TitleText.Content) > 0 {
				return elm.TitleText.Content
			}
		}
	}
	return book.Summary.Series
}

//Label returns string of book label
func (book *Book) Label() string {
	if !book.IsValid() {
		return ""
	}
	if book.Onix.DescriptiveDetail.Collection.TitleDetail != nil {
		for _, elm := range book.Onix.DescriptiveDetail.Collection.TitleDetail.TitleElement {
			if elm.TitleElementLevel == "02" && len(elm.TitleText.Content) > 0 {
				return elm.TitleText.Content
			}
		}
	}
	return ""
}

//ImageURL returns string of book cover image URL
func (book *Book) ImageURL() string {
	if !book.IsValid() {
		return ""
	}
	if book.Onix.CollateralDetail != nil {
		for _, elm := range book.Onix.CollateralDetail.SupportingResource {
			if elm.ResourceContentType == "01" {
				for _, v := range elm.ResourceVersion {
					if v.ResourceForm == "02" && len(v.ResourceLink) > 0 {
						return v.ResourceLink
					}
				}
			}
		}
	}
	return book.Summary.Cover
}

//Authors returns strings of Contributors
func (book *Book) Authors() []string {
	authors := []string{}
	if !book.IsValid() {
		return authors
	}
	for _, elm := range book.Onix.DescriptiveDetail.Contributor {
		if len(elm.PersonName.Content) > 0 {
			authors = append(authors, elm.PersonName.Content)
		}
	}
	if len(authors) > 0 {
		return authors
	}
	return []string{book.Summary.Author}
}

//Publisher returns strings of Publisher
func (book *Book) Publisher() string {
	if !book.IsValid() {
		return ""
	}
	pub := book.Onix.PublishingDetail.Imprint.ImprintName
	if len(pub) == 0 {
		pub = book.Onix.PublishingDetail.Publisher.PublisherName
	}
	if len(pub) == 0 {
		pub = book.Summary.Publisher
	}
	return pub
}

//PublicationDate returns Date of Publication
func (book *Book) PublicationDate() Date {
	if !book.IsValid() {
		return NewDate(time.Time{})
	}
	for _, pubdate := range book.Onix.PublishingDetail.PublishingDate {
		if pubdate.PublishingDateRole == "01" && !pubdate.Date.IsZero() {
			return pubdate.Date
		}
	}
	if !book.Hanmoto.DatePublished.IsZero() {
		return book.Hanmoto.DatePublished
	}
	return book.Summary.PubDate
}

//Description returns strings of book description or table of content
func (book *Book) Description() string {
	if !book.IsValid() {
		return ""
	}
	desc := ""
	if book.Onix.CollateralDetail != nil {
		for _, content := range book.Onix.CollateralDetail.TextContent {
			switch content.TextType {
			case "02", "04": //brief or table of content
				if len(desc) == 0 {
					desc = content.Text
				}
			case "03": //description
				desc = content.Text
			}
		}
	}
	return desc
}

//DecodeBook returns Book instance from byte buffer
func DecodeBook(b []byte) (*Book, error) {
	book := Book{}
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&book); err != nil {
		return &book, errs.Wrap(err, "error in DecodeBook() function")
	}
	return &book, nil
}

//DecodeBooks returns array of Book instance from byte buffer
func DecodeBooks(b []byte) ([]Book, error) {
	var books []Book
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&books); err != nil {
		return books, errs.Wrap(err, "error in DecodeBooks() function")
	}
	return books, nil
}

//EncodeBook returns bytes encoded from Book instance
func EncodeBook(book *Book) ([]byte, error) {
	return json.Marshal(book)
}

//EncodeBooks returns bytes encoded from list of Book
func EncodeBooks(books []Book) ([]byte, error) {
	return json.Marshal(books)
}

/* Copyright 2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
