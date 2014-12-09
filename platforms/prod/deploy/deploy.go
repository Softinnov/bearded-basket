package main

import "github.com/spf13/cobra"

var (
	flagInit   bool
	flagSIP    string
	flagMIP    string
	flagSSHKey string

	ipServer string
)

var rootCmd = &cobra.Command{
	Use:  "deploy",
	Long: `deploy [COMMAND] [--init]`,
}

func init() {
	rootCmd.AddCommand(cmdMaster, cmdSlave)
	rootCmd.PersistentFlags().BoolVar(&flagInit, "init", false, "server initialisation")
	rootCmd.PersistentFlags().StringVar(&flagSSHKey, "key", "", "ssh key (.pub)")
}

func main() {
	rootCmd.Execute()
}
