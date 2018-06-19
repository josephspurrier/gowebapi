package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"app/webapi/pkg/securegen"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("cliapp", "A command-line application to perform tasks for the webapi.")

	cGenerate = app.Command("generate", "Generate 256 bit (32 byte) base64 encoded JWT.")
)

func main() {
	argList := os.Args[1:]
	arg := kingpin.MustParse(app.Parse(argList))

	switch arg {
	case cGenerate.FullCommand():
		b, err := securegen.Bytes(32)
		if err != nil {
			log.Fatal(err)
		}

		enc := base64.StdEncoding.EncodeToString(b)
		fmt.Println(enc)
	}
}
