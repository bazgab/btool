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
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log/slog"
	"os/exec"
	"os"
	"strings"
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

	
	type Config struct {
		Database struct {
			Engine   string `yaml:"engine"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Host     string `yaml:"host"`
		} `yaml:"database"`
		Dump struct {
			Path         string   `yaml:"path"`
			Type         string   `yaml:"type"`
			DatabaseName []string `yaml:"database_name"`
			Tables       []string `yaml:"tables"`
		} `yaml:"dump"`
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
	
	// Check for mariadb engine usage
	if confValues.Database.Engine != "mariadb" {
		slog.Info("Database engine not compatible with this version of btool");
		os.Exit(1);
	}
	
	uL := confValues.Database.User
	uP := confValues.Database.Password 
	loginDetails := "--user=" + uL + " " + "-p" + uP
	
	if confValues.Dump.Type != "all_databases" {
		if confValues.Dump.Tables[0] != "all_tables" {
			
		// Runs the command with database_name and specific table name
		
		for i := 0; i < len(confValues.Dump.DatabaseName); i++ {
		fmt.Printf("Creating dump for : %s\n", confValues.Dump.DatabaseName[i])
		dumpName := confValues.Dump.DatabaseName[i] + "-dump.sql"
		arg := []string{loginDetails, "--databases", confValues.Dump.DatabaseName[i], confValues.Dump.Tables[i], ">", dumpName }
		fmt.Sprintf("Running command with following argument: %s", strings.Join(arg, " "))
		out, err := exec.Command("usr/bin/mariadb-dump", strings.Join(arg, " ")).Output()
		if err != nil {
			slog.Info("Error when executing dump command")
			os.Exit(1)
		} 
		fmt.Printf("Output: %s", out)
		}
		
		} else {
			
			// Config for running just database_name in command
			for i := 0; i < len(confValues.Dump.DatabaseName); i++ {
			fmt.Printf("Creating dump for : %s\n", confValues.Dump.DatabaseName[i])
			dumpName := confValues.Dump.DatabaseName[i] + "-dump.sql"
			arg := []string{loginDetails, "--databases", confValues.Dump.DatabaseName[i], ">", dumpName }
			fmt.Sprintf("Running command with following argument: %s", strings.Join(arg, " "))
			out, err := exec.Command("usr/bin/mariadb-dump", strings.Join(arg, " ")).Output()
			if err != nil {
			slog.Info("Error when executing dump command")
			os.Exit(1)
			}
			fmt.Printf("Output: %s", out); 
		}
		
		} 
	} else {
		
		// Config for running command with "--all-databases" hardcoded
		for i := 0; i < len(confValues.Dump.DatabaseName); i++ {
		fmt.Printf("Creating dump for : %s\n", confValues.Dump.DatabaseName[i])
		dumpName := confValues.Dump.DatabaseName[i] + "-dump.sql"
		arg := []string{loginDetails ,"--all-databases", ">", dumpName }
		fmt.Sprintf("Running command with following argument: %s", strings.Join(arg, " "))
		out, err := exec.Command("usr/bin/mariadb-dump", strings.Join(arg, " ")).Output()
		if err != nil {
			slog.Info("Error when executing dump command")
			os.Exit(1)
		} 
		fmt.Printf("Output: %s", out)
		}
	}
	
}

