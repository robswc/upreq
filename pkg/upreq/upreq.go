package upreq

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func GetReqs(path string) []string {
	// open the file
	file, err := os.Open(path)

	// handle any errors that occur from opening the file via os.Open
	if err != nil {
		log.Fatal(err)
	}

	// close the file when we're done
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// create a new scanner and read the file line by line
	scanner := bufio.NewScanner(file)
	var reqs []string
	for scanner.Scan() {
		reqs = append(reqs, scanner.Text())
	}

	// handle any errors that occur from scanning the file
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// return the requirements as a slice of strings
	return reqs
}

func WipeFile(path string) {
	err := os.Truncate(path, 0)
	if err != nil {
		fmt.Println(err)
	}
}

func WriteReqs(path string, reqs []string) []string {
	// open the file
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	// handle any errors that occur from opening the file via os.Open
	if err != nil {
		log.Fatal(err)
	}

	// close the file when we're done
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	// iterate over reqs and write them to the file
	for _, req := range reqs {
		_, err := file.WriteString(req + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	return GetReqs(path)
}

func DiffCheck(oldReqs []string, newReqs []string) []string {
	var diff []string
	for _, req := range newReqs {
		if !Contains(oldReqs, req) {
			diff = append(diff, "[+] "+req)
		}
	}
	return diff
}

func DisplayDiff(diff []string, strip string) {
	if strip == "true" {
		for _, req := range diff {
			fmt.Println(req[4:])
		}
	} else {
		for _, req := range diff {
			fmt.Println(req)
		}
	}
}

func Contains(reqs []string, req string) bool {
	for _, r := range reqs {
		if r == req {
			return true
		}
	}
	return false
}

func GetEnvReqs() []string {
	freeze := exec.Command("pip", "freeze")
	out, err := freeze.Output()
	if err != nil {
		fmt.Println(err)
	}
	reqs := strings.Fields(string(out))
	return reqs
}
