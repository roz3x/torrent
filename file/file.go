package file

import (
	"bytes"
	"crypto/sha1"
	"os"

	bg "github.com/jackpal/bencode-go"
)

//Info for file info section
type Info struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

//BitTorrent for the whole file
type BitTorrent struct {
	Announce string `bencode:"announce"`
	Info     Info   `bencode:"info"`
	PeerID   string
	Port     string
	InfoHash string
}

//Read gives out the bencode format of the torrnet file
func Read(filename string) (*BitTorrent, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	bencodeFormat := &BitTorrent{}
	bg.Unmarshal(file, bencodeFormat)
	return bencodeFormat, nil
}

//InfoHash gives the infohash which will be used to authenticate
//for the correct torrent file
func (b *Info) InfoHash() ([20]byte, error) {
	var buffer bytes.Buffer
	err := bg.Marshal(&buffer, *b)
	if err != nil {
		return [20]byte{}, err
	}
	return sha1.Sum(buffer.Bytes()), nil
}
