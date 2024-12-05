package main

import (
	"context"
	"fmt"
	"log"
	"os"

	. "github.com/openfga/go-sdk/client"
)

func main() {
	getTeams("contre")
	getTeams("becki")
}

func getTeams(user string) {
	fgaClient, err := NewSdkClient(&ClientConfiguration{
		ApiUrl:  "http://localhost:8080", // required, e.g. https://api.fga.example
		StoreId: os.Args[1],
		// AuthorizationModelId: os.Getenv("FGA_MODEL_ID"), // Optional, can be overridden per request
	})
	if err != nil {
		log.Fatal(err)
	}

	body := ClientListObjectsRequest{
		User:     fmt.Sprintf("user:%s", user),
		Relation: "ic",
		Type:     "team",
	}
	data, err := fgaClient.ListObjects(context.Background()).
		Body(body).
		Execute()
	log.Println(data.Objects)

	// body1 := ClientListUsersRequest{
	// 	Object:           openfga.FgaObject{},
	// 	Relation:         "",
	// 	UserFilters:      []openfga.UserTypeFilter{},
	// 	ContextualTuples: []openfga.TupleKey{},
	// 	Context:          &map[string]interface{}{},
	// }
	//
	// data1, err := fgaClient.ListUsers(context.Background()).Body(body1).Execute()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(data1.Users)

}
