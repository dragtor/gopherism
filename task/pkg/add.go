package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strings"
)

func Database(filepath string) *DB {
	return &DB{
		filepath: filepath,
	}
}

func AddNewTask(args []string) error {
	taskstmt := strings.Join(args, " ")
	fmt.Printf("Added \"%s\" to your task list.\n", taskstmt)
	filepath := "my.db"
	db := Database(filepath)
	db.Store("Hello")
	value, err := db.Retrive("key")
	if err != nil {
		panic(err)
	}
	fmt.Printf("value stored in db : %+v\n", value)

	return nil
}

//type storageAccessor interface {
//    Store(string) error
//    Read() []task
//}

//type task struct {

//}

type DB struct {
	//DbDetails *DbDetails
	filepath string
	//storageAccessor
}

//type DbDetails struct{
//    dataFile string
//}

type Task struct {
	ID       int    `json:"id"`
	Status   string `json:"status"`
	TaskDesc string `json:"taskDesc"`
}

func (d *DB) Store(data string) error {
	db, err := bolt.Open(d.filepath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte("Task"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		t := &Task{
			ID:       1,
			Status:   "PENDING",
			TaskDesc: "Visit Farm",
		}
		buf, err := json.Marshal(t)
		if err != nil {
			panic(err)
		}
		err = b.Put([]byte("ans"), buf)
		if err != nil {
			panic(err)
		}
		return err
	})
	return nil
}

func (d *DB) Retrive(key string) (*Task, error) {
	db, err := bolt.Open(d.filepath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var t Task
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		v := b.Get([]byte("ans"))
		fmt.Printf("in func retrive %s \n", v)
		err := json.Unmarshal(v, &t)
		if err != nil {
			panic(err)
		}
		fmt.Printf("struct : %+v\n", t)
		return nil
	})

	return &t, nil
}
