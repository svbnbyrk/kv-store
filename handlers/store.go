package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/svbnbyrk/kv-store/internal"
)

type Stores struct {
	l   *log.Logger
	kvs *internal.Store
}

//dependency injection
func NewStores(l *log.Logger, kvs *internal.Store) *Stores {
	return &Stores{l, kvs}
}

func (p *Stores) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.l.Println("GET", r.URL.Path)
		p.getValue(r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:], rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.l.Println("POST", r.URL.Path)

		p.setValue(r.URL.Path[strings.LastIndex(r.URL.Path, "/")+1:], rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Stores) getValue(key string, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET value")
	lp := p.kvs.Get(key)
	if lp == "" {
		http.Error(rw, "object not found", http.StatusNotFound)
	}
	err := ToJSON(rw, lp)
	if err != nil {
		http.Error(rw, "unable json writer", http.StatusInternalServerError)
	}
}

func (p *Stores) setValue(id string, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST key-value")
	kv := &SetModel{}

	err := kv.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
	}

	p.kvs.Post(kv.Key, kv.Value)

	if err != nil {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}
}
