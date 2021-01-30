package main

import (
	"reflect"
	"testing"
)

func TestMinMaxLessThan(t *testing.T) {
	lower, upper := minMax(1, 2)

	if lower != 1 {
		t.Fatalf("Expected lower to be 1, got %v\n", lower)
	}

	if upper != 2 {
		t.Fatalf("Expected upper to be 2, got %v\n", upper)
	}
}

func TestMinMaxGreaterThan(t *testing.T) {
	lower, upper := minMax(2, 1)

	if lower != 1 {
		t.Fatalf("Expected lower to be 1, got %v\n", lower)
	}

	if upper != 2 {
		t.Fatalf("Expected upper to be 2, got %v\n", upper)
	}
}

func TestMinMaxEqual(t *testing.T) {
	lower, upper := minMax(3, 3)

	if lower != 3 {
		t.Fatalf("Expected lower to be 3, got %v\n", lower)
	}

	if upper != 3 {
		t.Fatalf("Expected upper to be 3, got %v\n", upper)
	}
}

func TestAddEdgeOneEdgeOneComponent(t *testing.T) {
	// One edge between two vertices
	cc := NewConnectedComponents()
	cc.AddEdge(EntityPair{EntityID1: "e-1", EntityID2: "e-2"})

	if cc.numberConnectedComponents != 1 {
		t.Fatalf("Expected 1 connected component, got %v\n", cc.numberConnectedComponents)
	}

	if cc.nextConnectedComponentID != 1 {
		t.Fatalf("Expected next connected component ID to be 1, got %v\n", cc.nextConnectedComponentID)
	}

	// Check the vertex to connected component assignment
	expectedVertexToComponent := map[string]int{
		"e-1": 0,
		"e-2": 0,
	}

	if !reflect.DeepEqual(expectedVertexToComponent, cc.vertexToConnectedComponent) {
		t.Fatalf("Expected %v, got %v\n", expectedVertexToComponent, cc.vertexToConnectedComponent)
	}

	// Check the connected component to slice of vertices
	expectedComponentToVertices := map[int][]string{
		0: []string{"e-1", "e-2"},
	}

	if !reflect.DeepEqual(expectedComponentToVertices, cc.connectedComponentToVertices) {
		t.Fatalf("Expected %v, got %v\n", expectedComponentToVertices, cc.connectedComponentToVertices)
	}
}

func TestAddEdgeDuplicateEdge(t *testing.T) {
	// One edge between two vertices
	cc := NewConnectedComponents()
	cc.AddEdge(EntityPair{EntityID1: "e-1", EntityID2: "e-2"})
	cc.AddEdge(EntityPair{EntityID1: "e-2", EntityID2: "e-1"})

	if cc.numberConnectedComponents != 1 {
		t.Fatalf("Expected 1 connected component, got %v\n", cc.numberConnectedComponents)
	}

	if cc.nextConnectedComponentID != 1 {
		t.Fatalf("Expected next connected component ID to be 1, got %v\n", cc.nextConnectedComponentID)
	}

	// Check the vertex to connected component assignment
	expectedVertexToComponent := map[string]int{
		"e-1": 0,
		"e-2": 0,
	}

	if !reflect.DeepEqual(expectedVertexToComponent, cc.vertexToConnectedComponent) {
		t.Fatalf("Expected %v, got %v\n", expectedVertexToComponent, cc.vertexToConnectedComponent)
	}

	// Check the connected component to slice of vertices
	expectedComponentToVertices := map[int][]string{
		0: []string{"e-1", "e-2"},
	}

	if !reflect.DeepEqual(expectedComponentToVertices, cc.connectedComponentToVertices) {
		t.Fatalf("Expected %v, got %v\n", expectedComponentToVertices, cc.connectedComponentToVertices)
	}
}

func TestAddEdgeTwoEdgesOneComponent1(t *testing.T) {
	// Two edges, one connected component
	cc := NewConnectedComponents()
	cc.AddEdge(EntityPair{EntityID1: "e-1", EntityID2: "e-2"})
	cc.AddEdge(EntityPair{EntityID1: "e-2", EntityID2: "e-3"}) // previously seen node comes first

	if cc.numberConnectedComponents != 1 {
		t.Fatalf("Expected 1 connected component, got %v\n", cc.numberConnectedComponents)
	}

	if cc.nextConnectedComponentID != 1 {
		t.Fatalf("Expected next connected component ID to be 1, got %v\n", cc.nextConnectedComponentID)
	}

	// Check the vertex to connected component assignment
	expectedVertexToComponent := map[string]int{
		"e-1": 0,
		"e-2": 0,
		"e-3": 0,
	}

	if !reflect.DeepEqual(expectedVertexToComponent, cc.vertexToConnectedComponent) {
		t.Fatalf("Expected %v, got %v\n", expectedVertexToComponent, cc.vertexToConnectedComponent)
	}

	// Check the connected component to slice of vertices
	expectedComponentToVertices := map[int][]string{
		0: []string{"e-1", "e-2", "e-3"},
	}

	if !reflect.DeepEqual(expectedComponentToVertices, cc.connectedComponentToVertices) {
		t.Fatalf("Expected %v, got %v\n", expectedComponentToVertices, cc.connectedComponentToVertices)
	}
}

