package main

import (
	"fmt"
	"strings"
)

// Глобальные переменные следует избегать
// var justString string

// Пример из задания.
// Тут будет утечка памяти, т.к. при v[:100] происходит создание нового экземпляра строки с ссылкой на исходный массив байтов v
// Т.е. исходный массив продолжает существование, а значит попадает в heap => GC его не может забрать.
/*func someFunc() {
  v := createHugeString(1 << 10) // вначале не понял, что за lt, а потом как поняял....
  justString = v[:100]
}*/

// Вот так будет лучше, можно использовать buffer + copy, но предпочту strings
func someFunc() string {
    v := createHugeString(1 << 10)
    return strings.Clone(v[:100])
}

func createHugeString(size int) string{
	return strings.Repeat("x", size)
}

func main() {
	fmt.Println(someFunc())
}