package handshake

import (
	"errors"
	"io"
	"net"
	"time"

	"github.com/roz3x/torrent/file"
	"github.com/roz3x/torrent/pears"
)

const (
	//PSTR for bitTorrent protocol
	PSTR = "BitTorrent protocol"
)

var (
	timeUp   = make(chan bool, 1)
	connChan = make(chan net.Conn, 1)
)

//Handshake  for handshaking protocol
//as per  bittorrent protocol
func Handshake(p pears.Pear, bT *file.BitTorrent) (string, error) {
	url := p.URL()
	go timer()
	go getConn(url)
	select {
	case <-timeUp:
		return "", errors.New("request takes too long")
	case conn := <-connChan:
		//using a concurrent approch
		conn.SetDeadline(time.Now().Add(2 * time.Second))
		bufferLen := 49 + len(PSTR)
		buffer := make([]byte, bufferLen)
		buffer[0] = byte('\x13')
		copy(buffer[1:], PSTR)
		/*
			leaving 8 zero empty
		*/
		copy(buffer[1+19+8:], bT.InfoHash)
		copy(buffer[1+19+8+20:], bT.PeerID)
		_, err := conn.Write(buffer)

		if err != nil {
			return "", err
		}
		payload := make([]byte, 49)
		_, err = io.ReadFull(conn, payload)
		if err != nil {
			return "", nil
		}
		return string(payload), nil
	}

}

func timer() {
	time.Sleep(1 * time.Second)
	timeUp <- false
}

func getConn(url string) {
	conn, err := net.Dial("tcp", url)
	if err != nil {
		return
	}
	connChan <- conn
}
