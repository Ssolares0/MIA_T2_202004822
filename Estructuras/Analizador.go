package Estructuras

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	size_int := 5 * 1024 * 1024

	//crear el archivo binario
	_, err := os.Stat("Hard_disk.sdk")

	if os.IsNotExist(err) {
		file, err := os.Create("Hard_disk.sdk")

		if err != nil {
			fmt.Println("error al crear el archivo")
		}

		data := make([]byte, size_int)

		//llenamos el archivo en bytes
		WriteInBytes(file, data)

		file.Close()
		NewMBR(int64(size_int))

	} else {
		// El archivo ya existe, así que lo abrimos para su reemplazo
		file, err := os.Create("Hard_disk.sdk")
		if err != nil {
			fmt.Println("Error al reemplazar el archivo:", err)
			return
		}

		defer file.Close()

		data := make([]byte, size_int)

		// Llenamos el archivo con bytes
		WriteInBytes(file, data)

		NewMBR(int64(size_int))
	}

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
	fmt.Println(list_tokens)

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
