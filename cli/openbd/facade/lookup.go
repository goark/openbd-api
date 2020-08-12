package facade

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/openbd-api"
)

//newLookupBookCmd returns cobra.Command instance for show sub-command
func newLookupCmd(ui *rwi.RWI) *cobra.Command {
	lookupCmd := &cobra.Command{
		Use:   "lookup [flags] <book id>",
		Short: "Lookup book data by openBD API",
		Long:  "Lookup book data by openBD API",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errs.New("book id", errs.WithCause(os.ErrInvalid))
			}

			if rawFlag {
				resp, err := openbd.DefaultClient().LookupBooksRaw(args)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			}

			bks, err := openbd.DefaultClient().LookupBooks(args)
			if err != nil {
				return debugPrint(ui, err)
			}
			b, err := openbd.EncodeBooks(bks)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.OutputBytes(b))
		},
	}
	//options

	return lookupCmd
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
