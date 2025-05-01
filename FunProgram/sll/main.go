// Creating a singly linked list program 

package main

import "fmt"

// here what we have to do 
type Node struct {
    Data int
    Next *Node
}

// Now we have to create a method associated with the Node struct and also with the another SLL struct ok.

type CLL struct {
    Head *Node
}

func (cll *CLL) insertAtFirst(data int){
    if cll.Head == nil{
        node := Node{
            Data: data,
            Next: nil,
        }
        cll.Head = &node
    } else {
        node := Node{
            Data: data,
            Next: cll.Head,
        }
        cll.Head = &node
    }
}

func (cll *CLL) insertAtLast(data int){
    node := Node{
       Data: data,
       Next: nil,
    }
    if cll.Head == nil{
        cll.Head = &node
    } else{
        temp := cll.Head
        for temp.Next != nil{
           temp = temp.Next
        }
        temp.Next = &node
    }
}

func (cll *CLL) insertAfter(refNodeData int, data int){
    node := cll.search(refNodeData)
    if node != nil{
        newNode := Node{
            Data: data,
            Next: node.Next,
        }
        node.Next = &newNode
    }
    
}

// If we have to search a particular node

func (cll CLL) search(data int) *Node {
    if cll.Head != nil{
        temp := cll.Head
        for temp != nil{
            if temp.Data == data{
                return temp
            }
            temp = temp.Next
        }
        return nil
    }
    return nil
}

func (cll CLL) traverse(){
    if cll.Head != nil{
        temp := cll.Head
        for temp != nil{
            fmt.Printf("%d\t", temp.Data)
            temp = temp.Next
        }
    }
}

func main(){
    mylist := CLL{
        Head: nil,
    }
    mylist.insertAtFirst(10)
    mylist.insertAtFirst(20)
    mylist.insertAtFirst(30)
    mylist.insertAtLast(5)
    mylist.insertAfter(20, 100)
    mylist.traverse()
    fmt.Println(node)
}
