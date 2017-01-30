package main

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/st4rl00rd/dynamics/context"
	// import postgress Driver (pq)
	_ "github.com/lib/pq"
)

// Movement struct
type Movement struct {
	ID          int64      `db:"id"`
	Quantity    int        `db:"quantity"`
	Ended       bool       `db:"ended"`
	StoreFromID int64      `db:"store_from_id"`
	StoreToID   int64      `db:"store_to_id"`
	UserFromID  int64      `db:"user_from_id"`
	UserToID    int64      `db:"user_to_id"`
	Details     DetailsMap `db:"properties"`
}

// GetProperties returns all the properties from a table on your db
func GetProperties(ctx *context.Context, table string, idPrm int) {
	idCnv := strconv.Itoa(idPrm)
	fmt.Println("SELECT id FROM " + table + " WHERE id=" + idCnv)
	rows, err := ctx.Db.Query("SELECT id FROM " + table + " WHERE id=" + idCnv)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d\n", id)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

// DetailsMap to get the Movements Details
type DetailsMap map[string]interface{}

// Value = This one implements the driver.Value to JSONB
func (p DetailsMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

// Scan = This one implements the sql.Scanner to JSONB
func (p *DetailsMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}
