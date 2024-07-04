package main

import (
	"fmt"
	"log"

	"github.com/suzuki-shunsuke/go-ci-env/v3/cienv"
)

func main() {
	if err := core(); err != nil {
		log.Fatal(err)
	}
}

func core() error {
	p := cienv.Get(nil)
	n, err := p.PRNumber()
	if err != nil {
		return err
	}
	fmt.Printf("PRNumber: %v\n", n)
	return nil
}
