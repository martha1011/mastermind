package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	Yellow = "<FONT style=\"font-family:monospace; BACKGROUND-COLOR: yellow\"><span style=\"color:yellow\">Yellow</span></FONT>"
	White  = "<span style=\"width:10px;height:10px;border:1px solid black;\"><span style=\"font-family:monospace; color:white\">White_</span></span>"
	Blue   = "<span style=\"color:blue\"><FONT style=\"font-family:monospace; BACKGROUND-COLOR: blue\">Blue__</FONT></span>"
	Green  = "<span style=\"color:green\"><FONT style=\"font-family:monospace; BACKGROUND-COLOR: green\">Green_</FONT></span>"
	Black  = "<span style=\"color:black\"><FONT style=\"font-family:monospace; BACKGROUND-COLOR: black\">Black_</FONT></span>"
	Red    = "<span style=\"color:red\"><FONT style=\"font-family:monospace; BACKGROUND-COLOR: red\">Red___</FONT></span>"
)

var coloursList = []string{
	Yellow,
	White,
	Blue,
	Green,
	Black,
	Red,
}

var coloursListtochange = []string{
	Yellow,
	White,
	Blue,
	Green,
	Black,
	Red,
}

//var solutionList = []string{}

func storesolution() []string {
	var i int
	var tempcoloursList []string
	tempcoloursList = append(tempcoloursList, coloursListtochange...)
	var solutionList = []string{}
	for i = 0; i < 4; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		c := rand.Intn(len(tempcoloursList))
		solutionList = append(solutionList, tempcoloursList[c])
		//fmt.Println(tempcoloursList, c, solutionList)
		newtempcoloursList := tempcoloursList[:c]
		//fmt.Println("newtempcolourslist", newtempcoloursList)
		if c != len(tempcoloursList)-1 {
			//for z := c + 1; z < len(tempcoloursList); z++ {
			var tempvar = tempcoloursList[c+1:]
			//fmt.Println("newtempcoloursList part deux", tempvar)
			newtempcoloursList = append(newtempcoloursList,
				tempvar...)

			//}
		}
		tempcoloursList = newtempcoloursList
	}
	return solutionList
	//fmt.Println(solutionList)
}

func printlist(s []string) {
	//for i := range s {
		//fmt.Println(i, s[i])
	//}
}

func getanswer(thinglist []string) (v int) {
getthing:
	var buffer string
	_, err := fmt.Scanln(&buffer)
	if err != nil {
		fmt.Println("Oops - error! You must put in a number between 0 and",
			(len(thinglist) - 1), "\nTry again!")
		goto getthing
	}
	v, err = strconv.Atoi(buffer)
	if err != nil {
		fmt.Println("Oops - error! You must put in a number between 0 and",
			(len(thinglist) - 1), "\nTry again!")
		goto getthing
	}
	if v >= len(thinglist) || v < 0 {
		fmt.Println("Sorry, you must put in a number between 0 and",
			(len(thinglist) - 1))
		goto getthing
	}
	return v
}

func getguesses() (c1, c2, c3, c4 int) {
	fmt.Println("Guess your 1st colour from the available colours below - remember each colour will only appear once in the 4 colour sequence!")
	printlist(coloursList)
	c1 = getanswer(coloursList)
	fmt.Println("Guess your 2nd colour from the available colours below - remember each colour will only appear once in the 4 colour sequence!")
	printlist(coloursList)
	c2 = getanswer(coloursList)
	fmt.Println("Guess your 3rd colour from the available colours below - remember each colour will only appear once in the 4 colour sequence!")
	printlist(coloursList)
	c3 = getanswer(coloursList)
	fmt.Println("Guess your 4th colour from the available colours below - remember each colour will only appear once in the 4 colour sequence!")
	printlist(coloursList)
	c4 = getanswer(coloursList)
	return c1, c2, c3, c4
}

