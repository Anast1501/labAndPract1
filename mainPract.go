package main

import (
	"fmt"
)

func main() {

stack := Stack{} 
queue := Queue{}
hashTable := HashMap{}
set := NewSet()

for {
	var input string
	var key string
	fmt.Scanf("%s", &input)
	switch input {
		//Стек
	case "SPUSH":
		fmt.Scanf("%s", &input)
		stack.SPush(input)
	case "SPOP":
		fmt.Println(stack.SPop())
		//Очередь
	case "QPUSH":
		fmt.Scanf("%s", &input)
		queue.Qpush(input)
	case "QPOP":
		fmt.Println(queue.Qpop())
		//Хэш-таблица
	case "HPUSH":
		fmt.Scanf("%s %s", &key, &input)
		hashTable.Insert(key, input)
	case "HGET":                        //вывод по ключу 
		fmt.Scanf("%s", &key)
		fmt.Println(hashTable.HGet(key))
	case "HDEL":
		fmt.Scanf("%s", &key)
		hashTable.HDel(key)
		//Множество
	case "SETADD":
		fmt.Scanf("%s", &input)
		set.SetAdd(input)
	case "SETREMOVE":
		fmt.Scanf("%s", &input)
		set.SetRemove(input)
	case "SETCONTAINS":          //проверка содержания какого-либо элемента в множестве
		fmt.Scanf("%s", &input)
		fmt.Println(set.SetContains(input))
	case "SETPRINT":
		set.SetPrint()
		fmt.Println("")
	}
	}
}
