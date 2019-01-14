package main

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"github.com/popstk/olddriver/core"
)

const startURL = "http://z.thzdz.com/"
const asiaUncensored = "forum-181-1.html"

func mainPage() (*url.URL, error) {
	var u string
	c := colly.NewCollector()
	c.OnHTML("body > div.main > div:nth-child(3) #newurllink > a", func(e *colly.HTMLElement) {
		u = e.Attr("href")
	})
	c.Visit(startURL)

	if u == "" {
		return nil, errors.New("Can not get main page")
	}

	rsp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	return rsp.Request.URL, nil
}

func main() {
	u, err := mainPage()
	if err != nil {
		panic(err)
	}

	timeR, err := core.NewTimeRange("2006-01-02", "")
	if err != nil {
		panic(err)
	}

	lasttime, err := core.GetLastTime("taohua")
	if err != nil {
		log.Printf("%T ", err)
		panic(err)
	}

	collection, err := core.Collection("taohua")
	if err != nil {
		panic(err)
	}

	opt := options.FindOneAndUpdate()
	opt.SetUpsert(true)

	log.Print("Main page is ", u.String())
	next := ""

	c := colly.NewCollector(
		colly.AllowedDomains(u.Hostname()),
	)

	c.OnHTML("#threadlisttableid tbody tr th > a:nth-child(3)", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if href == "" {
			log.Print("No href: ", e.Name)
			return
		}

		parts := strings.Split(e.Text, " ")
		if len(parts) == 0 {
			log.Print("Invalid text: ", e.Text)
			return
		}

		t, err := time.Parse(timeR.Layout, parts[0])
		if err != nil {
			log.Print("Invalid time: ", e.Text)
			return
		}

		timeR.Add(t)

		collection.FindOneAndUpdate(nil, bson.M{"href": href}, bson.M{"$set": &core.Item{
			Href:  href,
			Title: e.Text,
			Time:  t,
		}}, opt)

		log.Print(e.Text)
	})

	c.OnHTML("#fd_page_bottom > div > a.nxt", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		next = e.Request.AbsoluteURL(link)
	})

	c.OnScraped(func(r *colly.Response) {
		if timeR.BeforeMin(lasttime) {
			log.Print("DeadLine")
			return
		}

		c.Visit(next)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Print("Visiting ", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Print("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	u.Path = asiaUncensored
	c.Visit(u.String())
}
