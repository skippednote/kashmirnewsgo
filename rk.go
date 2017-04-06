package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getRK() news {
	res, err := http.Post("http://www.risingkashmir.com/home/get_newsupdates", "application/json", nil)
	if err != nil {
		log.Fatal(err)
	}

	r := &rkNews{}
	json.NewDecoder(res.Body).Decode(r)

	ns := make(news, 0)
	for _, item := range r.Result {
		n := newsItem{}
		n.Title = item.Title
		n.Content = item.Content
		n.Link = item.Link
		n.Source = item.Source

		ns = append(ns, n)
	}

	return ns
}
