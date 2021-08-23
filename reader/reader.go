package reader

import (
	"encoding/json"
	"fmt"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var articles []Article
var articleFolder = "./articles"

type Article struct {
	Title string `json:"Title"`
	Content string `json:"Content"`
	Credit string `json:"Credit"`
	Writer string `json:"Writer"`
	Url string `json:"Url"`
	Date string `json:"Date"`
}

func Reader() {
	files, err := ioutil.ReadDir("./articles")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileDirectory := fmt.Sprintf("./articles/%s", file.Name())

		jsonFile, err := os.Open(fileDirectory)
		defer jsonFile.Close()
		if err != nil {
			log.Fatal(err)
		}

		jsonData, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Fatal(err)
		}

		article := Article{}
		err = json.Unmarshal(jsonData, &article)

		articles = append(articles, article)
		//articles = append(articles, file.Name())
	}

	templates := &promptui.SelectTemplates{
		Label: "{{ . }}?",
		Active:   "\U0000261E{{ .Title | cyan }} ({{ .Date | red }})",
		Inactive: "  {{ .Title | cyan }} ({{ .Date | red }})",
		Selected: "\U0000261E{{ .Title | red | cyan }}",
		Details:`
--------- Article Details ----------
{{ "Name:" | faint }}	{{ .Title }}
{{ "Date:" | faint }}	{{ .Date }}
{{ "Url:" | faint }}	{{ .Url }}`,
	}

	prompt := promptui.Select{
		Label:    "Articles",
		Items:    articles,
		Searcher: searcher,
		Size:     20,
		Templates: templates,
	}

	index, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fileName := fmt.Sprintf("%s.json", articles[index].Title)
	displayArticle(fileName)
}

// TODO update searcher for date
func searcher(input string, index int) bool {
	article := articles[index]
	name := strings.Replace(strings.ToLower(article.Title), " ", "", -1)
	input = strings.Replace(strings.ToLower(input), " ", "", -1)
	return strings.Contains(name, input)
}

func displayArticle(name string) {
	var article Article
	file := fmt.Sprintf("%s/%s", articleFolder, name)

	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(byteValue, &article)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(article.Content)
}