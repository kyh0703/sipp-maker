package cmd

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/kyh0703/sipp-maker/flags"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "make register scenario",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		flag := new(flags.Uac)
		flag.SetFlags(cmd.Flags())

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create(fmt.Sprintf("%s/register.xml", dir))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		tmpl, err := template.ParseFiles(
			"templates/register/base.tpl",
			"templates/register/register_sip.tpl",
			"templates/register/register_tel.tpl",
			"templates/register/auth_sip.tpl",
			"templates/register/auth_tel.tpl",
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
	rootCmd.AddCommand(registerCmd)
}
