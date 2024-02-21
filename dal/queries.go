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
)

func GetAllSettingsFromDb(ctx *context.Context) (map[string]string, *error) {
	result, err := d.execWithTypedReturn(getAllOptionsForAllServicesQuery, ctx)

	if err != nil {
        return nil, err
    }

	return *result, nil
}

func GetSettingsFromDb(serviceName *string, ctx *context.Context) (*string, *error) {
	var query string = fmt.Sprintf(getAllOptionsQuery, (*serviceName))
	
	//d.initPool()
	result, err := d.execWithReturn(query, ctx)

	if err != nil {
        return nil, err
    }

	return &(*result)[0], err
}

func GetConcreteOptionFromDb(serviceName *string, optionPath *string, ctx *context.Context) (*string, *error) {
	var pgJsonPath string = "'" + strings.ReplaceAll((*optionPath), ",", "'->'") + "'"

	var query string = fmt.Sprintf(getConcreteOptionQuery, pgJsonPath, (*serviceName))

	//d.initPool()
	queryResult, err := d.execWithReturn(query, ctx)

	if err != nil {
		return nil, err
    }

	r := (*queryResult)[0]

	result := fmt.Sprintf("%v", r)

	return &result, nil
}

func InsertNewOptionsToDb(serviceName *string, options *string, ctx *context.Context) *error {
	args := pgx.NamedArgs{
    	"serviceName": serviceName,
    	"settings": options,
  	}

	err := d.execWithArgs(insertNewOptionsQuery, &args, ctx)

	if err != nil {
		return err
    }

	return nil
}

func DeleteSettingsFromDb(serviceName *string, ctx *context.Context) *error {
	args := pgx.NamedArgs{
    	"serviceName": serviceName,
  	}

	//d.initPool()
	err := d.execWithArgs(deleteOptionsQuery, &args, ctx)

	if err != nil {
		return err
    }

	return nil
}

func ReplaceOptionsInDb(serviceName *string, options *string, ctx *context.Context) *error {
	args := pgx.NamedArgs{
    	"serviceName": serviceName,
    	"settings": options,
  	}

	//d.initPool()
	err := d.execWithArgs(replaceOptionsQuery, &args, ctx)

	if err != nil {
		return err
    }

	return nil
}

func UpdateOptionInDb(serviceName *string, optionPath *string, optionValue *string, ctx *context.Context) error {
	var pgJsonPath string = "{" + strings.ReplaceAll((*optionPath), "/", ",") + "}"

	var q string

	if iv, err := strconv.Atoi(*optionValue); err == nil {
    	q = fmt.Sprintf(updateOptionQuery, pgJsonPath, iv, *serviceName)
	} else {
		q = fmt.Sprintf(updateOptionQuery, pgJsonPath, "\"" + *optionValue + "\"", *serviceName)
	}

	err := d.exec(q, ctx)
	
	if err != nil {
		return *err
    }

	return nil
}