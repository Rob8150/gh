package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cred struct {
	Name    string `json: "Name"`
	Gitcode string `json: "Gitcode"`
	SSH     string `json: "SSH"`
	Project string `json: "Project"`
	Year    int    `json: "Year"`
	Month   int    `json: "Month"`
	Day     int    `json: "Day"`
}

func main() {
	data := Cred{
		Name:    "",
		Gitcode: "",
		SSH:     "",
		Project: "",
		Year:    2022,
		Month:   1,
		Day:     1,
	}

	fnam := ".git.json"
	data.loajson(fnam)

	fmt.Println("Git Hub Repo Generator")

	var filename string = "README.md"
	wr(filename)

	filename = "README.txt"
	rmtxt(filename)

	filename = "License.md"
	lic(filename)

	filename = ".gitignore"
	gi(filename)

	filename = "git.sh"
	gb(filename)

}

func wr(filename string) {

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	proj := "go"

	if proj == "go" {
		fmt.Fprintln(file, "[![Go Version](https://img.shields.io/github/go-mod/go-version/Rob8150/gh)](https://tip.golang.org/doc/go1.18)")
		fmt.Fprintln(file, "[![GoDoc](https://godoc.org/github.com/rwxrob/gh?status.svg)](https://godoc.org/github.com/Rob8150/gh)")
		fmt.Fprintln(file, "[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)")
		fmt.Fprintln(file, "[![Go Report Card](https://goreportcard.com/badge/github.com/rwxrob/gh)](https://goreportcard.com/report/github.com/rwxrob/gh)")
		fmt.Fprintln(file, "[![GoInt](https://github.com/Rob8150/Rob8150/blob/main/golang.png)]")

	}

	if proj == "kotlin" {
		fmt.Fprintln(file, "[![TeamCity (simple build status)](https://img.shields.io/teamcity/http/teamcity.jetbrains.com/s/Kotlin_KotlinPublic_Compiler.svg)](https://teamcity.jetbrains.com/buildConfiguration/Kotlin_KotlinPublic_Compiler?branch=%3Cdefault%3E&buildTypeTab=overview&mode=builds)")

		fmt.Fprintln(file, "[![Maven Central](https://img.shields.io/maven-central/v/org.jetbrains.kotlin/kotlin-maven-plugin.svg)](https://search.maven.org/#search%7Cga%7C1%7Cg%3A%22org.jetbrains.kotlin%22)")

		fmt.Fprintln(file, "[![GitHub license](https://img.shields.io/badge/license-Apache%20License%202.0-blue.svg?style=flat)](https://www.apache.org/licenses/LICENSE-2.0)")

		fmt.Fprintln(file, "[![Revved up by Gradle Enterprise](https://img.shields.io/badge/Revved%20up%20by-Gradle%20Enterprise-06A0CE?logo=Gradle&labelColor=02303A)](https://ge.jetbrains.com/scans?search.rootProjectNames=Kotlin)")

		fmt.Fprintln(file, "[![GoInt](https://github.com/Rob8150/Rob8150/blob/main/kotlin.jpeg)]")

	}

	if proj == "bash" {
		fmt.Fprintln(file, "[![CI](https://github.com/scop/bash-completion/actions/workflows/ci.yaml/badge.svg)](https://github.com/scop/bash-completion/actions/workflows/ci.yaml)")
	}
	fmt.Println("Creating README.md file for Github")

	fmt.Fprintln(file, "# GitHub Utilities in Go 1.19+")
	fmt.Fprintln(file, "???? *under construction* ????")
	fmt.Fprintln(file, "")

	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "## Design Considerations")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "* **Create high-level functions for stuff in the `gh` tool.** No need to re-invent anything here, just pulling out all that existing code into a usable Go 1.19 module.")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "### Work in Progress")
	fmt.Fprintln(file, "* **30 % Complete**")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "__________________________________________________________________________________________________")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "")

	fmt.Fprintln(file, "### Contributor License Agreements")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "Before we can accept your pull requests you'll need to sign a Contributor License Agreement (CLA):")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "- **If you are an individual writing original source code** and **you own the")
	fmt.Fprintln(file, "intellectual property**, then you'll need to sign an client agreement.")

	fmt.Fprintln(file, "- **If you work for a company that wants to allow you to contribute your")
	fmt.Fprintln(file, "  work**, then you'll need to sign a corporate client agreement.")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "You can sign these via email by contacting the Author. After that,")
	fmt.Fprintln(file, "we'll be able to accept your pull requests.")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "Note not all projects may need contributions so just email robert.clark.1967@gmail.com")
	fmt.Fprintln(file, "for license see License.md in this REPO.")

}

