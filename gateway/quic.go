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
	"yt/proto"
	"yt/ytproto"

	ggproto "github.com/gogo/protobuf/proto"
	quic "github.com/lucas-clemente/quic-go"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

type pServerInfo struct{}

var mlog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

//QuicServer Start a server that echos all data on the first stream opened by the client
func QuicServer() {
	pch := make(chan *protocol.Message)
	mch := make(chan *protocol.Message)
	d := client.NewPeer2PeerDiscovery("tcp@"+pushServerAddr, "")
	dm := client.NewPeer2PeerDiscovery("tcp@"+managerServerAddr, "")
	pxclient := client.NewBidirectionalXClient("PushServer", client.Failtry, client.RandomSelect, d, client.DefaultOption, pch)
	mxclient := client.NewBidirectionalXClient("Manager", client.Failtry, client.RandomSelect, dm, client.DefaultOption, mch)
	defer pxclient.Close()
	defer mxclient.Close()
	go func() {
		for msg := range pch {
			fmt.Printf("receive pch msg from server: %s\n", msg.Payload)
		}
	}()
	go func() {
		var msgProto ytproto.Man
		for msg := range mch {
			fmt.Printf("receive mch msg from server: %s\n", msg.Payload)
			ggproto.Unmarshal(msg.Payload, &msgProto)
			userlist.RLock()
			defer userlist.RUnlock()
			for _, uid := range msgProto.GetUid() {
				sendToUser, ok := userlist.ul[uid]
				if ok && sendToUser != nil {
					sendToUser.Write(msg.Payload)
				}
			}
		}
	}()
	ar := &argsInfo{
		ID:       gatewayID,
		Listener: gatewayUDPListener,
	}
	g := &RegisteInfo{
		args: ar,
	}
	g.registe2manager(mxclient)
	listener, err := quic.ListenAddr(quicaddr, generateTLSConfig(), nil)
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
				var userRequest proto.ActionRequest
				for {
					n, err := stream.Read(readBuff)
					if err != nil {
						mlog.Println(err)
						return
					}
					err = ggproto.Unmarshal(readBuff[:n], &userRequest)
					if err != nil {
						mlog.Println(err)
						return
					}
					var gt = &gtInfo{
						action: &userRequest,
					}
					switch userRequest.GetActionID() {
					case 1:
						if reply, err := gt.connect(mxclient); err == nil {
							switch {
							case reply == 1 || reply == 2:
								sendStream, _ := sess.OpenUniStream()
								userlist.Lock()
								userlist.ul[gt.action.GetUid()] = sendStream
								userlist.Unlock()
							default:
								mlog.Println("?")
							}
						} else {
							break
						}
					case 3:
						gt.joinGroup(mxclient)
					case 5:
						gt.leaveGroup(mxclient)
					case 7:
						gt.holdMic(mxclient)
					case 9:
						gt.releaseMic(mxclient)
					case 11:
						gt.disconnect(mxclient)
					default:
						mlog.Println("--------")
					}
					bf, _ := ggproto.Marshal(gt.action)
					stream.Write(bf)
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
