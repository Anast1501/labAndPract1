package main

import(
	"fmt"
)
func main(){

	stack:=&Stack{} //создаём новый стек

	//Пример использования стека
	stack.Push("1")
	stack.Push("2")
	stack.Push("3")

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(" ")


	queue:=&Queue{} //создание новой очереди

	//Пример использования очереди
	queue.Enqueue("1")
	queue.Enqueue("2") 
	queue.Enqueue("3")

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())


	arr:=NewArray() //создание нового массива //правильный порядок массива и ограничение

		//Пример использования массива
		arr.Append("1")
		arr.Append("2")
		arr.Append("3")
	
		for i:=0; i<arr.Length(); i++{
			value, err:=arr.Get(i)
			if err!=nil{
				fmt.Println("Error:", err)
				break
			}
			fmt.Println("Element at index", i, ":", value)
		}
	
		//Изменение элемента по индексу (меняем значение по индексу)
		err:=arr.Set(1,"99")
		if err !=nil{
			fmt.Println("Error:", err)
		}
	
		//Вывод элемента после изменения 
		value, err:=arr.Get(1)
		if err!=nil {
			fmt.Println("Error:", err)
		} else{
			fmt.Println("Updated element at index 1:", value)
		}

		// Создаем экземпляр HashMap
	hmap := HashMap{}

	// Вставляем пару ключ-значение
	err := hmap.Insert("n1", "a")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	err = hmap.Insert("n2", "b")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	err = hmap.Insert("n3", "c")
	if err != nil {
		fmt.Println("Ошибка при вставке:", err)
	}

	// Выводим все элементы через цикл
	for i := 0; i < len(hmap.table); i++ {
		if hmap.table[i] != nil {
			fmt.Printf("Ключ: %s, Значение: %s\n", hmap.table[i].key, hmap.table[i].value)
		}
	}
	}