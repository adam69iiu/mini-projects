package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)


func ConnectDatabase(dbURL string) (*pgxpool.Pool,error){
	conn,err := pgxpool.New(context.Background(),dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create new database connection: %w",err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to reach database. is it running? :%w",err)
	}
  
	return conn,nil

}
