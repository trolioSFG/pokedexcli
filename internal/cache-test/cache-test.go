package main

import (
	"fmt"
	"github.com/trolioSFG/pokecache"
	"time"
)

func main() {
	c := pokecache.NewCache(5 * time.Second)

	/***
	var data [5][]byte
	data[0] = []byte("uno")
	data[1] = []byte("dos")
	data[2] = []byte("tres")
	data[3] = []byte("cuatro")
	data[4] = []byte("cinco")
	***/
	data := [][]byte{
		[]byte("uno"),
		[]byte("dos"),
		[]byte("tres"),
		[]byte("cuatro"),
		[]byte("cinco"),
	}

	// index := [5]string{"u", "d", "t", "c", "x"}
	index := []string{"u", "d", "t", "c", "x"}


	for i, v := range data {
		time.Sleep(2 * time.Second)
		c.Add(index[i],v)
		fmt.Println(string(v))
	}

	for i, v := range data {

		if val, found := c.Get(index[i]); found {
			fmt.Printf("%s %v\n", index[i], string(val))
		} else {
			fmt.Printf("%s Not found, should be %s\n", index[i], string(v))
		}
	}
	
	time.Sleep(5 * time.Second)

	for i, v := range data {

		if val, found := c.Get(index[i]); found {
			fmt.Printf("%s %v\n", index[i], string(val))
		} else {
			fmt.Printf("%s Not found, should be %s\n", index[i], string(v))
		}
	}


}

