package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

const (
	BASE_URL              = "https://go.dev"
	FLAG_BLOG_TITLE_ClASS = ".blogtitle"
)

type BlogLink struct {
	Title *string
	Link  *string
}

var (
	BlogLinks []*BlogLink
)

func extractBlogLink(doc *goquery.Document) {
	doc.Find(FLAG_BLOG_TITLE_ClASS).Each(func(i int, s *goquery.Selection) {
		link, _ := s.Find("a").Attr("href")
		if strings.EqualFold(link, "/blog/") {
			return
		}
		title := s.Find("a").Text()

		blogLink := &BlogLink{
			Title: &title,
			Link:  &link,
		}
		BlogLinks = append(BlogLinks, blogLink)
	})
}

func main() {
	blogAllUrl := fmt.Sprintf("%s/blog/all", BASE_URL)
	resp, err := http.Get(blogAllUrl)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 解析HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("解析HTML失败:", err)
		return
	}

	extractBlogLink(doc)

	if len(BlogLinks) == 0 {
		fmt.Println("未找到博客链接")
		return
	}

	for _, blogLink := range BlogLinks {
		fmt.Printf("标题: %s\n链接: %s\n", *blogLink.Title, *blogLink.Link)
	}

}
