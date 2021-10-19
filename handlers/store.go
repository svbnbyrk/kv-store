package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/svbnbyrk/kv-store/internal"
)

type Store struct {
	l   *log.Logger
	kvs *internal.Store
}

//dependency injection
func NewStore(l *log.Logger, kvs *internal.Store) *Store {
	return &Store{l, kvs}
}

func (p *Store) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
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

func (p *Store) getValue(key string, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET value")
	lp := p.kvs.Get(key)
	if lp == "" {
		http.Error(rw, "object not found", http.StatusNotFound)
		return
	}

	kv := &SetModel{Key: key, Value: lp}

	err := kv.ToJSON(rw, lp)
	if err != nil {
		http.Error(rw, "unable json writer", http.StatusInternalServerError)
		return
	}
}

func (p *Store) setValue(id string, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST key-value")
	kv := &SetModel{}

	err := kv.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
		return
	}

	p.kvs.Post(kv.Key, kv.Value)

	if err != nil {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}
}

func (p *Store) FlushStore(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET flush store")
	err := p.kvs.Save(p.l)
	if err != nil {
		http.Error(rw, "flush failed ", http.StatusInternalServerError)
		return
	}
}
