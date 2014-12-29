package main

import "github.com/spf13/cobra"

var (
	flagInit   bool
	flagSIP    string
	flagMIP    string
	flagSSHKey string
	flagDir    string

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
	rootCmd.PersistentFlags().StringVar(&flagDir, "dir", ".", "scripts directory")
}

func main() {
	rootCmd.Execute()
}
