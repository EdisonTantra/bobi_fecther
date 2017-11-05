package main 

import (
	"github.com/kataras/iris/context"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"net/http"
	"fmt"
)

type Theater struct {
	Id string
	Name string
	Address string
	Phone_number string
	Link string
}

func theatersHandler(ctx context.Context) {
	var theaters_data []Theater
	resp, err := http.Get("http://21cineplex.com/theaters")
	if err != nil {
		panic(err)
	}

	raw_data, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	
	matcher := func(n *html.Node) bool {
		if n.DataAtom == atom.Tr && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n.Parent.Parent, "id") == "tb_theater"
		}
		return false
	}

	theaters := scrape.FindAll(raw_data, matcher)
	for _, theater := range theaters {
		node_a := theater.FirstChild.FirstChild
		if node_a != nil && node_a.DataAtom == atom.A {
			theater_id := scrape.Attr(theater, "data-city")
			link := scrape.Attr(node_a, "href")
			name := scrape.Text(node_a)
			addr := scrape.Attr(node_a, "rel")
			phone_number :=  scrape.Text(theater.LastChild)

			theater_json := Theater{ theater_id, name, addr, phone_number, link }
			theaters_data = append(theaters_data, theater_json)
		}
	}
	ctx.JSON(theaters_data)
	fmt.Println("Theaters: ok")
}