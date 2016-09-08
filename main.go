package main

import (
	"github.com/emt/jira-helper/service"
	"net/http"
	"github.com/emt/jira-helper/container"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/emt/jira-helper/db"
	"fmt"
	"time"
)

const propertiesFileName = "jira-helper.yml"

func main() {

	data, err := ioutil.ReadFile("/home/emt/" + propertiesFileName)
	if (err != nil) {
		panic(err);
	}

	properties := container.JiraProperties{}

	prErr := yaml.Unmarshal(data, &properties)

	if (prErr != nil) {
		panic(prErr);
	}

	fmt.Printf("Repo list %s\n", properties.Repos)
	client := service.NewClient(http.DefaultClient);

	//2. return fie durumunu belirtir eger dosya yoksa db yi olustur

	db, _ := db.NewDbConnection(properties.DbLocation);

	defer db.Connection.Close();

	container := &container.JiraAppContainer{
		//Wrapper for all services
		Client: client,
		//properties from yml
		Properties: &properties,
		//db connection instance
		DbAdaptor: db};


	//res, err := db.Connection.Exec(sqlStmt);

	if (err != nil) {
		panic(err);
	}
	//fmt.Println(res)

	container.Init();

	const windowsInaccuracy = 17 * 1000
	const delay = 700
	//
	for {
		container.CheckCommits();
		time.Sleep(1000 * windowsInaccuracy * delay)
	}
}




