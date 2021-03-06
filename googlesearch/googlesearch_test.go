package googlesearch_test

import (
	"fmt"
	"github.com/sundowndev/dorkgen/googlesearch"
	"net/url"
	"testing"

	assertion "github.com/stretchr/testify/assert"
)

var dork *googlesearch.GoogleSearch

func TestInit(t *testing.T) {
	assert := assertion.New(t)

	t.Run("should convert to URL correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Site("example.com").
			URL()

		assert.Equal(result, "https://www.google.com/search?q=site%3Aexample.com", "they should be equal")
	})

	t.Run("should convert to string correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := fmt.Sprint(dork.Site("example.com"))

		assert.Equal(result, "site:example.com", "they should be equal")
	})

	t.Run("should handle site tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Site("example.com").
			String()

		assert.Equal(result, "site:example.com", "they should be equal")
	})

	t.Run("should handle intext tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			InText("text").
			String()

		assert.Equal(result, "intext:\"text\"", "they should be equal")
	})

	t.Run("should handle inurl tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			InURL("index.php").
			String()

		assert.Equal(result, "inurl:\"index.php\"", "they should be equal")
	})

	t.Run("should handle filetype tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			FileType("pdf").
			String()

		assert.Equal(result, "filetype:\"pdf\"", "they should be equal")
	})

	t.Run("should handle cache tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Cache("www.google.com").
			String()

		assert.Equal(result, "cache:\"www.google.com\"", "they should be equal")
	})

	t.Run("should handle related tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Related("www.google.com").
			String()

		assert.Equal(result, "related:\"www.google.com\"", "they should be equal")
	})

	t.Run("should handle ext tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Ext("(doc | pdf | xls | txt | xml)").
			String()

		assert.Equal(result, "ext:(doc | pdf | xls | txt | xml)", "they should be equal")
	})

	t.Run("should handle exclude tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Exclude(googlesearch.New().Plain("html")).
			Exclude(googlesearch.New().Plain("htm")).
			Exclude(googlesearch.New().Plain("php")).
			Exclude(googlesearch.New().Plain("md5sums")).
			String()

		assert.Equal(result, "-html -htm -php -md5sums", "they should be equal")
	})

	t.Run("should handle 'OR' tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Site("facebook.com").
			Or().
			Site("twitter.com").
			String()

		assert.Equal(result, "site:facebook.com | site:twitter.com", "they should be equal")
	})

	t.Run("should handle 'AND' tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			InTitle("facebook").
			And().
			InTitle("twitter").
			String()

		assert.Equal(result, "intitle:\"facebook\" + intitle:\"twitter\"", "they should be equal")
	})

	t.Run("should handle group tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Site("linkedin.com").
			Group(googlesearch.New().InText("1").Or().InText("2")).
			String()

		assert.Equal(result, "site:linkedin.com (intext:\"1\" | intext:\"2\")", "they should be equal")
	})

	t.Run("should handle group tag correctly", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Site("linkedin.com").
			Group(googlesearch.New().InText("1").Or().InText("2")).
			InTitle("jordan").
			String()

		assert.Equal(result, "site:linkedin.com (intext:\"1\" | intext:\"2\") intitle:\"jordan\"", "they should be equal")
	})

	t.Run("should return URL values", func(t *testing.T) {
		dork = googlesearch.New()

		result := dork.
			Site("linkedin.com").
			Group(googlesearch.New().InText("1").Or().InText("2")).
			InTitle("jordan").
			QueryValues()

		assert.Equal(url.Values{
			"q": []string{"site:linkedin.com (intext:\"1\" | intext:\"2\") intitle:\"jordan\""},
		}, result, "they should be equal")
	})
}
