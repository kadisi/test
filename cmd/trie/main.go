package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"
)

var filename string

func Help() {
	fmt.Println("Usage:")
	flag.PrintDefaults()
	fmt.Println()
}

func PrintCount(tree *Node) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file %v error %v\n", filename, err)
	}
	defer f.Close()

	scan := bufio.NewScanner(f)
	scan.Split(SplitEnglishWords)

	fmt.Println("----  Compute Englist words counts ----")
	for scan.Scan() {
		fmt.Printf("[%s] count %d\n", Lower(scan.Text()), GetCounts(tree, scan.Text()))
	}
}

func Process() {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file %v error %v\n", filename, err)
	}
	defer f.Close()
	tree := CreateTree(f)
	//PrintCount(tree)

	fmt.Println("-------- Travel Tree --------")
	TravelTree(tree, NewStack())

}
func main() {
	flag.StringVar(&filename, "file", "./test.txt", "file name, suggest to use abs path")
	flag.Parse()

	Help()
	go Process()
	http.ListenAndServe(":8080", nil)
}
