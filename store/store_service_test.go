package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	_store := InitializeStore()
	testStoreService = _store
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.amazon.se/s?k=www+ama&language=sv_SE&adgrpid=117754408488&hvadid=602638401989&hvdev=c&hvlocphy=9062383&hvnetw=g&hvqmt=b&hvrand=16209796131505399838&hvtargid=kwd-308989583418&hydadcr=15412_2297153&tag=textstdgledes-21&ref=pd_sl_3tvmga4j5y_b"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	SaveUrl(shortURL, initialLink, userUUId)

	assert.Equal(t, initialLink, RetrieveInitialUrl(shortURL))
}
