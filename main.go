package main

import (
	"fmt"
	"os"
)

func main() {
	// VS Code と dlv を使って、デバッグを行い、
	// fmt.Println の中で何が起こっているかを確認した
	// めっちゃ楽しい
	f, err := os.Create("hello.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(f, "私は %d 歳です", 25)
}
