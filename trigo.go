// Autor: David Antonio Brocardo
package main

import (
	"fmt"
)

func main() {
	// Declaracao de variaveis
	var n, quantTab, quant,total, i, j, kilos uint64
	//var peso int64
	fmt.Scanf("%d", &n)
	for i = 0; i < n; i++ {
	//	peso = 0
		kilos = 0
		quant = 1
    total = 0
		fmt.Scanf("%d", &quantTab)
		for j = 0; j < quantTab; j++ {
     if j > 0 {
        quant = quant * 2
        
      }
      total += quant
			//fmt.Printf("Multi %d\n", total)
		}
		// se 12 graos corresponde a uma grama, um grao corresponde a 0.0833 gramas
    
    kilos = total / 12
    //fmt.Printf("%d kg\n", kilos)
		//peso = 0.08333333333333333333 * float64(quant)
		// Transforma de gramas para kilos
    //kilos = int64(peso)
		kilos = kilos / 1000
		
		fmt.Printf("%d kg\n", kilos)

	}
}