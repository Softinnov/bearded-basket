package main

import "github.com/spf13/cobra"

var cmdMaster = &cobra.Command{
	Use:   "master",
	Short: "deploy in the master (sql) side",
	Long:  `deploy in the master (sql) side`,
	Run:   runMaster,
}

func init() {
	cmdMaster.PersistentFlags().StringVar(&flagMIP, "ip", "", "server ip address")
}

func runMaster(cmd *cobra.Command, args []string) {
	ipServer = flagMIP
	if flagInit {
		Init(false)
	}
	Deploy(false)
}
