package main

//Golang remake (with lots of excessive documentation)
//Sorry Kat

import (
	"bufio"
	crand "crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ogier/pflag"
)

var lastIndex = -1

//Generates 'cryptographically secure' random intergers
func SecureIntn(min, max int) int {
	bi, _ := crand.Int(crand.Reader, big.NewInt(int64(max-min)))
	return min + int(bi.Int64())
}

//Allows for "go get" by installing detecting gopath
func goDir() string {
	if e := os.Getenv("GOPATH"); e != "" {
		return e
	}
	return os.Getenv("HOME") + "/go"
}

func getGoDiceWords() string {
	return goDir() + "/src/github.com/pavona/dice/godicewords.txt"
}

//Reads file and finds last index
//Allows for variable sized text lists
func lastindex() int {
	f, err := os.Open(getGoDiceWords())
	if err != nil {
		panic(err)
	}

	rd := bufio.NewReader(f)

	l := 0
	for {
		_, err := rd.ReadString('\n')
		if err != nil {
			break
		}
		l++
	}
	return l
}

//Reads file using random int to find stopping index
//Returns word back to main function
func GetWord() string {
	if lastIndex < 0 {
		lastIndex = lastindex()
	}

	EndOfFile := lastIndex

	SecureNum := SecureIntn(0, EndOfFile)

	f, err := os.Open(getGoDiceWords())
	if err != nil {
		panic(err)
	}

	rd := bufio.NewReader(f)

	l := 0
	for {
		l++
		line, err := rd.ReadString('\n')
		if err != nil {
			break
		}

		if l == SecureNum {
			line = strings.TrimRight(line, "\n")
			return line
		}

		if err != nil {
			break
		}
	}
	return ""
}

//Main function
func main() {
	var lenraw = pflag.IntP("length", "l", 6, "Declares number of words in password (default: 6")
	pflag.Parse()

	//Byte to variable
	passlen := *lenraw

	//Init. empty string that will be the password
	s := ""

	l := 0

	//Iterate until password length matches requested length
	for {
		if l < passlen {
			s += GetWord()
			s += " "
		} else {
			fmt.Println("\n", s, "\n")
			break
		}
		l++
	}

}
