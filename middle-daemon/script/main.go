package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Listening to certificate request...")
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			words := strings.Fields(scanner.Text())
			count, err := strconv.Atoi(words[1])
			check(err)

			domain_name := words[0]
			fmt.Println("Delivering precertificate...")
			sourceFile := "../../CA-middle-daemon-storage/Precertificate/" + domain_name + "/precert.pem"
			destinationfolder := "../../middle-daemon-website-daemon-storage/Precertificate/" + domain_name
			if _, err := os.Stat(destinationfolder); os.IsNotExist(err) {
				os.Mkdir(destinationfolder, 0744)
			}
			copy_file(sourceFile, destinationfolder+"/precert.pem")
			fmt.Println("Decrypting certificates...")
			main_decrypt(domain_name, count)

			fmt.Println("Building merkle tree and proof...")
			ex, err := os.Executable()
			check(err)
			exPath := filepath.Dir(ex)

			c := exec.Command("python", exPath+"/MDincproof.py", domain_name, strconv.Itoa(count))
			if err := c.Run(); err != nil {
				check(err)
			}
		}

		if scanner.Err() != nil {
			panic(scanner.Err())
		}
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func copy_file(sourceFile string, destinationFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(destinationFile, input, 0744)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return
	}
}
