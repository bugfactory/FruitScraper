package main

import "github.com/patito/FruitScraper/crawler"

func main() {

	url := "http://hiring-tests.s3-website-eu-west-1.amazonaws.com/2015_Developer_Scrape/5_products.html"
	c := crawler.New(url)

	c.Init()
	c.Start()
	c.Print()
}
