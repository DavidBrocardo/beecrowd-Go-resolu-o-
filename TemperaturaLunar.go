// Questao 2 - Trabalho 2 
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
	var i, j, maior, menor, soma, cont, tests int
	var temperatura []int

	cont = 1
	tests = 1
	menor = 1000
	maior = -1000
	// Abre o arquivo de entrada
	file, ferror := os.Open("lua.in") 
	// testa o erro ocorrido
	if ferror != nil {
		fmt.Println("Erro:", ferror) 
	} else {
		reader := bufio.NewReader(file) 
		// Cria um arquivo para escrita(saida)
		arquivo, err := os.Create("lua.out")
		if err != nil {
			fmt.Println("Erro ao criar o arquivo:", err)
			return
		}
		defer arquivo.Close()

		// Comeca a leitura do arquivo de entrada
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
		
			// Converte a string para int8
			n, err := strconv.ParseInt(str1, 10, 64)
			m, err := strconv.ParseInt(str2, 10, 64)
			if err != nil {
				fmt.Println("1 Erro ao converter a string para int8:", err)
				return
			}
			if n == 0 && m == 0 {
				break
			} else {
				temperatura = make([]int, n)
				for i = 0; i < int(n); i++ {

					temp, errorT := reader.ReadString('\n')
					if errorT != nil {
						fmt.Println("2 Erro ao ler string:")
						return
					}
					// Remove quebra de linha
					temp = strings.TrimRight(temp, "\n")
					
					//Converte a string para int64
					t, err := strconv.ParseInt(temp, 10, 64)
					temperatura[i] = int(t)
					if err != nil {
						fmt.Println("2 Erro ao converter a string para int64:", err)
						return
					}
				}

				for i = 0; i < int(n); i++ {
					soma = 0
					cont = 0
					
					for j = 0; j < int(m); j++ {
						if (i + j) < int(n) {
							soma += temperatura[i+j]
							cont++
						}
					}
					if cont == int(m) {
						soma = soma / int(m)

						if maior < soma {
							maior = soma
						}
						if menor > soma {
							menor = soma
						}
					}
				}
		
				// Concatena os valores em uma única string
				escritaTeste := fmt.Sprintf("Teste %d\n", tests)
				tests++
				MenorMaior := fmt.Sprintf("%d %d\n\n", menor, maior)
				// Escreve uma string no arquivo
				_, err = arquivo.Write([]byte(escritaTeste))
				if err != nil {
					fmt.Println("Erro ao escrever no arquivo:", err)
					return
				}
				_, err = arquivo.Write([]byte(MenorMaior))
				if err != nil {
					fmt.Println("Erro ao escrever no arquivo:", err)
					return
				}
				cont++
				menor = 1000
				maior = -1000
			}

		}
	}
	// Fecha o Arquivo
	if err := file.Close(); err != nil {
		panic(err)
	}

}
