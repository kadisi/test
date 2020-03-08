package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

func parseCrt() {
	certPEMBlock, err := ioutil.ReadFile("./server.crt")
	if err != nil {
		log.Fatalf("read server.crt error %v", err)
	}
	certDERBlock, _ := pem.Decode(certPEMBlock)
	x509Cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	fmt.Println(x509Cert.IPAddresses)

}

func main() {
}
