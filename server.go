package Fakedb

import (
	"bytes"
	"fmt"
	"net"
)

type Server struct {
	Listener   net.Listener
	ConnAmount int
	FakeDB     *DB
}

func NewDB(port string) error {
	s := &Server{
		ConnAmount: 0,
	}
	s.FakeDB.Tables = make(map[string]*table)
	var err error
	s.Listener, err = net.Listen("tcp", ":" + port)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer s.Listener.Close()
	fmt.Println("Server is listening...")
	for {
		conn, err := s.Listener.Accept()
		s.ConnAmount++
		addr := conn.RemoteAddr().String()
		fmt.Printf("Client %s connected\n", addr)
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handler(conn, s.FakeDB, addr)
	}
}

func handler(conn net.Conn, FakeDB *DB, addr string) {
	defer conn.Close()
	for {
		req := make([]byte, 1024*8)
		n, err := conn.Read(req)
		if n == 0 || err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%s> %s\n", addr, string(req[0:n]))
		//CreateTable(...)
		source := bytes.SplitN(req[0:n], []byte("|"), 2)
		if len(source) != 2 {
			fmt.Println("Wrong input")
			break
		}
		res := []byte("OK")
		str := source[1]
		switch string(source[0]) {
		case "CREATE_TABLE":
			err = FakeDB.CreateTable(string(str))
		case "DELETE_TABLE":
			err = FakeDB.DeleteTable(string(str))
		case "INSERT":
			err = FakeDB.Insert(string(str))
		case "SELECT":
			var t []*rowType
			t, err = FakeDB.Select(string(str))
			res = []byte(fmt.Sprint(t))
		default:
			res, err = nil, fmt.Errorf("unexpected command")
		}
		if err != nil {
			fmt.Println("Error:", err)
			conn.Write([]byte(fmt.Sprintf("Error: %v", err)))
			continue
		}
		conn.Write(res)
	}
}
