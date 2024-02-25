package zoomg

import "encoding/xml"

type Rss struct {
	XMLName xml.Name `xml:"rss" json:"rss,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	Version string   `xml:"version,attr" json:"version,omitempty"`
	Channel struct {
		Text          string `xml:",chardata" json:"text,omitempty"`
		Atom          string `xml:"atom,attr" json:"atom,omitempty"`
		Title         string `xml:"title"`
		Link          string `xml:"link"`
		Description   string `xml:"description"`
		Language      string `xml:"language"`
		Copyright     string `xml:"copyright"`
		LastBuildDate string `xml:"lastBuildDate"`
		Item          []struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			Guid struct {
				Text        string `xml:",chardata" json:"text,omitempty"`
				IsPermaLink string `xml:"isPermaLink,attr" json:"ispermalink,omitempty"`
			} `xml:"guid" json:"guid,omitempty"`
			Link        string   `xml:"link"`
			Category    []string `xml:"category"`
			Title       string   `xml:"title"`
			Description string   `xml:"description"`
			PubDate     string   `xml:"pubDate"`
		} `xml:"item" json:"item,omitempty"`
	} `xml:"channel" json:"zoomg,omitempty"`
}
