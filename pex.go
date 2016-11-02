package torrent

import "github.com/lovedboy/torrent/util"

type peerExchangeMessage struct {
	Added      util.CompactIPv4Peers `bencode:"added"`
	AddedFlags []byte                `bencode:"added.f"`
	Dropped    util.CompactIPv4Peers `bencode:"dropped"`
}
