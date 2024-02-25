package carir

import "time"

type CarIR struct {
	Items []struct {
		Brand struct {
			ID                int    `json:"id,omitempty"`
			TitleFa           string `json:"titleFa,omitempty"`
			TitleEn           string `json:"titleEn,omitempty"`
			Slug              string `json:"slug,omitempty"`
			Logo              string `json:"logo,omitempty"`
			PriceListOrdering int    `json:"priceListOrdering,omitempty"`
		} `json:"brand,omitempty"`
		Items []struct {
			Tip struct {
				ID           int    `json:"id,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayTitle string `json:"displayTitle,omitempty"`
			} `json:"tip,omitempty"`
			Brand struct {
				ID      int    `json:"id,omitempty"`
				TitleFa string `json:"titleFa,omitempty"`
			} `json:"brand,omitempty"`
			BodyClass struct {
				ID    int    `json:"id,omitempty"`
				Title string `json:"title,omitempty"`
			} `json:"bodyClass,omitempty"`
			MarketPrice       int `json:"marketPrice,omitempty"`
			CompanyPrice      int `json:"companyPrice,omitempty"`
			MarketPriceStatus struct {
				ID    int    `json:"id,omitempty"`
				Title string `json:"title,omitempty"`
			} `json:"marketPriceStatus,omitempty"`
			CompanyPriceStatus struct {
				ID    int    `json:"id,omitempty"`
				Title string `json:"title,omitempty"`
			} `json:"companyPriceStatus,omitempty"`
			UpdateDate time.Time `json:"updateDate,omitempty"`
		} `json:"items,omitempty"`
	} `json:"items,omitempty"`
}
