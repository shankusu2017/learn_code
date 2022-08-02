package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net"
)

func HandleClientConnect(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Receive Connect Request From ", conn.RemoteAddr().String())
	buffer := make([]byte, 10240)
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			log.Println(err.Error())
			break
		}
		fmt.Printf("Receive Data len:%d, len:%d Content:[%s]\n", len(buffer), length, string(buffer[:length]))
		//发送给客户端
		_, err = conn.Write([]byte("服务器收到数据:" + string(buffer[:length])))
		if err != nil {
			break
		}
	}
	fmt.Println("Client " + conn.RemoteAddr().String() + " Connection Closed.....")
}

func ClientSayHello(info *tls.ClientHelloInfo) (*tls.Config, error) {
	fmt.Println("22222222")
	return nil, nil
}

func main() {
	tlsConfig := &tls.Config{}

	crt, err := tls.LoadX509KeyPair("server.crt", "server.key")
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}

	//tlsConfig.GetConfigForClient= ClientSayHello
	tlsConfig.Certificates = []tls.Certificate{crt}

	// Time returns the current time as the number of seconds since the epoch.
	// If Time is nil, TLS uses time.Now.
	//tlsConfig.Time = time.Now
	// Rand provides the source of entropy for nonces and RSA blinding.
	// If Rand is nil, TLS uses the cryptographic random reader in package
	// crypto/rand.
	// The Reader must be safe for use by multiple goroutines.
	tlsConfig.Rand = rand.Reader
	l, err := tls.Listen("tcp", ":4444", tlsConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		} else {
			go HandleClientConnect(conn)
		}
	}

}
