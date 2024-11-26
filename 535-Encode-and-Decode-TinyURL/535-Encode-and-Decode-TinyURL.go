package leetcode

import (
	"math/rand/v2"
	"strings"
)

type Codec struct {
	lookup        map[string]string
	reverseLookup map[string]string
}

func Constructor() Codec {
	return Codec{
    lookup: make(map[string]string),
    reverseLookup: make(map[string]string),
  }
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if val, exists := this.reverseLookup[longUrl]; exists {
		return val
	}

	shortUrl := "https://silentFellow/" + generateRandom(9)
	for {
		if _, exists := this.reverseLookup[shortUrl]; !exists {
			break
		}
		shortUrl = "https://silentFellow/" + generateRandom(9)
	}

	this.lookup[shortUrl] = longUrl
	this.reverseLookup[longUrl] = shortUrl

	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	val, exists := this.lookup[shortUrl]
	if !exists {
		return ""
	}
	return val
}

func generateRandom(length int) string {
	var random strings.Builder

	for range length {
		valTyps := rand.IntN(3)

		switch valTyps {
		case 0:
    random.WriteByte(byte(0 + rand.IntN(10)))
    case 1:
    random.WriteByte(byte('a' + rand.IntN(26)))
    case 2:
    random.WriteByte(byte('A' + rand.IntN(26)))
		}
	}

	return random.String()
}
