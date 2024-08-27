package shortener

import (
	"math/rand"
	"urlshortner/internal/storage"
)

type Shortener struct {
	storage *storage.PostgresStorage
}

func NewShortener(storage *storage.PostgresStorage) *Shortener {
	return &Shortener{storage: storage}
}

func (s *Shortener) Shorten(url string) (string, error) {
	shortURL := generateShortURL()
	err := s.storage.Save(shortURL, url)
	if err != nil {
		return "", err
	}
	return shortURL, nil
}

func (s *Shortener) Resolve(shortURL string) (string, error) {
	return s.storage.Load(shortURL)
}

func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortURL := make([]byte, 6)
	for i := range shortURL {
		shortURL[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortURL)
}
