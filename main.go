package main

import(
	"fmt"
)
func main(){

	stack:=Stack{} //создаём новый стек

	//Пример использования стека
	stack.SPush("1")
	stack.SPush("2")
	stack.SPush("3")

	fmt.Println(stack.SPop())
	fmt.Println(stack.SPop())
	fmt.Println(stack.SPop())
	fmt.Println(stack.SPop())
	fmt.Println(" ")


	queue:=Queue{} //создание новой очереди

	//Пример использования очереди
	queue.Qpop("1")
	queue.Qpop("2") 
	queue.Qpop("3")

	fmt.Println(queue.Qdel())
	fmt.Println(queue.Qdel())
	fmt.Println(queue.Qdel())
	fmt.Println(queue.Qdel())


	arr:=NewArray() //создание нового массива //правильный порядок массива и ограничение

		//Пример использования массива
		arr.ArAdd("1")
		arr.ArAdd("2")
		arr.ArAdd("3")
	
		for i:=0; i<arr.ALength(); i++{
			value, err:=arr.AGet(i)
			if err!=nil{
				fmt.Println("Error:", err)
				break
			}
			fmt.Println("Element at index", i, ":", value)
		}
	
		//Изменение элемента по индексу (меняем значение по индексу)
		err:=arr.ASet(1,"99")
		if err !=nil{
			fmt.Println("Error:", err)
		}
	
		//Вывод элемента после изменения 
		value, err:=arr.AGet(1)
		if err!=nil {
			fmt.Println("Error:", err)
		} else{
			fmt.Println("Updated element at index 1:", value)
		}

		// Создаем экземпляр HashMap
	hmap := HashMap{}

	// Вставляем пару ключ-значение
	err = hmap.Insert("n1", "a")
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

	sll := LinkedList{} //создание нового односвязного списка

	// Пример использования односвязного списка
	sll.LSAdd("A")
	sll.LSAdd("B")
	sll.LSAdd("C")

	sll.LSDisplay()

	err = sll.LSDel("B")
	if err != nil {
		fmt.Println(err)
	}

	sll.LSDisplay() //отображение списка 


	dll := DoublyLinkedList{} //создание двусвязного списка

	// Пример использования двусвязного списка
	dll.LDAdd("X")
	dll.LDAdd("Y")
	dll.LDAdd("Z")

	dll.LDDisplay()

	err = dll.LDDel("Y")
	if err != nil {
		fmt.Println(err)
	}

	dll.LDDisplay()
	}