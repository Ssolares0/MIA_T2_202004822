package Estructuras

import (
	"fmt"
	"time"
)

type MBR struct {
	MBR_SIZE int64
	MBR_DATE [16]byte
	MBR_ID   int64
}

func NewMBR(size int64) {

	var nmb MBR
	var signature int64 = 1
	var data [16]byte

	//obtener la fecha actual
	now := time.Now()
	formatt := now.Format(time.RFC3339)

	//convertimos la cadena en bytes de logitud
	dataBytes := []byte(formatt)[:16]

	copy(data[:], dataBytes)
	//los datos del mbr

	fmt.Println(size)
	fmt.Println(data)
	fmt.Println(signature)

	nmb.MBR_SIZE = size
	nmb.MBR_DATE = data
	nmb.MBR_ID += signature
}
