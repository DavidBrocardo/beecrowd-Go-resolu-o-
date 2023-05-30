// Questao 5 - Trabalho 2
// Autor : David Antonio Brocardo

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var tests int
	tests = 1
	//abre o arquivo de entrada para leitura
	file, ferror := os.Open("aero.in")
	// testa o erro ocorrido
	if ferror != nil {
		fmt.Println("Erro:", ferror) // Imprime o erro ocorrido
	} else {

		reader := bufio.NewReader(file) // cria um objeto reader para o buffer file
		defer file.Close()              //ao final fecha o arquivo

		// Cria um arquivo para escrita(saida)
		arquivo, err := os.Create("aero.out")
		if err != nil {
			fmt.Println("1 Erro ao criar o arquivo:", err)
			return
		}
		defer arquivo.Close()

		for {
			for {
				line, rerror := reader.ReadString('\n')
				if rerror != nil {
					return
				}
				splitStr := strings.Split(line, " ")
				str1 := splitStr[0]
				str2 := splitStr[1]
				// Remove o caractere de quebra de linha
				str2 = strings.TrimRight(str2, "\n")

				// Converte a string para int64
				A, err := strconv.ParseInt(str1, 10, 64) //num de aeroportos
				V, err := strconv.ParseInt(str2, 10, 64) // num de voos
				if err != nil {
					fmt.Println("1 Erro ao converter a string para int64:", err)
					return
				}
				if A == 0 && V == 0 { //condicao de parada
					break
				}

				voos := make(map[int]int)
				maxVoos := 0

				//le as informaçoes dos voos
				for i := 0; i < int(V); i++ {
					line, rerror := reader.ReadString('\n')
					if rerror != nil {
						return
					}

					splitStr := strings.Split(line, " ")
					strx := splitStr[0]
					stry := splitStr[1]
					// Remove o caractere de quebra de linha
					stry = strings.TrimRight(stry, "\n")

					X, err := strconv.ParseInt(strx, 10, 64)
					Y, err := strconv.ParseInt(stry, 10, 64)
					if err != nil {
						fmt.Println("1 Erro ao converter a string para                 int64:", err)
						return
					}
					//incrementa o contador de voos para os aeroportos de origem e destino
					voos[int(X)]++
					voos[int(Y)]++

					//verifica se o numero de voos deste aeroporto e maior que o maximo atual
					if voos[int(X)] > maxVoos {
						maxVoos = voos[int(X)]
					}
					if voos[int(Y)] > maxVoos {
						maxVoos = voos[int(Y)]
					}
				}

				//lista os aeroportos com o maximo de voos

				escritaTeste := fmt.Sprintf("Teste %d\n", tests)
				tests++
				// Escreve uma teste no arquivo
				_, err = arquivo.Write([]byte(escritaTeste))
				if err != nil {
					fmt.Println("Erro ao escrever no arquivo:", err)
					return
				}

				var aero string
				for i := 1; i <= int(A); i++ {
					if voos[i] == maxVoos {
						aero = fmt.Sprintf("%d ", i)
						// escreve o aero no aero
						_, err = arquivo.Write([]byte(aero))
						if err != nil {
							fmt.Println("Erro ao escrever no arquivo:", err)
							return
						}

					}
				}
				// quebra linha
				aero = fmt.Sprintf("\n\n")
				_, err = arquivo.Write([]byte(aero))
				if err != nil {
					fmt.Println("Erro ao escrever no arquivo:", err)
					return
				}
			}
			// Fecha o Arquivo
			if err := file.Close(); err != nil {
				panic(err)
			}

		}
	}
}
