package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Relation represents the structure of a relation object
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Relations represents the structure of relations data
type Relations struct {
	Index []Relation `json:"index"`
}

// FetchRelations fetches relations data for a given group ID
func FetchRelations(groupID string) (goodGroup Relation) {
	// Make a GET request to fetch relation data from the API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("Error fetching relations: ", err)
		return
	}

	// Ensure the response body is closed after the function returns
	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {
			fmt.Println("Failed to close response body:", err1)
		}
	}(resp.Body)

	// Read the response body into a byte slice
	data, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("Error reading response body: ", err2)
		return
	}

	// Define a variable to hold the decoded JSON data
	var relations Relations

	// Unmarshal the JSON data into the relations struct
	if err3 := json.Unmarshal(data, &relations); err3 != nil {
		fmt.Println("Error parsing JSON: ", err3)
		return
	}

	// Convert groupID to an integer
	intID, err4 := strconv.Atoi(groupID)
	if err4 != nil {
		fmt.Println("Error converting groupID to int: ", err4)
		return
	}

	// Search for the relation with the given group ID
	for _, relation := range relations.Index {
		if relation.ID == intID {
			goodGroup = relation
			return
		}
	}

	return goodGroup
}
