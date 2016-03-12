package main

import (
	"fmt"

	parser "local/haya14busa/hatebu-ranking/parser"
)

func main() {
	es, err := parser.NewEntries("http://b.hatena.ne.jp/ranking/daily/")
	if err != nil {
		panic(err.Error())
	}
	for _, e := range es {
		fmt.Printf("rank:%d\t", e.Rank)
		fmt.Printf("url:%s\t", e.Url)
		fmt.Printf("title:%s\t", e.Title)
		fmt.Println("")
	}
}
