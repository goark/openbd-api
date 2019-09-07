package openbd

import (
	"net/http"
	"net/url"
	"testing"
)

// func TestLookupBooksRaw(t *testing.T) {
// 	testCases := []struct {
// 		ids []string
// 	}{
// 		{ids: nil},
// 		{ids: []string{}},
// 		{ids: []string{"foo"}},
// 		{ids: []string{"9784535585744"}},
// 		{ids: []string{"9784535585744", "9784797369915"}},
// 		{ids: []string{"9784535585744", "foo"}},
// 		{ids: []string{"9784535585744&isbn=9784797369915"}},
// 	}
//
// 	for _, tc := range testCases {
// 		b, err := DefaultClient().LookupBooksRaw(tc.ids)
// 		if err != nil {
// 			t.Errorf("Client.LookupBooksRaw() is \"%v\", want nil", err)
// 			fmt.Printf("error info: %+v\n", err)
// 			continue
// 		}
// 		fmt.Printf("response: %+v\n", string(b))
// 	}
// }
//
// func TestLookupBooks(t *testing.T) {
// 	testCases := []struct {
// 		ids   []string
// 		valid []bool
// 	}{
// 		{ids: nil, valid: []bool{false}},
// 		{ids: []string{}, valid: []bool{false}},
// 		{ids: []string{"foo"}, valid: []bool{false}},
// 		{ids: []string{"9784535585744"}, valid: []bool{true}},
// 		{ids: []string{"9784535585744", "9784797369915"}, valid: []bool{true, true}},
// 		{ids: []string{"9784535585744", "foo"}, valid: []bool{true, false}},
// 		{ids: []string{"9784535585744&isbn=9784797369915"}, valid: []bool{false}},
// 	}
//
// 	for _, tc := range testCases {
// 		bks, err := DefaultClient().LookupBooks(tc.ids)
// 		if err != nil {
// 			t.Errorf("Client.LookupBooks() is \"%v\", want nil", err)
// 			fmt.Printf("error info: %+v\n", err)
// 			continue
// 		}
// 		if len(bks) != len(tc.valid) {
// 			t.Errorf("Count of Client.LookupBooks() is %v, want %v", len(bks), len(tc.valid))
// 			continue
// 		}
// 		for i, bk := range bks {
// 			if bk.Valid() != tc.valid[i] {
// 				t.Errorf("Book[%d] is %v, want %v", i, bk.Valid(), tc.valid[i])
// 				continue
// 			}
// 			if bk.Valid() {
// 				id := bk.Id()
// 				if id != tc.ids[i] {
// 					t.Errorf("Book[%d] is %v, want %v", i, id, tc.ids[i])
// 				}
// 			}
// 		}
// 	}
// }

func TestMakeLookupCommand(t *testing.T) {
	testCases := []struct {
		v   url.Values
		str string
	}{
		{v: nil, str: "https://api.openbd.jp/v1/get"},
		{v: url.Values{}, str: "https://api.openbd.jp/v1/get"},
		{v: url.Values{"isbn": {"foo"}}, str: "https://api.openbd.jp/v1/get?isbn=foo"},
		{v: url.Values{"isbn": {"foo,bar"}}, str: "https://api.openbd.jp/v1/get?isbn=foo%2Cbar"},
		{v: url.Values{"isbn": {"foo&isbn=bar"}}, str: "https://api.openbd.jp/v1/get?isbn=foo%26isbn%3Dbar"},
	}

	for _, tc := range testCases {
		u := DefaultClient().MakeLookupCommand(tc.v)
		if u.String() != tc.str {
			t.Errorf("Client.MakeLookupCommand() is \"%v\", want \"%v\"", u.String(), tc.str)
		}
		u = (*Server)(nil).CreateClient(nil, &http.Client{}).MakeLookupCommand(tc.v)
		if u.String() != tc.str {
			t.Errorf("Client.MakeLookupCommand() is \"%v\", want \"%v\"", u.String(), tc.str)
		}
	}
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
