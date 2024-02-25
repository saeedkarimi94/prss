package tasnim

import "encoding/xml"

type Rss struct {
	XMLName xml.Name `xml:"rss" json:"rss,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	A10     string   `xml:"a10,attr" json:"a10,omitempty"`
	Version string   `xml:"version,attr" json:"version,omitempty"`
	Channel struct {
		Text  string `xml:",chardata" json:"text,omitempty"`
		Media string `xml:"media,attr" json:"media,omitempty"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			Rel  string `xml:"rel,attr" json:"rel,omitempty"`
			Type string `xml:"type,attr" json:"type,omitempty"`
			Href string `xml:"href,attr" json:"href,omitempty"`
		} `xml:"link" json:"link,omitempty"`
		Description   string `xml:"description"`
		Language      string `xml:"language"`
		Copyright     string `xml:"copyright"`
		LastBuildDate string `xml:"lastBuildDate"`
		Image         struct {
			Text  string `xml:",chardata" json:"text,omitempty"`
			URL   string `xml:"url"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"image" json:"image,omitempty"`
		Item []struct {
			Text string `xml:",chardata" json:"text,omitempty"`
			Guid struct {
				Text        string `xml:",chardata" json:"text,omitempty"`
				IsPermaLink string `xml:"isPermaLink,attr" json:"ispermalink,omitempty"`
			} `xml:"guid" json:"guid,omitempty"`
			Link        string `xml:"link"`
			Category    string `xml:"category"`
			Title       string `xml:"title"`
			Description string `xml:"description"`
			PubDate     string `xml:"pubDate"`
			Thumbnail   struct {
				Text   string `xml:",chardata" json:"text,omitempty"`
				URL    string `xml:"url,attr" json:"url,omitempty"`
				Width  string `xml:"width,attr" json:"width,omitempty"`
				Height string `xml:"height,attr" json:"height,omitempty"`
			} `xml:"thumbnail" json:"thumbnail,omitempty"`
			Content struct {
				Text   string `xml:",chardata" json:"text,omitempty"`
				URL    string `xml:"url,attr" json:"url,omitempty"`
				Width  string `xml:"width,attr" json:"width,omitempty"`
				Height string `xml:"height,attr" json:"height,omitempty"`
			} `xml:"content" json:"content,omitempty"`
		} `xml:"item" json:"item,omitempty"`
	} `xml:"channel" json:"tasnim,omitempty"`
}
