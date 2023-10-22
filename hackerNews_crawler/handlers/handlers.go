package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/chukwuka-emi/hackerNews_comments_sentiments/hackerNews_crawler/messagebroker"
	"github.com/gocolly/colly"
)

type comment struct {
	PostURL string
	Author  string `selector:"a.hnuser"`
	URL     string `selector:".age a[href]" attr:"href"`
	Comment string `selector:".comment"`
}

type stringArray []string 


func (arr stringArray) Contains(str string) bool{ 
  isMatch := false
  for _, item := range arr{
    if strings.Contains(item,str){
     isMatch = true
	 break;
	}
  }
  return isMatch
}

// CrawlNews ...
func CrawlNews(){
	newsCollector := colly.NewCollector()
	newsCollector.OnHTML("table tr td.subtext .subline",func(e * colly.HTMLElement){
		if strings.Contains(e.ChildText("a"),"comments"){
			crawlComments(e)
		}

	})
	newsCollector.Visit("https://news.ycombinator.com/")
}

func crawlComments(e * colly.HTMLElement){
	commentsCollector := colly.NewCollector()
	href := e.DOM.ChildrenFiltered(".age").ChildrenFiltered("a").First().AttrOr("href","")
	if len(href)==0{
		return
	}
	// comments := make([]*comment,0)
	newsLink := fmt.Sprintf("https://news.ycombinator.com/%s",href)

	commentsCollector.OnHTML(".comment-tree tr.athing",func(h *colly.HTMLElement) {
		commentObj := &comment{Author: "",URL: "",Comment: ""}
		if err := h.Unmarshal(commentObj); err !=nil{
			log.Fatalln("error", err.Error())
		}
		commentObj.Comment = strings.TrimSpace(commentObj.Comment[:len(commentObj.Comment)-5])
		commentObj.URL = fmt.Sprintf("https://news.ycombinator.com/%s",commentObj.URL)
        commentObj.PostURL = newsLink

		jsonString, _ := json.Marshal(commentObj)
		messagebroker.InitBroker(string(jsonString))
	})
	commentsCollector.Visit(newsLink)
}