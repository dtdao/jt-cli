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

var articles []string
var articleFolder = "./articles"

type Article struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

func Reader() {
	files, err := ioutil.ReadDir("./articles")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		articles = append(articles, file.Name())
	}

	prompt := promptui.Select{
		Label:    "Select articles",
		Items:    articles,
		Searcher: searcher,
		Size:     20,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
	displayArticle(result)
}

func searcher(input string, index int) bool {
	article := articles[index]
	name := strings.Replace(strings.ToLower(article), " ", "", -1)
	input = strings.Replace(strings.ToLower(input), " ", "", -1)
	return strings.Contains(name, input)
}

func displayArticle(name string) {

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
	var article Article
	err = json.Unmarshal(byteValue, &article)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(article.Content)
}