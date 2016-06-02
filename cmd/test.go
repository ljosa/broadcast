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
