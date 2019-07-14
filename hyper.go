
package hyper

import (
	"net"
	"encoding/json"
	"strconv"
	"bytes"
)

type Client struct {
	conn net.Conn
}

type Server struct {
	listener net.Listener
}

func (c *Client) Connect(addr string, port uint16) {
	// Construt an address that we can use.
	address := addr + ":" + strconv.Itoa(int(port))

	// Create a net.Conn connection that we'll use to send
	//  data over.
	c.conn, _ = net.Dial("tcp", address)
}

func (c *Client) Recieve() map[string]interface{} {
	buffer := make([]byte, 8192)

	c.conn.Read(buffer)

	dataResult := make(map[string]interface{})

	json.Unmarshal(bytes.Trim(buffer, "\x00"), &dataResult)

	return dataResult
}

func (c *Client) Send(data map[string]interface{}) {
	dataBytes, _ := json.Marshal(data)

	c.conn.Write(dataBytes)
}

func (c *Client) Close() {
	defer c.conn.Close()
}

func (c *Server) Host(port uint16, onClient func(Client)) {
	address := ":" + strconv.Itoa(int(port))

	c.listener, _ = net.Listen("tcp", address)

	for {
		conn, _ := c.listener.Accept()

		client := Client{}
		client.conn = conn

		go onClient(client)
	}
}

func (c *Server) Send(client Client, data map[string]interface{}) {
	dataBytes, _ := json.Marshal(data)

	client.conn.Write(dataBytes)
}

func (c *Server) Recieve(client Client) map[string]interface{} {
	buffer := make([]byte, 8192)

	client.conn.Read(buffer)

	dataResult := make(map[string]interface{})

	json.Unmarshal(bytes.Trim(buffer, "\x00"), &dataResult)

	return dataResult
}
