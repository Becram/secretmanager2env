package main

import (
	"flag"
	"fmt"
	"secretmanager2env/aws"
)

func main() {
	secret := flag.String("s", "secret", "Secret To Fetch")
	version := flag.String("v", "version", "Version of secret To Fetch")
	flag.Parse()
	if *secret == "secret" {
		fmt.Println("You must specify a secret name to fetch")
		return
	}
	aws.GetSecret(secret, version)
	// aws.Check_env()
}
