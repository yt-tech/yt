package gateway

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"

	ggproto "github.com/gogo/protobuf/proto"
	quic "github.com/lucas-clemente/quic-go"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

type pServerInfo struct {
}

const addr = "localhost:4242"

const message = "foobar"

var mlog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

//QuicServer Start a server that echos all data on the first stream opened by the client
func QuicServer() {
	ch := make(chan *protocol.Message)
	chm := make(chan *protocol.Message)
	d := client.NewPeer2PeerDiscovery("tcp@"+pushServerAddr, "")
	dm := client.NewPeer2PeerDiscovery("tcp@"+ManagerServerAddr, "")
	xclient := client.NewBidirectionalXClient("GatewayRoot", client.Failtry, client.RandomSelect, d, client.DefaultOption, ch)
	xmclient := client.NewBidirectionalXClient("Manager", client.Failtry, client.RandomSelect, dm, client.DefaultOption, chm)
	defer xclient.Close()
	defer xmclient.Close()
	go func() {
		for msg := range ch {
			fmt.Printf("receive msg from server: %s\n", msg.Payload)
		}
	}()
	gwRegiste(xclient)
	listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if err != nil {
		panic(err)

	}
	for {
		sess, err := listener.Accept(context.Background())
		if err != nil {
			panic(err)
		}
		go func() {
			for {
				stream, err := sess.AcceptStream(context.Background())
				if err != nil {
					panic(err)
				}
				readBuff := make([]byte, 1024)
				for {
					n, err := stream.Read(readBuff)
					if err != nil {
						mlog.Println(err)
					}
					var gt *gtInfo
					ggproto.Unmarshal(readBuff[:n], gt.action)
					switch gt.action.GetActionID() {
					case 1:
						gt.connectAction(xmclient)
						bf, _ := ggproto.Marshal(gt.action)
						stream.Write(bf)
					case 2:
					case 3:
						reply := loginAction(a, xmclient)
						a.ActionID = 2
						a.Ack = int32(reply)
						bf, _ := ggproto.Marshal(&a)
						stream.Write(bf)
					case 4:
					case 5:
						reply := loginAction(a, xmclient)
						a.ActionID = 2
						a.Ack = int32(reply)
						bf, _ := ggproto.Marshal(&a)
						stream.Write(bf)
					case 6:
					case 7:
						reply := loginAction(a, xmclient)
						a.ActionID = 2
						a.Ack = int32(reply)
						bf, _ := ggproto.Marshal(&a)
						stream.Write(bf)
					case 8:
					case 9:
						reply := loginAction(a, xmclient)
						a.ActionID = 2
						a.Ack = int32(reply)
						bf, _ := ggproto.Marshal(&a)
						stream.Write(bf)
					case 10:
					case 11:
						reply := loginAction(a, xmclient)
						a.ActionID = 2
						a.Ack = int32(reply)
						bf, _ := ggproto.Marshal(&a)
						stream.Write(bf)
					case 12:
					}
				}
			}
		}()
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
