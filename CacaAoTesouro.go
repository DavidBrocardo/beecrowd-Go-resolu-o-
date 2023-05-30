// Questao 3 - Trabalho 2
// Autor : David Antonio Brocardo

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var i, cont, tests, somaX, somaY, cont1 int
	cont1 = 0
	var joias []int
	tests = 1
	var saida string
	// Abre o arquivo de entrada
	file, ferror := os.Open("tesouro.in")
	// testa o erro ocorrido
	if ferror != nil {
		fmt.Println("Erro:", ferror)
	} else {
		reader := bufio.NewReader(file)
		// Cria um arquivo para escrita(saida)
		arquivo, err := os.Create("tesouro.out")
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
			str3 := splitStr[2]
			// Remove o caractere de quebra de linha
			str3 = strings.TrimRight(str3, "\n")
			// Converte a string para int64
			x, err := strconv.ParseInt(str1, 10, 64)
			y, err := strconv.ParseInt(str2, 10, 64)
			n, err := strconv.ParseInt(str3, 10, 64)
			quantidade := int(n) - 1
			joias = make([]int, n)
			if err != nil {
				fmt.Println("1 Erro ao converter a string para int8:", err)
				return
			}
			if x == 0 && y == 0 && n == 0 {
				break
			} else {

				for i = 0; i < int(n); i++ {
					// leitura da string referente ao valor da joia
					temp, errorT := reader.ReadString('\n')
					if errorT != nil {
						fmt.Println("2 Erro ao ler string:")
						return
					}
					// retira \n
					temp = strings.TrimRight(temp, "\n")

					//Converte a string para int64
					joia, err := strconv.ParseInt(temp, 10, 64)
					joias[i] = int(joia)
					//imprimi mensagem de erro
					if err != nil {
						fmt.Println("2 Erro ao converter a string para int64:", err)
						return
					}
				}
				// ordena de maneira decresente os valores das joias
				sort.Sort(sort.Reverse(sort.IntSlice(joias)))
				// zera valor inicial de joias e quantidade de intems
				for {
					somaX = 0
					somaY = 0
					itensX := 0
					itensY := 0
					// atribui ao valor inicia a diferenca entre o menor e o maior
					if x > y {
						somaX = int(x - y)
					}
					if y > x {
						somaY = int(y - x)
					}
					// realizar a alocação das joias
					for i = 0; i < int(n); i++ {

						if itensX < int(x) && (somaX <= somaY) {
							somaX += joias[i]
							itensX++
						} else {
							somaY += joias[i]
							itensY++
						}
					}
					// verifica se foi possivel
					if somaX == somaY || cont1 == quantidade-1 {
						break
					} else { // caso contrario troca o valor inicial
						c := cont1 + 1
						aux := joias[cont1]
						joias[cont1] = joias[c]
						joias[c] = aux
						cont1++
					}
				}
				cont1 = 0
				// Concatena a string e o valor inteiro em uma única string
				escritaTeste := fmt.Sprintf("Teste %d\n", tests)
				tests++
				if somaX == somaY {
					saida = fmt.Sprintf("S\n\n")
				} else {
					saida = fmt.Sprintf("N\n\n")
				}

				// Escreve uma string no arquivo
				_, err = arquivo.Write([]byte(escritaTeste))
				if err != nil {
					fmt.Println("Erro ao escrever no arquivo:", err)
					return
				}
				_, err = arquivo.Write([]byte(saida))
				if err != nil {
					fmt.Println("Erro ao escrever no arquivo:", err)
					return
				}
				cont++

			}

		}
	}
	// Fecha o Arquivo
	if err := file.Close(); err != nil {
		panic(err)
	}

}
