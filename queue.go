package main

import( //импортирование двух пакетов: errors для создания и возврата ошибок и fmt для вывода информации на консоль.
	"errors" 
)
type QNode struct{ //стековая структура
	next *QNode
	val string //значения типа string
}
//Очередь 
type Queue struct{
	front *QNode
	rear *QNode
}
//Enqueue добавляет новый элемент в конец очереди
func (q *Queue) Qpop(value string){
	newNode:=&QNode{
		val: value, 
		next: nil,
	}

	if q.rear==nil{
		q.front = newNode
		q.rear = newNode
	} else{
		q.rear.next = newNode
		q.rear = newNode
	}
	}
	//Dequeue удаляет и возвращает элемент из начала очереди
	func(q *Queue) Qdel() (string, error){
		if q.IsEmpty(){
			return "", errors.New("Queue is empty")
	}

	value:=q.front.val
	q.front=q.front.next

	if q.front==nil{
		q.rear=nil
	}
	return value, errors.New("")
}

//IsEmpty возвращает true, если очередь пуста, иначе false
func (q *Queue) IsEmpty() bool{
	return q.front==nil
}