func rmtxt(filename string) {

	data := Cred{
		Name:    "",
		Gitcode: "",
		SSH:     "",
		Project: "",
		Year:    2022,
		Month:   1,
		Day:     1,
	}

	fnam := ".git.json"
	data.loajson(fnam)
	currentTime := time.Now()
	t1 := Date(data.Year, data.Month, data.Day)
	t2 := currentTime
	//t1 := Date(2022, 9, 1)
	days := t2.Sub(t1).Hours() / 24

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Println("Creating README.txt file private same as ./play")

	fmt.Fprintln(file, "Todo Create Readme.txt")

	fmt.Fprintln(file, "YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Fprintln(file, "Welcome "+data.Name)
	fmt.Fprint(file, "Elapsed time in Days : ")
	fmt.Fprintln(file, math.Round(days))
	fmt.Fprintln(file, "You are working on project "+data.Project)
	fmt.Fprintln(file, " ")
	fmt.Fprintln(file, ".git.json")
	fmt.Fprintln(file, "----GO MOD --------------------------------------")

	fmt.Fprint(file, "")
	fmt.Fprintln(file, "cd ~/gocode/src/github.com/Rob8150/"+data.Project)
	fmt.Fprintln(file, "go mod init github.com/Rob8150/"+data.Project)
	fmt.Fprintln(file, "go get github.com/charmbracelet/lipgloss")
	fmt.Fprintln(file, " ")

	fmt.Fprintln(file, "----GitHub---------------------------------------")
	fmt.Fprintln(file, "git branch -m master main")
	fmt.Fprintln(file, "git remote remove origin")
	fmt.Fprintln(file, "git remote add origin https://"+data.Gitcode+"@github.com/Rob8150/"+data.Project)
	fmt.Fprintln(file, "git push --set-upstream origin main")

	fmt.Fprintln(file, "cd gocode/src/github.com/Rob8150/"+data.Project)
	fmt.Fprintln(file, "-----Transport Remote-----------------------------")
	fmt.Fprintln(file, " ")
	fmt.Fprintln(file, "-----ZIP------------------------------------------")
	fmt.Fprintln(file, "cd ~/gocode/src/github.com/Rob8150/")
	fmt.Fprintln(file, "tar -cvzf "+data.Project+".tar.gz "+data.Project)

	fmt.Fprintln(file, " ")
	fmt.Fprintln(file, "-----Upstream ------------------------------------")
	fmt.Fprintln(file, "scp ~/gocode/src/github.com/Rob8150/"+data.Project+".tar.gz "+data.SSH+":gocode/src/github.com/Rob8150/"+data.Project+".tar.gz")

	fmt.Fprintln(file, " ")
	fmt.Fprintln(file, "-----DownStream-----------------------------------")
	fmt.Fprintln(file, "scp "+data.SSH+":gocode/src/github.com/Rob8150/"+data.Project+".tar.gz ~/gocode/src/github.com/Rob8150/"+data.Project+".tar.gz")

	fmt.Fprintln(file, "ssh "+data.SSH)
	fmt.Fprintln(file, "cd ~/gocode/src/github.com/Rob8150/")
	fmt.Fprintln(file, "Extract")
	fmt.Fprintln(file, "tar -zxvf "+data.Project+".tar.gz")

}

func (dat *Cred) savjson(d Cred, fnam string) {
	dat.Name = d.Name
	dat.Gitcode = d.Gitcode
	dat.SSH = d.SSH
	dat.Project = d.Project
	dat.Year = d.Year
	dat.Month = d.Month
	dat.Day = d.Day

	file, _ := json.MarshalIndent(dat, "", " ")
	_ = ioutil.WriteFile(fnam, file, 0644)
}

func (dat *Cred) loajson(fnam string) {
	raw, err := ioutil.ReadFile(fnam)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &dat)
}

