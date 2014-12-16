package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var cmdMaster = &cobra.Command{
	Use:   "master",
	Short: "deploy in the master (sql) side",
	Long:  `deploy in the master (sql) side`,
	Run:   runMaster,
}

func init() {
	cmdMaster.Flags().StringVar(&flagMIP, "ip", "", "server ip address")
}

func runMaster(cmd *cobra.Command, args []string) {
	ipServer = flagMIP
	if flagInit {
		if ipServer == "" || flagSSHKey == "" {
			fmt.Println(ipServer, flagSSHKey)
			log.Fatal("Error usage: initialisation require <ip> and <ssh key>")
		}
	}
	if ipServer == "" {
		log.Fatal("Error usage: <ip> is required")
	}
	if flagInit {
		Init("")
	}
	Deploy("")
}
