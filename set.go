package main

import "fmt"

//Определение структуры Set
type Set struct {
	elements map[string]bool //Используем map для хранения элементов множества
}

//Операция для создания нового множества 
func NewSet() *Set {
	s := &Set{}
	s.elements = make(map[string]bool)
	return s
	}

//Операция для создания нового множества 
func (s *Set) SetAdd(element string) {
	s.elements[element] = true
}

//Операция для удаления элемента из множества
func (s *Set) SetRemove(element string) {
	delete(s.elements, element)
}

//Операция для проверки наличия элемента в множестве
func (s *Set) SetContains(element string) bool {
	return s.elements[element]
}

//Операция для получения размера множества 
func (s *Set) SetSize() int {
	return len(s.elements)
}

//Операция для вывода элементов множества
func (s *Set) SetPrint() {
	for element := range s.elements {
		fmt.Print(" ",element)
	}
}










/*
func main() {

	//Создаём новое множество
	mySet := NewSet()

	//Добавляем элементы
	mySet.SetAdd(1)
	mySet.SetAdd(2)
	mySet.SetAdd(3)



	//Выводим элементы множества
	fmt.Println("Elements in the set:")
	mySet.SetPrint()
	fmt.Println("")
	//Проверяем наличие элемента 
	fmt.Println("Contains 2:", mySet.SetContains(2))

	//Удаление элемента
	mySet.SetRemove(2)

	//Выводим обновлённые элементы
	fmt.Println("Elements in the set after removal:")
	mySet.SetPrint()
	fmt.Println("")

	//Получаем размер множества
	fmt.Println("Size of the set: ", mySet.SetSize())

}
*/