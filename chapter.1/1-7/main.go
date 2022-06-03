package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	defaultLanguage = kingpin.Flag("default-language", "Default Language").String()

	generateCmd = kingpin.Command("create-index", "Generate Index")
	inputFolder = generateCmd.Arg("INPUT", "Input Folder").Required().String()

	searchCmd  = kingpin.Command("search", "Search")
	inputFile  = searchCmd.Flag("input", "Input index file").Short('i').String()
	searchWord = searchCmd.Arg("WORDS", "Search words").Strings()
)

func main() {
	switch kingpin.Parse() {
	case generateCmd.FullCommand():
		fmt.Println(*defaultLanguage)
		fmt.Println(*inputFolder)
		fmt.Println(*inputFile)
		fmt.Println(*searchWord)
	case searchCmd.FullCommand():
		fmt.Println(*defaultLanguage)
		fmt.Println(*inputFolder)
		fmt.Println(*inputFile)
		fmt.Println(*searchWord)
	}
}
