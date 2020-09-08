package main

import "timewise/db"

func main() {
	var sql = new(db.SQL)
	sql.Connect()
	defer sql.Close()
}
