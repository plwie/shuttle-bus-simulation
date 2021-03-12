package main

import (
	"fmt"
	// "math"
)

type Graph struct{
	pointer []*Point
	Edges []*Edge
	Point map[*Point]bool

}

type Point struct{
	key int
	adjacent []*Point
	// edge int *Edge
}
type Edge struct {
	From *Point
	To  *Point
	Cost int

}

// type Coordinate struct {
// 	X int
// 	Y int
// }

// type myVertex struct {
// 	key    int
// 	outTo  map[string]float64
// 	inFrom map[string]float64
// }

// type myEdge struct {
// 	from   int
// 	to     int
// 	weight float64
// }




// type Edge struct {
// 	Source      *Vertex
// 	Destination *Vertex
// }


// func New(x int, y int) Coordinate {
// 	return Coordinate{x, y}
// }


// func (edge *Edge) getCapacity() int{
// 	return edge.capacity
// }

const Infinity = int(^uint(0) >> 1)

// func (g *Graph) AddCoordinate(k int, i int){ //test.AddDistance(1, 45, 60)
// 	if contains(g.pointer,k){
// 		g.pointer = append(g.pointer, &Point{key:k})


// 	} else{
// 	err := fmt.Errorf("Point %v Not exist", k)
// 		fmt.Println(err.Error())
// 	}

// }



func (g *Graph) AddPoint(k *Point){ 
	if contains(g.pointer,k.key){
		err := fmt.Errorf("Point %v not added the key already exist", k)
		fmt.Println(err.Error())
	} else{
	g.pointer = append(g.pointer, &Point{key:k.key}) 

	}
}

func (g*Graph) getPoint(k int) *Point{
	for i, v := range g.pointer{
		if v.key == k{
			return g.pointer[i]
		}
	}
	return nil
}


func (g *Graph) AddPath(from, to *Point, cost int){  //test.AddEdge(1,2)
	// get vertex or road
	edge := &Edge{
		
		From: from,
		To:  to,
		Cost:   cost,
	}
	
	g.Edges = append(g.Edges, edge)
	g.AddPoint(from)
	g.AddPoint(to)
}

	// 	fromPoint := g.getPoint(from)
	// 	toPoint := g.getPoint(to)
	// }
	
	// // if capacity == math.Inf(-1) {
	// // 	return fmt.Errorf("-inf weight is reserved for internal usage")
	// // }

	// if fromPoint == nil || toPoint == nil{
	// 	err := fmt.Errorf("Invalid (%v-->%v)", from, to)
	// 	fmt.Println(err.Error())
	// } else if contains(fromPoint.adjacent,to) {
	// 	err := fmt.Errorf("Path already Exist(%v-->%v)", from, to)
	// 	fmt.Println(err.Error())
	// } else{
	// fromPoint.adjacent = append(fromPoint.adjacent,toPoint)
	// // fmt.Println("%v", fromVertex)
	// }

// }

func contains(s []*Point, k int) bool{
	for _, v := range s{
		if k == v.key{
			return true
		}
	}
	return false
}

func (g*Graph) Print(){
	for _, v := range g.pointer {	
		if v.adjacent == nil{
		break
		}
	// fmt.Printf("Available Destination")
		fmt.Printf("\nPoint %v (%v) --> Point " , v.key)
		
		for _, r := range v.adjacent{

			fmt.Printf("%v (%v)", r.key)
			// fmt.Printf("%v",v)
		}
	}
	fmt.Println()
}

// var m map[string]Vertex

func main() {
	// test := &Graph{}
	a := &Point{1, nil}
	b := &Point{2, nil}
	c := &Point{3, nil}
	d := &Point{4, nil}
	e := &Point{5, nil}

	

	// for i:=1; i<7; i++{
	test := &Graph{}
	// test.AddPoint(a)
	// test.AddPoint(b)
	// test.AddPoint(c)
	// test.AddPoint(d)
	// test.AddPoint(e)
		// test.AddEdge(i,i+1)

	// }

	// test.AddDistance(1, 45, 60)
	// test.AddDistance(2, 90, 120)
	// test.AddDistance(3,140, 200)
	// test.AddDistance(4,170, 250)
	// test.AddDistance(5,200, 300)

	// m = make(map[string]Vertex)
	// m["HM Stops"] = Vertex{
	// 	1,
	// 	40.68433, -74.39967,
	// }
	// test.AddCoordinate(1, -2.0, 2)
	// test.AddCoordinate(2, -2, -2)
	// test.AddCoordinate(3,0, 3)
	// test.AddCoordinate(4,2, -1)
	// test.AddCoordinate(5,2, 3)
	// // g.getVertex(1)
	
	test.AddPath(a,b,5)
	test.AddPath(a,c,7)
	test.AddPath(b,c,2)
	test.AddPath(b,d,8)
	test.AddPath(c,d,6)
	test.AddPath(c,e,10)
	// test.AddPath(1,2)
	

	// test.AddDistance(1, 45, 60)
	// test.AddDistance(2, 90, 120)
	// test.AddDistance(3,140, 200)
	// test.AddDistance(4,170, 250)
	// test.AddDistance(5,200, 300)

	// // // f := New(36,78)

	test.Print()
	// fmt.Println(m["HM Stops"])
	// fmt.Println(g.getVertex(s1))
}