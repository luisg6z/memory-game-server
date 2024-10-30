package main

func main() {
	grpcServer := NewGRPCServer(":8080")
	grpcServer.Run()
}
