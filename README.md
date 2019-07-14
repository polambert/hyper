# hyper
An all around networking library for go, as complicated as you want it to be.

Hyper can be used for networking anything from games to backends to writing your own HTTP server.
We don't recommend trying to write your own HTTP server. Please don't use hyper for that. Please.

Hyper works by sending `map[string]interface{}` with JSON encoding over Go's `net` system.

Here's how you can create a simple client-server system:

The system has a client connect to a server, attempt a username and password, and the server responds with whether or not it was correct.

```
// The Client
client := hyper.Client{}

client.Connect("localhost", 8000)

data := make(map[string]interface{})
data["username"] = "myusername"
data["pin"] = 2900

client.Send(data)

data = client.Recieve()

if data["success"] {
	fmt.Println("Login was a success!")
} else {
	fmt.Println("Failed login attempt.")
}
```

```
// The Server
server := hyper.Server{}

// Hyper uses a function to handle each client.
server.Host(8000, func(client hyper.Client) {
	// Client is simply a struct containing a net.Conn.
	// You can write your own custom data straight to the net.Conn if you want.
	// The net.Conn is accessible as <client.conn>.
	data := server.Recieve(client)
	
	// Create the map that we'll be return to the client, and
	//  give it a default failure for the login attempt.
	returnData := map[string]interface{
		"success": false
	}
	
	// Here, we check that the login was succesfull.
	if data["username"] == "myusername" {
		if data["pin"] == 2900 {
			returnData["success"] = true
		}
	}
	
	server.Send(client, returnData)
})
```

If you have any questions about hyper, support will be continued indefinitely.

Contact me at:

- parkerthelamb@gmail.com
- Park#3337

