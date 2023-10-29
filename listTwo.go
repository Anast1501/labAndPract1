//Двусвязный список 
package main

import (
	"errors"
	"fmt"
)

// LDNode представляет узел двусвязного списка.
type LDNode struct {
    data string
    next *LDNode //ссылка на след. элемент
    prev *LDNode //ссылка на предыдущий элемент
}

// DoublyLinkedList представляет двусвязный список.
type DoublyLinkedList struct {
    head *LDNode
    tail *LDNode
}

// Add добавляет новый узел с заданными данными в конец списка.
func (dll *DoublyLinkedList) LDAdd(data string) {
    // Создаем новый узел
    newNode := &LDNode{data: data, next: nil, prev: dll.tail}
    // Если список пуст, устанавливаем новый узел как головной и хвостовой узел
    if dll.head == nil {
        dll.head = newNode
        dll.tail = newNode
    } else {
        dll.tail.next = newNode
        dll.tail = newNode
    }
}

// Remove удаляет узел с указанными данными из списка.
func (dll *DoublyLinkedList) LDDel(data string) error {
    current := dll.head
    for current != nil {
        if current.data == data {
            // Если удаляемый узел - головной, переносим голову на следующий узел
            if current.prev != nil {
                current.prev.next = current.next
            } else {
                dll.head = current.next
            }
            // Если удаляемый узел - хвостовой, обновляем хвостовой узел
            if current.next != nil {
                current.next.prev = current.prev
            } else {
                dll.tail = current.prev
            }
            return nil
        }
        current = current.next
    }

    return errors.New("Data not found in the list")
}

// Display выводит содержимое списка.
func (dll *DoublyLinkedList) LDDisplay() {
    current := dll.head
    for current != nil {
        fmt.Print(current.data, " <-> ")
        current = current.next
    }
    fmt.Println("nil")
}

