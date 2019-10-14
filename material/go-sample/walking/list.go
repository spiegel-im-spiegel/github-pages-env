package walking

import (
	"sort"
	"sync"
)

type PathList struct {
	mutex *sync.Mutex
	list  []string
	count int
}

func NewList() *PathList {
	return &PathList{mutex: &sync.Mutex{}, list: make([]string, 0, 10240), count: 0}
}

func (p *PathList) Append(path string) {
	p.mutex.Lock()
	p.list = append(p.list, path)
	p.count++
	p.mutex.Unlock()
}

func (p *PathList) List() []string {
	p.mutex.Lock()
	sort.Strings(p.list)
	p.mutex.Unlock()
	return p.list
}

func (p *PathList) Count() int {
	return p.count
}
