package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"io"
	"os"
)

var Schema = `
CREATE TABLE phone_number(
    id  int NOT NULL AUTO_INCREMENT,
    number text,
    PRIMARY KEY (id)
);
`

func IsTableExists(db *sqlx.DB, tableName string) bool {
	_, err := db.Query(fmt.Sprintf("select * from %s", tableName))
	if err != nil {
		return false
	}
	return true
}

func DropTable(db *sqlx.DB, tableName string) error {
	_, err := db.Query(fmt.Sprintf("drop table %s", tableName))
	if err != nil {
		return nil
	}
	return nil
}

func validateFlags() {
	if *dbsource == "" {
		fmt.Printf("Error : must set db flag")
	}
}

var (
	dbsource *string
)

func init() {
	dbsource = flag.String("db", "", "Data source")
	flag.Parse()
	validateFlags()
}

func (s *Sql) InsertPhoneNumberRecordInTable(phoneNumber string) error {
	query := `INSERT INTO phone_number (number) values ("%s")`
	finalquery := fmt.Sprintf(query, phoneNumber)
	fmt.Printf("query : %s\n", finalquery)
	_, err := s.db.Query(fmt.Sprintf(query, phoneNumber))
	return err
}

type Sql struct {
	db *sqlx.DB
}

func NewSql(db *sqlx.DB) *Sql {
	return &Sql{
		db: db,
	}
}

func main() {
	db, err := sqlx.Connect("mysql", *dbsource)
	if err != nil {
		panic(err)
	}
	s := NewSql(db)

	if IsTableExists(db, "phone_number") {
		DropTable(db, "phone_number")
	}
	if !IsTableExists(db, "phone_number") {
		db.MustExec(Schema)
	}
	filePath := "./test/phone_number.csv"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(file)
	for records, err := csvReader.Read(); err != io.EOF; records, err = csvReader.Read() {
		phoneNumber := records[0]
		err = s.InsertPhoneNumberRecordInTable(phoneNumber)
		if err != nil {
			fmt.Printf("Error : failed to insert phonenumer %s", phoneNumber)
			continue
		}
	}

    //alter table : Add new column normalized number varchar(10)
    alterTableQuery := `ALTER TABLE phone_number 
                        ADD normalized_number varchar(10) unique`

    _, err = s.db.Query(alterTableQuery)
    if err != nil {
        panic(err)
    }

    // iterate over all number 
    // normalize number in format & insert into table
    // if it returns error then delete that row


}
