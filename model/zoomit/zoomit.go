package zoomit

import "encoding/xml"

type Rss struct {
	XMLName xml.Name `xml:"rss" json:"rss,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	Dc      string   `xml:"dc,attr" json:"dc,omitempty"`
	Content string   `xml:"content,attr" json:"content,omitempty"`
	Version string   `xml:"version,attr" json:"version,omitempty"`
	Channel struct {
		Text        string `xml:",chardata" json:"text,omitempty"`
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Image       struct {
			Text   string `xml:",chardata" json:"text,omitempty"`
			URL    string `xml:"url"`
			Title  string `xml:"title"`
			Link   string `xml:"link"`
			Width  string `xml:"width"`
			Height string `xml:"height"`
		} `xml:"image" json:"image,omitempty"`
		Copyright     string `xml:"copyright"`
		LastBuildDate string `xml:"lastBuildDate"`
		Item          []struct {
			Text        string   `xml:",chardata" json:"text,omitempty"`
			Title       string   `xml:"title"`
			PubDate     string   `xml:"pubDate"`
			Link        string   `xml:"link"`
			Guid        string   `xml:"guid"`
			Description string   `xml:"description"`
			Category    []string `xml:"category"`
		} `xml:"item" json:"item,omitempty"`
	} `xml:"channel" json:"zoomit,omitempty"`
}
