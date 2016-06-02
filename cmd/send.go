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
	"github.com/ljosa/broadcast/lib"
	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send the email to all recipients",
	Long:  `Send the email to all recipients in the spec`,
	RunE: func(cmd *cobra.Command, args []string) error {
		spec, err := lib.LoadSpec(specFilename)
		if err != nil {
			return err
		}
		return lib.Send(spec, smtpServer, smtpPort)
	},
}

func init() {
	RootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVarP(&specFilename, "spec", "", "spec.json", "File name of spec to write (use --force to overwrite)")
	sendCmd.Flags().StringVarP(&smtpServer, "smtp-server", "", "localhost", "SMTP server hostname")
	sendCmd.Flags().IntVarP(&smtpPort, "smtp-port", "", 25, "SMTP port")
}
