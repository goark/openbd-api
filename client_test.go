package openbd

import (
	"context"
	"net/http"
	"net/url"
	"testing"
)

func TestmakeLookupCommand(t *testing.T) {
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
		u := DefaultClient().makeLookupCommand(tc.v)
		if u.String() != tc.str {
			t.Errorf("Client.makeLookupCommand() is \"%v\", want \"%v\"", u.String(), tc.str)
		}
		u = (*Server)(nil).CreateClient(WithContext(context.Background()), WithHttpClient(&http.Client{})).makeLookupCommand(tc.v)
		if u.String() != tc.str {
			t.Errorf("Client.makeLookupCommand() is \"%v\", want \"%v\"", u.String(), tc.str)
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
