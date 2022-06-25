package collection

// NewStringMap new StringMap
func NewStringMap(cap ...int) StringMap {
	switch len(cap) {
	case 1:
		if cap[0] > 0 {
			return make(StringMap, cap[0])
		}
	}

	return make(StringMap)
}

// InitStringMap init StringMap by given map
func InitStringMap(ctx map[string]string) StringMap {
	sm := make(StringMap)
	if ctx != nil {
		for k, v := range ctx {
			sm[k] = v
		}
	}

	return sm
}

func (sm StringMap) Put(key, value string) StringMap {
	sm[key] = value

	return sm
}

func (sm StringMap) Get(key string) (string, bool) {
	value, ok := sm[key]

	return value, ok
}

func (sm StringMap) Remove(key string) {
	if sm.Has(key) {
		delete(sm, key)
	}
}

func (sm StringMap) Has(key string) bool {
	_, ok := sm[key]

	return ok
}

func (sm StringMap) Size() int {
	return len(sm)
}
