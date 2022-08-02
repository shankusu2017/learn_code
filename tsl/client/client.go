package main

import (
	"crypto/tls"
	"fmt"
	_ "github.com/apparentlymart/go-openvpn-mgmt/openvpn"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type fakeConn struct {
	conn net.Conn
}

func (fc *fakeConn) Read(b []byte) (n int, err error) {
	return fc.conn.Read(b)
	fmt.Println("tls.req.Read")
	return 0, nil
}

func (fc *fakeConn) Write(b []byte) (n int, err error) {
	return fc.conn.Write(b)
}
func (fc *fakeConn) Close() error {
	return fc.conn.Close()
	//return nil
}
func (fc *fakeConn) LocalAddr() net.Addr {
	return fc.conn.LocalAddr()
	//return nil
}
func (fc *fakeConn) RemoteAddr() net.Addr {
	return fc.conn.RemoteAddr()
	//return nil
}
func (fc *fakeConn) SetDeadline(t time.Time) error {
	return fc.conn.SetDeadline(t)
	//return nil
}
func (fc *fakeConn) SetReadDeadline(t time.Time) error {
	return fc.conn.SetReadDeadline(t)
	//return nil
}
func (fc *fakeConn) SetWriteDeadline(t time.Time) error {
	return fc.conn.SetWriteDeadline(t)
	//return nil
}

func main() {
	var err error

	fd, _ := os.OpenFile("client.log", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	conf := &tls.Config{
		InsecureSkipVerify: true, //这里是跳过证书验证，因为证书签发机构的CA证书是不被认证的
		KeyLogWriter:       fd,
	}
	//注意这里要使用证书中包含的主机名称

	//tcpCli := &fakeConn{}
	//tcpCli.conn, _ = net.Dial("tcp", "127.0.0.1:44444")
	tcpCli, err := net.Dial("tcp", "127.0.0.1:4444")
	if err != nil {
		log.Fatalln("net.Dial fail(%s)", err.Error())
	}
	conn := tls.Client(tcpCli, conf)
	err = conn.Handshake()
	// conn, err := tls.Dial("udp", "127.0.0.1:4444", conf)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println("connect tls server done!")
	}
	defer conn.Close()

	log.Println("Client Connect To ", conn.RemoteAddr())
	status := conn.ConnectionState()
	fmt.Printf("%#v\n", status)
	buf := make([]byte, 1024)
	ticker := time.NewTicker(1 * time.Millisecond * 500)
	for {
		select {
		case <-ticker.C:
			{
				_, err = io.WriteString(conn, "hello")
				if err != nil {
					log.Fatalln(err.Error())
				}
				len, err := conn.Read(buf)
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println("Receive From Server:", string(buf[:len]))
				}
			}
		}
	}
}
