/*
cfssl is the command line tool to issue/sign/bundle client certificate. It's
also a tool to start a HTTP server to handle web requests for signing, bundling
and verification.

Usage:
	cfssl command [-flags] arguments

	The commands are

	bundle	 create a certificate bundle
	sign	 signs a certificate signing request (CSR)
	serve	 starts a HTTP server handling sign and bundle requests
	version	 prints the current cfssl version
	genkey   generates a key and an associated CSR
	gencert  generates a key and a signed certificate
	selfsign generates a self-signed certificate

Use "cfssl [command] -help" to find out more about a command.
*/
package main

import (
	"flag"
	"os"

	"github.com/jazznetworks/cfssl/cli"
	"github.com/jazznetworks/cfssl/cli/bundle"
	"github.com/jazznetworks/cfssl/cli/certinfo"
	"github.com/jazznetworks/cfssl/cli/crl"
	"github.com/jazznetworks/cfssl/cli/gencert"
	"github.com/jazznetworks/cfssl/cli/gencrl"
	"github.com/jazznetworks/cfssl/cli/genkey"
	"github.com/jazznetworks/cfssl/cli/info"
	"github.com/jazznetworks/cfssl/cli/ocspdump"
	"github.com/jazznetworks/cfssl/cli/ocsprefresh"
	"github.com/jazznetworks/cfssl/cli/ocspserve"
	"github.com/jazznetworks/cfssl/cli/ocspsign"
	"github.com/jazznetworks/cfssl/cli/printdefault"
	"github.com/jazznetworks/cfssl/cli/revoke"
	"github.com/jazznetworks/cfssl/cli/scan"
	"github.com/jazznetworks/cfssl/cli/selfsign"
	"github.com/jazznetworks/cfssl/cli/serve"
	"github.com/jazznetworks/cfssl/cli/sign"
	"github.com/jazznetworks/cfssl/cli/version"

	_ "github.com/go-sql-driver/mysql" // import to support MySQL
	_ "github.com/lib/pq"              // import to support Postgres
	_ "github.com/mattn/go-sqlite3"    // import to support SQLite3
)

// main defines the cfssl usage and registers all defined commands and flags.
func main() {
	// Add command names to cfssl usage
	flag.Usage = nil // this is set to nil for testabilty
	// Register commands.
	cmds := map[string]*cli.Command{
		"bundle":         bundle.Command,
		"certinfo":       certinfo.Command,
		"crl":            crl.Command,
		"sign":           sign.Command,
		"serve":          serve.Command,
		"version":        version.Command,
		"genkey":         genkey.Command,
		"gencert":        gencert.Command,
		"gencrl":         gencrl.Command,
		"ocspdump":       ocspdump.Command,
		"ocsprefresh":    ocsprefresh.Command,
		"ocspsign":       ocspsign.Command,
		"ocspserve":      ocspserve.Command,
		"selfsign":       selfsign.Command,
		"scan":           scan.Command,
		"info":           info.Command,
		"print-defaults": printdefaults.Command,
		"revoke":         revoke.Command,
	}

	// If the CLI returns an error, exit with an appropriate status
	// code.
	err := cli.Start(cmds)
	if err != nil {
		os.Exit(1)
	}
}
