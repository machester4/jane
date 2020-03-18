package provider

import (
	"errors"
	"github.com/machester4/jane/bktree"
	"github.com/machester4/jane/helpers"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	provider *Handler
	once sync.Once
	ErrorNotFoundTree = errors.New("not provider bk-tree found")
	ErrorProviderHandlerNotInitialized = errors.New("provider handler not initialized")
)

func getWordsFromFile(name string) []string {
	path, err := os.Getwd()
	helpers.CheckError(err)

	words, err := ioutil.ReadFile(path + "/dict/" + name + ".dic")
	helpers.CheckError(err)

	return strings.Split(string(words), "\n")
}

func (p *Handler) GetTree(provider string) BKTree {
	b, found := p.storage.Get(provider)
	if !found {
		helpers.CheckError(ErrorNotFoundTree)
	}
	return b.(BKTree)
}

func GetHandler() *Handler {
	if provider == nil {
		helpers.CheckError(ErrorProviderHandlerNotInitialized)
	}
	return provider
}

func CreateHandler(providers ...string) {
	// Create handler instance
	once.Do(func() {
		// Create cache storage for BK-TREES
		provider = &Handler{storage: cache.New(5*time.Minute, 10*time.Minute)}

		// Get all word from providers and create BK-TREES
		for _, p := range providers {
			var b bktree.BKTree
			for _, w := range getWordsFromFile(p) {
				b.Add(bktree.Word(w))
			}
			provider.storage.Set(p, b, cache.NoExpiration)
		}
	})
}