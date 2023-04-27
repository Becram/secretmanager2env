package main

import (
	"flag"
	"fmt"
	"secretmanager2env/aws"
)

func main() {
	secret := flag.String("s", "secret", "Secret To Fetch")
	output := flag.String("o", "env", "Output type")
	filename := flag.String("n", "changeme", "Output type")
	version := flag.String("v", "version", "Version of secret To Fetch")
	flag.Parse()
	if *secret == "secret" {
		fmt.Println("You must specify a secret name to fetch")
		return
	}
	if *output == "json" && *filename == "changeme" {
		fmt.Println("Pass json output filename")
		return
	}

	aws.GetSecret(secret, version, output, filename)
	// aws.Check_env()
}
