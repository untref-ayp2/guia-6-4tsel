package ejercicios

import (
	"guia6/dictionary"
	"strings"
)

func Traducir(texto string, dic dictionary.Dictionary[string, string]) string{
	
	arrayS := strings.Split(texto, " ")
	resultado := ""

	for i, v := range arrayS {
		
		if dic.Contains(v) && i == len(arrayS)-1{
			
			resultado += dic.Get(v)
		} else if dic.Contains(v){

			resultado += dic.Get(v) + " "
		} else if !dic.Contains(v) && i == len(arrayS)-1{
			
			resultado += "error"
		} else {

			resultado += "error "
		}
	}

	return resultado
}