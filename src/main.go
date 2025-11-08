package main

import (
	"fmt"
	"integrity-check/src/hash"
	"os"
	"path/filepath"
)

func initialize(directory string) {
	if !hash.IsDir(directory) {
		fmt.Println("The specified path is not a directory:", directory)
		return
	}
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			hashValue, err := hash.ComputeFileHash(path)
			if err != nil {
				panic(err)
			}
			hash.WriteKey(path, hashValue)
		}
		return nil
	})
}

func checkDirectory(directory string) {
	hasUnmached := false
	files := []string{}
	if !hash.IsDir(directory) {
		fmt.Println("The specified path is not a directory:", directory)
		return
	}
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			hashValue, err := hash.ComputeFileHash(path)
			if err != nil {
				panic(err)
			}
			origin, exist := hash.GetKey(path)
			if !exist {
				fmt.Println("No stored hash for file:", path)
				return nil
			}
			if origin != hashValue {
				files = append(files, path)
				hasUnmached = true
			}
		}
		return nil
	})
	if hasUnmached {
		fmt.Println("Status: Modified (Hash mismatch)")
		for _, file := range files {
			fmt.Println(" -", file)
		}
	} else {
		fmt.Println("Status: Unmodified")
	}
}

func checkFile(file string) {
	if !hash.IsFile(file) {
		fmt.Println("The specified path is not a file:", file)
		return
	}
	hashValue, err := hash.ComputeFileHash(file)
	if err != nil {
		panic(err)
	}
	origin, exist := hash.GetKey(file)
	if !exist {
		fmt.Println("No stored hash for file:", file)
	} else if origin == hashValue {
		fmt.Println("Status: Unmodified")
	} else {
		fmt.Println("Status: Modified (Hash mismatch)")
	}
}

func updateStore(directory string) {
	if !hash.IsDir(directory) {
		fmt.Println("The specified path is not a directory:", directory)
		return
	}
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			hashValue, err := hash.ComputeFileHash(path)
			if err != nil {
				panic(err)
			}
			hash.WriteKey(path, hashValue)
		}
		return nil
	})
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: integrity-check <command> <value>")
		return
	}

	command := os.Args[1]
	value := os.Args[2]

	switch command {
	case "init":
		initialize(value)
	case "check":
		checkDirectory(value)
	case "-check":
		checkFile(value)
	case "update":
		updateStore(value)
	default:
		fmt.Println("Unknown command:", command)
	}
}
