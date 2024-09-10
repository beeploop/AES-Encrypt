package cli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/beeploop/aes-encrypt/encrypt"
)

type CLI struct {
	key string
}

func NewCLI(key string) *CLI {
	return &CLI{
		key: key,
	}
}

func (c *CLI) Start() {
	file := flag.String("file", "", "Input file")
	mode := flag.String("mode", "encrypt", "operation mode encrypt/decrypt")
	output := flag.String("output", "output", "Output file")

	flag.Parse()
	if *file == "" {
		fmt.Println("Invalid usage")
		flag.PrintDefaults()
		os.Exit(1)
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, *file)
	source, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	crypto, err := encrypt.New(c.key)
	if err != nil {
		panic(err)
	}

	var outputSource []byte
	switch *mode {
	case "encrypt":
		if encrypted, err := crypto.Encrypt(source); err != nil {
			panic(err)
		} else {
			outputSource = encrypted
		}
	case "decrypt":
		if decrypted, err := crypto.Decrypt(source); err != nil {
			panic(err)
		} else {
			outputSource = decrypted
		}
	default:
		panic("unsupported mode")
	}

	ext := filepath.Ext(*file)
	basename := strings.Split(*output, ".")[0]
	outputFile := basename + ext
	outputPath := filepath.Join(wd, outputFile)
	if err := os.WriteFile(outputPath, outputSource, 0666); err != nil {
		panic(err)
	}

	fmt.Println("done...")
}
