package rs

import (
	"encoding/json"
	"math"
	"math/rand"
	"sort"
	"strconv"
	//  "fmt"
	//  "time"
	//  rs "rs/lib"
)

type Graph struct {
	Edges    []*Edge
	StopList []*BusStop
}

type Edge struct {
	Parent  *BusStop
	Child   *BusStop
	Density float64
	Level   int
	Cost    int
}

type Building struct {
	BusStopList []struct {
		Source      string `json:"source"`
		Destination string `json:"destination"`
		Distance    int    `json:"distance"`
	} `json:"busStopList"`
}

const Infinity = int(^uint(0) >> 1)

// AddEdge adds an Edge to the Graph
func (g *Graph) AddEdge(parent, child *BusStop, cost int) {
	edge := &Edge{
		Parent: parent,
		Child:  child,
		Cost:   cost,
	}

	g.Edges = append(g.Edges, edge)
	g.AddNode(parent)
	g.AddNode(child)
}

// AddNode adds a Node to the Graph list of Nodes, if the node is not already existed
// g.Nodes is a map for better caching reasons
func (g *Graph) AddNode(node *BusStop) {
	for _, gNode := range g.StopList {
		if node.Name == gNode.Name {
			return
		}
	}
	g.StopList = append(g.StopList, node)
}

// String returns a string representation of the Graph
func (g *Graph) String() string {
	var s string

	s += "Edges:\n"
	for _, edge := range g.Edges {
		s += edge.Parent.Name + " -> " + edge.Child.Name + " = " + strconv.Itoa(edge.Cost)
		s += "\n"
	}
	s += "\n"

	s += "stopList: "
	i := 0
	for _, node := range g.StopList {
		if i == len(g.StopList)-1 {
			s += node.Name
		} else {
			s += node.Name + ", "
		}
		i++
	}
	s += "\n"

	return s
}

// Dijkstra implements THE Dijkstra algorithm
// Returns the shortest path from startNode to all the other Nodes
func (g *Graph) Dijkstra(startNode *BusStop) (costTable map[*BusStop]int) {

	// First, we instantiate a "Cost Table", it will hold the information:
	// "From startNode, what's is the cost to all the other Nodes?"
	// When initialized, It looks like this:
	// NODE  COST
	//  A     0    // The startNode has always the lowest cost to itself, in this case, 0
	//  B    Inf   // the distance to all the other Nodes are unknown, so we mark as Infinity
	//  C    Inf
	// ...
	costTable = g.NewCostTable(startNode)

	// An empty list of "visited" Nodes. Everytime the algorithm runs on a Node, we add it here
	var visited []*BusStop

	// A loop to visit all Nodes
	for len(visited) != len(g.StopList) {

		// Get closest non visited Node (lower cost) from the costTable
		node := getClosestNonVisitedNode(costTable, visited)

		// Mark Node as visited
		visited = append(visited, node)

		// Get Node's Edges (its neighbors)
		nodeEdges := g.GetNodeEdges(node)

		for _, edge := range nodeEdges {

			// The distance to that neighbor, let's say B is the cost from the costTable + the cost to get there (Edge cost)
			// In the first run, the costTable says it's "0"
			// Plus the actual cost, let's say "5"
			// The distance becomes "5"
			distanceToNeighbor := costTable[node] + edge.Cost

			// If the distance above is lesser than the distance currently in the costTable for that neighbor
			if distanceToNeighbor < costTable[edge.Child] {

				// Update the costTable for that neighbor
				costTable[edge.Child] = distanceToNeighbor
			}
		}
	}

	return costTable
}

// NewCostTable returns an initialized cost table for the Dijkstra algorithm work with
// by default, the lowest cost is assigned to the startNode – so the algorithm starts from there
// all the other Nodes in the Graph receives the Infinity value
func (g *Graph) NewCostTable(startNode *BusStop) map[*BusStop]int {
	costTable := make(map[*BusStop]int)
	costTable[startNode] = 0

	for _, node := range g.StopList {
		if node != startNode {
			costTable[node] = 100
		}
	}

	return costTable
}

// GetNodeEdges returns all the Edges that start with the specified Node
// In other terms, returns all the Edges connecting to the Node's neighbors
func (g *Graph) GetNodeEdges(node *BusStop) (edges []*Edge) {
	for _, edge := range g.Edges {
		if edge.Parent == node {
			edges = append(edges, edge)
		}
	}

	return edges
}

