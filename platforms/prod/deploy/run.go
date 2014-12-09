package main

import (
	"fmt"
	"log"
	"os/exec"
)

func Init(slaveMode string) {
	if ipServer == "" || flagSSHKey == "" {
		log.Fatal("Error usage: initialisation require <ip> and <ssh key>")
	}
	var step2, step3 string

	step1 := fmt.Sprintf("./init.sh %s %s", ipServer, flagSSHKey)
	step2 := fmt.Sprintf("./data%s.sh %s", slaveMode, ipServer)
	step3 := fmt.Sprintf("./launch.sh %s scripts/install%s.sh", slaveMode, ipServer)

	o, e := exec.Command(
		"sh", "-c",
		fmt.Sprintf("%s && %s && %s", step1, step2, step3),
	).Output()

	fmt.Printf("%s", o)
	if e != nil {
		log.Fatal(e)
	}
}

func Deploy(slaveMode string) {
	if ipServer == "" {
		log.Fatal("Error usage: <ip> is required")
	}
	var step4, step5, step7 string

	step4 := fmt.Sprintf("./build%s.sh", slaveMode)
	step5 := fmt.Sprintf("./save%s.sh", slaveMode)
	step6 := fmt.Sprintf("./upload.sh %s", ipServer)

	if slaveMode {
		step7 = fmt.Sprintf("./launch.sh %s scripts/update_slave.sh %s", ipServer, flagMIP)
	} else {
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
