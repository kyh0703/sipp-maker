package cmd

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/kyh0703/sipp-maker/flags"
	"github.com/spf13/cobra"
)

var uacCmd = &cobra.Command{
	Use:   "uac",
	Short: "make uac scenario",
	Run: func(cmd *cobra.Command, args []string) {
		flag := new(flags.Uac)
		flag.SetFlags(cmd.Flags())

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create(fmt.Sprintf("%s/uac.xml", dir))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		tmpl, err := template.ParseFiles(
			"templates/uac/base.tpl",
			"templates/uac/invite_sip.tpl",
			"templates/uac/invite_tel.tpl",
			"templates/uac/sdp.tpl",
			"templates/uac/ack_sip.tpl",
			"templates/uac/ack_tel.tpl",
		)
		if err != nil {
			log.Fatal(err)
		}

		if err := tmpl.Execute(file, &flag); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	uacCmd.PersistentFlags().StringP("uri", "u", "sip", "invite mode 'sip' or 'tel'")
	uacCmd.PersistentFlags().BoolP("sdp", "s", true, "invite sdp 'true' or 'false'")
	rootCmd.AddCommand(uacCmd)
}
