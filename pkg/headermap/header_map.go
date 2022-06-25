package headermap

import (
	"net/http"
)

type HeaderMap struct {
	keyMap map[string]string
	ctx    http.Header
}

func NewHeaderMap() HeaderMap {
	return HeaderMap{
		keyMap: make(map[string]string),
		ctx:    make(http.Header),
	}
}

func (hm HeaderMap) Put(key, value string) HeaderMap {
	hm.keyMap[key] = ""
	hm.ctx.Set(key, value)

	return hm
}

func (hm HeaderMap) Get(key string) (string, bool) {
	if ok := hm.Has(key); ok {
		return hm.ctx.Get(key), true
	}

	return "", false
}

func (hm HeaderMap) Add(key, value string) HeaderMap {
	if ok := hm.Has(key); ok {
		hm.ctx.Add(key, value)
	} else {
		hm.Put(key, value)
	}

	return hm
}

func (hm HeaderMap) Remove(key string) HeaderMap {
	if ok := hm.Has(key); ok {
		hm.ctx.Del(key)
		delete(hm.keyMap, key)
	}

	return hm
}

func (hm HeaderMap) Has(key string) bool {
	_, ok := hm.keyMap[key]

	return ok
}

func (hm HeaderMap) Values() http.Header {
	return hm.ctx
}

func (hm HeaderMap) Clean() HeaderMap {
	for k := range hm.keyMap {
		hm.Remove(k)
	}

	return hm
}

func (hm HeaderMap) Length() int {
	return len(hm.keyMap)
}
