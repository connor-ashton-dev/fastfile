package p2p

type HandshakeFunc func(Peer) error

func NOPEHandshakeFunc(Peer) error { return nil }
