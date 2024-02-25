package tabnak

import "encoding/xml"

type Rss struct {
	XMLName xml.Name `xml:"rss" json:"rss,omitempty"`
	Text    string   `xml:",chardata" json:"text,omitempty"`
	Version string   `xml:"version,attr" json:"version,omitempty"`
	Channel struct {
		Text           string `xml:",chardata" json:"text,omitempty"`
		Title          string `xml:"title"`
		Link           string `xml:"link"`
		Description    string `xml:"description"`
		ManagingEditor string `xml:"managingEditor"`
		WebMaster      string `xml:"webMaster"`
		LastBuildDate  string `xml:"lastBuildDate"`
		Generator      string `xml:"generator"`
		Item           []struct {
			Text        string `xml:",chardata" json:"text,omitempty"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			Description string `xml:"description"`
			Author      string `xml:"author"`
			Enclosure   struct {
				Text   string `xml:",chardata" json:"text,omitempty"`
				URL    string `xml:"url,attr" json:"url,omitempty"`
				Type   string `xml:"type,attr" json:"type,omitempty"`
				Length string `xml:"length,attr" json:"length,omitempty"`
			} `xml:"enclosure" json:"enclosure,omitempty"`
			PubDate string `xml:"pubDate"`
		} `xml:"item" json:"item,omitempty"`
	} `xml:"channel" json:"tabnak,omitempty"`
}
