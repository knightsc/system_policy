package main

import (
	"flag"
	"log"

	osquery "github.com/kolide/osquery-go"

	"github.com/knightsc/system_policy/osquery/table/kextpolicy"
	"github.com/knightsc/system_policy/osquery/table/legacyexec"
)

func main() {
	flSocket := flag.String("socket", "", "")
	flag.Int("timeout", 0, "")
	flag.Int("interval", 0, "")
	flag.Bool("verbose", false, "")
	flag.Parse()

	if *flSocket == "" {
		log.Fatalln("--socket flag cannot be empty")
	}

	server, err := osquery.NewExtensionManagerServer("system_policy", *flSocket)
	if err != nil {
		log.Fatalf("Error creating osquery extension server: %s\n", err)
	}

	server.RegisterPlugin(legacyexec.TablePlugin(), kextpolicy.TablePlugin())
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
