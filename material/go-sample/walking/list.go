package walking

import (
	"sort"
	"sync"
)

type PathList struct {
	mutex *sync.Mutex
	list  []string
}

func NewList() *PathList {
	return &PathList{mutex: &sync.Mutex{}, list: make([]string, 0, 10240)}
}

func (p *PathList) Init() {
	p.mutex.Lock()
	p.list = p.list[:]
	p.mutex.Unlock()
}

func (p *PathList) Append(path string) {
	p.mutex.Lock()
	p.list = append(p.list, path)
	p.mutex.Unlock()
}

func (p *PathList) List() []string {
	p.mutex.Lock()
	list := make([]string, len(p.list))
	copy(list, p.list)
	sort.Strings(list)
	p.mutex.Unlock()
	return list
}

func (p *PathList) Count() int {
	return len(p.List())
}
