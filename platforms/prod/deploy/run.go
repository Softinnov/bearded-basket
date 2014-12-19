package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Init(slaveMode string) {
	step1 := fmt.Sprintf("%s/init.sh %s %s", flagDir, ipServer, flagSSHKey)
	step2 := fmt.Sprintf("%s/data.sh %s", flagDir, ipServer)
	step3 := fmt.Sprintf("%s/launch.sh %s %[1]s/install.sh", flagDir, ipServer)

	fmt.Println(fmt.Sprintf("%s && %s && %s", step1, step2, step3))
	c := exec.Command(
		"sh", "-c",
		fmt.Sprintf("%s && %s && %s", step1, step2, step3),
	)
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr

	e := c.Run()
	if e != nil {
		log.Fatal(e)
	}
}

func Deploy(slaveMode string) {
	var step7 string

	step4 := fmt.Sprintf("%s/build%s.sh", flagDir, slaveMode)
	step5 := fmt.Sprintf("%s/save%s.sh", flagDir, slaveMode)
	step6 := fmt.Sprintf("%s/upload.sh %s", flagDir, ipServer)

	if slaveMode != "" {
		step7 = fmt.Sprintf("%s/launch.sh %s %[1]s/update_slave.sh %[3]s", flagDir, ipServer, flagMIP)
	} else {
		step7 = fmt.Sprintf("%s/launch.sh %s %[1]s/update.sh", flagDir, ipServer)
	}
	c := exec.Command(
		"sh", "-c",
		fmt.Sprintf("%s && %s && %s && %s", step4, step5, step6, step7),
	)
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr

	e := c.Run()
	if e != nil {
		log.Fatal(e)
	}
}
