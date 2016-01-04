package crawler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Fruit struct {
	Title     string `json:"title"`
	Size      string `json:"date"`
	UnitPrice string `json:"unit_price"`
	//Description string `json:"description"`
}

type Crawler struct {
	Fruits []Fruit
	Url    string
	Doc    *goquery.Document
	Total  float64
}

var (
	tagTitle     string = ".productTitleDescriptionContainer h1"
	tagUnitPrice string = ".pricePerUnit"
	tagProduct   string = ".productInfo"
	//tagDescription string = ""
)

// Save the fruit in the slice os fruits
func (c *Crawler) SaveFruit(i int, s *goquery.Selection) {

	// Link to the fruit/product
	link, _ := s.Find("h3 a").Attr("href")
	title := GetFruitInfo(link, tagTitle)
	size := IntToString(UrlSize(link))

	// String with garbage (filtering)
	unitPrice := GetFruitInfo(link, tagUnitPrice)[3:7]

	c.Total += StringToFloat(unitPrice)

	fruit := Fruit{title, size, unitPrice}
	c.Fruits = append(c.Fruits, fruit)
}

func GetFruitInfo(link string, tag string) string {

	doc, err := goquery.NewDocument(link)
	if err != nil {
		log.Fatal(err)
	}

	info := ""
	doc.Find(tag).Each(func(i int, s *goquery.Selection) {
		info = s.Text()
	})

	return info
}

// Get ao posts in the main page
func (c *Crawler) Start() {
	c.Doc.Find(tagProduct).Each(c.SaveFruit)
}

// Return the URL size
func UrlSize(url string) int {

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return len(body) / 1024
}

// Convert int to string and add in the end "kb"
func IntToString(num int) string {
	return strconv.Itoa(num) + "kb"
}

// Initialize Crawler struct
func (c *Crawler) Init() {

	var err error
	c.Doc, err = goquery.NewDocument(c.Url)
	if err != nil {
		log.Fatal(err)
	}
}

// Create new Crawler object
func New(url string) *Crawler {

	c := &Crawler{}
	c.Fruits = make([]Fruit, 0)
	c.Url = url
	c.Doc = nil

	return c
}

func StringToFloat(num string) float64 {
	result, err := strconv.ParseFloat(num, 64)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

// Print the JSON data
func (c *Crawler) Print() {

	results := map[string]interface{}{
		"results": c.Fruits,
		"total":   c.Total,
	}

	b, err := json.MarshalIndent(results, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", b)
}
