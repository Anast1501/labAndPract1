package main

import( //импортирование двух пакетов: errors для создания и возврата ошибок и fmt для вывода информации на консоль.
	"errors" 
)

//Node - это структура, представляющая элемент (узел) стека. Узел содержит два поля: next, которое указывает на следующий узел в стеке, и val, которое хранит целочисленное значение элемента стека.
type Node struct{ //стековая структура
	next *Node
	val string //значения типа int
}

//Стек
type Stack struct{ 
	head *Node
}

//Push добавляет новый элемент в стек
func (s *Stack) Push(value string) {
	newNode := &Node{
		val:  value,
		next: s.head,
	}
	s.head = newNode
}
//Pop удаляет и возвращает верхний элемент стека
func (s *Stack) Pop() (string, error) {
	if s.IsEmpty(){
		return "", errors.New("Stack is empty") //вызов проверки на пустоту стека
	}
	value:=s.head.val
	s.head=s.head.next
	return value, errors.New("")
}

//IsEmpty возвращает true, если стек пуст, иначе false 
func (s *Stack) IsEmpty() bool{
	return s.head == nil
}

