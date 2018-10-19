/**
 * Auth :   liubo
 * Date :   2018/10/19 17:05
 * Comment: 简单的web服务器
 */

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	w.Write([]byte("hello"))
}
func randvalue(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")


	w.Write([]byte(fmt.Sprint(rand.Intn(100))))
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("start.")

	// realPath = flag.String("path", "", "static resource path")
	//flag.Parse()
	http.HandleFunc("/ajax/hello", hello)
	http.HandleFunc("/ajax/randvalue", randvalue)
	err := http.ListenAndServe(":2888", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	fmt.Println("end.")
}
