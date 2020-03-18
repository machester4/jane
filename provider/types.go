package provider

import (
	"github.com/machester4/jane/bktree"
	"github.com/patrickmn/go-cache"
)

type BKTree = bktree.BKTree
type Cache = cache.Cache

type Handler struct {
	storage *Cache
}