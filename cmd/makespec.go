package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/mail"
	"os"
	"path/filepath"

	"github.com/ljosa/broadcast/lib"
	"github.com/spf13/cobra"
)

var htmlFilename string
var textFilename string
var recipientsFilename string
var force bool
var senderName string
var senderEmail string
var subject string

// makespecCmd represents the makespec command
var makespecCmd = &cobra.Command{
	Use:   "makespec",
	Short: "Create a spec file",
	Long: `Creates a spec.json file in the current working directory

The spec file includes the list of recipients, the HTML and plain text version
of the contents, as well as other settings such as the sender's name and address.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if htmlFilename == "" && textFilename == "" {
			return errors.New("makespec requires at least one of --html and --text")
		}
		if senderEmail == "" {
			fmt.Println("Warning: you must edit the spec to set the sender's email address")
		}
		if recipientsFilename == "" {
			fmt.Println("Warning: no recipients")
		}
		var err error
		spec := lib.Spec{}
		spec.FromAddr = senderEmail
		if senderName != "" {
			spec.FromName = senderName
		}
		spec.Subject = subject
		if htmlFilename != "" {
			html, err := ioutil.ReadFile(htmlFilename)
			if err != nil {
				return err
			}
			spec.Html = string(html)
		}
		if textFilename != "" {
			text, err := ioutil.ReadFile(textFilename)
			if err != nil {
				return err
			}
			spec.Text = string(text)
		}
		spec.Recipients, err = readRecipients()
		if err != nil {
			return err
		}

		err = writeSpec(spec)
		if err != nil {
			return err
		}

		return nil
	},
}

func readRecipients() ([]lib.Recipient, error) {
	file, err := os.Open(recipientsFilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	recipients := make([]lib.Recipient, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			addr, err := mail.ParseAddress(line)
			if err != nil {
				return nil, err
			}
			recipients = append(recipients, lib.Recipient{Name: addr.Name, Addr: addr.Address})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return recipients, nil
}

func writeSpec(spec lib.Spec) error {
	data, err := json.MarshalIndent(&spec, "", "\t")
	f, err := ioutil.TempFile(filepath.Dir(specFilename), specFilename)
	if err != nil {
		return err
	}
	defer os.Remove(f.Name())
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	fn := f.Name()
	err = f.Close()
	if err != nil {
		return err
	}
	if !force {
		if _, err := os.Stat(specFilename); err == nil {
			return errors.New("spec file already exists (use -f to overwrite): " + specFilename)
		}
	}
	err = os.Rename(fn, specFilename)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	RootCmd.AddCommand(makespecCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	makespecCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "Overwrite spec file if it exists")
	makespecCmd.PersistentFlags().StringVarP(&htmlFilename, "html", "", "", "File name of HTML contents")
	makespecCmd.PersistentFlags().StringVarP(&recipientsFilename, "recipients", "", "", "Name of file containing recipients, one per line")
	makespecCmd.PersistentFlags().StringVarP(&senderEmail, "sender", "", "", "Sender's email address")
	makespecCmd.PersistentFlags().StringVarP(&senderName, "sender-name", "", "", "Sender's email address")
	makespecCmd.PersistentFlags().StringVarP(&specFilename, "spec", "", "spec.json", "File name of spec to write (use --force to overwrite)")
	makespecCmd.PersistentFlags().StringVarP(&subject, "subject", "", "", "Subject of the email")
	makespecCmd.PersistentFlags().StringVarP(&textFilename, "text", "", "", "File name of HTML contents")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makespecCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
