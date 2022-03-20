package facade

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"

	"github.com/goark/errs"
	"github.com/goark/gocli/exitcode"
	"github.com/goark/gocli/rwi"
	"github.com/goark/openbd-api/cli/openbd/ecode"
	"github.com/spf13/cobra"
)

var (
	//Name is applicatin name
	Name = "openbd"
)
var (
	debugFlag bool //debug flag
	rawFlag   bool //raw flag
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   Name,
		Short: "Search for books data",
		Long:  "Search for books data from openBD API",
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugPrint(ui, ecode.ErrNoCommand)
		},
	}
	rootCmd.SilenceUsage = true
	rootCmd.SetArgs(args)               //arguments of command-line
	rootCmd.SetIn(ui.Reader())          //Stdin
	rootCmd.SetOutput(ui.ErrorWriter()) //Stdout and Stderr
	rootCmd.AddCommand(newLookupCmd(ui))

	//global options (others)
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "for debug")
	rootCmd.PersistentFlags().BoolVarP(&rawFlag, "raw", "", false, "Output raw data from API")

	return rootCmd
}

func debugPrint(ui *rwi.RWI, err error) error {
	if debugFlag && err != nil {
		fmt.Fprintf(ui.ErrorWriter(), "Error: %+v\n", err)
		return nil
	}
	return errs.Cause(err)
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			_ = ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := newRootCmd(ui, args).ExecuteContext(ctx); err != nil {
		exit = exitcode.Abnormal
	}
	return
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
