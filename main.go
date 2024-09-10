package main

import "github.com/beeploop/aes-encrypt/frontend/cli"

const (
	key = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

func main() {
	frontend := cli.NewCLI(key)
	frontend.Start()
}
