package parser

import (
	"dicuz-crawler/config"
	"dicuz-crawler/model"
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func ParseSection(doc *goquery.Document, item model.Item) model.ParseResult {
	parseResult := model.ParseResult{}
	doc.Find(config.Crawler.Selector.Section).Each(func(i int, selection *goquery.Selection) {
		content, _ := selection.Html()
		url, _ := selection.Attr("href")
		//log.Printf("url: %s, content: %s", url, content)
		item.Title = content
		item.Url = url
		IdRe := regexp.MustCompile(`/thread-([\d]+)-1-1.html`)
		match := IdRe.FindSubmatch([]byte(url))
		if len(match) >= 2 {
			item.Id = string(match[1])
		}
		parseResult.Items = append(parseResult.Items, content)
		parseResult.Requests = append(parseResult.Requests, model.Request{
			Url:       url,
			ParseFunc: ParseContent,
			Deliver:   item,
		})
	})
	return parseResult
}