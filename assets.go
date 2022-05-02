package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	asset "cloud.google.com/go/asset/apiv1"
	"google.golang.org/api/iterator"
	assetpb "google.golang.org/genproto/googleapis/cloud/asset/v1"
)

// give org or project
// retunr json of sql instances links
func listSQLAssets(parent string) []Resourse {
	ctx := context.Background()
	client, err := asset.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	fmt.Println("Inside of function")
	assetType := "sqladmin.googleapis.com/Instance"
	req := &assetpb.ListAssetsRequest{
		Parent: parent,
		//Parent: fmt.Sprintf("organizations/%s",organizaitonID),
		AssetTypes:  []string{assetType},
		ContentType: assetpb.ContentType_RESOURCE,
	}
	it := client.ListAssets(ctx, req)
	var sqlList []Resourse

	for {
		response, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		sqlList = append(sqlList,
			Resourse{
				getProjectNameFromSlefLink(response.Name),
				getResourceNameFromSelfLink(response.Name)})

	}

	return sqlList
}

func writeData2File(resource []Resourse) {
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	file, _ := json.MarshalIndent(resource, "", " ")
	_ = ioutil.WriteFile("users", file, 0644)
}

func readData4File() []Resourse {
	jsonFile, err := os.Open("users.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var res []Resourse
	json.Unmarshal(byteValue, &res)
	return res
}
