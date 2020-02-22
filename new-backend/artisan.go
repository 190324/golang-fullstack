package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	switch os.Args[1] {
	case "gql":
		commandGql()
	case "migrate":
		fmt.Println("migrate")
	default:
		panic(fmt.Errorf("%s command not found", os.Args[1]))
	}
}

func commandGql() {
	var configPath string
	switch os.Args[2] {
	case "api", "admin":
		configPath = "config/gqlgen/" + os.Args[2] + ".yml"
	default:
		panic(fmt.Errorf("%s %s command not found", os.Args[1], os.Args[2]))
	}

	fmt.Println("config: ", configPath)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("go", "run", "scripts/gqlgen.go", "-c", configPath)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(stdout.String(), stderr.String())
		fmt.Fprintln(os.Stderr, err.Error())
	} else {
		fmt.Println("success!!")
	}

}
