package lib

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/patrickmn/go-cache"
)

var (
	provider                           *handler
	errorNotFoundTree                  = errors.New("not provider bk-tree found")
	errorProviderHandlerNotInitialized = errors.New("provider handler not initialized")
)

// handler is a words provider handler
type handler struct {
	storage *cache.Cache
}

func getWordsFromFile(name string) []string {
	path, err := os.Getwd()
	checkError(err)

	words, err := ioutil.ReadFile(path + "/dict/" + name + ".dic")
	checkError(err)

	return strings.Split(string(words), "\n")
}

// getTree - return bk-tree from cache
func (p *handler) getTree(provider string) (bktree, error) {
	b, found := p.storage.Get(provider)
	if found == false {
		return bktree{}, errorNotFoundTree
	}
	return b.(bktree), nil
}

// getProviderHandler - return instance of provider handler
func getProviderHandler() (*handler, error) {
	if provider == nil {
		return provider, errorProviderHandlerNotInitialized
	}
	return provider, nil
}
