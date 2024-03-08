package dal

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
)

const (
	getAllOptionsForAllServicesQuery string = "SELECT serviceName, settings FROM settings"
	getAllOptionsQuery string = "SELECT settings FROM settings WHERE serviceName = '%s'"
	getConcreteOptionQuery string = "SELECT settings->%s FROM settings WHERE servicename = '%s'"
	insertNewOptionsQuery string = "INSERT INTO settings(serviceName, settings) VALUES (@serviceName, @settings)"
	deleteOptionsQuery string = "DELETE from settings WHERE servicename = @serviceName"
	replaceOptionsQuery string = "UPDATE settings SET settings = @settings WHERE servicename = @serviceName"
	updateOptionQuery string = "UPDATE settings SET settings = jsonb_set(settings, '%s', '%v', TRUE) WHERE servicename = '%s'"
	deleteConcreteOptionQuery string = "UPDATE settings SET settings = settings::jsonb #- '{%s}' WHERE servicename = '%s'"
)

func GetAllSettingsFromDB(ctx *context.Context, db *Orm) (*map[string]string, error) {
	return db.execWithTypedReturn(getAllOptionsForAllServicesQuery, ctx)
}

func GetSettingsFromDB(serviceName *string, ctx *context.Context, db *Orm) (*string, error) {
	query := fmt.Sprintf(getAllOptionsQuery, (*serviceName))

	result, err := db.execWithReturn(query, ctx)

	if err != nil {
        return nil, err
    }

	return &(*result)[0], err
}

func InsertNewOptionsToDB(serviceName *string, options *string, ctx *context.Context, db *Orm) error {
	args := pgx.NamedArgs{
    	"serviceName": serviceName,
    	"settings": options,
  	}

	err := db.execWithArgs(insertNewOptionsQuery, &args, ctx)

	return err
}

func DeleteSettingsFromDB(serviceName *string, ctx *context.Context, db *Orm) error {
	args := pgx.NamedArgs{
    	"serviceName": serviceName,
  	}

	err := db.execWithArgs(deleteOptionsQuery, &args, ctx)

	return err
}

func ReplaceOptionsInDB(serviceName *string, options *string, ctx *context.Context, db *Orm) error {
	args := pgx.NamedArgs{
    	"serviceName": serviceName,
    	"settings": options,
  	}

	err := db.execWithArgs(replaceOptionsQuery, &args, ctx)

	return err
}

func UpdateOptionInDB(serviceName *string, optionPath *string, optionValue *string, ctx *context.Context, db *Orm) error {
	pgJSINPath := "{" + strings.ReplaceAll((*optionPath), "/", ",") + "}"

	var q string

	if iv, err := strconv.Atoi(*optionValue); err == nil {
    	q = fmt.Sprintf(updateOptionQuery, pgJSINPath, iv, *serviceName)
	} else {
		q = fmt.Sprintf(updateOptionQuery, pgJSINPath, "\"" + *optionValue + "\"", *serviceName)
	}

	err := db.exec(q, ctx)

	return err
}

func GetConcreteOptionFromDB(serviceName *string, optionPath *string, ctx *context.Context, db *Orm) (*string, error) {
	pgJSINPath := "'" + strings.ReplaceAll((*optionPath), ",", "'->'") + "'"

	query := fmt.Sprintf(getConcreteOptionQuery, pgJSINPath, (*serviceName))
	
	queryResult, err := db.execWithReturn(query, ctx)

	if err != nil {
		return nil, err
    }

	r := (*queryResult)[0]

	result := fmt.Sprintf("%v", r)

	return &result, nil
}

func DeleteConcreteOptionFromDB(serviceName *string, optionPath *string, ctx *context.Context, db *Orm) error {
	query := fmt.Sprintf(deleteConcreteOptionQuery, *optionPath, (*serviceName))
	err := db.exec(query, ctx)
	return err
}