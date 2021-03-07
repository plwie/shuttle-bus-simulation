package main

import (
	"fmt"
	// "math"
)

type Graph struct{
	pointer []*Point

}

type Point struct{
	key int
	adjacent []*Point
	Lat, Long float64
}

// type Coordinate struct {
// 	X int
// 	Y int
// }

// func New(x int, y int) Coordinate {
// 	return Coordinate{x, y}
// }

func (g *Graph) AddDistance(k int, i float64, j float64){ //test.AddDistance(1, 45, 60)
	if contains(g.pointer,k){
		g.pointer = append(g.pointer, &Point{Lat:i, Long: j, key:k}) 

	} else{
	err := fmt.Errorf("Point %v Not exist", k)
		fmt.Println(err.Error())
	}

}


// func contains(s []*Point, k int) bool{
// 	for _, v := range s{
// 		if k == v.key{
// 			return true
// 		}
// 	}
// 	return false
// }

func (g *Graph) AddPoint(k int){
	if contains(g.pointer,k){
		err := fmt.Errorf("Point %v not added the key already exist", k)
		fmt.Println(err.Error())
	} else{
	g.pointer = append(g.pointer, &Point{key:k}) 

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

func (g *Graph) AddPath(from, to int){  //test.AddEdge(1,2)
	// get vertex or road
	fromPoint := g.getPoint(from)
	toPoint := g.getPoint(to)

	if fromPoint == nil || toPoint == nil{
		err := fmt.Errorf("Invalid (%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromPoint.adjacent,to) {
		err := fmt.Errorf("Path already Exist(%v-->%v)", from, to)
		fmt.Println(err.Error())
	} else{
	fromPoint.adjacent = append(fromPoint.adjacent,toPoint)
	// fmt.Println("%v", fromVertex)
	}

}

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
		// fmt.Println(v.Lat)
		fmt.Printf("\nPoint %v (%f,%f) --> Point " , v.key , v.Lat, v.Long)
		
		for _, r := range v.adjacent{

			fmt.Printf("%v (%f,%f)", r.key,r.Lat,r.Long)
			// fmt.Printf("%v",v)
		}
	}
	fmt.Println()
}

// var m map[string]Vertex

func main() {
	test := &Graph{}

	

	for i:=1; i<7; i++{
		test.AddPoint(i)
		// test.AddEdge(i,i+1)

	}

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
	// test.AddDistance(1, 45, 60)
	// test.AddDistance(2, 90, 120)
	// test.AddDistance(3,140, 200)
	// test.AddDistance(4,170, 250)
	// test.AddDistance(5,200, 300)
	// g.getVertex(1)
	test.AddPath(1,2)
	test.AddPath(2,3)
	test.AddPath(3,4)
	test.AddPath(4,5)
	test.AddPath(5,6)
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