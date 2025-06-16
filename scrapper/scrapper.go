package scrapper

import (
	"fmt"

	"github.com/gocolly/colly"
)

type UserAgents struct {
	UserAgentKey   string
	UserAgentValue string
}

type AmazonScrapper struct {
	c         *colly.Collector
	AgentUser UserAgents
}

func (amazon_scrapper *AmazonScrapper) CreateCollector() {
	amazon_scrapper.c = colly.NewCollector(
		colly.AllowedDomains("www.amazon.com.br"),
	)
	amazon_scrapper.AgentUser.UserAgentKey = "User-Agent"
	amazon_scrapper.AgentUser.UserAgentValue = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.6422.76 Safari/537.36"
}
func (amazon_scrapper *AmazonScrapper) ScrapPage(target_url string) {

	amazon_scrapper.c.OnRequest(func(r *colly.Request) {
		fmt.Printf("\nVisiting My %s Url", target_url)
		r.Headers.Set(amazon_scrapper.AgentUser.UserAgentKey, amazon_scrapper.AgentUser.UserAgentValue)
	})
	amazon_scrapper.c.OnResponse(func(r *colly.Response) {
		fmt.Println("Status Code:", r.StatusCode)
		r.Save("./pagina_amazon_capturada.html")
	})

	amazon_scrapper.c.OnHTML("div.a-section.a-spacing-small.puis-padding-left-small.puis-padding-right-small", func(e *colly.HTMLElement) {
		nomeProduto := e.ChildText("h2.a-size-base-plus.a-spacing-none.a-color-base.a-text-normal span")
		precoProduto := e.ChildText("span.a-price-whole")

		if precoProduto == "" {
			precoProduto = "NaN"
		}
		fmt.Println("Produto:", nomeProduto)
		fmt.Println("Preco Produto: R$", precoProduto)
		fmt.Println("#################################")
	})

	amazon_scrapper.c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Erro na requisição: %v\n", err)
		fmt.Printf("Status Code: %d\n", r.StatusCode)
	})
	amazon_scrapper.c.Visit(target_url)
}
