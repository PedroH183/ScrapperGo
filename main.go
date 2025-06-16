package main

import (
	"goscrapper.com/m/scrapper"
)

func main() {
	amazon_scrapper := scrapper.AmazonScrapper{}
	amazon_scrapper.CreateCollector()

	link_teclados := "https://www.amazon.com.br/s?k=logitech+mx+keys+s&i=computers"
	amazon_scrapper.ScrapPage(link_teclados)
}
