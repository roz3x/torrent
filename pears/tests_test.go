package pears

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/roz3x/torrent/file"
)

const (
	fileLocation    = "../testdata/file.torrent"
	payloadLocation = "../testdata/payload"
	port            = "55000"
)

func TestPears(t *testing.T) {
	fatal := func(e error) {
		if e != nil {
			t.Fatal(e.Error())
		}
	}
	log := func(a interface{}) {
		t.Logf("%s", a)
	}
	bT, err := file.Read(fileLocation)
	fatal(err)
	bT.Port = port
	url, err := bT.URL()
	fatal(err)
	log(url)
	respose, err := http.Get(url)
	fatal(err)
	payload, err := ioutil.ReadAll(respose.Body)
	fatal(err)
	pears, err := GetPears(payload)
	fatal(err)
	log(pears)
}

func TestParsing(t *testing.T) {
	file, err := os.Open(payloadLocation)
	if err != nil {
		t.Fatal(err.Error())
	}
	payload, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err.Error())
	}
	pears, err := GetPears(payload)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("%v", pears)
}
