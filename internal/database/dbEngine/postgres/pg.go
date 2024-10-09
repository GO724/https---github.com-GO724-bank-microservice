package pg

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
	"gopkg.in/ini.v1"
)

type pgConfig struct {
	dbName   string
	host     string
	port     int
	user     string
	password string
}

type postgres struct {
	db  *pgxpool.Pool
	ctx *context.Context
	err error
}

var (
	pgInstance *postgres
	pgOnce     sync.Once
)

func NewPG(ctx context.Context) (*postgres, error) {
	// singleton
	pgOnce.Do(func() {
		cfg, err := ini.Load("ini/pg.ini")
		if err != nil {
			log.Fatal(err)
		}

		// Copy to config{}
		pgConfig := pgConfig{}
		pgConfig.dbName = cfg.Section("postgres").Key("dbName").String()
		pgConfig.host = cfg.Section("postgres").Key("host").String()
		pgConfig.port = cfg.Section("postgres").Key("port").MustInt(5432)
		pgConfig.user = cfg.Section("postgres").Key("user").String()
		pgConfig.password = cfg.Section("postgres").Key("password").String()

		ctx := context.Background()

		connString := "postgres://" + pgConfig.user + ":" + pgConfig.password + "@" + pgConfig.host + ":" + strconv.Itoa(pgConfig.port) + "/" + pgConfig.dbName

		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			err = fmt.Errorf("unable to create connection pool[%s]: %w", connString, err)
			fmt.Println("can't connect to " + connString)
		} else {
			fmt.Println("connected to postgres." + pgConfig.dbName + "@" + pgConfig.host + ":" + strconv.Itoa(pgConfig.port))
		}

		pgInstance = &postgres{db, &ctx, err}
	})

	return pgInstance, pgInstance.err
}

func (pg *postgres) Ping(ctx context.Context) error {
	return pg.db.Ping(ctx)
}

func (pg *postgres) Close() {
	pg.db.Close()
}

func (pg *postgres) GetInitError() error {
	return pg.err
}

func (pg *postgres) GetInitContext() *context.Context {
	return pg.ctx
}

func (pg *postgres) ExecSql(sqlQuery string) error {
	res, err := pg.db.Exec(*pg.ctx, sqlQuery)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to execute query: %s: %v\n", sqlQuery, err)
	} else {
		fmt.Fprintf(os.Stderr, "execute query[%v] : %v rows inserted\n", sqlQuery, res.RowsAffected())
	}
	return err
}
