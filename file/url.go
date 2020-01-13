package file

import (
	"fmt"
	"net/url"
)

//URL will use the fields in the
//BencodeInfo struct and make appropriate
//url with that
func (bT BitTorrent) URL() (string, error) {
	baseURL, err := url.Parse(bT.Announce)
	if err != nil {
		return "", err
	}

	if bT.PeerID == "" {
		bT.GenPeerID()
		infoHash, err := bT.Info.InfoHash()
		if err != nil {
			return "", err
		}
		bT.InfoHash = fmt.Sprintf("%s", infoHash)
	}

	parameters := url.Values{
		"info_hash":  []string{bT.InfoHash},
		"peer_id":    []string{bT.PeerID},
		"port":       []string{bT.Port},
		"uploaded":   []string{"0"},
		"downloaded": []string{"0"},
		"compact":    []string{"1"},
		"left":       []string{fmt.Sprintf("%v", bT.Info.Length)},
	}
	baseURL.RawQuery = parameters.Encode()
	return baseURL.String(), nil
}

//GenPeerID will generate a random 20 byte
//peerID , helpful for identifing ourself
//among other peers
func (bT *BitTorrent) GenPeerID() {
	// peerid := make([]byte, 20)
	// rand.Read(peerid)
	bT.PeerID = fmt.Sprintf("11111111111111111111")
}