func TestAddEdgeTwoEdgesOneComponent2(t *testing.T) {
	// Two edges, one connected component
	cc := NewConnectedComponents()
	cc.AddEdge(EntityPair{EntityID1: "e-1", EntityID2: "e-2"})
	cc.AddEdge(EntityPair{EntityID1: "e-3", EntityID2: "e-2"}) // previously seen node comes second

	if cc.numberConnectedComponents != 1 {
		t.Fatalf("Expected 1 connected component, got %v\n", cc.numberConnectedComponents)
	}

	if cc.nextConnectedComponentID != 1 {
		t.Fatalf("Expected next connected component ID to be 1, got %v\n", cc.nextConnectedComponentID)
	}

	// Check the vertex to connected component assignment
	expectedVertexToComponent := map[string]int{
		"e-1": 0,
		"e-2": 0,
		"e-3": 0,
	}

	if !reflect.DeepEqual(expectedVertexToComponent, cc.vertexToConnectedComponent) {
		t.Fatalf("Expected %v, got %v\n", expectedVertexToComponent, cc.vertexToConnectedComponent)
	}

	// Check the connected component to slice of vertices
	expectedComponentToVertices := map[int][]string{
		0: []string{"e-1", "e-2", "e-3"},
	}

	if !reflect.DeepEqual(expectedComponentToVertices, cc.connectedComponentToVertices) {
		t.Fatalf("Expected %v, got %v\n", expectedComponentToVertices, cc.connectedComponentToVertices)
	}
}

func TestAddEdgeTwoEdgesTwoComponents(t *testing.T) {
	// Two edges, two connected components
	cc := NewConnectedComponents()
	cc.AddEdge(EntityPair{EntityID1: "e-1", EntityID2: "e-2"})
	cc.AddEdge(EntityPair{EntityID1: "e-3", EntityID2: "e-4"})

	if cc.numberConnectedComponents != 2 {
		t.Fatalf("Expected 2 connected component, got %v\n", cc.numberConnectedComponents)
	}

	if cc.nextConnectedComponentID != 2 {
		t.Fatalf("Expected next connected component ID to be 2, got %v\n", cc.nextConnectedComponentID)
	}

	// Check the vertex to connected component assignment
	expectedVertexToComponent := map[string]int{
		"e-1": 0,
		"e-2": 0,
		"e-3": 1,
		"e-4": 1,
	}

	if !reflect.DeepEqual(expectedVertexToComponent, cc.vertexToConnectedComponent) {
		t.Fatalf("Expected %v, got %v\n", expectedVertexToComponent, cc.vertexToConnectedComponent)
	}

	// Check the connected component to slice of vertices
	expectedComponentToVertices := map[int][]string{
		0: []string{"e-1", "e-2"},
		1: []string{"e-3", "e-4"},
	}

	if !reflect.DeepEqual(expectedComponentToVertices, cc.connectedComponentToVertices) {
		t.Fatalf("Expected %v, got %v\n", expectedComponentToVertices, cc.connectedComponentToVertices)
	}
}

func TestAddEdgeThreeEdges(t *testing.T) {
	// Three edges, two connected components
	cc := NewConnectedComponents()
	cc.AddEdge(EntityPair{EntityID1: "e-1", EntityID2: "e-2"})
	cc.AddEdge(EntityPair{EntityID1: "e-4", EntityID2: "e-5"})
	cc.AddEdge(EntityPair{EntityID1: "e-2", EntityID2: "e-3"})

	if cc.numberConnectedComponents != 2 {
		t.Fatalf("Expected 2 connected component, got %v\n", cc.numberConnectedComponents)
	}

	if cc.nextConnectedComponentID != 2 {
		t.Fatalf("Expected next connected component ID to be 2, got %v\n", cc.nextConnectedComponentID)
	}

	// Check the vertex to connected component assignment
	expectedVertexToComponent := map[string]int{
		"e-1": 0,
		"e-2": 0,
		"e-3": 0,
		"e-4": 1,
		"e-5": 1,
	}

	if !reflect.DeepEqual(expectedVertexToComponent, cc.vertexToConnectedComponent) {
		t.Fatalf("Expected %v, got %v\n", expectedVertexToComponent, cc.vertexToConnectedComponent)
	}

	// Check the connected component to slice of vertices
	expectedComponentToVertices := map[int][]string{
		0: []string{"e-1", "e-2", "e-3"},
		1: []string{"e-4", "e-5"},
	}

	if !reflect.DeepEqual(expectedComponentToVertices, cc.connectedComponentToVertices) {
		t.Fatalf("Expected %v, got %v\n", expectedComponentToVertices, cc.connectedComponentToVertices)
	}
}

