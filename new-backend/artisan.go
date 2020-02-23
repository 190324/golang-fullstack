package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

var (
	db = []string{"solomo", "member"}
)

func main() {
	switch os.Args[1] {
	case "gql":
		commandGql()
	case "migrate":
		commandMigrate()
	case "migration":
		commandMigration()
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

	runCommand("go", "run", "scripts/gqlgen.go", "-c", configPath)
}

func commandMigration() {
	if len(os.Args) < 4 || stringInSlice(os.Args[3], db) == -1 {
		fmt.Printf("which database: %v \n", db)
		os.Exit(1)
	}

	migrationPath := "migrations/" + os.Args[3]

	runCommand("soda", "g", "sql", os.Args[2], "-p", migrationPath, "-e", os.Args[3])
}

func commandMigrate() {

	postion := stringInSlice("-e", os.Args)
	if postion == -1 {
		fmt.Printf("which database: %v \n", db)
		os.Exit(1)
	}

	if stringInSlice(os.Args[postion+1], db) == -1 {
		fmt.Printf("which database: %v \n", db)
		os.Exit(1)
	}

	_db := os.Args[postion+1]
	migrationPath := "migrations/" + _db

	if len(os.Args) == 2 || os.Args[2] == "up" {
		runCommand("soda", "migrate", "-p", migrationPath, "-e", _db)
	} else if os.Args[2] == "down" {
		runCommand("soda", "migrate", "down", "-p", migrationPath, "-e", _db)
	}
}

func runCommand(c ...string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(c[0], c[1:]...)
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

// in_array
func stringInSlice(a string, list []string) int {
	for i, b := range list {
		if b == a {
			return i
		}
	}
	return -1
}
