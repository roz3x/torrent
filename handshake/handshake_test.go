package handshake

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"testing"

	"github.com/roz3x/torrent/file"
	"github.com/roz3x/torrent/pears"
)

const (
	payloadLocation = "../testdata/payload"
	torrentLocation = "../testdata/file.torrent"
)

func TestHandshake(t *testing.T) {
	fatal := func(err error) {
		if err != nil {
			t.Fatal(err.Error())
		}
	}
	t.Logf("test started")
	filePayload, err := os.Open(payloadLocation)
	fatal(err)
	payload, err := ioutil.ReadAll(filePayload)
	p, err := pears.GetPears(payload)
	fatal(err)
	bT, err := file.Read(torrentLocation)
	fatal(err)

	var wg sync.WaitGroup
	wg.Add(len(p))
	//
	// causing problem for rapid prototyping
	//
	t.Parallel()
	t.Logf("starting the parallel thing")
	shakeAll := func(p pears.Pear) {
		handshake, err := Handshake(p, bT)
		fatal(err)
		fmt.Printf("%v", len(handshake))
		wg.Done()
	}
	for _, t := range p {
		go shakeAll(t)
	}
	wg.Wait()

}

func TestFullTorrentFile(t *testing.T) {

}
