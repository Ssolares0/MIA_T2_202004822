package Estructuras

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Analyze(command string) {
	//variables
	//var flagExit = false

	//aqui al comando, le quitamos los espacios y lo devolvemos como un token

	token_ := strings.Split(command, " ")

	//CONVERTIMOS EL TOKEN EN MINUSCULAS
	if len(token_) > 0 {
		token_[0] = strings.ToLower(token_[0])
	}

	fmt.Println(token_[0])

	switch token_[0] {
	case "mkdisk":
		//estamos aca

		Analyze_Mkdisk(token_[1:])
	case "execute":

		Analyze_execute(token_[1:])

	case "rep":
		Analyze_rep(token_[1:])

	case "exit":
		//flagExit = true
		fmt.Println("gracias por usar el programa")

	default:
		fmt.Println("error al leer el comando")

	}

	/*
		if !flagExit {
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Ingresa un comando: ")

			comando, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error al leer la entrada:", err)
				return
			}
			Analyze(comando)
		}
	*/
}

func Analyze_Mkdisk(list_tokens []string) {

	fmt.Println(list_tokens)

	//variables del mkdisk

	//tamano de 5 mb
	size_int := 5 * 1024

	//Si no existe el directorio Discos, entonces crearlo
	if _, err := os.Stat("Discos"); os.IsNotExist(err) {
		err = os.Mkdir("Discos", 0664)
		if err != nil {
			fmt.Println("Error al crear el directorio Discos: ", err)
			return
		}
	}
	//Contar la cantidad de discos para asignar el nombre
	archivos, err := ioutil.ReadDir("Discos")
	if err != nil {
		fmt.Println("Error al leer el directorio: ", err)
		return
	}

	letter := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//crear nombre del disco a partir de la cantidad de discos que hay en la carpeta
	nameDisk := string(letter[len(archivos)])

	//crear el archivo binario
	file, err := os.Create("Discos/" + "Hard_Disk" + nameDisk + ".dsk")

	if err != nil {
		fmt.Println("error al crear el disco", err)
		return
	}
	defer file.Close()

	//en este apartado emepzamos en la creacion del MBR en el disk

	randomNum := rand.Intn(99) + 1
	var disk MBR

	dateNow := time.Now()
	date := dateNow.Format("2006-01-02 15:04:05")
	disk.MBR_SIZE = int64(size_int)
	disk.MBR_ID = (int64(randomNum))
	copy(disk.MBR_DATE[:], date)

	//llenamos el archivo en bytes
	bufer := new(bytes.Buffer)
	for i := 0; i < 1024; i++ {
		bufer.WriteByte(0)
	}

	var totalBytes int = 0
	for totalBytes < int(size_int) {
		c, err := file.Write(bufer.Bytes())
		if err != nil {
			fmt.Println("Error al escribir en el archivo: ", err)
			return
		}
		totalBytes += c
	}
	fmt.Println("Archivo llenado con 0s")
	//Escribir el MBR en el disco
	file.Seek(0, 0)
	err = binary.Write(file, binary.LittleEndian, &disk)
	if err != nil {
		fmt.Println("Error al escribir el MBR en el disco: ", err)
		return
	}
	fmt.Println("Disco", nameDisk, "creado con exito")

}

func Analyze_execute(list_tokens []string) {
	//fmt.Println(list_tokens)
	tokens := strings.Split(list_tokens[0], "=")

	//fmt.Println(tokens)
	readFile, err := os.Open(tokens[1])
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	for _, line := range lines {
		var token []string

		//obtenemos los valores de linea por linea

		if line == "mkdisk" {
			token := append(token, line)
			Analyze_Mkdisk(token)
		}
		if line == "rep" {
			token := append(token, line)
			Analyze_rep(token)

		}

	}

}

func Analyze_rep(list_tokens []string) {
	//Abrir el disco A
	archivo, err := os.Open("Discos/Hard_DiskA.dsk")
	if err != nil {
		fmt.Println("Error al abrir el disco: ", err)
		return
	}
	defer archivo.Close()
	disk := NewMBR()
	archivo.Seek(int64(0), 0)
	err = binary.Read(archivo, binary.LittleEndian, &disk)
	if err != nil {
		fmt.Println("Error al leer el MBR del disco: ", err)
		return
	}
	fmt.Println("Tamaño: ", disk.MBR_SIZE)
	fmt.Println("Fecha: ", string(disk.MBR_DATE[:]))
	fmt.Println("Signature: ", disk.MBR_ID)

}

func Confirmacion(msg string) bool {
	fmt.Println(msg + "(y/n)")
	//var respuesta string
	return true

}

// WriteInBytes es un ayudante para escribir un conjunto de bytes en el archivo.
func WriteInBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}
