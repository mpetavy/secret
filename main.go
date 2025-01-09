package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/mpetavy/common"
	"os"
)

//go:embed go.mod
var resources embed.FS

var (
	txt = flag.String("t", "", "secret value")
)

func init() {
	common.Init("", "", "", "", "", "", "", "", &resources, nil, nil, run, 0)
}

func run() error {
	key := os.Getenv("SECRETKEY")
	if key == "" {
		return fmt.Errorf("SECRETKEY environment variable not set")
	}

	if common.IsStringEnrypted(*txt) {
		m, err := common.DecryptString([]byte(key), *txt)
		if common.Error(err) {
			return err
		}

		fmt.Printf("%s\n", m)
	} else {
		m, err := common.EncryptString([]byte(key), *txt)
		if common.Error(err) {
			return err
		}

		fmt.Printf("%s\n", m)
	}

	return nil
}

func main() {
	common.Run([]string{"t"})
}
