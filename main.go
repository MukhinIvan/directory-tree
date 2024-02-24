package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	return printDir(out, path, printFiles, "")
}

func printDir(out io.Writer, path string, printFiles bool, indent string) error {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	if !printFiles {
		tmpSlice := []fs.DirEntry{}
		for _, dirEntry := range dirEntries {
			if dirEntry.IsDir() {
				tmpSlice = append(tmpSlice, dirEntry)
			}
		}
		dirEntries = tmpSlice
	}

	sort.Slice(dirEntries, func(i, j int) bool {
		return dirEntries[i].Name() < dirEntries[j].Name()
	})

	for i, dirEntry := range dirEntries {
		var pref1 string
		if i+1 != len(dirEntries) {
			pref1 = "├───"
		} else {
			pref1 = "└───"
		}

		if dirEntry.IsDir() {
			fmt.Fprintf(out, "%s%s\n", indent+pref1, dirEntry.Name())
			var pref2 string
			if i+1 != len(dirEntries) {
				pref2 += indent + "│\t"
			} else {
				pref2 += indent + "\t"
			}
			printDir(out, filepath.Join(path, dirEntry.Name()), printFiles, pref2)
		} else if printFiles {
			info, err := dirEntry.Info()
			if err != nil {
				return err
			}
			var size string
			if info.Size() == 0 {
				size = "empty"
			} else {
				size = fmt.Sprintf("%db", info.Size())
			}
			fmt.Fprintf(out, "%s%s (%s)\n", indent+pref1, dirEntry.Name(), size)
		}
	}
	return nil
}
