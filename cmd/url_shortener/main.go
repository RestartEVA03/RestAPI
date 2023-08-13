package main

import (
	"first_rest_api_v1/internal/config"
	"fmt"
)

func main() {

	//init config: cleanenv
	conf := config.Must_LoadConf()
	fmt.Println(conf)

	//TODO : init logger: slog

	//TODO : init storage: sqlite3

	//TODO : init router: chi, "chi render"

	//TODO : run server

}