func (dat *Cred) display() {
	fmt.Println(dat.Name)
	fmt.Println(dat.Gitcode)
	fmt.Println(dat.SSH)
	fmt.Println(dat.Project)
	fmt.Println(dat.Year)
	fmt.Println(dat.Month)
	fmt.Println(dat.Day)

}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func getInput(askfor string) string {
	scanner := bufio.NewScanner(os.Stdin)
	var text string
	fmt.Print(askfor)
	scanner.Scan()
	text = scanner.Text()
	return text
}

// getNum convert string to float64
func getNum(given string) float64 {
	var clean string
	var gotnum float64
	clean = strings.TrimSpace(given)
	gotnum, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		log.Fatal(err)
	}
	return gotnum
}

// -----------------------------------------------------------------------------------------------------------------
func lic(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	fmt.Println("Creating License.md for Github")

	fmt.Fprintln(file, "## Copyright 2022 Robert Ian Clark")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "#### Licensed under the Apache License, Version 2.0 (the \"License\")")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "you may not use this software except in compliance with the License.")
	fmt.Fprintln(file, "You may obtain a copy of the License at")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "## http://www.apache.org/licenses/LICENSE-2.0")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "Unless required by applicable law or agreed to in writing, software")
	fmt.Fprintln(file, "distributed under the License is distributed on an **\"AS IS\"** BASIS,")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "**WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND**, either express or implied.")
	fmt.Fprintln(file, "See the License for the specific language governing permissions and")
	fmt.Fprintln(file, "limitations under the License.")
	fmt.Fprintln(file, "all Software and files within this REPO are under the afore said (the \"License\")")

	fmt.Fprintln(file, "# Contributor License Agreements")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "Before we can accept your pull requests you'll need to sign a Contributor License Agreement (CLA):")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "- **If you are an individual writing original source code** and **you own the")
	fmt.Fprintln(file, "intellectual property**, then you'll need to sign an client agreement.")

	fmt.Fprintln(file, "- **If you work for a company that wants to allow you to contribute your")
	fmt.Fprintln(file, "  work**, then you'll need to sign a corporate client agreement.")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "You can sign these via email by contacting the Author. After that,")
	fmt.Fprintln(file, "we'll be able to accept your pull requests.")
	fmt.Fprintln(file, "")
	fmt.Fprintln(file, "Note not all projects may need contributions so just email robert.clark.1967@gmail.com")
	fmt.Fprintln(file, "")
}

// ------------------------------------------------------------------------------------------------------------------
func gi(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	fmt.Println("Creating .gitignore file")

	fmt.Fprintln(file, "README.txt")
	fmt.Fprintln(file, "git.sh")
	fmt.Fprintln(file, ".git.json")
	fmt.Fprintln(file, "secrets.txt")

}

//-----------------------------------------------------------------------------------------------------------------

func gb(filename string) {

	data := Cred{
		Name:    "",
		Gitcode: "",
		SSH:     "",
		Project: "",
		Year:    2022,
		Month:   1,
		Day:     1,
	}

	fnam := ".git.json"
	data.loajson(fnam)

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	fmt.Println("Creating bash script for Github initialization ")
	fmt.Println("echo Setup Git and GitHub via git.sh")
	fmt.Println("echo be sure to chmod + x git.sh")
	fmt.Println("echo be sure Github Repo has correct name")

	fmt.Fprintln(file, "git init")
	fmt.Fprintln(file, "git branch -m master main")
	fmt.Fprintln(file, "git remote add origin https://"+data.Gitcode+"@github.com/Rob8150/"+data.Project)
	fmt.Fprintln(file, "git add .")
	fmt.Fprintln(file, "git commit -m \"Initial Commit\"")
	fmt.Fprintln(file, "git push --set-upstream origin main")
	fmt.Fprintln(file, "git status")
}

//-------EOF------------------------------------------------------------------------------------------------------
