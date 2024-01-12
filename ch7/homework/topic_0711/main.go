// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

//!+main

func main() {
	db := database{
		RWMutex: sync.RWMutex{},
		kv:      map[string]dollars{"shoes": 50.0, "socks": 5.0},
	}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	sync.RWMutex
	kv map[string]dollars
}

func (db *database) Range() map[string]dollars {
	db.Lock()
	defer db.Unlock()
	return db.kv
}

func (db *database) Get(key string) (dollars, bool) {
	db.RLock()
	defer db.RUnlock()
	val, ok := db.kv[key]
	return val, ok
}

func (db *database) Set(key string, value string) error {
	db.Lock()
	defer db.Unlock()
	price, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return err
	}
	db.kv[key] = dollars(price)
	return nil
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db.Range() {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.Get(item); ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if _, ok := db.Get(item); ok {
		db.Set(item, price)
		fmt.Fprintf(w, "%s: %s\n", item, price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}
