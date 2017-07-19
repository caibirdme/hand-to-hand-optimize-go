package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	http.HandleFunc("/test", handler)
	log.Fatal(http.ListenAndServe(":9876", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if nil != err {
		w.Write([]byte(err.Error()))
		return
	}
	log.Println(r.Form)
	doSomeThingOne(10000)
	buff := genSomeBytes()
	b, err := ioutil.ReadAll(buff)
	if nil != err {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(b)
}

func doSomeThingOne(times int) {
	var inner = int(math.Log2(float64(times)))
	for i := 0; i < times; i++ {
		for j := 0; j < inner; j++ {

		}
	}
}

func genSomeBytes() *bytes.Buffer {
	var buff bytes.Buffer
	for i := 1; i < 20000; i++ {
		buff.Write([]byte{'0' + byte(i%10)})
	}
	return &buff
}
