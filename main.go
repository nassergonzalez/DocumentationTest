package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte("FixDataHere")); err != nil {
		log.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		log.Fatal(err)
	}
	str := base64.StdEncoding.EncodeToString(b.Bytes())
	fmt.Print(str)
}
