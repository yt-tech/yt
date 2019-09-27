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
	managerproto "yt/rpcproto"
	"yt/ytproto"

	ggproto "github.com/gogo/protobuf/proto"
	quic "github.com/lucas-clemente/quic-go"
	"google.golang.org/grpc"
)

type pServerInfo struct{}

var mlog = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

//QuicServer Start a server that echos all data on the first stream opened by the client
func QuicServer() {
	conn, err := grpc.Dial(managerServerAddr, grpc.WithInsecure())
	if err != nil {
		mlog.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	grpcClient := managerproto.NewDataClient(conn)

	listener, err := quic.ListenAddr(quicaddr, generateTLSConfig(), nil)
	if err != nil {
		panic(err)

	}
	go func() {
		var r = &managerproto.BroadcastRegiste{}
		gr, _ := grpcClient.Broadcast(context.Background(), r)
		userlist.RLock()
		defer userlist.RUnlock()
		for k, qs := range userlist.ul {
			in, _ := gr.Recv()
			bf, _ := in.Marshal()
			mlog.Println(k, bf, qs)
			// qs.Write(bf)
		}
	}()
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
				var userRequest ytproto.ActionRequest
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
						if err == nil {
							var reply = 1
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
						gt.joinGroup(grpcClient)
					case 5:
						// gt.leaveGroup(mxclient)
					case 7:
						// gt.holdMic(mxclient)
					case 9:
						// gt.releaseMic(mxclient)
					case 11:
						// gt.disconnect(mxclient)
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
