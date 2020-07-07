package ShanMap

import (
	"fmt"
	"github.com/tansoz/ShanObject"
	"sync"
)

type SafeMap struct {
	ShanMap
	sync.RWMutex
	data map[interface{}]ShanObject.Object
}

func NewSafeMap() ShanMap {
	return &SafeMap{
		data: make(map[interface{}]ShanObject.Object),
	}
}

func (this *SafeMap) Get(name interface{}) ShanObject.Object {
	this.Lock()
	tmp := this.data[name]
	if tmp == nil {
		tmp = ShanObject.NewObject(nil)
	}
	this.Unlock()
	return tmp
}

func (this *SafeMap) Set(name interface{}, v interface{}) {
	this.Lock()
	this.data[name] = ShanObject.NewObject(v)
	this.Unlock()
}
func (this *SafeMap) Del(name interface{}) {
	this.Lock()
	delete(this.data, name)
	this.Unlock()
}
func (this *SafeMap) String() string {
	return fmt.Sprint(this.data)
}
func (this *SafeMap) Len() int {
	this.Lock()
	l := len(this.data)
	this.Unlock()
	return l
}
func (this *SafeMap) Map() map[interface{}]ShanObject.Object {
	return this.data
}
