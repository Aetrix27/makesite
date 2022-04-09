package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Page holds all the information we need to generate a new
// HTML page from a text file on the filesystem.
type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func main() {
	directory := "./"
	files, err := ioutil.ReadDir(directory)

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		current_file := file.Name()
		//fmt.Println(current_file)
		if strings.Contains(current_file, ".txt") {
			fmt.Println(current_file)
			filenames := flag.String(current_file, current_file, "")
			flag.Parse()
			save(*filenames)
		}

	}

	//save(*filename)
}

func save(filename string) {
	fileContents, err := ioutil.ReadFile(filename)
	trimmed := strings.TrimSuffix(filename, ".txt")
	fmt.Println(trimmed)
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	fmt.Print(string(fileContents))

	page := Page{
		TextFilePath: "./" + trimmed,
		TextFileName: filename,
		HTMLPagePath: trimmed + ".html",
		Content:      string(fileContents),
	}

	// Create a new template in memory named "template.tmpl".
	// When the template is executed, it will parse template.tmpl,
	// looking for {{ }} where we can inject content.
	t := template.Must(template.New(trimmed + ".tmpl").ParseFiles(trimmed + ".tmpl"))

	// Create a new, blank HTML file.
	newFile, err := os.Create(trimmed + ".html")
	if err != nil {
		panic(err)
	}

	// Executing the template injects the Page instance's data,
	// allowing us to render the content of our text file.
	// Furthermore, upon execution, the rendered template will be
	// saved inside the new file we created earlier.
	t.Execute(newFile, page)
}
