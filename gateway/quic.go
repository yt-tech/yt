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
	command "yt/ytproto/cmd"

	ggproto "github.com/gogo/protobuf/proto"
	quic "github.com/lucas-clemente/quic-go"
	"google.golang.org/grpc"
)

type pServerInfo struct{}

var mlog = log.New(os.Stdout, "gateway ", log.LstdFlags|log.Lshortfile)

//QuicServer Start a server that echos all data on the first stream opened by the client
func QuicServer() {
	conn, err := grpc.Dial(managerServerAddr, grpc.WithInsecure())
	if err != nil {
		mlog.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	listener, err := quic.ListenAddr(quicaddr, generateTLSConfig(), nil)
	if err != nil {
		panic(err)

	}
	grpcClient := command.NewManagerClient(conn)
	go func() {
		var r = new(command.Msg)
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
				// go func() {
				// 	audioStream, err := sess.AcceptUniStream(context.Background())
				// 	if err != nil {
				// 		mlog.Println(err)
				// 	}
				// 	audioBufferr := make([]byte, 512)
				// 	for {
				// 		n, err := audioStream.Read(audioBufferr)
				// 		mlog.Println(n, err)
				// 	}
				// }()
				readBuff := make([]byte, 1024)
				var requestMsg command.Msg
				for {
					n, err := stream.Read(readBuff)
					if err != nil {
						mlog.Println(err)
						return
					}
					err = ggproto.Unmarshal(readBuff[:n], &requestMsg)
					if err != nil {
						mlog.Println(err)
						return
					}
					var gt = &gateway{
						cmdMsg: &command.Msg{},
					}
					switch requestMsg.GetCtype() {
					case command.CommandType_ConnectRequest:
						connectRequestUserID := requestMsg.Request.Connect.GetUid()
						if err == nil {
							var reply = 1
							switch {
							case reply == 1 || reply == 2:
								connectResponse, err := gt.connect(grpcClient, connectRequestUserID)
								if err != nil {
									sess.Close()
									mlog.Println(err)
								} else {
									sendStream, _ := sess.OpenUniStream()
									userlist.Lock()
									userlist.ul[connectRequestUserID] = sendStream
									userlist.Unlock()
								}
								rs := &command.Msg{
									Ctype: command.CommandType_ConnectResponse,
									Response: &command.Response{
										ConnectAck: connectResponse,
									},
								}
								bf, _ := rs.Marshal()
								stream.Write(bf)
							default:
								mlog.Println("?")
							}
						} else {
							sess.Close()
							break
						}
					case command.CommandType_SubscribeTopicRequest:
						subscribeTopicResponse, _ := gt.subscribeTopic(grpcClient, requestMsg.Request.Subscribe)
						cm := &command.Msg{
							Ctype: command.CommandType_SubscribeTopicResponse,
							Response: &command.Response{
								SubscribeAck: subscribeTopicResponse,
							},
						}
						bf, _ := cm.Marshal()
						stream.Write(bf)
					case command.CommandType_HoldMicRequest:
					case command.CommandType_ReleaseMicRequest:
					case command.CommandType_DisconnectRequest:
					default:
						mlog.Println("--------")
					}
					bf, _ := ggproto.Marshal(gt.cmdMsg)
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
