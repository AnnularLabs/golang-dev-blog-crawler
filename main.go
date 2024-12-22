package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type BlogLink struct {
	Title string
	Link  string
}

var (
	blogLinks []*BlogLink
)

// 提取文本内容的函数
func extractBlogLink(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" && strings.HasPrefix(attr.Val, "/blog/") {
				if strings.EqualFold(attr.Val, "/blog/") {
					break
				}
				blogLinks = append(blogLinks, &BlogLink{
					Title: n.FirstChild.Data,
					Link:  "https://go.dev" + attr.Val,
				})
				break
			}
		}
	}
	for nextNode := n.FirstChild; nextNode != nil; nextNode = nextNode.NextSibling {
		extractBlogLink(nextNode)
	}
}

func main() {
	resp, err := http.Get("https://go.dev/blog/all")
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 解析HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("解析HTML失败:", err)
		return
	}

	// 提取博客链接
	extractBlogLink(doc)

	// 保存到文件
	if len(blogLinks) == 0 {
		fmt.Println("没有获取到博客链接")
		return
	}

	for _, link := range blogLinks {
		fmt.Printf("标题: %s\n链接: %s\n\n", link.Title, link.Link)
	}
}
