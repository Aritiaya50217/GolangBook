package main

import "fmt"

type Node struct {
	data string
	prev *Node
	next *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func initial() {
	head.data = "A"
	head.prev = nil
	head.next = nil

	var nodeB *Node = &Node{data: "B", prev: head, next: nil}
	head.next = nodeB

	var nodeC *Node = &Node{data: "C", prev: nodeB, next: nil}
	nodeB.next = nodeC

	tail.data = "D"
	tail.prev = nodeC
	tail.next = head

	nodeC.next = tail
	head.prev = tail

}

func removeNode(removePosition int) {
	var p = head
	var i = 0
	for {
		if p.next == nil || i >= removePosition-1 {
			break
		}
		p = p.next
		i++
	}
	// save the node you want to delete
	var temp = p.next
	// previous node next point to next of delete the node
	p.next = p.next.next
	p.next.prev = p
	// set the delete node next to null
	temp.next = nil
	temp.prev = nil
}

func output() {
	var p = head
	for {
		fmt.Printf("%s-> ", p.data)
		p = p.next
		if p == head {
			break
		}
	}
	fmt.Printf("%s ", p.data)
	fmt.Printf("End\n")

	p = tail
	for {
		fmt.Printf("%s ->", p.data)
		p = p.prev
		if p == tail {
			break
		}
	}
	fmt.Printf("%s ", p.data)
	fmt.Printf("Start\n\n")
}

func main() {
	initial()
	fmt.Printf("Delete a new node C at index 2 : \n")
	removeNode(2)
	output()
}

/*


			   head									p
								- next ->
				A	 								B
			|		|			<- prev -		|		|
		  next	   prev						  prev	   next
			|		|							|		|
								- prev ->
	tail		D									C
								<- next -


	 		  head									p
								- next ->
				A	 								B
			|		|			<- prev -		|		|
		  next	   prev						  prev	   next
			|		|							|		|
								- prev ->				    Node temp = p.next
	tail		D									C	--	temp
								<- next -


			   head						p
						- next ->
				A	 					B  		p.next = p.next.next (ย้ายเส้น next จาก C ไปยัง D )
			|		|	<- prev -	/	|
		  next	   prev			next 	prev
			|		|			/		|
						- prev ->
	tail		D						C	--  temp
						<- next -

		   head							p
						- next ->
				A	 					B  		p.next.prev = p (ย้ายเส้น prev จาก C ไปยัง D )
			|		|	<- prev -	/		/
		  next	   prev			next 	prev
			|		|			/		/

	tail		D						C	--  temp
						<- next -

	 		  head						p
						- next ->
				A	 					B
			|		|	<- prev -	/		/
		  next	   prev			next 	prev
			|		|			/		/
												temp.next = nil
												temp.prev = nil
	tail		D						C	--  	temp  (ถูกลบ)


*/
