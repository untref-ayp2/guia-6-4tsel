package ejercicios

import (
	"guia6/dictionary"
)

func Frecuencia(s string) *dictionary.Dictionary[string, int]{
	
	d := dictionary.NewDictionary[string, int]()

	for _, v := range s {
		
		aux := string(v)

		if d.Contains(aux) {

			count := d.Get(aux)
			d.Put(string(v), count+1)
		} else {

			d.Put(string(v), 1)
		}
	}

	return &d
}