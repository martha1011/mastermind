package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type userGame struct {
	feedbackgameslice [][]string
	solutionList      []string
}

var gameIDslice []*userGame
var gameIDsliceMutex sync.Mutex

func main() {
	fmt.Println("Please open a web browser and head on over to localhost:8080 to play your game!")

	http.Handle("/", http.FileServer(http.Dir("./pagecontent")))

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html><head><title>Mastermind</title><link rel=\"stylesheet\" type=\"text/css\" href=\"avenir-white.css\"><!-- https://github.com/jasonm23/markdown-css-themes/blob/gh-pages/avenir-white.css --></head>")
		fmt.Fprintf(w, "<body>")
		f, _ := os.Open("./pagecontent/setup.html")
		io.Copy(w, f)
		fmt.Fprintf(w, "</body></html>")
	})

	http.HandleFunc("/setup", func(w http.ResponseWriter, r *http.Request) {
		ug := &userGame{}
		gameIDsliceMutex.Lock()
		gameIDslice = append(gameIDslice, ug)
		ID := len(gameIDslice) - 1
		gameIDsliceMutex.Unlock()
		ug.solutionList = storesolution()
		fmt.Fprintf(w, "<html>")
		fmt.Fprintf(w, "<body>")
		f, _ := os.Open("./pagecontent/setup.html")
		io.Copy(w, f)
		//fmt.Fprintf(w, "%s, %d", ug.solutionList, ID)
		fmt.Fprintf(w, "<p><a href='/guess?ID=%d'>Time to start guessing!</a></p></body></html>", ID)
		fmt.Fprintf(w, "<!--%s-->", ug.solutionList)
	})

	http.HandleFunc("/processguess", func(w http.ResponseWriter, r *http.Request) {
		formColour1, _ := strconv.Atoi(r.FormValue("Colour1"))
		formColour2, _ := strconv.Atoi(r.FormValue("Colour2"))
		formColour3, _ := strconv.Atoi(r.FormValue("Colour3"))
		formColour4, _ := strconv.Atoi(r.FormValue("Colour4"))
		formID, _ := strconv.Atoi(r.FormValue("ID"))
		ug := gameIDslice[formID]
		fmt.Fprintf(w, `
<html><head><title>Mastermind</title><link rel="stylesheet" type="text/css" href="avenir-white.css"><!-- https://github.com/jasonm23/markdown-css-themes/blob/gh-pages/avenir-white.css --></head>
<body>`)
		fmt.Fprintf(w, "<p>Your Guess: %s; %s; %s; %s</p>", coloursList[formColour1], coloursList[formColour2], coloursList[formColour3], coloursList[formColour4])
		var guessstring = []string{coloursList[formColour1], coloursList[formColour2], coloursList[formColour3], coloursList[formColour4]}
		//fmt.Fprintf(w, "<p>%s</p>", guessstring)
		//fmt.Fprintf(w, "<p>Temporary Debug Solution: %s</p>", ug.solutionList)
		guessstring = append(guessstring, "~~~")
		switch ug.solutionList[0] {
		case guessstring[0]:
			guessstring = append(guessstring, White)
		case guessstring[1], guessstring[2], guessstring[3]:
			guessstring = append(guessstring, Black)
		default:
			guessstring = append(guessstring, "_____")
		}
		switch ug.solutionList[1] {
		case guessstring[1]:
			guessstring = append(guessstring, White)
		case guessstring[0], guessstring[2], guessstring[3]:
			guessstring = append(guessstring, Black)
		default:
			guessstring = append(guessstring, "_____")
		}
		switch ug.solutionList[2] {
		case guessstring[2]:
			guessstring = append(guessstring, White)
		case guessstring[0], guessstring[1], guessstring[3]:
			guessstring = append(guessstring, Black)
		default:
			guessstring = append(guessstring, "_____")
		}
		switch ug.solutionList[3] {
		case guessstring[3]:
			guessstring = append(guessstring, White)
		case guessstring[0], guessstring[1], guessstring[2]:
			guessstring = append(guessstring, Black)
		default:
			guessstring = append(guessstring, "_____")
		}
		if strings.Join(guessstring[5:], "") == strings.Join([]string{White, White, White, White}, "") {
			fmt.Fprintf(w, "Woop you win! It'll be a new solution next time, so please <a href='/setup'>click here</a> to play again :)</p>")
		} else {
			fmt.Fprintf(w, "<p>White means right colour, right position.</p>")
			fmt.Fprintf(w, "<p>Black means right colour, wrong position.</p>")
			fmt.Fprintf(w, "<p> _____ means wrong colour, wrong position!</p>")
			fmt.Fprintf(w, "Your feedback: %s", guessstring)
			ug.feedbackgameslice = append(ug.feedbackgameslice, guessstring)
			fmt.Fprintf(w, "<p>I'm afraid you haven't got the solution yet. <a href='/guess?ID=%d'>Click here</a> to guess again!</p>", formID)
		}
		fmt.Fprintf(w, "</html></body>")
	})

	http.HandleFunc("/guess", func(w http.ResponseWriter, r *http.Request) {
		ID, _ := strconv.Atoi(r.FormValue("ID"))
		//fmt.Println("DEBUG", ID, err)
		ug := gameIDslice[ID]
		fmt.Fprintf(w, "<html><body>")
		f, _ := os.Open("./pagecontent/guess.html")
		io.Copy(w, f)
		for i, fb := range ug.feedbackgameslice {
			fmt.Fprintf(w, "<p>Your Guess Number %d:%s</p>", i, fb)
		}
		fmt.Fprintf(w, "<form action=\"/processguess?ID=%d\" method=\"post\">", ID)
		fmt.Fprintf(w, "<input type=\"hidden\" name=\"ID\" value=\"%d\">", ID)
		fmt.Fprintf(w, `<div>
<div style="float:left; width:101px; height:auto;">
    <div style="width:200px; float:left;">
        Colour 1
    </div>
    <div style="width:200px; float:left;">`)
		fmt.Fprintf(w, "<p><select name=\"Colour1\">")
		for c1f := range coloursList {
			fmt.Fprintf(w, "<option value=\"%d\">%s</option>", c1f, coloursList[c1f])
		}

		fmt.Fprintf(w, `</select></p></div>
</div>
    <div style="float:left; width:101px; height:auto;">
    <div style="width:200px; float:left;">
        Colour 2
    </div>
    <div style="width:200px; float:left;">`)
		fmt.Fprintf(w, "<p><select name=\"Colour2\">")
		for c2f := range coloursList {
			fmt.Fprintf(w, "<option value=\"%d\">%s</option>", c2f, coloursList[c2f])
		}
		fmt.Fprintf(w, `</select></p></div>
</div>
<div style="float:left; width:101px; height:auto;">
    <div style="width:200px; float:left;">
        Colour 3
    </div>
    <div style="width:200px; float:left;">`)
		fmt.Fprintf(w, "<p><select name=\"Colour3\">")
		for c3f := range coloursList {
			fmt.Fprintf(w, "<option value=\"%d\">%s</option>", c3f, coloursList[c3f])
		}
		fmt.Fprintf(w, `</select></p></div>
</div>
<div style="float:left; width:101px; height:auto;">
    <div style="width:200px; float:left;">
        Colour 4
    </div>
    <div style="width:200px; float:left;">`)
		fmt.Fprintf(w, "<p><select name=\"Colour4\">")
		for c4f := range coloursList {
			fmt.Fprintf(w, "<option value=\"%d\">%s</option>", c4f, coloursList[c4f])
		}
		fmt.Fprintf(w, "</select></p></div></div></div>")
		fmt.Fprintf(w, `<input type="submit" value="Submit"></form>`)
		fmt.Fprintf(w, "</body></html>")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
