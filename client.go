package openbd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/spiegel-im-spiegel/errs"
)

const (
	APIVersion = "v1"
)

//Client is http.Client for Aozora API Server
type Client struct {
	server *Server
	client *http.Client
}

//LookupBooksRaw gets books data (raw data)
func (c *Client) LookupBooksRaw(ids []string) ([]byte, error) {
	params := url.Values{}
	params.Set("isbn", strings.Join(ids, ","))
	return c.get(c.MakeLookupCommand(params))
}

//LookupBooks gets books data (struct data)
func (c *Client) LookupBooks(ids []string) ([]Book, error) {
	b, err := c.LookupBooksRaw(ids)
	if err != nil {
		return nil, errs.Wrapf(err, "error in Client.LookupBooks() function")
	}
	return DecodeBooks(b)
}

//MakeLookupCommand returns URI for lookup command
func (c *Client) MakeLookupCommand(v url.Values) *url.URL {
	u := c.server.URL()
	u.Path = fmt.Sprintf("/%v/%v", APIVersion, "get")
	u.RawQuery = v.Encode()
	return u
}

func (c *Client) get(u *url.URL) ([]byte, error) {
	resp, err := c.client.Get(u.String())
	if err != nil {
		return nil, errs.Wrapf(err, "error in Client.get(%v) function", u)
	}
	defer resp.Body.Close()

	if !(resp.StatusCode != 0 && resp.StatusCode < http.StatusBadRequest) {
		return nil, errs.Wrapf(ErrHTTPStatus, "%v (in %v)", resp.Status, u)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, errs.Wrapf(err, "error in Client.get(%v) function", u)
	}
	return body, nil
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
