package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {

	myAmazonUrl := "https://www.amazon.com.br/s?k=mesas&__mk_pt_BR=%C3%85M%C3%85%C5%BD%C3%95%C3%91"
	// Meu colletor basico, preciso de no minimo um para coletar
	// um determinado site por vez

	c := colly.NewCollector(
		colly.AllowedDomains("www.amazon.com.br"),
	)
	// roda quando eu faço uma request
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("\nVisiting My %s Url", myAmazonUrl)
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.76 Safari/537.36")
	})
	// Roda quando eu recebo o html solicitado
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status Code:", r.StatusCode)

		r.Save("./pagina_amazon.html")
	})
	// Evento que faz o lifecycle com o conteudo do html
	c.OnHTML("div.a-section.a-spacing-small.puis-padding-left-small.puis-padding-right-small", func(e *colly.HTMLElement) {
		nomeProduto := e.ChildText("h2.a-size-base-plus.a-spacing-none.a-color-base.a-text-normal span")
		precoProduto := e.ChildText("span.a-price-whole")

		if precoProduto == "" {
			precoProduto = "NaN"
		}
		fmt.Println("Produto:", nomeProduto)
		fmt.Println("Preco Produto: R$", precoProduto)
		fmt.Println("#################################")
	})
	// Handler Error
	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Erro na requisição: %v\n", err)
		fmt.Printf("Status Code: %d\n", r.StatusCode)
	})
	c.Visit(myAmazonUrl)
}
