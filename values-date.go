package openbd

import (
	"strconv"
	"strings"
	"time"

	"github.com/goark/errs"
)

//Time is wrapper class of time.Time
type Date struct {
	time.Time
}

//NewDate returns Time instance
func NewDate(tm time.Time) Date {
	return Date{tm}
}

var timeTemplate1 = []string{
	"2006-01-02 15:04:05",
	time.RFC3339,
}
var timeTemplate2 = []string{
	"2006-01-02",
	"2006-01",
}
var timeTemplate3 = []string{
	"20060102",
	"200601",
	"2006",
}

//UnmarshalJSON returns result of Unmarshal for json.Unmarshal()
func (t *Date) UnmarshalJSON(b []byte) error {
	s := string(b)
	if ss, err := strconv.Unquote(s); err == nil {
		s = ss
	}
	if len(s) == 0 || strings.ToLower(s) == "null" {
		*t = Date{time.Time{}}
		return nil
	}
	var lastErr error
	if strings.Contains(s, "-") {
		if strings.Contains(s, ":") {
			for _, tmplt := range timeTemplate1 {
				if tm, err := time.Parse(tmplt, s); err != nil {
					lastErr = errs.Wrap(err, errs.WithContext("time_string", s), errs.WithContext("time_template", tmplt))
				} else {
					*t = Date{tm}
					return nil
				}
			}
			return lastErr
		}
		for _, tmplt := range timeTemplate2 {
			if tm, err := time.Parse(tmplt, s); err != nil {
				lastErr = errs.Wrap(err, errs.WithContext("time_string", s), errs.WithContext("time_template", tmplt))
			} else {
				*t = Date{tm}
				return nil
			}
		}
		return lastErr
	}
	for _, tmplt := range timeTemplate3 {
		if tm, err := time.Parse(tmplt, s); err != nil {
			lastErr = errs.Wrap(err, errs.WithContext("time_string", s), errs.WithContext("time_template", tmplt))
		} else {
			*t = Date{tm}
			return nil
		}
	}
	return lastErr
}

//MarshalJSON returns time string with RFC3339 format
func (t *Date) MarshalJSON() ([]byte, error) {
	if t == nil || t.IsZero() {
		return []byte("\"\""), nil
	}
	return []byte(strconv.Quote(t.String())), nil
}

func (t Date) String() string {
	return t.Format("2006-01-02")
}

/* Copyright 2019,2020 Spiegel
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
