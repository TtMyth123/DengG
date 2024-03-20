package main

import (
	"fmt"
	"github.com/TtMyth123/DengG/Lottery61BS/sysInit"
	"path/filepath"
)

func main() {
	path, e := filepath.Abs("conf/app.conf")
	fmt.Println(path, e)
	sysInit.Init()
	fmt.Println("Hello, World!11")
}
