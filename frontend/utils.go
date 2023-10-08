package main

import "fmt"

func getFromArray[T any](array []T, index int) *T {
	if len(array) <= index {
		return nil
	} else {
		return &array[index]
	}
}

func logFromArray[T any](array []T) {
	n := len(array)
	if n == 0 {
		infoLog("Array is empty")
	} else {
        infoLog("Printing array")
        for _, elem := range array {
            fmt.Print(elem)
            fmt.Print(" ")
        }
        fmt.Println()
	}
}

func Map[T any, M any](vs []T, f func(T) M) []M {
  vsm := make([]M, len(vs))
  for i, v := range vs {
    vsm[i] = f(v)
  }
  return vsm
}

func Filter[T any](target []T, f func(T) bool) []T {
    var store []T
    for i, v := range target {
        if (f(v)) {
            store = append(store, target[i])
        }
    }
    return store
}


