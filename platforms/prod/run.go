package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/pipe.v2"
)

var (
	flagInit   bool
	flagSIP    string
	flagMIP    string
	flagSSHKey string
)

func init() {
	flag.BoolVar(&flagInit, "init", false, "server initialisation")
	flag.StringVar(&flagSIP, "slave-ip", "", "slave ip address (slave configuration)")

	flag.StringVar(&flagMIP, "master-ip", "", "server ip address (master configuration)")
	flag.StringVar(&flagSSHKey, "key", "", "ssh key (.pub)")
}

func main() {
	var slaveMode bool = false
	var ipServer string
	flag.Parse()

	if flagSIP != "" {
		ipServer = flagSIP
		slaveMode = true
	} else {
		ipServer = flagMIP
	}
	if flagInit && ipServer != "" && flagSSHKey != "" {
		var step2 string
		var step3 string

		step1 := fmt.Sprintf("./scripts/init.sh %s %s", ipServer, flagSSHKey)

		if slaveMode {
			step2 = fmt.Sprintf("./scripts/data_slave.sh %s", ipServer)
			step3 = fmt.Sprintf("./scripts/launch.sh %s scripts/install_slave.sh", ipServer)
		} else {
			step2 = fmt.Sprintf("./scripts/data.sh %s", ipServer)
			step3 = fmt.Sprintf("./scripts/launch.sh %s scripts/install.sh", ipServer)
		}
		e := pipe.Run(pipe.Line(pipe.Script(
			pipe.System(step1),
			pipe.System(step2),
			pipe.System(step3),
		), pipe.Write(os.Stdout)))
		if e != nil {
			log.Fatal(e)
		}
	} else {
		fmt.Println("Usage: initialisation require <ip> and <ssh key>")
		return
	}

	if ipServer != "" {
		var step4 string
		var step5 string
		var step7 string

		step6 := fmt.Sprintf("./scripts/upload.sh %s", ipServer)

		if slaveMode {
			step4 = "./scripts/build_slave.sh"
			step5 = "./scripts/save_slave.sh"
			step7 = fmt.Sprintf("./scripts/launch.sh %s scripts/update_slave.sh %s", ipServer, flagMIP)
		} else {
			step4 = "./scripts/build.sh"
			step5 = "./scripts/save.sh"
			step7 = fmt.Sprintf("./scripts/launch.sh %s scripts/update.sh", ipServer)
		}
		e := pipe.Run(pipe.Line(pipe.Script(
			pipe.System(step4),
			pipe.System(step5),
			pipe.System(step6),
			pipe.System(step7),
		), pipe.Write(os.Stdout)))
		if e != nil {
			log.Fatal(e)
		}
	} else {
		fmt.Println("Usage: <ip> is required")
		return
	}
}
