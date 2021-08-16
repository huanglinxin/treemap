package treemap

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	n := 100
	mymap := TreeMap{}
	records := [100]bool{}
	for i := 0; i < n; i++ {
		records[i] = false
	}
	for i := 0; i < n; i++ {
		key := rand.Int() % n
		value := key
		records[key] = true
		fmt.Println("insert "+"key "+strconv.Itoa(key), "value "+strconv.Itoa(value))
		mymap.Insert("key "+strconv.Itoa(key), "value "+strconv.Itoa(value))
	}
	fmt.Println("showall:")
	mymap.ShowAll()
	fmt.Println("findish")
	for i := 0; i < n; i++ {
		if records[i] == true && strings.Compare(mymap.GetValue("key "+strconv.Itoa(i)), "value "+strconv.Itoa(i)) != 0 {
			fmt.Println(mymap.GetValue("key "+strconv.Itoa(i)) + " : " + "value " + strconv.Itoa(i))
		}
	}
	fmt.Println("pass")
	fmt.Println("----------------------------------------------------------")
	for i := 0; i < 50; i++ {
		fmt.Println("delete "+"key "+strconv.Itoa(i), "value "+strconv.Itoa(i))
		mymap.Delete("key " + strconv.Itoa(i))
		records[i] = false
	}
	fmt.Println("showall:")
	mymap.ShowAll()
	fmt.Println("findish")
	for i := 0; i < n; i++ {
		if records[i] == true && strings.Compare(mymap.GetValue("key "+strconv.Itoa(i)), "value "+strconv.Itoa(i)) != 0 {
			fmt.Println(mymap.GetValue("key "+strconv.Itoa(i)) + " : " + "value " + strconv.Itoa(i))
		}
	}
	fmt.Println("pass")
	fmt.Println("-------------------------")
	return
}