// func NewCar(stopList []*BusStop) *Car {
// 	var p *Car
// 	p = new(Car)
// 	p.Parent = "a"
// 	p.Child = "b"
// 	return p
// }

func (g *Graph) GetSpeed(parent *BusStop, child *BusStop) (speed int) {
	for _, edge := range g.Edges {
		if edge.Parent == parent && edge.Child == child {
			return 50 - (edge.Level * 10)
		}
	}
	return -1
}

func (g *Graph) GenerateTraffic(randomCar []*Car, parent *BusStop, child *BusStop) {
	// fmt.Println("GenerateTraffic")
	// random1 car capacity
	// child is nil, only one node as a parameter (Generate traffic near parent node)
	if child == nil {
		for _, edge := range g.Edges {
			if edge.Child == parent {
				temp := rand.Intn(65) + len(randomCar)
				edge.Density = float64(temp / 2)
				edge.Level = int(math.Floor(edge.Density / 10))
			}
		}
		// parent and child are nil, there is no node as a parameter (Generate traffic the whole road)
		if parent == nil {
			for _, edge := range g.Edges {
				temp := rand.Intn(65) + len(randomCar)
				edge.Density = float64(temp / 2)
				edge.Level = int(math.Floor(edge.Density / 10))
			}
		}
	} else {
		for _, edge := range g.Edges {
			if edge.Parent == parent && edge.Child == child {
				temp := rand.Intn(65) + len(randomCar)
				edge.Density = float64(temp / 2)
				edge.Level = int(math.Floor(edge.Density / 10))
			}
		}
	}
}

// getClosestNonVisitedNode returns the closest Node (with the lower cost) from the costTable
// **if the node hasn't been visited yet**
func getClosestNonVisitedNode(costTable map[*BusStop]int, visited []*BusStop) *BusStop {
	type CostTableToSort struct {
		Node *BusStop
		Cost int
	}
	var sorted []CostTableToSort

	// Verify if the Node has been visited already
	for node, cost := range costTable {
		var isVisited bool
		for _, visitedNode := range visited {
			if node == visitedNode {
				isVisited = true
			}
		}
		// If not, add them to the sorted slice
		if !isVisited {
			sorted = append(sorted, CostTableToSort{node, cost})
		}
	}

	// We need the Node with the lower cost from the costTable
	// So it's important to sort it
	// Here I'm using an anonymous struct to make it easier to sort a map
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Cost < sorted[j].Cost
	})

	return sorted[0].Node
}

func (g *Graph) GenerateBuildingBusStop(stopList *[]*BusStop, buildingInputJson string) {
	var buildingInput Building
	err := json.Unmarshal([]byte(buildingInputJson), &buildingInput)
	if err != nil {
		panic(err)
	}
	initBuildingList := make(map[string]*BusStop)
	isSourceInit := false
	isDestinationInit := false
	for _, building := range buildingInput.BusStopList {
		count := 0
		for initBuilding := range initBuildingList {
			if building.Source == initBuilding {
				isSourceInit = true
				count++
			}
			if building.Destination == initBuilding {
				isDestinationInit = true
				count++
			}
			if count == 2 {
				break
			}
		}
		if isSourceInit != true {
			tempSourceBuilding := BusStop{Name: building.Source}
			initBuildingList[building.Source] = &tempSourceBuilding
			*stopList = append(*stopList, &tempSourceBuilding)
		}
		if isDestinationInit != true {
			tempDestinationBuilding := BusStop{Name: building.Destination}
			initBuildingList[building.Destination] = &tempDestinationBuilding
			*stopList = append(*stopList, &tempDestinationBuilding)
		}
		// Add edge
		g.AddEdge(initBuildingList[building.Source], initBuildingList[building.Destination], building.Distance)
		g.AddEdge(initBuildingList[building.Destination], initBuildingList[building.Source], building.Distance)
		g.GenerateTraffic(CarGroup(), initBuildingList[building.Source], initBuildingList[building.Destination])
		g.GenerateTraffic(CarGroup(), initBuildingList[building.Destination], initBuildingList[building.Source])
	}
}
