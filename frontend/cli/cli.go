package cli

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/beeploop/aes-encrypt/encrypt"
)

type CLI struct{}

func NewCLI() *CLI {
	return &CLI{}
}

func (c *CLI) Start() {
	key := flag.String("key", "", "32 character AES key. You can pass in a txt file")
	file := flag.String("file", "", "Input file")
	mode := flag.String("mode", "encrypt", "operation mode encrypt/decrypt")
	output := flag.String("output", "output", "Output file")

	flag.Parse()
	if *file == "" || *key == "" {
		fmt.Println("Invalid usage")
		flag.PrintDefaults()
		os.Exit(1)
	}

	inputFile := c.readInputFile(*file)
	crypto := c.InitEncryptor(key)

	var outputSource []byte
	switch *mode {
	case "encrypt":
		if encrypted, err := crypto.Encrypt(inputFile); err != nil {
			panic(err)
		} else {
			outputSource = encrypted
		}
	case "decrypt":
		if decrypted, err := crypto.Decrypt(inputFile); err != nil {
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
	c.saveFile(outputSource, outputFile)

	fmt.Println("done...")
}

func (c *CLI) InitEncryptor(key *string) *encrypt.Encrypt {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var encryptionKey []byte
	ext := filepath.Ext(*key)
	if ext == ".txt" {
		key, err := os.ReadFile(filepath.Join(wd, *key))
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewReader(bytes.NewReader(key))
		line, _, err := scanner.ReadLine()
		if err != nil {
			panic(err)
		}

		encryptionKey = []byte(line)
	} else {
		encryptionKey = []byte(*key)
	}

	crypto, err := encrypt.New(encryptionKey)
	if err != nil {
		panic(err)
	}

	return crypto
}

func (c *CLI) readInputFile(src string) []byte {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, src)
	source, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return source
}

func (c *CLI) saveFile(src []byte, output string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	outputPath := filepath.Join(wd, output)
	if err := os.WriteFile(outputPath, src, 0666); err != nil {
		panic(err)
	}
}
