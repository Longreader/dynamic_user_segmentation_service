package repository

import (
	"sync"
)

type Pull struct {
	sync.Mutex // mutex for lock
	pull       map[string]struct{}
}

func New() *Pull {

	return &Pull{
		pull: make(map[string]struct{}),
	}
}

func (st *Pull) set(key string) {
	st.pull[key] = struct{}{}
}

func (st *Pull) get(key string) bool {
	if st.count() > 0 {
		_, ok := st.pull[key]
		if !ok {
			return false
		}
		return ok
	}
	return false
}

func (st *Pull) delete(key string) {
	delete(st.pull, key)
}

func (st *Pull) count() int {
	return len(st.pull)
}

func (st *Pull) Set(key string) {
	st.Lock()
	defer st.Unlock()
	st.set(key)
}

func (st *Pull) Get(key string) bool {
	st.Lock()
	defer st.Unlock()
	return st.get(key)
}

func (st *Pull) Delete(key string) {
	st.Lock()
	defer st.Unlock()
	st.delete(key)
}

func (st *Pull) Count() int {
	st.Lock()
	defer st.Unlock()
	return st.count()
}
