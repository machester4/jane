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

func (p *handler) getTree(provider string) bktree {
	b, found := p.storage.Get(provider)
	if !found {
		checkError(errorNotFoundTree)
	}
	return b.(bktree)
}

func getProviderHandler() *handler {
	if provider == nil {
		checkError(errorProviderHandlerNotInitialized)
	}
	return provider
}
