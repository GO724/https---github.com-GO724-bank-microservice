package database

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"gopkg.in/ini.v1"
)

type Database struct {
	DB *pgxpool.Pool
}

type dbConfig struct {
	dbName   string
	host     string
	port     int
	user     string
	password string
}

func ReadConfig(fileName string) (dbConfig, error) { // load config from ini

	cfg, err := ini.Load(fileName)
	if err != nil {
		err = fmt.Errorf("unable to read configuration file[%s]: %w", fileName, err)
		return dbConfig{}, err
	}

	// сopy ini to config{}
	return dbConfig{
		dbName:   cfg.Section("postgres").Key("dbName").String(),
		host:     cfg.Section("postgres").Key("host").String(),
		port:     cfg.Section("postgres").Key("port").MustInt(5432),
		user:     cfg.Section("postgres").Key("user").String(),
		password: cfg.Section("postgres").Key("password").String(),
	}, nil
}

func (pgConfig *dbConfig) GetConnectionString() string { // get connection string from pgConfig
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		pgConfig.user,
		pgConfig.password,
		pgConfig.host,
		pgConfig.port,
		pgConfig.dbName,
	)
}

func New(ctx context.Context) (*Database, error) { // create new db connection

	var (
		pgxPoolInstance *Database
		pgOnce          sync.Once
		err             error
	)

	pgOnce.Do(func() { // singleton

		pgConfig, err := ReadConfig("ini/pg.ini")
		if err != nil {
			err = fmt.Errorf("can't get param to connect to database: %w", err)
			log.Fatal(err)
		}

		connectionString := pgConfig.GetConnectionString()

		dbConn, err := pgxpool.New(ctx, connectionString)
		if err != nil {
			log.Fatal(fmt.Errorf("unable to create connection pool[%s]: %w", connectionString, err))
		} else {
			fmt.Printf("connected to postgres.%s@%s:%d\n", pgConfig.dbName, pgConfig.host, pgConfig.port)
		}

		pgxPoolInstance = &Database{dbConn}
	})

	return pgxPoolInstance, err
}
