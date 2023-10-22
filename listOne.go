//Односвязный список 
package main

import (
	"errors"
	"fmt"
)

// Node представляет узел односвязного списка.
 type LNode struct {
	data string
	next *LNode
}

// LinkedList представляет односвязный список.
type LinkedList struct {
	head *LNode
}

// Add добавляет новый узел с заданными данными в конец списка.
func (ll *LinkedList) LSAdd(data string) {
    // Создаем новый узел
    newNode := &LNode{data: data, next: nil}
    // Если список пуст, устанавливаем новый узел как головной узел
    if ll.head == nil {
        ll.head = newNode
        return
    }

    // Иначе, находим последний узел и устанавливаем новый узел как следующий за последним
    current := ll.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}

// Remove удаляет узел с указанными данными из списка.
func (ll *LinkedList) LSDel(data string) error {
    // Проверяем, не пуст ли список
    if ll.head == nil {
        return errors.New("List is empty")
    }

    // Если удаляемый узел - головной, переносим голову на следующий узел
    if ll.head.data == data {
        ll.head = ll.head.next
        return nil
    }

    // Иначе, ищем узел для удаления и переключаем указатели
    current := ll.head
    for current.next != nil {
        if current.next.data == data {
            current.next = current.next.next
            return nil
        }
        current = current.next
    }

    return errors.New("Data not found in the list")
}

// Display выводит содержимое списка.
func (ll *LinkedList) LSDisplay() {
    current := ll.head
    for current != nil {
        fmt.Print(current.data, " -> ")
        current = current.next
    }
    fmt.Println("nil")
}

