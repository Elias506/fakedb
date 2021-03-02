package driver

import (
	"net"
)

type FakeClient struct {
	Conn  net.Conn
}

func NewFakeConn(port string) (*FakeClient, error) {
	conn, err := net.Dial("tcp", ":" + port)
	if err != nil {
		return nil, err
	}
	return &FakeClient{
		Conn: conn,
	}, nil
}

func (fc *FakeClient) Send(msg string) (string, error) {
	n, err := fc.Conn.Write([]byte(msg))
	if n == 0 || err != nil {
		return "", nil
	}
	buff := make([]byte, 1024)
	n, err = fc.Conn.Read(buff)
	if err != nil {
		return "", err
	}

	return string(buff[0:n]), nil
}


