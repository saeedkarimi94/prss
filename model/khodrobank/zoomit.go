package khodrobank

import "encoding/xml"

type Rss struct {
	XMLName xml.Name `xml:"rss" json:"rss,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	Version string   `xml:"version,attr" json:"version,omitempty"`
	Channel struct {
		Text string `xml:",chardata" json:"text,omitempty"`
		Item []struct {
			Text        string `xml:",chardata" json:"text,omitempty"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
			PubDate     string `xml:"pubDate"`
		} `xml:"item" json:"item,omitempty"`
	} `xml:"channel" json:"khodrobank,omitempty"`
}
