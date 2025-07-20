/*
Copyright 2025 The btool Authors All rights reserved

DISCLAIMER:

btool is free software; you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free
Software Foundation

btool is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or
FITNESS FOR A PARTICULAR PURPOSE.

The full license can be seen in the file ./LICENSE.  If not see
<http://www.gnu.org/licenses/>.
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "btool",
	Short:   "A general-purpose backup tool",
	Version: "0.1 (Pre-Alpha)",
	Long: `A general-purpose backup tool aiming to facilitate basic aspects of creating, managing and deleting backups and dumps.

For Documentation, as well as a better understanding of the architecture, see the project on github: https://github.com/bazgab/btool
`,
}


// Execute - This call needs to only happen once
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.CompletionOptions.DisableDefaultCmd = true
}


// This provides a custom template for the help page, or when running just the 'btool' command
func CustomHelpTemplate() string {
	
	return `{{.UseLine}}

{{.Short}}

Usage:
  {{.UseLine}}{{if .HasAvailableSubCommands}}

Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}

{{if .HasAvailableLocalFlags}}Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}

{{if .HasAvailablePersistentFlags}}Global Flags:
{{.PersistentFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}

{{if .HasAvailableSubCommands}}Flags by Command:
{{range .Commands}}{{if .HasAvailableFlags}}
  {{.Name}}:
{{.LocalFlags.FlagUsages | indent 4}}{{end}}{{end}}{{end}}

Use "{{.CommandPath}} [command] --help" for more information about a command.`

}
