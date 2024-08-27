package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale

	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT product, volume, date FROM sales WHERE id = :id", sql.Named("id", client))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := Sale{}

		err := rows.Scan(&s.Product, &s.Volume, &s.Date)
		if err != nil {
			return nil, err
		}
		sales = append(sales, s)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
