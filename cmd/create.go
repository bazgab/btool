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
	
	"github.com/spf13/cobra"
	"log/slog"
	"gopkg.in/yaml.v3"
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
	c, err := cmd.Flags().GetString("conf")
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("The following file has been selected for config values: " + c)
	
	
	// Default values in case the user does not provide these values
	// Notice: attributes password and path do not have a default value, as it is critical to provide this information
	const (
	
		defaultUser = "root"
		defaultEngine = "mariadb"
		defaultHost = "localhost"
		defaultType = "all_databases"
		defaultTables = "all_tables"
	)	
	
	type TestYAML struct {
		User     string `yaml:"user"`
		Password []string `yaml:"password"`
	}


	var yamlValues TestYAML

	yamlFile, err := os.ReadFile("test.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	
	replaced := os.ExpandEnv(string(yamlFile))

	err = yaml.Unmarshal([]byte(replaced), &yamlValues)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Reading from YAML FILE:")
	
	// Checking if we can parse multiple values 
	fmt.Printf("Password field has %d values\n", len(yamlValues.Password))
	
	for i := 0; i < len(yamlValues.Password); i++ {
		fmt.Printf("Value %d : %s\n", i, yamlValues.Password[i])
		
	}


	// Now testing if we can parse an env variable
	fmt.Println("User: ", yamlValues.User)
	fmt.Println("Checking if user is the same as /home/bazgab...")
	
	//Check user
	if yamlValues.User  == "" {
		fmt.Println("User check - No selected user. Setting option to default value: ", defaultUser)
	} else if yamlValues.User == "/home/gabbaz"{
		fmt.Println("Warning - user is /home/gabbaz")
	} else {
		fmt.Println("Error - invalid value")
	}

	
}
