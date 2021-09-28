package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndFetchOriginalUrl(t *testing.T) {
	longUrl := "https://covid-19.ontario.ca/proof-covid-19-vaccination"
	userID := "abcde"
	shortUrl := "cajkdcjqojcoa"

	SaveUrlMapping(shortUrl, longUrl, userID)

	fetchedLongUrl := FetchOriginalUrl(shortUrl, userID)

	assert.Equal(t, longUrl, fetchedLongUrl)
}
