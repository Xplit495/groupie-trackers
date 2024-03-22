package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Index []Relation `json:"index"`
}

func FetchRelations(groupID string) (goodGroup Relation) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("Error fetching relations: ", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err1 := Body.Close()
		if err1 != nil {

		}
	}(resp.Body)

	data, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println("Error reading response body: ", err1)
		return
	}

	var relations Relations

	if err2 := json.Unmarshal(data, &relations); err2 != nil {
		fmt.Println("Error parsing JSON: ", err2)
		return
	}

	intID, err3 := strconv.Atoi(groupID)
	if err3 != nil {
		fmt.Println("Error converting groupID to int: ", err3)
		return
	}

	for _, relation := range relations.Index {
		if relation.ID == intID {
			goodGroup = relation
			return
		}
	}

	return goodGroup
}
