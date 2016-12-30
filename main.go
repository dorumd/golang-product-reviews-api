package main

import (
	"github.com/dorumd/golang-product-reviews-api/product"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var env string

	flag.StringVar(&env, "env", "dev", "ENV")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [arguments] <command> \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		log.Fatal("Command argument required")
	}
	cmd := flag.Arg(0)

	// Configure Server
	service, err := product.NewProductService(env)
	if err != nil {
		log.Fatal(err)
	}

	// Run Main App
	switch cmd {
	case "serve":
		if err := service.Run(); err != nil {
			log.Fatal(err)
		}
	case "migrate-db":
		if err := service.MigrateDb(); err != nil {
			log.Fatal(err)
		}
	case "load-fixtures":
		if err := service.LoadFixtures(); err != nil {
			log.Fatal(err)
		}
	default:
		flag.Usage()
		log.Fatalf("Unknown Command: %s", cmd)
	}
}
