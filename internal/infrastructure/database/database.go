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

type pgConfig struct {
	dbName   string
	host     string
	port     int
	user     string
	password string
}

func (pgConfig *pgConfig) ReadConfig(fileName string) error { // load config from ini

	cfg, err := ini.Load(fileName)
	if err != nil {
		err = fmt.Errorf("unable to read configuration file[%s]: %w", fileName, err)
		return err
	}

	// сopy ini to config{}
	pgConfig.dbName = cfg.Section("postgres").Key("dbName").String()
	pgConfig.host = cfg.Section("postgres").Key("host").String()
	pgConfig.port = cfg.Section("postgres").Key("port").MustInt(5432)
	pgConfig.user = cfg.Section("postgres").Key("user").String()
	pgConfig.password = cfg.Section("postgres").Key("password").String()
	return err
}

func (pgConfig *pgConfig) GetConnectionString() string { // get connection string from pgConfig
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
		pgConfig        *pgConfig
		pgxPoolInstance *Database
		pgOnce          sync.Once
		err             error
	)

	pgOnce.Do(func() { // singleton

		if pgConfig.ReadConfig("ini/pg.ini") != nil {
			err = fmt.Errorf("can't get param to connect to database: %w", err)
			log.Fatal(err)
		}

		connectionString := pgConfig.GetConnectionString()

		dbConn, err := pgxpool.New(ctx, connectionString)
		if err != nil {
			log.Fatal(fmt.Errorf("unable to create connection pool[%s]: %w", connectionString, err))
		} else {
			fmt.Printf("connected to postgres.%s@%s:%d", pgConfig.dbName, pgConfig.host, pgConfig.port)
		}

		pgxPoolInstance = &Database{dbConn}
	})

	return pgxPoolInstance, err
}
