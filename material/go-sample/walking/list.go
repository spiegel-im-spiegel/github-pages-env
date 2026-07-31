package walking

import (
	"sync"
)

type PathList struct {
	mutex *sync.Mutex
	list  []string
}

func New() *PathList {
	return &PathList{mutex: &sync.Mutex{}, list: make([]string, 0, 10240)}
}

func (p *PathList) Init() {
	p.mutex.Lock()
	p.list = p.list[:0]
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
	p.mutex.Unlock()
	return list
}
