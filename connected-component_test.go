package main

import (
	"reflect"
	"testing"
)

func TestReadEntityPairsFromFile1(t *testing.T) {
	actual := *readEntityPairsFromFile("./test/unipartite_1.csv")
	expected := []EntityPair{
		{
			EntityID1: "e-1",
			EntityID2: "e-2",
		},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v, got %v\n", expected, actual)
	}
}

func TestReadEntityPairsFromFile2(t *testing.T) {
	actual := *readEntityPairsFromFile("./test/unipartite_2.csv")
	expected := []EntityPair{
		{
			EntityID1: "e-1",
			EntityID2: "e-2",
		},
		{
			EntityID1: "e-3",
			EntityID2: "e-4",
		},
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected %v, got %v\n", expected, actual)
	}
}

func TestConnectedComponentsOneComponent(t *testing.T) {
	entityPairs := []EntityPair{
		{
			EntityID1: "e-1",
			EntityID2: "e-2",
		},
	}
	actual := connectedComponents(&entityPairs)
	expected := map[string]int{
		"e-1": 0,
		"e-2": 0,
	}

	if !reflect.DeepEqual(expected, *actual) {
		t.Fatalf("Expected %v, got %v\n", expected, actual)
	}
}

func TestConnectedComponentsOneComponent2(t *testing.T) {
	entityPairs := []EntityPair{
		{
			EntityID1: "e-1",
			EntityID2: "e-2",
		},
		{
			EntityID1: "e-1",
			EntityID2: "e-3",
		},
	}
	actual := connectedComponents(&entityPairs)
	expected := map[string]int{
		"e-1": 0,
		"e-2": 0,
		"e-3": 0,
	}

	if !reflect.DeepEqual(expected, *actual) {
		t.Fatalf("Expected %v, got %v\n", expected, actual)
	}
}

func TestConnectedComponentsTwoComponents(t *testing.T) {
	entityPairs := []EntityPair{
		{
			EntityID1: "e-1",
			EntityID2: "e-2",
		},
		{
			EntityID1: "e-3",
			EntityID2: "e-4",
		},
	}
	actual := connectedComponents(&entityPairs)
	expected := map[string]int{
		"e-1": 0,
		"e-2": 0,
		"e-3": 1,
		"e-4": 1,
	}

	if !reflect.DeepEqual(expected, *actual) {
		t.Fatalf("Expected %v, got %v\n", expected, actual)
	}
}

func TestConnectedComponentsTwoComponents2(t *testing.T) {
	entityPairs := []EntityPair{
		{
			EntityID1: "e-1",
			EntityID2: "e-2",
		},
		{
			EntityID1: "e-3",
			EntityID2: "e-4",
		},
		{
			EntityID1: "e-1",
			EntityID2: "e-5",
		},
	}
	actual := connectedComponents(&entityPairs)
	expected := map[string]int{
		"e-1": 0,
		"e-2": 0,
		"e-3": 1,
		"e-4": 1,
		"e-5": 0,
	}

	if !reflect.DeepEqual(expected, *actual) {
		t.Fatalf("Expected %v, got %v\n", expected, actual)
	}
}

func TestConnectedComponentsTwoComponents3(t *testing.T) {
	entityPairs := []EntityPair{
		{
			EntityID1: "e-1",
			EntityID2: "e-2",
		},
		{
			EntityID1: "e-3",
			EntityID2: "e-4",
		},
		{
			EntityID1: "e-5",
			EntityID2: "e-6",
		},
		{
			EntityID1: "e-1", // this should cause a merge
			EntityID2: "e-6",
		},
	}
	actual := connectedComponents(&entityPairs)
	expected := map[string]int{
		"e-1": 0,
		"e-2": 0,
		"e-3": 1,
		"e-4": 1,
		"e-5": 0,
		"e-6": 0,
	}

	if !reflect.DeepEqual(expected, *actual) {
		t.Fatalf("Expected %v, got %v\n", expected, actual)
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
