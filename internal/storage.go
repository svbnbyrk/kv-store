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

//Store is db
type Store struct {
	db map[string]string
}

//NewStore is constractor
func NewStore() *Store {
	store := &Store{db: map[string]string{}}
	return store
}

//Get value
func (k Store) Get(key string) string {
	return k.db[key]
}

//Post is update key-value
func (k Store) Post(key string, value string) {
	k.db[key] = value
}

func (k Store) Delete() {
	k.db = make(map[string]string)
}

//Save is saving map to json file
func (k Store) Save(l *log.Logger) error {

	jsonStr, err := json.Marshal(k.db)
	if err != nil {
		l.Printf("Error: %s", err)
		return err
	}

	files, err := ioutil.ReadDir("tmp/.")
	if err != nil {
		l.Fatal(err)
		return err
	}

	if len(files) > 0 {
		err := os.Remove("tmp/" + files[0].Name())
		if err != nil {
			l.Fatal(err)
			return err
		}
	}

	fp := createTimeStamp(time.Now())

	err = ioutil.WriteFile(fp, jsonStr, 0644)

	if err != nil {
		l.Fatal(err)
		return err
	}
	return nil
}

//Read is reading tmp/TIMESTAMP-data.json directory and fill db map
func (k Store) Read(l *log.Logger) {
	//tmp folder is creating if not exist
	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		os.Mkdir("tmp", os.ModePerm)
	}

	files, err := ioutil.ReadDir("tmp/.")
	if err != nil {
		log.Fatal(err)
	}

	//tmp folder is empty
	if len(files) == 0 {
		return
	}
	// Open our jsonFile
	jsonFile, err := os.Open("tmp/" + files[0].Name())
	// if we os.Open returns an error then handle it
	if err != nil {
		l.Println(err)
		return
	}
	l.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		l.Println(err)
		return
	}

	var result map[string]string
	json.Unmarshal([]byte(byteValue), &result)

	for key, value := range result {
		k.db[key] = value
	}
}

func createTimeStamp(t time.Time) string {
	s := fmt.Sprintf("store-%s-data.json", t.Format("2006-01-02 15_04_05"))
	return filepath.Join("tmp", s)
}
