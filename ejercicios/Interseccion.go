package ejercicios

import (
	"fmt"
	"guia6/dictionary"
	"guia6/linkedlist"
)

func Interseccion(s1 []string, s2 []string) linkedlist.LinkedList[string] {

	dic := dictionary.NewDictionary[string, int]()
	list := linkedlist.NewLinkedList[string]()

	for i := 0; i < len(s1); i++ {

		dic.Put(s1[i], i)
	}

	fmt.Println(dic)

	for _, v := range s2 {
		
		if dic.Contains(v){

			list.InsertAt(v, list.Size())
		}
	}

	return *list
}
