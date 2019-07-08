package main

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/popstk/olddriver/core"
	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	startURL   = "http://taohuale.us/"
	spiderName = "taohua"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

func mainPage() (*url.URL, error) {
	var u string
	c := colly.NewCollector()
	c.OnHTML("body > div > div.categorythr > div:nth-child(3) > ul > li:nth-child(1) > a", func(e *colly.HTMLElement) {
		u = e.Attr("href")
		if u == "" {
			log.Print("Save main page")
			e.Response.Save("mainpage.txt")
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Print("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
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

	u.Path = conf.Tag
	c.Visit(u.String())
	conf.Last = timeR.Max
	return nil
}

func main() {
	confs, err := core.GetSpiderConfig(spiderName)
	if err == core.ErrNoDocuments {
		log.Print("No Documents for ", spiderName)
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	c := cron.New()
	for _, conf := range confs {
		log.Print("Add Tag ", conf.Tag, ", cron is ", conf.Cron)
		_ = c.AddFunc(conf.Cron, func() {
			if err = crawl(conf); err != nil {
				log.Print(err)
			}
		})
	}

	c.Start()
	core.WaitForExit()
	c.Stop()
}
