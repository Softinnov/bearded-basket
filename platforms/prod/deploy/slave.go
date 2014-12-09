package main

import "github.com/spf13/cobra"

var cmdSlave = &cobra.Command{
	Use:   "slave",
	Short: "deploy on the slave (sql) side",
	Long:  `deploy on the slave (sql) side`,
	Run:   runSlave,
}

func init() {
	cmdSlave.Flags().StringVar(&flagSIP, "ip", "", "server ip address")
	cmdSlave.Flags().StringVar(&flagMIP, "master", "", "master ip address")
}

func runSlave(cmd *cobra.Command, args []string) {
	ipServer = flagMIP
	if flagInit {
		Init("_slave")
	}
	Deploy("_slave")
}
