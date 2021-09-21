package scraper

import (
	"better-av-tool/pkg/archive"
	myclient "better-av-tool/pkg/client"

	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DefaultScraper struct {
	doc       *goquery.Document
	cookie    *http.Cookie
	isArchive bool
}

func (DefaultScraper) FetchDoc(query string) (err error) {
	return nil
}

func (DefaultScraper) GetPlot() string {
	return ""
}

func (DefaultScraper) GetTitle() string {
	return ""
}

func (DefaultScraper) GetDirector() string {
	return ""
}

func (DefaultScraper) GetRuntime() string {
	return ""
}

func (DefaultScraper) GetTags() []string {
	return []string{}
}

func (DefaultScraper) GetMaker() string {
	return ""
}

func (DefaultScraper) GetActors() []string {
	return []string{}
}

func (DefaultScraper) GetLabel() string {
	return ""
}

func (DefaultScraper) GetNumber() string {
	return ""
}

func (DefaultScraper) GetFormatNumber() string {
	return ""
}

func (DefaultScraper) GetCover() string {
	return ""
}

func (s *DefaultScraper) GetWebsite() string {
	if s.doc == nil {
		return ""
	}
	return s.doc.Url.String()
}

func (DefaultScraper) GetPremiered() string {
	return ""
}

func (DefaultScraper) GetYear() string {
	return ""
}

func (DefaultScraper) GetSeries() string {
	return ""
}

func (DefaultScraper) GetType() string {
	return ""
}

func (DefaultScraper) NeedCut() bool {
	return false
}

func (s *DefaultScraper) GetDocFromURL(u string) (err error) {
	log.Infof("fetching %s", u)
	if s.cookie == nil {
		s.cookie = &http.Cookie{}
	}
	res, err := client.Get(u, s.cookie)
	if err != nil {
		return err
	}

	utfBody, _ := myclient.DecodeHTMLBody(res.Body, "")
	s.doc, err = goquery.NewDocumentFromReader(utfBody)
	if err != nil {
		return err
	}
	s.doc.Url = res.Request.URL
	return nil
}

// Download ...
func Download(url, filename string, progress func(current, total int64)) error {
	return client.Download(url, filename, progress)
}

func (s *DefaultScraper) GetAvailableUrl(orginUrl string) (string, error) {

	resp := &archive.AvailableResp{}
	err := client.GetJSON(fmt.Sprintf(archive.GetAvailableUrl, orginUrl), resp)
	if err != nil {
		return "", err
	}

	return resp.ArchivedSnapshots.Closest.URL, nil
}

// GetOutputPath ...
func GetOutputPath(s Scraper, conf string) string {
	p := strings.Replace(conf, "{year}", s.GetYear(), 1)
	if len(s.GetActors()) > 0 {
		p = strings.Replace(p, "{actor}", s.GetActors()[0], 1)
	} else {
		p = strings.Replace(p, "{actor}", "", 1)
	}
	p = strings.Replace(p, "{maker}", s.GetMaker(), 1)
	p = strings.Replace(p, "{num}", s.GetNumber(), 1)

	return strings.Replace(p, "//", "/", -1)
}