func TestAddEdgeFourEdges(t *testing.T) {
	// Four edges, two connected components
	cc := NewConnectedComponents()
	cc.AddEdge(EntityPair{EntityID1: "e-1", EntityID2: "e-2"})
	cc.AddEdge(EntityPair{EntityID1: "e-5", EntityID2: "e-6"})
	cc.AddEdge(EntityPair{EntityID1: "e-3", EntityID2: "e-4"})
	cc.AddEdge(EntityPair{EntityID1: "e-4", EntityID2: "e-1"}) // causes a merge

	if cc.numberConnectedComponents != 2 {
		t.Fatalf("Expected 2 connected component, got %v\n", cc.numberConnectedComponents)
	}

	if cc.nextConnectedComponentID != 3 {
		t.Fatalf("Expected next connected component ID to be 3, got %v\n", cc.nextConnectedComponentID)
	}

	// Check the vertex to connected component assignment
	expectedVertexToComponent := map[string]int{
		"e-1": 0,
		"e-2": 0,
		"e-3": 0,
		"e-4": 0,
		"e-5": 1,
		"e-6": 1,
	}

	if !reflect.DeepEqual(expectedVertexToComponent, cc.vertexToConnectedComponent) {
		t.Fatalf("Expected %v, got %v\n", expectedVertexToComponent, cc.vertexToConnectedComponent)
	}

	// Check the connected component to slice of vertices
	expectedComponentToVertices := map[int][]string{
		0: []string{"e-1", "e-2", "e-3", "e-4"},
		1: []string{"e-5", "e-6"},
	}

	if !reflect.DeepEqual(expectedComponentToVertices, cc.connectedComponentToVertices) {
		t.Fatalf("Expected %v, got %v\n", expectedComponentToVertices, cc.connectedComponentToVertices)
	}
}

func TestAddEdgeSixEdgesTwoComps(t *testing.T) {
	// Six edges, two connected components
	cc := NewConnectedComponents()
	cc.AddEdge(EntityPair{EntityID1: "e-1", EntityID2: "e-2"})
	cc.AddEdge(EntityPair{EntityID1: "e-2", EntityID2: "e-3"})
	cc.AddEdge(EntityPair{EntityID1: "e-10", EntityID2: "e-11"})
	cc.AddEdge(EntityPair{EntityID1: "e-4", EntityID2: "e-5"})
	cc.AddEdge(EntityPair{EntityID1: "e-5", EntityID2: "e-6"})
	cc.AddEdge(EntityPair{EntityID1: "e-6", EntityID2: "e-1"}) // causes a merge

	if cc.numberConnectedComponents != 2 {
		t.Fatalf("Expected 2 connected component, got %v\n", cc.numberConnectedComponents)
	}

	if cc.nextConnectedComponentID != 3 {
		t.Fatalf("Expected next connected component ID to be 3, got %v\n", cc.nextConnectedComponentID)
	}

	// Check the vertex to connected component assignment
	expectedVertexToComponent := map[string]int{
		"e-1":  0,
		"e-2":  0,
		"e-3":  0,
		"e-4":  0,
		"e-5":  0,
		"e-6":  0,
		"e-10": 1,
		"e-11": 1,
	}

	if !reflect.DeepEqual(expectedVertexToComponent, cc.vertexToConnectedComponent) {
		t.Fatalf("Expected %v, got %v\n", expectedVertexToComponent, cc.vertexToConnectedComponent)
	}

	// Check the connected component to slice of vertices
	expectedComponentToVertices := map[int][]string{
		0: []string{"e-1", "e-2", "e-3", "e-4", "e-5", "e-6"},
		1: []string{"e-10", "e-11"},
	}

	if !reflect.DeepEqual(expectedComponentToVertices, cc.connectedComponentToVertices) {
		t.Fatalf("Expected %v, got %v\n", expectedComponentToVertices, cc.connectedComponentToVertices)
	}
}

