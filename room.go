package main

type room struct {
	//forwardは他のクライアントに転送するためのメッセージを保持するチャネルです
	forward chan []byte
	//joinはチャットルームに参加しようとしているクライアントのためのチャネルです。
	join chan *client
	//leaveはチャットルームから退室しようとしているクライアントのためのチャネルです。
	leave chan *client
	//clientsには材質しているすべてのクライアントが保持されます。
	clients map[*client]bool
}
