package file

import (
	"testing"
)

const (
	torrentLocation = "../testdata/file.torrent"
	port            = "55000"
)

//TestGetURL ....
func TestGetURL(t *testing.T) {
	bT, err := Read(torrentLocation)
	if err != nil {
		t.Fatalf("%v\n", err.Error())
	}
	bT.Port = port
	url, err := bT.URL()
	if err != nil {
		t.Fatalf("%v\n", err.Error())
	}
	t.Logf("url final:%v\n", url)
}
