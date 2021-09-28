package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const userID = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortUrlGenerator(t *testing.T) {
	longUrl_1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortUrl_1 := GenerateShortUrl(longUrl_1, UserID)

	longUrl_2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortUrl_2 := GenerateShortUrl(longUrl_2, UserID)

	longUrl_3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortUrl_3 := GenerateShortUrl(longUrl_3, UserID)

	assert.Equal(t, shortUrl_1, "jTa4L57P")
	assert.Equal(t, shortUrl_2, "d66yfx7N")
	assert.Equal(t, shortUrl_3, "dhZTayYQ")
}
