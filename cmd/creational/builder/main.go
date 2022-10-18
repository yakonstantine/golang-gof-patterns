package main

import (
	"fmt"
	b "gof/src/creational/builder"
)

func main() {
	var bld b.QueryBuilder = b.NewSQLQueryBuilder()
	bld = bld.
		SetFields("id", "name", "age", "sex").
		SetFrom("people").
		SetWhere("age > 18").
		SetAnd("sex = 'M'").
		SetOrderBy("age", false)

	res, err := bld.Build()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)
}
