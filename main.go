package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const fileRandomNums = "random_nums.txt"
const fileRandomWords = "random_words.txt"

const numsRandom = 10000
const numsMin = 1
const numsMax = 1000
const stringsRandom = 1000
const letters = "abcdefghigklmnopqrstuvwxyz"
const stringMin = 1
const stringMax = 16

var random = rand.New(rand.NewSource(time.Now().Unix()))

func main() {
	var msg string

	nums := randNum()
	msg = fromArrToStr(nums)
	writeFile(fileRandomNums, msg)
	words := randString()
	msg = fromArrToStr(words)
	writeFile(fileRandomWords, msg)
}

func fromArrToStr(randValues []string) string {
	result := ""
	for i := range randValues {
		result += randValues[i] + "\n"
	}
	return result
}

func randString() []string {
	var randStrings []string
	for i := 0; i < stringsRandom; i++ {
		word := ""
		randLenStr := random.Intn(stringMax-stringMin+1) + stringMin
		for k := 0; k < randLenStr; k++ {
			randLetter := letters[random.Intn(len(letters))]
			word += string(randLetter)
		}
		randStrings = append(randStrings, word)
	}

	return randStrings
}

func randNum() []string {
	var randNums []string
	for i := 0; i < numsRandom; i++ {
		num := random.Intn(numsMax-numsMin+1) + numsMin
		strNum := strconv.Itoa(num)
		randNums = append(randNums, strNum)
	}
	return randNums
}

func writeFile(fileName string, msg string) {
	err := ioutil.WriteFile(fileName, []byte(msg), 0755)

	if err != nil {
		fmt.Println("I can't write a file.", err)
		os.Exit(1)
	}

	fmt.Printf("Success: The content was written to a file \"%s\".\n", fileName)
}

func readFile(fileName string) string {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("I can't read a file.")
	}
	return string(data)
}