func TestAddEdgeFiveEdgesThreeComps(t *testing.T) {
	// Five edges, three connected components
	cc := NewConnectedComponents()
	cc.AddEdge(EntityPair{EntityID1: "e-1", EntityID2: "e-2"})
	cc.AddEdge(EntityPair{EntityID1: "e-2", EntityID2: "e-3"})
	cc.AddEdge(EntityPair{EntityID1: "e-10", EntityID2: "e-11"})
	cc.AddEdge(EntityPair{EntityID1: "e-4", EntityID2: "e-5"})
	cc.AddEdge(EntityPair{EntityID1: "e-5", EntityID2: "e-6"})

	if cc.numberConnectedComponents != 3 {
		t.Fatalf("Expected 3 connected component, got %v\n", cc.numberConnectedComponents)
	}

	if cc.nextConnectedComponentID != 3 {
		t.Fatalf("Expected next connected component ID to be 3, got %v\n", cc.nextConnectedComponentID)
	}

	// Check the vertex to connected component assignment
	expectedVertexToComponent := map[string]int{
		"e-1":  0,
		"e-2":  0,
		"e-3":  0,
		"e-4":  2,
		"e-5":  2,
		"e-6":  2,
		"e-10": 1,
		"e-11": 1,
	}

	if !reflect.DeepEqual(expectedVertexToComponent, cc.vertexToConnectedComponent) {
		t.Fatalf("Expected %v, got %v\n", expectedVertexToComponent, cc.vertexToConnectedComponent)
	}

	// Check the connected component to slice of vertices
	expectedComponentToVertices := map[int][]string{
		0: []string{"e-1", "e-2", "e-3"},
		1: []string{"e-10", "e-11"},
		2: []string{"e-4", "e-5", "e-6"},
	}

	if !reflect.DeepEqual(expectedComponentToVertices, cc.connectedComponentToVertices) {
		t.Fatalf("Expected %v, got %v\n", expectedComponentToVertices, cc.connectedComponentToVertices)
	}
}

func TestConnectedComponentsFromFile1(t *testing.T) {

	// Calculate connected components in file
	numRowsRead, cc := connectedComponentsFromFile("./test/unipartite_1.csv")

	if numRowsRead != 2 {
		t.Fatalf("Expected to read 2 rows, read %v rows", numRowsRead)
	}

	// Check the vertex to connected component assignment
	expectedVertexToComponent := map[string]int{
		"e-1": 0,
		"e-2": 0,
	}

	if !reflect.DeepEqual(expectedVertexToComponent, cc.vertexToConnectedComponent) {
		t.Fatalf("Expected %v, got %v\n", expectedVertexToComponent, cc.vertexToConnectedComponent)
	}
}

func TestConnectedComponentsFromFile2(t *testing.T) {

	// Calculate connected components in file
	numRowsRead, cc := connectedComponentsFromFile("./test/unipartite_2.csv")

	if numRowsRead != 3 {
		t.Fatalf("Expected to read 3 rows, read %v rows", numRowsRead)
	}

	// Check the vertex to connected component assignment
	expectedVertexToComponent := map[string]int{
		"e-1": 0,
		"e-2": 0,
		"e-3": 1,
		"e-4": 1,
	}

	if !reflect.DeepEqual(expectedVertexToComponent, cc.vertexToConnectedComponent) {
		t.Fatalf("Expected %v, got %v\n", expectedVertexToComponent, cc.vertexToConnectedComponent)
	}
}

func TestResultsHeader(t *testing.T) {
	actual := resultsHeader(",")
	expected := "Entity ID,Component ID"

	if expected != actual {
		t.Fatalf("Expected %v, got %v\n", expected, actual)
	}
}

func TestBuildResultsLine(t *testing.T) {
	actual := buildResultsLine("e-100", 4, "|")
	expected := "e-100|4"

	if expected != actual {
		t.Fatalf("Expected %v, got %v\n", expected, actual)
	}
}

func TestSortedListVertices(t *testing.T) {
	m := map[string]int{
		"e-4": 2,
		"e-1": 0,
		"e-2": 0,
	}

	actual := sortedListVertices(&m)
	expected := []string{"e-1", "e-2", "e-4"}

	if !reflect.DeepEqual(expected, *actual) {
		t.Fatalf("Expected %v, got %v\n", expected, actual)
	}
}

func TestCalculateConnectedComponents1(t *testing.T) {

	// Calculate the connected components
	calculateConnectedComponents("./test/test-1/edge_list.csv", "./test/test-1/actual.csv", ",")

	// Read the actual and expected results
	if !FilesHaveSameContent("./test/test-1/actual.csv", "./test/test-1/expected.csv") {
		t.Fatal("Actual results differ from expected results")
	}
}

func TestCalculateConnectedComponents2(t *testing.T) {

	// Calculate the connected components
	calculateConnectedComponents("./test/test-2/edge_list.csv", "./test/test-2/actual.csv", ",")

	// Read the actual and expected results
	if !FilesHaveSameContent("./test/test-2/actual.csv", "./test/test-2/expected.csv") {
		t.Fatal("Actual results differ from expected results")
	}
}
