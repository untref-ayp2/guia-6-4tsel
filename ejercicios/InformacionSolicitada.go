package ejercicios

import (
	"fmt"
	"guia6/dictionary"
)

func InformacionSolicitada(entrada dictionary.Dictionary[string, []string]) *dictionary.Dictionary[string, []string]{
	
	informacion := dictionary.NewDictionary[string, []string]()

	entradas := entrada.GetValues()
	fmt.Println(entradas)
	claves := entrada.GetKeys()
	fmt.Println(claves)

	for i := 0; i < len(entradas); i++ {
		
		for j := 0; j < len(entradas[i]); j++ {
			
			if !informacion.Contains(entradas[i][j]) {

				informacion.Put(entradas[i][j], []string{claves[i]})
			} else {

				aux := informacion.Get(entradas[i][j])
				dia := claves[i]
				aux = append(aux, dia)
				
				informacion.Put(entradas[i][j], aux)
			}
		}
	}

	return &informacion
}