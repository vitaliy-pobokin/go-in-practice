package main

import (
	"fmt"
	"gopkg.in/ini.v1"
)

func main() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}

	err := ini.MapTo(&config, "conf.ini")
	if err != nil {
		fmt.Printf("Failed to parse config file: %s\n", err)
	}
	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)
}
