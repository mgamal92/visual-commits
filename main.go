package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strings"
)

// scan given a path crawls it and its subfolders
// searching for Git repositories
func scan(folder string) {
	fmt.Printf("Found folders:\n")
	repositories := recursiveScanFolder(folder)
	filePath := getDotFilePath()
	addNewSliceElementsToFile(filePath, repositories)
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

func getDotFilePath() string {
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + "/.gogitlocalstats"
	
	return dotFile
}


func addNewSliceElementsToFile(filePath string, newRepos []string) {
	existingRepos := parseFileLinesToSlice(filePath)
    repos := joinSlices(newRepos, existingRepos)
    dumpStringsSliceToFile(repos, filePath)
}

// given a file path string, gets the content of each line and parses it to a slice of strings.
func parseFileLinesToSlice (filePath string) []string{
	f := openFile(filePath)
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
        if err != io.EOF {
            // panic(err)
        }
	}
	
	return lines
}

func openFile (filePath string) *os.File {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 755)

	if err != nil {
		if os.IsNotExist(err) {
			_, err = os.Create(filePath)

			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	return f
}

func joinSlices(new []string, existing []string) []string {
	for _, i := range new {
		if !sliceContains(existing, i) {
			existing = append(existing, i)
		}
	}

	return existing
}

func sliceContains(slice []string, value string) bool {
	for _, v := range slice {
		if (v == value ){
			return true
		}
	}

	return false
}


func dumpStringsSliceToFile(repos []string, filePath string) {
	content := strings.Join(repos, "\n")
	ioutil.WriteFile(filePath, []byte(content), 0755)
}
/***************Part 2************************ */
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