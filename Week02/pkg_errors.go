package main

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	// 相当于api在调用
	err := service()
	if errors.Cause(err) == sql.ErrNoRows {
		//fmt.Printf("data not found: %v\n", err)
		fmt.Printf("%+v\n", err)
		return
	}

	if err != nil {
		// unknown err
		fmt.Printf("unknown err:%+v\n", err)
	}

	return
}

func dao() error {
	return errors.Wrapf(sql.ErrNoRows, "dao failed")
}

func service() error {
	return errors.WithMessage(dao(), "service failed")
}
