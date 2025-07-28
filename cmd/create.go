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
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"log/slog"
	"gopkg.in/yaml.v3"
)

// The following values will be used in case the user does not provide a value for these parameters
// Notice: the fields 'password' and 'path' do not have default values, making them required
const (
	
		defaultUser = "root"
		defaultEngine = "mariadb"
		defaultHost = "localhost"
		defaultType = "all_databases"
		defaultTables = "all_tables"
	)
	
	
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a dump",
	Long: `Creates a dump from a config file. 

All values can be assigned directly into the config file. For templates and a more detailed reference spec sheet, see: 
https://github.com/bazgab/btool/etc/etc`,
	Run: runCreate,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("config", "C", "", "config file with dump parameters (required)")
}

func runCreate(cmd *cobra.Command, _ []string) {
	c, err := cmd.Flags().GetString("config")
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("The following file has been selected for config values: " + c)
	
	// Testing declaring nested nodes
	/*
	type Database struct {
		User     string `yaml:"user"`
		Password []string `yaml:"password"`
	}

	type Dump struct {
		Path string `yaml:"path"`
		Type string `yaml:"type"`
	}
	*/
	type Config struct {
		Database struct {
			User     string `yaml:"user"`
			Password string `yaml:"password"`
		}`yaml:"database"`
		Dump struct {
			Path string `yaml:"path"`
			Type string `yaml:"type"`
		}`yaml:"dump"`
		
	}
	
	confFile, err := os.ReadFile(c)
	if err != nil {
		fmt.Println(err.Error())
	}

	
	replaced := os.ExpandEnv(string(confFile))
	
	var confValues Config
	err = yaml.Unmarshal([]byte(replaced), &confValues)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	
	/*
	for i := 0; i < len(confValues.Password); i++ {
		fmt.Printf("Value %d : %s\n", i, confValues.Password[i])
		
	}
	*/
	
	// Now testing if we can parse an env variable
	//fmt.Println("User: ", confValues.User)
	
	//Check user
	if confValues.Database.User  == "" {
		fmt.Println("User check - No selected user. Setting option to default value: ", defaultUser)
	} else {
		fmt.Printf("User value: %s\n", confValues.Database.User)
	}

	
}
