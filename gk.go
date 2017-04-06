package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

func getGK() news {
	doc, err := goquery.NewDocument("http://www.greaterkashmir.com/")
	if err != nil {
		log.Fatal(err)
	}
	ns := make(news, 0)

	block := doc.Find("#ctl00_ContentPlaceHolder1_mostRead1_dvMostPopular").Find("h2")
	if block.Text() == "Latest News" {
		news := doc.Find(".latestNews").First().Find(".Latesthead a")
		news.Each(func(i int, s *goquery.Selection) {
			title := s.Text()
			link, exists := s.Attr("href")

			if exists {
				link = link
			}
			n := newsItem{
				Title:   title,
				Link:    link,
				Source:  "GreaterKashmir",
				Content: "",
			}
			ns = append(ns, n)
		})
	}
	return ns
}
