package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

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
	//fmt.Println("Inside of function")
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

func writeData2File(filename string, resource []Resourse) {

	file, _ := json.MarshalIndent(resource, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)
}

func readData4File(filename string) []Resourse {
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully read from %s\n", filename)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var res []Resourse
	json.Unmarshal(byteValue, &res)
	return res
}

func checkFileTimeStemp(path string) (t time.Time) {
	file, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	//mfiletime := file.ModTime()
	return file.ModTime()
}
