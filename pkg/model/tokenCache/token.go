package tokencache

import (
	"fmt"
	"time"
)

type BlacklistedToken struct {
	tokens map[string]int64
}

func (b *BlacklistedToken) Set(token string, expTime int64) {
	b.tokens[token] = expTime
}
func (b *BlacklistedToken) Remove(token string) {
	delete(b.tokens, token)
}
func (b *BlacklistedToken) IsPresent(token string) bool {
	_, ok := b.tokens[token]
	return ok
}
func (b *BlacklistedToken) GetExpTime(token string) (int64, error) {
	expTime, ok := b.tokens[token]
	if !ok {
		return 0, fmt.Errorf("unable to find token on cache")
	}
	return expTime, nil
}

func (b *BlacklistedToken) Clean() {
	for key, value := range b.tokens {
		if value < (time.Now().Unix()) {
			delete(b.tokens, key)
		}
	}
}

var blacklistedTokenCache *BlacklistedToken

func GetBlacklistTokenCache() *BlacklistedToken {
	if blacklistedTokenCache != nil {
		blacklistedTokenCache.Clean()
		return blacklistedTokenCache
	}
	blacklistedTokenCache = &BlacklistedToken{
		tokens: make(map[string]int64),
	}
	return blacklistedTokenCache
}
