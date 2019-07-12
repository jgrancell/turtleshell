package main

import (
	"fmt"
	"os"
)

func historySave(text string, conf Configuration) {
	// Determining if our HistFile exists
	if results := createFile(conf.HistFile); results == true {
		file, err := os.OpenFile(conf.HistFile, os.O_APPEND|os.O_WRONLY, 0600)

		if err != nil {
			fmt.Println("Unable to access tsh history file", conf.HistFile)
		} else {
			defer file.Close()

			if _, err = file.WriteString(text); err != nil {
				fmt.Println("Unable to write to the tsh history file", conf.HistFile)
			}
		}
	}
}
