# Go-P2P

Goでの基本的なP2P接続のサンプル
<br>
1つのピアがサーバーとして、
<br>
もう1つのピアがクライアントとして動作する

### ピア1: サーバー

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Waiting for connection...")
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Printf("Connection from %s\n", conn.RemoteAddr().String())
	conn.Write([]byte("Hello, Peer!"))
}
```

### ピア2: クライアント

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:5000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println("Received:", string(buffer[:n]))
}
```
