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
	"github.com/robfig/cron"
)

const (
	startURL       = "http://z.thzdz.com/"
	asiaUncensored = "forum-181-1.html"
	spiderName     = "taohua"
)

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

func crawl(conf *core.SpiderConfig) error {
	defer func() {
		log.Print("Save spider config...")
		if err := core.SaveSpiderConfig(spiderName, conf); err != nil {
			log.Print(err)
		}
	}()

	u, err := mainPage()
	if err != nil {
		return err
	}

	timeR, err := core.NewTimeRange("2006-01-02")
	if err != nil {
		return err
	}

	collection, err := core.Collection(spiderName)
	if err != nil {
		return err
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

		collection.FindOneAndUpdate(nil, bson.M{
			"href": href,
		}, bson.M{
			"$set": &core.Item{
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
		// 严格小于
		if timeR.Min.Before(conf.Last) {
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
	conf.Last = timeR.Max
	return nil
}

func main() {
	conf, err := core.GetSpiderConfig(spiderName)
	if err != nil {
		panic(err)
	}

	if err = crawl(conf); err != nil {
		log.Print(err)
	}

	c := cron.New()
	c.AddFunc(conf.Cron, func() {
		if err = crawl(conf); err != nil {
			log.Print(err)
		}
	})

	c.Start()
	core.WaitForExit()
	c.Stop()
}
