package main

import(
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

var opts struct {
	Name string `short:"n" long:"name" default:"World" description:"A name to say hello to."`
	Spanish bool `short:"s" long:"spanish" description:"Use Spanish Language"`
}

func main() {
	p := flags.NewParser(&opts, flags.Options(0))
	_, err := p.Parse()

	if err != nil {
		p.WriteHelp(os.Stdout)
	}

	if opts.Spanish == true {
		fmt.Printf("Hola %s!\n", opts.Name)
	} else {
		fmt.Printf("Hello %s!\n", opts.Name)
	}
}
