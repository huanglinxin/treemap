package treemap

import (
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	mymap := TreeMap{}
	for i := 0; i < 100; i++ {
		mymap.Insert("key "+strconv.Itoa(i), "value "+strconv.Itoa(i))
	}
	mymap.ShowAll()
	return
}
