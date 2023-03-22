package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.amazon.se/AMD-Ryzen-5600X-l%C3%A5da-Processor/dp/B08166SLDF/ref=sr_1_5?crid=1ZPKRAE79H5BN&keywords=amd&qid=1679516710&sprefix=amd%2Caps%2C217&sr=8-5"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://www.amazon.se/STRIX-NVIDIA-GEFORCE-GAMING-90YV0FR7-M0NA0/dp/B098324J34/ref=sr_1_5?keywords=nvidia+rtx+3070&qid=1679516766&sprefix=nvidi%2Caps%2C97&sr=8-5"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	initialLink_3 := "https://www.amazon.se/ASUS-ATX-moderkort-effektsteg-M-2-kortplatser-kylfl%C3%A4nsar/dp/B0BF5B847D/ref=sr_1_8?keywords=motherboard+x670&qid=1679516805&sprefix=motherboard+x67%2Caps%2C99&sr=8-8"
	shortLink_3 := GenerateShortLink(initialLink_3, UserId)

	assert.Equal(t, shortLink_1, "ZkqZM4ds")
	assert.Equal(t, shortLink_2, "Z3AxryUK")
	assert.Equal(t, shortLink_3, "7whHAH28")
}
