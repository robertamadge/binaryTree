package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
	Len  int
}

//String só para saber o valor do node e converter para string
func (n Node) String() string {
	return strconv.Itoa(n.Value)
}

//Searching a BinanySearchTree e convertendo para string, se não vai aparecer o valor na memoria
func (b BinarySearchTree) String() string {
	//sb = string builder, qualquer mudança que eu fizer ela reflete de volta no sb
	sb := strings.Builder{}
	b.inAscendOrder(&sb)
	//return o resultado da search inAscendOrderByNode
	return sb.String()
}

func (b BinarySearchTree) inAscendOrder(sb *strings.Builder) {
	b.inAscendOrderByNode(sb, b.Root)
}

//Onde acontece a search pela arvore e monta na ordem crescente
func (b BinarySearchTree) inAscendOrderByNode(sb *strings.Builder, root *Node) {
	//Se for nil ja estamos no final da arvore
	if root == nil {
		return
	}
	//Procura do lado esquerdo, valores menores
	b.inAscendOrderByNode(sb, root.Left)
	sb.WriteString(fmt.Sprintf("%s ", root))
	//Direito
	b.inAscendOrderByNode(sb, root.Right)
}

//Adicionar primeiro como nó o primeiro
func (b *BinarySearchTree) add(value int) {
	b.Root = b.addByNode(b.Root, value)
	b.Len++
}

//Médoto para adicionar
func (b *BinarySearchTree) addByNode(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value}
	}

	if value < root.Value {
		root.Left = b.addByNode(root.Left, value)
	} else {
		root.Right = b.addByNode(root.Right, value)
	}

	return root
}

//Search por um numero especifico, preorder
func (b BinarySearchTree) Search(value int) (*Node, bool) {
	return b.SearchByNode(b.Root, value)
}

//Função Search para procurar, metodo recursivo
func (b BinarySearchTree) SearchByNode(root *Node, value int) (*Node, bool) {
	if root == nil {
		return nil, false
	}

	if value == root.Value {
		return root, true
	} else if value < root.Value {
		return b.SearchByNode(root.Left, value)
	} else {
		return b.SearchByNode(root.Right, value)
	}
}

//Remove, primeiro eu procuro o valor, substituo  e deleto ele de onde ele estava antes
func (b *BinarySearchTree) Remove(value int) {
	b.RemoveByNode(b.Root, value)
}

func (b *BinarySearchTree) RemoveByNode(root *Node, value int) *Node {
	//Find the node to delete
	if root == nil {
		return root
	}

	if value > root.Value {
		root.Right = b.RemoveByNode(root.Right, value)
	} else if value < root.Value {
		root.Left = b.RemoveByNode(root.Left, value)
	} else {
		//Se não existe a arvore da esquerda
		if root.Left == nil {
			return root.Right
		} else {
			//max value
			temp := root.Left
			for temp.Right != nil {
				temp = temp.Right
			}
			//Substitui pelo maior valor que eu achei do lado esquerdo
			root.Value = temp.Value

			root.Left = b.RemoveByNode(root.Left, value)
		}
	}
	return root
}

func main() {
	n := &Node{3, nil, nil}
	n.Left = &Node{2, nil, nil}
	n.Right = &Node{4, nil, nil}

	//Testanto search em ordem crescente
	bst := BinarySearchTree{
		Root: n,
		Len:  1,
	}
	fmt.Println(bst)

	//Testanto adicionar
	bst.add(9)
	bst.add(7)
	bst.add(1)
	fmt.Println(bst)

	//Testando search elemento especifico
	fmt.Println(bst.Search(9))
	fmt.Println(bst.Search(8))

	//Remover
	bst.Remove(4)
	fmt.Println(bst)
	bst.Remove(10)
	fmt.Println(bst)
}
