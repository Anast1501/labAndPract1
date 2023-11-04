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

//LDAddHead добавляет новый узел с данными в начало списка 
func (dll *DoublyLinkedList) LDAddHead(data string) {
    //Создание нового узла 
    newNode:=&LDNode{data: data, next: dll.head, prev: nil}
    //Если список пуст, устанавливаем новый узел как головой и хвостовой узел
    if dll.head == nil {
        dll.head = newNode
        dll.tail = newNode
    } else {
        dll.head.prev = newNode
        dll.head = newNode
    }
}

//LDDelHead удаляет узел из начала списка
func (dll *DoublyLinkedList) LDDelHead() error {
    if dll.head==nil {
        return errors.New("list is empty")
}

//Если удаляемый узел является единственный в списке
if dll.head == dll.tail {
    dll.head=nil
    dll.tail=nil
    return nil
}

dll.head = dll.head.next
dll.head.prev = nil
return nil
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

//LDDelTail удаляет узел с конца списка
func (dll *DoublyLinkedList) LDDelTail() error {
    if dll.tail == nil {
        return errors.New("List is empty")
    }

    // Если удаляемый узел - единственный в списке
    if dll.head == dll.tail {
        dll.head = nil
        dll.tail = nil
        return nil
    }

    // Удаление последнего элемента
    prevTail := dll.tail.prev
    prevTail.next = nil
    dll.tail = prevTail
    return nil
}

// Remove удаляет узел с указанными данными из списка. (по значению)
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

