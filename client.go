package openbd

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/spiegel-im-spiegel/errs"
)

const (
	defaultAPIVersion = "v1"
)

//Client is http.Client for Aozora API Server
type Client struct {
	server *Server
	client *http.Client
	ctx    context.Context
}

//LookupBooksRaw gets books data (raw data)
func (c *Client) LookupBooksRaw(ids []string) ([]byte, error) {
	params := url.Values{}
	params.Set("isbn", strings.Join(ids, ","))
	b, err := c.get(c.makeLookupCommand(params))
	return b, errs.Wrap(err)
}

//LookupBooks gets books data (struct data)
func (c *Client) LookupBooks(ids []string) ([]Book, error) {
	b, err := c.LookupBooksRaw(ids)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	books, err := DecodeBooks(b)
	return books, errs.Wrap(err)
}

func (c *Client) makeLookupCommand(v url.Values) *url.URL {
	u := c.server.URL()
	u.Path = fmt.Sprintf("/%v/%v", c.apiDir(), "get")
	u.RawQuery = v.Encode()
	return u
}

func (c *Client) apiDir() string {
	return defaultAPIVersion
}

func (c *Client) get(u *url.URL) ([]byte, error) {
	req, err := http.NewRequestWithContext(c.ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", u.String()))
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", u.String()))
	}
	defer resp.Body.Close()

	if !(resp.StatusCode != 0 && resp.StatusCode < http.StatusBadRequest) {
		return nil, errs.Wrap(ErrHTTPStatus, errs.WithContext("url", u.String()), errs.WithContext("status", resp.Status))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return body, errs.Wrap(err, errs.WithContext("url", u.String()))
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
