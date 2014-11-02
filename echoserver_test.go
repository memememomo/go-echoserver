package echoserver

import (
	"bufio"
	"fmt"
	"github.com/lestrrat/go-tcptest"
	"log"
	"net"
	"strconv"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	echo := func(port int) {
		server := &Server{Host: "localhost", Port: port}
		err := server.Run()
		if err != nil {
			t.Error("Error")
		}
	}

	server, err := tcptest.Start(echo, 30*time.Second)
	if err != nil {
		t.Error("Failed to start echoserver: %s", err)
	}

	log.Printf("echoserver started on port %d", server.Port())

	conn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(server.Port()))
	if err != nil {
		t.Error("Failed to connect to echoserver")
	}
	fmt.Fprintf(conn, "test hogehoge\n")
	res, err := bufio.NewReader(conn).ReadString('\n')
	if res != "test hogehoge\n" {
		t.Error("Wrong Response")
	}
}
