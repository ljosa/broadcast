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
