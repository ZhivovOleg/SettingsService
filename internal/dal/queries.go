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

func GetAllSettingsFromDb(ctx *context.Context, db *Orm) (*map[string]string, *error) {
	return db.execWithTypedReturn(getAllOptionsForAllServicesQuery, ctx)
}

func GetSettingsFromDb(serviceName *string, ctx *context.Context, db *Orm) (*string, *error) {
	var query string = fmt.Sprintf(getAllOptionsQuery, (*serviceName))

	result, err := db.execWithReturn(query, ctx)

	if err != nil {
        return nil, err
    }

	return &(*result)[0], err
}

func InsertNewOptionsToDb(serviceName *string, options *string, ctx *context.Context, db *Orm) *error {
	args := pgx.NamedArgs{
    	"serviceName": serviceName,
    	"settings": options,
  	}

	err := db.execWithArgs(insertNewOptionsQuery, &args, ctx)

	return err
}

func DeleteSettingsFromDb(serviceName *string, ctx *context.Context, db *Orm) *error {
	args := pgx.NamedArgs{
    	"serviceName": serviceName,
  	}

	err := db.execWithArgs(deleteOptionsQuery, &args, ctx)

	return err
}

func ReplaceOptionsInDb(serviceName *string, options *string, ctx *context.Context, db *Orm) *error {
	args := pgx.NamedArgs{
    	"serviceName": serviceName,
    	"settings": options,
  	}

	err := db.execWithArgs(replaceOptionsQuery, &args, ctx)

	return err
}

func UpdateOptionInDb(serviceName *string, optionPath *string, optionValue *string, ctx *context.Context, db *Orm) *error {
	var pgJsonPath string = "{" + strings.ReplaceAll((*optionPath), "/", ",") + "}"

	var q string

	if iv, err := strconv.Atoi(*optionValue); err == nil {
    	q = fmt.Sprintf(updateOptionQuery, pgJsonPath, iv, *serviceName)
	} else {
		q = fmt.Sprintf(updateOptionQuery, pgJsonPath, "\"" + *optionValue + "\"", *serviceName)
	}

	err := db.exec(q, ctx)

	return err
}

func GetConcreteOptionFromDb(serviceName *string, optionPath *string, ctx *context.Context, db *Orm) (*string, *error) {
	var pgJsonPath string = "'" + strings.ReplaceAll((*optionPath), ",", "'->'") + "'"

	var query string = fmt.Sprintf(getConcreteOptionQuery, pgJsonPath, (*serviceName))
	
	queryResult, err := db.execWithReturn(query, ctx)

	if err != nil {
		return nil, err
    }

	r := (*queryResult)[0]

	result := fmt.Sprintf("%v", r)

	return &result, nil
}

func DeleteConcreteOptionFromDb(serviceName *string, optionPath *string, ctx *context.Context, db *Orm) *error {
	var query string = fmt.Sprintf(deleteConcreteOptionQuery, *optionPath, (*serviceName))
	err := db.exec(query, ctx)
	return err
}