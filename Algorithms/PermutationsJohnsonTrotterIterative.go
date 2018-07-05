package main

import (
	"fmt"
	"strings"
)

type node struct {
	idx  int
	val  string
	dir  string
	move string
}

var Nodes []node
var permutation []string

func Generate(val string) {
	var largest, largestidx int
	var Node node
	n := len(val)
	movable := true
	permutation = append(permutation, val)
	if n == 1 {
	    fmt.Println("length:",len(permutation),"\nPermutations are:", permutation)
		return
	}
	for i := 0; i < n; i++ { //O(n)
		Node.idx = i
		Node.val = string(val[i])
		Node.dir = "L"

		if i == 0 {
			Node.move = "N"
		} else {
			Node.move = "Y"
		}
		Nodes = append(Nodes, Node)
	}
	fmt.Println("Nodes are:", Nodes)
	for movable {
		//find largest among the movables
		largestidx = -1
		largest = -1
		for i := 0; i < n; i++ {
			if Nodes[i].idx > largest && Nodes[i].move == "Y" {
				largest = Nodes[i].idx
				largestidx = i
			}
		}
		temp := node{}
		//fmt.Println("largest is:", largest, ", largestindex:", largestidx)
		//Change the position
		if largestidx != -1 {
			//find direction of Move
			if Nodes[largestidx].dir == "L" {
				temp = Nodes[largestidx-1]
				Nodes[largestidx-1] = Nodes[largestidx]
				Nodes[largestidx] = temp
			} else {
				temp = Nodes[largestidx+1]
				Nodes[largestidx+1] = Nodes[largestidx]
				Nodes[largestidx] = temp
			}
			//get the result
			permutation = append(permutation, getString(Nodes))
			//get ready for next iteration
			//flip direction and manipulate movable status
			for i := 0; i < n; i++ {

				if Nodes[i].idx > largest {
					if Nodes[i].dir == "L" {
						Nodes[i].dir = "R"
					} else {
						Nodes[i].dir = "L"
					}
				}
				if i == 0 && Nodes[i].dir == "L" {
					Nodes[i].move = "N"
				} else if i == n-1 && Nodes[i].dir == "R" {
					Nodes[i].move = "N"
				} else {
					if Nodes[i].dir == "L" && Nodes[i-1].idx > Nodes[i].idx {
						Nodes[i].move = "N"
					} else if Nodes[i].dir == "R" && Nodes[i+1].idx > Nodes[i].idx {
						Nodes[i].move = "N"
					} else {
						Nodes[i].move = "Y"
					}
				}

			}

		} else {
			movable = false
		}
	}

	fmt.Println("length:",len(permutation),"\nPermutations are:", permutation)

}

func getString(Nodes []node) string {
	val := ""
	for i := range Nodes {
		val = strings.Join([]string{val, Nodes[i].val}, "")
	}
	return val
}

func main() {
	val := "abcde"
	Generate(val)
}

