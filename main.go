package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// scan given a path crawls it and its subfolders
// searching for Git repositories
func scan(folder string) {
	fmt.Printf("Found folders:\n")
	recursiveScanFolder(folder)
	// filePath := getDotFilePath()
	// addNewSliceElementsToFile(filePath, repositories)
	fmt.Printf("\nSuccessfully added\n")
}

func scanGitFolders(folders []string, folder string) []string {

	// remove "/" character from the end of folder name
	folder = strings.TrimSuffix(folder, "/")

	// open folder
	f, err := os.Open(folder)

	if err != nil {
		log.Fatal(err)
	}

	// fetch files from folder
	files, err := f.Readdir(-1)

	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	var path string

	for _, file := range files {
		if file.IsDir() {
			path = folder + "/" + file.Name()
			if file.Name() == ".git" {
				path = strings.TrimSuffix(path, "/.git")

				fmt.Println(path)
				folders = append(folders, path)
				continue
			}

			if file.Name() == "vendor" || file.Name() == "node_modules" {
				continue
			}

			folders = scanGitFolders(folders, path)
		}
	}

	return folders
}



func recursiveScanFolder (folder string) []string {
	return scanGitFolders(make([]string, 0), folder)
}


// stats generates a nice graph of your Git contributions
func stats(email string) {
	print("stats")
}

func main() {
	var folder string
	var email string

	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "mg.dev1992@gmail.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}

	stats(email)
}