package main

import (
	"fmt"
	"log"
	"os/exec"
)

func Init(slaveMode bool) {
	if ipServer == "" || flagSSHKey == "" {
		log.Fatal("Error usage: initialisation require <ip> and <ssh key>")
	}
	var step2, step3 string

	step1 := fmt.Sprintf("./init.sh %s %s", ipServer, flagSSHKey)

	if slaveMode {
		step2 = fmt.Sprintf("./data_slave.sh %s", ipServer)
		step3 = fmt.Sprintf("./launch.sh %s scripts/install_slave.sh", ipServer)
	} else {
		step2 = fmt.Sprintf("./data.sh %s", ipServer)
		step3 = fmt.Sprintf("./launch.sh %s scripts/install.sh", ipServer)
	}
	o, e := exec.Command(
		"sh", "-c",
		fmt.Sprintf("%s && %s && %s", step1, step2, step3),
	).Output()

	fmt.Printf("%s", o)
	if e != nil {
		log.Fatal(e)
	}
}

func Deploy(slaveMode bool) {
	if ipServer == "" {
		log.Fatal("Error usage: <ip> is required")
	}
	var step4, step5, step7 string

	step6 := fmt.Sprintf("./upload.sh %s", ipServer)

	if slaveMode {
		step4 = "./build_slave.sh"
		step5 = "./save_slave.sh"
		step7 = fmt.Sprintf("./launch.sh %s scripts/update_slave.sh %s", ipServer, flagMIP)
	} else {
		step4 = "./build.sh"
		step5 = "./save.sh"
		step7 = fmt.Sprintf("./launch.sh %s scripts/update.sh", ipServer)
	}
	o, e := exec.Command(
		"sh", "-c",
		fmt.Sprintf("%s && %s && %s && %s", step4, step5, step6, step7),
	).Output()

	fmt.Printf("%s", o)
	if e != nil {
		log.Fatal(e)
	}
}
