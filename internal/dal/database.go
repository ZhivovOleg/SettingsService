package dal

import (
	"context"
	"fmt"
	"sync"

	"github.com/ZhivovOleg/SettingsService/internal/options"
	"github.com/ZhivovOleg/SettingsService/internal/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Struct for access to database
type Orm struct {
	pool *pgxpool.Pool
	locker *sync.Once
	isInited bool
}

func (o *Orm) Init(connString string) error {
	o.isInited = false
	o.locker = new(sync.Once)

	var initPoolerr error
	o.locker.Do(func() {
		pool, err := pgxpool.New(context.Background(), *options.ServiceSetting.DBConnectionString)

		if err != nil {
			o.isInited = false
			initPoolerr = fmt.Errorf("Can't init database pool: " + err.Error())
			return
		}		

		pingDBErr := pool.Ping(context.Background())	
		if pingDBErr != nil {
			utils.Logger.Error("Can't connect with database: " + pingDBErr.Error())
			panic("Can't connect with database: " + pingDBErr.Error())
		}

		o.pool = pool
		o.isInited = true
	})
	return initPoolerr
}

func (o *Orm) execWithReturn(query string, ctx *context.Context) (*[]string, error) {
	rows, err := o.pool.Query(*ctx, query)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []string

	for rows.Next() {
		var currStr string
		err = rows.Scan(&currStr)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
    	}

		result = append(result, currStr)
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
    }

	return &result, nil
}

func (o *Orm) execWithArgs(query string, args *pgx.NamedArgs, ctx *context.Context) error {
	_, err := o.pool.Exec(*ctx, query, *args)
	
	if err != nil {
		return err
	}

	return nil
}

func (o *Orm) exec(query string, ctx *context.Context) error {
	_, err := o.pool.Exec(*ctx, query)
	
	if err != nil {
		return err
	}

	return nil
}

func (o *Orm) execWithTypedReturn(query string, ctx *context.Context) (*map[string]string, error) {
	rows, err := o.pool.Query(*ctx, query)
	
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	result := make(map[string]string)

	for rows.Next() {
		var currName string
		var currOpt string
		err = rows.Scan(&currName, &currOpt)

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
    	}

		result[currName] = currOpt
	}

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
    }

	return &result, nil
}