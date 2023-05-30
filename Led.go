// Autor: David Antonio Brocardo
package main

import (	
    "fmt"    
    "strconv"
)

func main() {
  // Declaracao de variaveis 
	var n, i, contLed int
  var valor string
  
	fmt.Scanf("%d", &n)
	for i = 0; i < n; i++ {
		fmt.Scanf("%s", &valor)
		for i := 0; i < len(valor); i++ {
			digito, _ := strconv.Atoi(string(valor[i])) // Funcao responsavel por realizar a leitura de um digito por vez retornando um inteiro
      
			if digito == 1 {
        
				contLed = contLed + 2
			}
			if digito == 2 {
				contLed = contLed + 5
			}
			if digito == 3 {
				contLed = contLed + 5
			}
			if digito == 4 {
				contLed = contLed + 4
			}
			if digito == 5 {
				contLed = contLed + 5
			}
			if digito == 6 {
				contLed = contLed + 6
			}
			if digito == 7 {
				contLed = contLed + 3
			}
			if digito == 8 {
				contLed = contLed + 7
			}
			if digito == 9 {
				contLed = contLed + 6
			}
			if digito == 0 {
				contLed = contLed + 6
			}

		}
    fmt.Printf("%d leds\n", contLed)
    contLed = 0
	}
}
