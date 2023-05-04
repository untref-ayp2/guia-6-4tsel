package ejercicios

import (
	"guia6/dictionary"
)

func NotaFinal(dic dictionary.Dictionary[string, []int]) dictionary.Dictionary[string, float32]{

	promedios := dictionary.NewDictionary[string, float32]()
	alumnos := dic.GetKeys()
	notas := dic.GetValues()

	for i := 0; i < len(alumnos); i++ {
		
		var sumaNota float32

		for _, v := range notas[i] {
			
			sumaNota += float32(v)
		}

		promedio := sumaNota / float32(len(notas[i]))
		promedios.Put(alumnos[i], promedio)
	}

	return promedios
}