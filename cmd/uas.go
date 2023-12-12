package cmd

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/kyh0703/sipp-maker/flags"
	"github.com/spf13/cobra"
)

var uasCmd = &cobra.Command{
	Use:   "uas",
	Short: "make uas scenario",
	Run: func(cmd *cobra.Command, args []string) {
		flag := new(flags.Uas)
		flag.SetFlags(cmd.Flags())

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create(fmt.Sprintf("%s/uas.xml", dir))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		tmpl, err := template.ParseFiles(
			"templates/uas/base.tpl",
			"templates/uas/404.tpl",
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
	uasCmd.PersistentFlags().IntP("code", "c", 0, "response sip code then terminate")
	rootCmd.AddCommand(uasCmd)
}
