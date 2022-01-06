/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/temporalio/background-checks/api"
	"github.com/temporalio/background-checks/temporal"
	"go.temporal.io/sdk/client"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run API Server",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := temporal.NewClient(client.Options{})
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

		srv := &http.Server{
			Handler: api.Router(c),
			Addr:    "0.0.0.0:8081",
		}

		log.Fatal(srv.ListenAndServe())
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
