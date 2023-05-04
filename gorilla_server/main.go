package main

func gorillaServer() error {
	s := Server{
		Config: Config{
			Port: "8052",
		},
	}

	return s.Start()
}

func main() {
	panic(gorillaServer())
}
