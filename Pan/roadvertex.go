package main

import (
	"fmt")

type Graph struct{
	vertices []*Vertex

}

type Vertex struct{
	key int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int){
	if contains(g.vertices,k){
		err := fmt.Errorf("Vertex %v not added the key already exist", k)
		fmt.Println(err.Error())
	} else{
	g.vertices = append(g.vertices, &Vertex{key:k}) 
	}
}

func (g *Graph) AddEdge(from, to int){
	// get vertex or road
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)

	if fromVertex == nil || toVertex == nil{
		err := fmt.Errorf("Invalid (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent,to) {
		err := fmt.Errorf("Edge already Exist(%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else{
	fromVertex.adjacent = append(fromVertex.adjacent,toVertex)
	}

}

func (g*Graph) getVertex(k int) *Vertex{
	for i, v := range g.vertices{
		if v.key == k{
			return g.vertices[i]
		}
	}
	return nil
}
	

func contains(s []*Vertex, k int) bool{
	for _, v := range s{
		if k == v.key{
			return true
		}
	}
	return false
}


func (g*Graph) Print(){
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %v : ", v.key)
		for _, v := range v.adjacent{
			fmt.Printf("%v", v.key)
			// fmt.Printf(v)
		}
	}
	fmt.Println()
}


func main() {
	test := &Graph{}

	for i:=1; i<7; i++{
		test.AddVertex(i)

	}

	test.AddEdge(1,2)
	test.AddEdge(2,3)
	test.AddEdge(3,4)
	test.AddEdge(4,5)
	test.AddEdge(5,6)

	test.Print()
}