package openbd

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/fetch"
)

const (
	defaultAPIVersion = "v1"
)

//Client is http.Client for Aozora API Server
type Client struct {
	server *Server
	client fetch.Client
}

//LookupBooksRaw gets books data (raw data)
func (c *Client) LookupBooksRaw(ids []string) ([]byte, error) {
	return c.LookupBooksRawContext(context.Background(), ids)
}

//LookupBooks gets books data (struct data)
func (c *Client) LookupBooks(ids []string) ([]Book, error) {
	return c.LookupBooksContext(context.Background(), ids)
}

//LookupBooksRawContext gets books data with context.Context. (raw data)
func (c *Client) LookupBooksRawContext(ctx context.Context, ids []string) ([]byte, error) {
	params := url.Values{}
	params.Set("isbn", strings.Join(ids, ","))
	b, err := c.get(ctx, c.makeLookupCommand(params))
	return b, errs.Wrap(err)
}

//LookupBooksRawContext gets books data with context.Context. (struct data)
func (c *Client) LookupBooksContext(ctx context.Context, ids []string) ([]Book, error) {
	b, err := c.LookupBooksRawContext(ctx, ids)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	books, err := DecodeBooks(b)
	return books, errs.Wrap(err)
}

func (c *Client) makeLookupCommand(v url.Values) *url.URL {
	if c == nil {
		return nil
	}
	u := c.server.URL()
	u.Path = fmt.Sprintf("/%v/%v", c.apiDir(), "get")
	u.RawQuery = v.Encode()
	return u
}

func (c *Client) apiDir() string {
	return defaultAPIVersion
}

func (c *Client) get(ctx context.Context, u *url.URL) ([]byte, error) {
	if c == nil {
		return nil, errs.Wrap(fetch.ErrNullPointer)
	}
	resp, err := c.client.Get(u, fetch.WithContext(ctx))
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", u.String()))
	}
	return resp.DumpBodyAndClose()
}

/* Copyright 2019-2021 Spiegel
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
