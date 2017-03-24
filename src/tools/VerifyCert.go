package main

import (
	"crypto/x509"
	"encoding/pem"
	"flag"
        "os"
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	certFile := flag.String("certFile", "", "provide X509 cert file name")
	certPEMFl  := flag.String("certPEM", "", "provide X509 PEM file name")

	flag.Parse()
        if len(*certFile) == 0 || len(*certPEMFl) == 0 {
	     flag.PrintDefaults()
             os.Exit(0)
        } 
        certPEM, err := ioutil.ReadFile(*certFile)
        if err != nil {
		log.Fatal(err)
	}
        rootPEM, err := ioutil.ReadFile(*certPEMFl)
        if err != nil {
		log.Fatal(err)
	}
	roots := x509.NewCertPool()
        ok := roots.AppendCertsFromPEM(rootPEM)
	if !ok {
		panic("failed to parse root certificate")
	}
 	block, _ := pem.Decode(certPEM)
	if block == nil {
		panic("failed to parse certificate PEM")
	}
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		panic("failed to parse certificate: " + err.Error())
	}

	opts := x509.VerifyOptions{
		//DNSName: "mail.google.com",
		Roots:   roots,
	}

	if _, err := cert.Verify(opts); err != nil {
		panic("failed to verify certificate: " + err.Error())
	}
	fmt.Printf("Verified oh cert file: %s , PEM : %s\n",*certFile,*certPEMFl)

}
