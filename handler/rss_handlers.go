package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/einkaaf/prss/constants"
	"github.com/einkaaf/prss/model/digiato"
	"github.com/einkaaf/prss/model/tabnak"
	"github.com/einkaaf/prss/model/tasnim"
	"github.com/einkaaf/prss/model/yjc"
	"github.com/einkaaf/prss/model/zoomg"
	"github.com/einkaaf/prss/model/zoomit"
	"github.com/gin-gonic/gin"
)

func fetchData(c *gin.Context, url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch the RSS feed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %v", err)
	}

	if err := xml.Unmarshal(responseBody, target); err != nil {
		return fmt.Errorf("invalid XML format: %v", err)
	}

	return nil
}

func ZoomitHandler(c *gin.Context) {
	var data zoomit.Rss
	if err := fetchData(c, constants.ZoomitFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to convert data to JSON")
		return
	}

	c.JSON(http.StatusOK, string(res))
}

func DigiatoHandler(c *gin.Context) {
	var data digiato.Rss
	if err := fetchData(c, constants.DigiatoFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to convert data to JSON")
		return
	}

	c.JSON(http.StatusOK, string(res))
}

func TasnimHandler(c *gin.Context) {
	var data tasnim.Rss
	if err := fetchData(c, constants.TasnimFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to convert data to JSON")
		return
	}

	c.JSON(http.StatusOK, string(res))
}

func TabnakHandler(c *gin.Context) {
	var data tabnak.Rss
	if err := fetchData(c, constants.TabnakFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to convert data to JSON")
		return
	}

	c.JSON(http.StatusOK, string(res))
}

func YJCHandler(c *gin.Context) {
	var data yjc.Rss
	if err := fetchData(c, constants.YJCFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to convert data to JSON")
		return
	}

	c.JSON(http.StatusOK, string(res))
}
func ZoomgHandler(c *gin.Context) {
	var data zoomg.Rss
	if err := fetchData(c, constants.ZoomgFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	res, err := json.Marshal(data)
	if err != nil {
		handleError(c, http.StatusInternalServerError, "Failed to convert data to JSON")
		return
	}

	c.JSON(http.StatusOK, string(res))
}

func handleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}
