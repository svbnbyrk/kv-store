package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Store struct {
	db            map[string]string
	lasttimestamp time.Time
}

//dependency injection
func NewStore() *Store {
	store := &Store{db: map[string]string{}}
	return store
}
func (k Store) Get(key string) string {
	return k.db[key]
}

func (k Store) Post(key string, value string) {
	k.db[key] = value
}

func (k Store) Save(l *log.Logger) {
	jsonStr, err := json.Marshal(k.db)
	if err != nil {
		l.Printf("Error: %s", err)
	}
	t := time.Now()
	
	files, err := ioutil.ReadDir("tmp/.")
	if err != nil {
	    log.Fatal(err)
	}

	// Open our jsonFile
	if len(files) > 0{
		err := os.RemoveAll("tmp/"+files[0].Name())
		if err != nil {
			log.Fatal(err)
		 }
	}

	//set timesstamp
	k.lasttimestamp = t

	if err != nil {
		l.Fatal(err)
	}
	s := fmt.Sprintf("store-%s-data.json", t.Format("2006-01-02 15_04_05"))
	fp := filepath.Join("tmp", s)
	
	// created temp directory if not exist temp directory
	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		os.Mkdir(fp, os.ModePerm)
	}

	err = ioutil.WriteFile(fp, jsonStr, 0644)

	if err != nil {
		l.Fatal(err)
	}
}

func (k Store) Read(l *log.Logger) {


	files, err := ioutil.ReadDir("tmp/.")
	if err != nil {
	    log.Fatal(err)
	}

	if len(files) == 0{
		return
	}
	// Open our jsonFile
	jsonFile, err := os.Open("tmp/"+files[0].Name())
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]string
	json.Unmarshal([]byte(byteValue), &result)
	
	for key, value := range result{        
		k.db[key] = value
	}
}
