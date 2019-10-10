package gateway

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"log"
	"math/big"
	"os"

	tp "github.com/henrylee2cn/teleport"
	quic "github.com/lucas-clemente/quic-go"
)

var mlog = log.New(os.Stdout, "gateway ", log.LstdFlags|log.Lshortfile)
var gt = &gateway{}

//QuicServer Start a server that echos all data on the first stream opened by the client
func quicServer() {
	listener, err := quic.ListenAddr(quicServeAddr(), generateTLSConfig(), nil)
	if err != nil {
		panic(err)

	}
	// rpcx
	defer tp.FlushLogger()
	tp.SetLoggerLevel("ERROR")

	cli := tp.NewPeer(tp.PeerConfig{})
	defer cli.Close()

	cli.RoutePush(new(Broadcast))

	rpcsess, terr := cli.Dial(":9090")
	if terr != nil {
		tp.Fatalf("%v", err)
	}
	broadcastRegister(rpcsess)
	for {
		newClientSession, err := listener.Accept(context.Background())
		if err != nil {
			mlog.Println(err)
			newClientSession.Close()
			break
		}
		var ytCli = new(ytClientInfo)
		ytCli.quicSession = newClientSession
		ytCli.tpSession = rpcsess
		go ytCli.process()
	}
}

// Setup a bare-bones TLS config for the server
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}
