package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"fmt"
	"github.com/emt/jira-helper/dto"
	"log"
)

type DbAdaptor struct {
	Connection *sql.DB
}

func (s DbAdaptor) GetRepo(repoName string) (*dto.Repo) {
	var query=fmt.Sprintf("select * from REPO_TABLE WHERE REPO_NAME = '%s'", repoName);
	rows, err := s.Connection.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var repo *dto.Repo;

	for rows.Next() {
		var id int
		var repoName string
		var time string
		err = rows.Scan(&id, &repoName, &time)
		if err != nil {
			log.Fatal(err)
		}
		repo = &dto.Repo{RepoName:repoName, LastCommit:time}

	}
	return repo;
}

func (s DbAdaptor) UpdateRepo(repoName string, time string) {
	deleteQuery := fmt.Sprintf("delete from REPO_TABLE where REPO_NAME = '%s';", repoName);
	s.Connection.Exec(deleteQuery);
	createQuery := fmt.Sprintf("insert into REPO_TABLE(REPO_NAME,LAST_COMMIT) VALUES ('%s','%s');", repoName, time);
	s.Connection.Exec(createQuery);
}

func NewDbConnection(location string) (*DbAdaptor, bool) {
	fileExistStatus := true;
	db := &DbAdaptor{Connection:createConnection(location)}

	if _, err := os.Stat(location); os.IsNotExist(err) {
		fileExistStatus = false;
	}

	return db, fileExistStatus;
}

func createConnection(location string) (*sql.DB) {
	db, err := sql.Open("sqlite3", location);
	if (err != nil) {
		panic(err)
	}
	return db;
}
