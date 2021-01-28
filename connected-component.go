package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

// EntityPair represents two entities connected by an edge
type EntityPair struct {
	EntityID1 string
	EntityID2 string
}

// readEntityPairsFromFile reads entity pairs from a CSV file
func readEntityPairsFromFile(filepath string) *[]EntityPair {

	fmt.Printf("[>] Reading graph from edge list file: %v\n", filepath)

	// Open the file for reading and ensure it is closed
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal("[!] Couldn't open CSV file ", err)
	}
	defer file.Close()

	// Initialise the slice of entity pairs
	var edges []EntityPair

	// Parse the input file
	r := csv.NewReader(file)
	numRowsRead := 0

	for {

		// Read a row from the file
		row, err := r.Read()
		numRowsRead++

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("[!] Error reading CSV file: ", err)
		}

		if len(row) != 2 {
			log.Fatal("[!] Invalid row: ", row)
		}

		entityPair := EntityPair{
			EntityID1: row[0],
			EntityID2: row[1],
		}

		edges = append(edges, entityPair)
	}

	fmt.Printf("[>] Read %v rows from file %v\n", numRowsRead, filepath)

	// Return a pointer to the slice of entity pairs
	return &edges
}

// minMax returns the (minimum, maximum) value of a pair of integers
func minMax(v1 int, v2 int) (int, int) {
	if v1 <= v2 {
		return v1, v2
	}
	return v2, v1
}

// connectedComponents calculates the connected components for a list of edges
func connectedComponents(edges *[]EntityPair) (*map[string]int, int) {

	// Initialise the map of vertex ID to the connected component ID
	vertexToConnectedComponent := map[string]int{}

	// Next connected component ID
	nextConnectedComponentID := 0

	// Number of connected components
	numberConnectedComponents := 0

	// Walk through each pair of entities
	for _, pair := range *edges {

		cc1, present1 := vertexToConnectedComponent[pair.EntityID1]
		cc2, present2 := vertexToConnectedComponent[pair.EntityID2]

		if present1 && present2 {
			// Both vertices have been seen before

			if cc1 == cc2 {
				continue
			}

			// Re-assign the highest connected component ID to merge components
			lowestCC, highestCC := minMax(cc1, cc2)
			for entity, cc := range vertexToConnectedComponent {
				if cc == highestCC {
					vertexToConnectedComponent[entity] = lowestCC
				}
			}
			numberConnectedComponents--

		} else if !present1 && present2 {
			// Only EntityID2 has been seen before
			vertexToConnectedComponent[pair.EntityID1] = cc2

		} else if present1 && !present2 {
			// Only EntityID1 has been seen before
			vertexToConnectedComponent[pair.EntityID2] = cc1

		} else {
			// Neither entity has been seen before
			vertexToConnectedComponent[pair.EntityID1] = nextConnectedComponentID
			vertexToConnectedComponent[pair.EntityID2] = nextConnectedComponentID
			nextConnectedComponentID++
			numberConnectedComponents++
		}
	}

	return &vertexToConnectedComponent, numberConnectedComponents
}

// resultsHeader builds the results file header
func resultsHeader(delimiter string) string {

	// Precondition
	if len(delimiter) == 0 {
		log.Fatal("Cannot use a blank delimiter")
	}

	return "Entity ID" + delimiter + "Component ID"
}

// buildResultsLine builds a line for the results file
func buildResultsLine(entityID string, component int, delimiter string) string {

	// Preconditions
	if len(entityID) == 0 {
		log.Fatal("Blank entity IDs are not valid")
	}

	if component < 0 {
		log.Fatal("Component IDs must be positive integers")
	}

	return entityID + delimiter + strconv.Itoa(component)
}

// sortedListVertices returns a sorted list of the vertices
func sortedListVertices(vertexToComponent *map[string]int) *[]string {

	// Get a slice of the keys
	keys := make([]string, len(*vertexToComponent))

	i := 0
	for k := range *vertexToComponent {
		keys[i] = k
		i++
	}

	// Sort the slice
	sort.Strings(keys)

	// Return the sorted list of keys
	return &keys
}

// writeConnectedComponentsToFile writes the vertex to connected component mapping to file
func writeVertexToConnectedComponentToFile(
	vertexToComponent *map[string]int,
	filepath string,
	delimiter string) {

	// Open the output CSV file for writing
	outputFile, err := os.Create(filepath)
	if err != nil {
		log.Fatalf("[!] Unable to open output file %v for writing: %v\n", filepath, err)
	}
	defer outputFile.Close()

	// Write the header
	fmt.Fprintln(outputFile, resultsHeader(delimiter))

	// Get a slice of sorted vertices
	sortedVertices := sortedListVertices(vertexToComponent)

	// Write each vertex to its connected component
	for _, vertex := range *sortedVertices {
		fmt.Fprintln(outputFile, buildResultsLine(vertex, (*vertexToComponent)[vertex], delimiter))
	}
}

// calculateConnectedComponents calculates the connected components from an edge list file
func calculateConnectedComponents(
	inputFilepath string,
	outputFilepath string,
	outputDelimiter string) {

	// Display a summary of the running parameters
	fmt.Printf("[>] Parameters\n")
	fmt.Printf("    Input file:            %v\n", inputFilepath)
	fmt.Printf("    Output file:           %v\n", outputFilepath)
	fmt.Printf("    Output file delimiter: %v\n", outputDelimiter)

	// Read the network into an in-memory data structure
	t0 := time.Now()
	edges := readEntityPairsFromFile(inputFilepath)

	fmt.Printf("[>] Read %v edges from file %v\n", len(*edges), inputFilepath)
	fmt.Printf("[>] Edges read in %v\n", time.Now().Sub(t0))

	// Calculate the connected components
	t1 := time.Now()
	vertexToConnectedComponent, numberConnectedComponents := connectedComponents(edges)
	fmt.Printf("[>] Connected components computed in %v\n", time.Now().Sub(t1))
	fmt.Printf("[>] Found %v connected components\n", numberConnectedComponents)

	// Write the connected components to a file
	t2 := time.Now()
	fmt.Printf("[>] Writing results to file %v\n", outputFilepath)
	writeVertexToConnectedComponentToFile(vertexToConnectedComponent, outputFilepath, outputDelimiter)
	fmt.Printf("[>] Vertex to connected component mapping written in %v\n", time.Now().Sub(t2))

	// Show the total execution time
	fmt.Printf("[>] Total time taken: %v\n", time.Now().Sub(t0))
}

func main() {

	// Command line arguments
	inputFilepath := flag.String("input", "unipartite.csv", "Location of the input CSV file of edges")
	outputFilepath := flag.String("output", "results.csv", "Location of the output CSV file of entity ID to connected component ID")
	delimiter := flag.String("delimiter", ",", "Delimiter for the CSV file of entity ID to connected component ID")
	flag.Parse()

	// Calculate the connected components given the command line arguments
	fmt.Println("Connected component calculator")
	calculateConnectedComponents(*inputFilepath, *outputFilepath, *delimiter)
}
