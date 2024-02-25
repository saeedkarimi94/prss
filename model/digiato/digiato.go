package digiato

import "encoding/xml"

type Rss struct {
	XMLName xml.Name `xml:"rss" json:"rss,omitempty"`
	Text    string   `xml:"," json:"text,omitempty"`
	Version string   `xml:"version,attr" json:"version,omitempty"`
	Content string   `xml:"content,attr" json:"content,omitempty"`
	Wfw     string   `xml:"wfw,attr" json:"wfw,omitempty"`
	Dc      string   `xml:"dc,attr" json:"dc,omitempty"`
	Atom    string   `xml:"atom,attr" json:"atom,omitempty"`
	Sy      string   `xml:"sy,attr" json:"sy,omitempty"`
	Slash   string   `xml:"slash,attr" json:"slash,omitempty"`
	Channel struct {
		Text  string `xml:"text" json:"text,omitempty"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:"text" json:"text,omitempty"`
			Href string `xml:"href,attr" json:"href,omitempty"`
			Rel  string `xml:"rel,attr" json:"rel,omitempty"`
			Type string `xml:"type,attr" json:"type,omitempty"`
		} `xml:"link" json:"link,omitempty"`
		Description     string `xml:"description"`
		LastBuildDate   string `xml:"lastBuildDate"`
		Language        string `xml:"language"`
		UpdatePeriod    string `xml:"updatePeriod"`
		UpdateFrequency string `xml:"updateFrequency"`
		Generator       string `xml:"generator"`
		Image           struct {
			Text   string `xml:"text" json:"text,omitempty"`
			URL    string `xml:"url"`
			Title  string `xml:"title"`
			Link   string `xml:"link"`
			Width  string `xml:"width"`
			Height string `xml:"height"`
		} `xml:"image" json:"image,omitempty"`
		Item []struct {
			Text      string `xml:"text" json:"text,omitempty"`
			Title     string `xml:"title"`
			Link      string `xml:"link"`
			Shortlink string `xml:"shortlink"`
			Image     struct {
				Text string `xml:"text" json:"text,omitempty"`
				URL  string `xml:"url"`
			} `xml:"image" json:"image,omitempty"`
			Comments string   `xml:"comments"`
			PubDate  string   `xml:"pubDate"`
			Creator  string   `xml:"creator"`
			Category []string `xml:"category"`
			Guid     struct {
				Text        string `xml:"text" json:"text,omitempty"`
				IsPermaLink string `xml:"isPermaLink,attr" json:"ispermalink,omitempty"`
			} `xml:"guid" json:"guid,omitempty"`
			Description string `xml:"description"`
			CommentRss  string `xml:"commentRss"`
		} `xml:"item" json:"item,omitempty"`
	} `xml:"channel" json:"digiato,omitempty"`
}
