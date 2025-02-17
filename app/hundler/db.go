package hundler

import (
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
)

type Client struct {
	USER       string
	PASS       string
	DBNAME     string
	ConnectDB  *sql.DB
}
func (client *Client) Init(user string, pass string, dbname string){
	client.USER   = user
	client.PASS   = pass
	client.DBNAME = dbname
}

func (client *Client) InitConnectDB(){
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", client.USER, client.PASS, client.DBNAME)
	var err error
	client.ConnectDB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
}

/*
func (client *Client) CreateTable(){
	_, err := client.ConnectDB.Exec(`CREATE TABLE "tasks" (id SERIAL PRIMARY KEY, 
									title TEXT NOT NULL, description TEXT, status TEXT 
									CHECK (status IN ('new','in_progress','done')) DEFAULT 'new',
									created_at TIMESTAMP DEFAULT now(), updated_at TIMESTAMP DEFAULT now())
	`)
	if err != nil {
		panic(err)
	}
}
	*/