package dal

import (
	"context"
	"fmt"
	"gisogd/SettingsService/options"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type db struct {
	pool *pgxpool.Pool
	locker *sync.Once
	isInited bool
}

var d *db

func InitPool() (*pgxpool.Pool, *error) {
	d = new(db)
	d.isInited = false
	d.locker = new(sync.Once)

	var initPoolerr error
	d.locker.Do(func() {
		pool, err := pgxpool.New(context.Background(), *options.ServiceSetting.DbConnectionString)

		if err != nil {
			d.isInited = false
			initPoolerr = fmt.Errorf("Can't init database pool: " + err.Error())
			return
		}		
		d.pool = pool
		d.isInited = true
	})

	if initPoolerr != nil {
		return nil, &initPoolerr
	}
	return d.pool, nil
}

func (d *db) execWithReturn(query string, ctx *context.Context) (*[]string, *error) {
	//d.initPool()
	rows, err := d.pool.Query(*ctx, query)
	
	if err != nil {
		return nil, &err
	}

	defer rows.Close()

	var result []string

	for rows.Next() {
		var currStr string
		err = rows.Scan(&currStr)

		if err != nil {
			fmt.Println(err.Error())
			return nil, &err
    	}

		result = append(result, currStr)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil, &err
    }

	return &result, nil
}

func (d *db) execWithArgs(query string, args *pgx.NamedArgs, ctx *context.Context) *error {
	_, err := d.pool.Exec(*ctx, query, *args)
	
	if err != nil {
		return &err
	}

	return nil
}

func (d *db) exec(query string, ctx *context.Context) *error {
	_, err := d.pool.Exec(*ctx, query)
	
	if err != nil {
		return &err
	}

	return nil
}

func (d *db) execWithTypedReturn(query string, ctx *context.Context) (*map[string]string, *error) {
	rows, err := d.pool.Query(*ctx, query)
	
	if err != nil {
		return nil, &err
	}

	defer rows.Close()

	result := make(map[string]string)

	for rows.Next() {
		var currName string
		var currOpt string
		err = rows.Scan(&currName, &currOpt)

		if err != nil {
			fmt.Println(err.Error())
			return nil, &err
    	}

		result[currName] = currOpt
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil, &err
    }

	return &result, nil
}