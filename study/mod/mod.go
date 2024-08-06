package mod

import "fmt"
import "rsc.io/quote"

/* go proxy */
// go env -w GOPROXY=https://goproxy.cn,direct

/* go mod */
// init: go mod init demo-go
// auto re-organize deps: go mod tidy
// install deps: go get xxx

func Mod() {
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}
