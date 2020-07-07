package ShanMap

import "github.com/tansoz/ShanObject"

type ShanMap interface {
	Get(name interface{}) ShanObject.Object
	Set(name interface{}, v interface{})
	Del(name interface{})
	Len() int
	Map() map[interface{}]ShanObject.Object
}
