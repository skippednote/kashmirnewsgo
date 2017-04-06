package main

type newsItem struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Source  string `json:"source"`
	Content string `json:"content"`
}

type news []newsItem

type rkNewsItem struct {
	Title   string `json:"title"`
	Link    string `json:"news_url"`
	Source  string
	Content string `json:"news_content"`
}

type rkNews struct {
	Result []rkNewsItem `json:"result"`
}
