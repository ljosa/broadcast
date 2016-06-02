// Copyright © 2016 Vebjorn Ljosa <vebjorn@ljosa.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"github.com/ljosa/broadcast/lib"
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Send the email to yourself",
	Long:  `Send an email to a single address`,
	RunE: func(cmd *cobra.Command, args []string) error {
		spec, err := lib.LoadSpec(specFilename)
		if err != nil {
			return err
		}
		if len(args) == 0 {
			return errors.New("No test recipients specified")
		}
		recipients := make([]lib.Recipient, 0)
		for i := 0; i < len(args); i++ {
			recipients = append(recipients, lib.Recipient{Addr: args[i]})
		}
		spec.Recipients = recipients
		return lib.Send(spec, smtpServer, smtpPort)
	},
}

func init() {
	RootCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(&specFilename, "spec", "", "spec.json", "File name of spec to read")
	testCmd.Flags().StringVarP(&smtpServer, "smtp-server", "", "localhost", "SMTP server hostname")
	testCmd.Flags().IntVarP(&smtpPort, "smtp-port", "", 25, "SMTP port")
}
