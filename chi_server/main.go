package main

func chiServer() error {
	s := Server{
		Config: Config{
			Port: "8052",
		},
	}

	return s.Start()
}

func main() {
	panic(chiServer())
}
