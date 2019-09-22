package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
	"github.com/popstk/olddriver/core"
)

const (
	startURL   = "http://taohuale.us/"
	spiderName = "taohua"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func mainPage() (*url.URL, error) {
	var u string
	c := colly.NewCollector()
	c.OnHTML("div.main div:nth-child(3) #newurllink a", func(e *colly.HTMLElement) {
		if u = strings.TrimSpace(e.Attr("href")); u == "" {
			log.Println("Save main page")
			_ = e.Response.Save("mainPage.html")
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Request URL %s failed with response: %s", r.Request.URL, err)
	})

	_ = c.Visit(startURL)

	if u == "" {
		return nil, errors.New("can not get main page")
	}

	log.Println("->", u)

	rsp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	return rsp.Request.URL, nil
}

func main() {
	flag.Parse()

	persist, err := core.NewPersist(spiderName)
	core.Must(err)

	conf, err := persist.Conf()
	core.Must(err)

	u, err := mainPage()
	core.Must(err)

	timeR, err := core.NewTimeRange("2006-1-2")
	core.Must(err)

	u.Path = conf.Forum
	log.Println("->", u.String())
	c := colly.NewCollector(colly.AllowedDomains(u.Hostname()))

	cc := c.Clone()
	cc.OnXML(`//*[@id="postlist"]/div[1]//p[@class="attnm"]`, func(e *colly.XMLElement) {
		torrent := e.ChildAttrs("./a", "href")
		for i, v := range torrent {
			u, err := url.Parse(core.JoinURL(e.Request.URL, v))
			if err != nil {
				continue
			}

			torrent[i] = fmt.Sprintf("%s://%s/forum.php?mod=attachment&aid=%s",
				u.Scheme, u.Host, u.Query().Get("aid"))
		}
		e.Request.Ctx.Put("torrent", torrent)
	})

	// note: css selector 从第二页开始查找不到，改用xpath
	c.OnXML(`//*[@id="threadlisttableid"]/tbody[@id="separatorline"]/following-sibling::tbody/tr`, func(e *colly.XMLElement) {
		title := e.ChildText(".//th/a[2]")
		if title == "" {
			return
		}

		href := e.ChildAttr(".//th/a[2]", "href")
		href = core.JoinURL(e.Request.URL, href)

		timeStr := e.ChildAttr(".//td[2]/em/span/span", "title")
		if timeStr == "" {
			timeStr = e.ChildText(".//td[2]/em/span")
		}

		t, err := timeR.AddTime(timeStr)
		if err != nil {
			log.Println(err)
			return
		}

		var torrent []string
		if err := cc.Request("GET", href, nil, e.Request.Ctx, nil); err != nil {
			log.Println(err)
			return
		}

		torrent, _ = e.Request.Ctx.GetAny("torrent").([]string)
		item := core.Item{
			Tag:   spiderName,
			URL:   href,
			Time:  t,
			Title: title,
			Link:  torrent,
		}

		log.Println("get item: ", item.URL)

		if err := persist.Insert(href, &item); err != nil {
			log.Println(err)
		}
	})

	c.OnHTML("#fd_page_bottom > div > a.nxt", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		next := e.Request.AbsoluteURL(link)
		e.Request.Ctx.Put("next", next)
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println("min time is ", timeR.Min)
		log.Println("conf.Last time is ", conf.Last.Time())

		if timeR.Min.Before(conf.Last.Time()) {
			log.Println("DeadLine")
			return
		}

		if next, ok := r.Ctx.GetAny("next").(string); ok {
			if err = c.Visit(next); err != nil {
				log.Println(err)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting ", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Print("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	if err = c.Visit(u.String()); err != nil {
		log.Println(err)
	}

	conf.Last.Set(timeR.Max)

	if err = persist.SaveConf(conf); err != nil {
		log.Println(err)
	}
}
