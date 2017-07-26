// Copyright 2017 Mobile Data Books, LLC. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()
	t1 := time.Now()

	t1 = time.Now()
	t01 := t1.Sub(t0)
	tStr1 := fmt.Sprintf("Hello World from Go in minimal Docker container (4.28MB) - tc-helloworld-go-ws - v.1.0, it took %v to run", t01)

	fmt.Fprintln(w, tStr1)

	fmt.Printf("%s", tStr1)

}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Started, serving at 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
