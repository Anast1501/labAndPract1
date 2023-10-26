package main

import (
	"fmt"
)

type TreeNode struct { //TreeNode представляет узел бинарного дерева
	Value       int
	Left, Right *TreeNode //каждый узел имеет ссылки на левого и правого потомка
}

type FullBinaryTree struct {  //FullBinaryTree содержит ссылку на корневой узел Root
	Root *TreeNode
}

//NewFullBinaryTree создаёт и возвращает новое пустое бинарное дерево
func NewFullBinaryTree() *FullBinaryTree { 
	return &FullBinaryTree{}
}

//Insert вставляет новое значение в бинарное дерево
func (t *FullBinaryTree) Insert(value int) {  	
	t.Root = insertRec(t.Root, value) //insertRec правильно вставляет (в логическом порядке)
}

func insertRec(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return &TreeNode{Value: value, Left: nil, Right: nil}
	}

	if value < root.Value {
		root.Left = insertRec(root.Left, value)
	} else if value > root.Value {
		root.Right = insertRec(root.Right, value)
	}

	return root
}

//Удаляет значение из бинарного дерева
func (t *FullBinaryTree) Delete(value int) {
	t.Root = deleteRec(t.Root, value)
}

func deleteRec(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return root
	}

	if value < root.Value {
		root.Left = deleteRec(root.Left, value)
	} else if value > root.Value {
		root.Right = deleteRec(root.Right, value)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		root.Value = minValue(root.Right)
		root.Right = deleteRec(root.Right, root.Value)
	}

	return root
}

//minValue находит наименьшее значение в поддереве, начиная с заданного корня (наименьший элемент ищется для того, чтобы поменять (перенести) его при удалении элемента)
func minValue(root *TreeNode) int {
	minValue := root.Value
	for root.Left != nil {
		minValue = root.Left.Value
		root = root.Left
	}
	return minValue
}

//Search поиск значения в бинарном дереве и возвращает true, если значение найдено, и false, если нет (поиск элемента есть ли такой элемент или нету)
func (t *FullBinaryTree) Search(value int) bool {
	return searchRec(t.Root, value)
}

func searchRec(root *TreeNode, value int) bool {
	if root == nil {
		return false
	}

	if value == root.Value {
		return true
	}

	if value < root.Value {
		return searchRec(root.Left, value)
	} else {
		return searchRec(root.Right, value)
	}
}

/*
func (t *FullBinaryTree) InOrderTraversal() {
	inOrderRec(t.Root)
	fmt.Println()
}

func inOrderRec(root *TreeNode) {
	if root != nil {
		inOrderRec(root.Left)
		fmt.Printf("%d ", root.Value)
		inOrderRec(root.Right)
	}
}
*/

//PrintTree вывод дерева
func PrintTree(node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		fmt.Println("Empty tree")
		return
	}

	if node.Right != nil {
		PrintTree(node.Right, prefix+mapBool(isLeft, "│   ", "    "), false)
	}

	fmt.Print(prefix)
	fmt.Print(mapBool(isLeft, "└── ", "┌── "))
	fmt.Println(node.Value)

	if node.Left != nil {
		PrintTree(node.Left, prefix+mapBool(isLeft, "    ", "│   "), true)
	}
}

//mapBool проверяет слева или справа находится предок (корень, след. значение)
func mapBool(b bool, trueVal, falseVal string) string {
	if b {
		return trueVal
	}
	return falseVal
}


//isFullBinaryTree проверка бинарности дерева на полноту
func isFullBinaryTree(root *TreeNode) bool {
    if root == nil {
        return true
    }

    // Если у узла нет дочерних узлов, он является листом
    if root.Left == nil && root.Right == nil {
        return true
    }

    // Если у узла есть оба дочерних узла, рекурсивно проверяем их
    if root.Left != nil && root.Right != nil {
        return isFullBinaryTree(root.Left) && isFullBinaryTree(root.Right)
    }

    // Если у узла есть только один дочерний узел, дерево не является полным бинарным
    return false
}
