package pkg

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strconv"
	"strings"
)

func Database(filepath string) *DB {
	return &DB{
		filepath: filepath,
	}
}

func AddNewTask(args []string) error {
	taskdesc := strings.Join(args, " ")
	filepath := "/tmp/task.db"
	db := Database(filepath)
	task := Task{
		Status:   TASK_STATUS_PENDING,
		TaskDesc: taskdesc,
	}
	err := db.Store(&task)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Added \"%s\" to your task list.\n", taskdesc)
	return nil
}

func ListTask() error {
	db := Database("/tmp/task.db")
	taskList := []string{TASK_STATUS_PENDING}
	tskList, err := db.AllTask(taskList)
	if err != nil {
		panic(err)
	}
	fmt.Printf("You have the following tasks:\n")
	for _, t := range tskList {
		fmt.Printf("%d. %s\n", t.ID, t.TaskDesc)
	}
	return nil
}

func MarkDone(args []string) error {
	if len(args) != 1 {
		fmt.Println("Invalid len of args")
		return nil
	}
	taskId, _ := strconv.ParseInt(args[0], 10, 64)
	db := Database("/tmp/task.db")
	err := db.MarkAsComplete(int(taskId))
	return err
}

func (d *DB) MarkAsComplete(taskId int) error {
	db := Database("/tmp/task.db")
	tsk := db.Get(taskId)
	tsk.Status = TASK_STATUS_COMPLETE
	db.Update(tsk)
	return nil
}

type DB struct {
	filepath string
}

type Task struct {
	ID       int    `json:"id"`
	Status   string `json:"status"`
	TaskDesc string `json:"taskDesc"`
}

const (
	TASKBUCKET           = "TaskBucket"
	TASK_STATUS_COMPLETE = "COMPLETED"
	TASK_STATUS_PENDING  = "PENDING"
)

func (d *DB) Get(id int) *Task {
	db, err := bolt.Open(d.filepath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var task Task
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		value := b.Get(itob(id))
		//fmt.Printf("%s", string(value))
		err := json.Unmarshal(value, &task)
		if err != nil {
			panic(err)
		}
		return nil
	})
	return &task
}

func (d *DB) Update(task *Task) error {
	db, err := bolt.Open(d.filepath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Task"))
		if err != nil {
			panic(err)
		}
		buf, err := json.Marshal(task)
		if err != nil {
			panic(err)
		}
		err = b.Put(itob(task.ID), buf)
		if err != nil {
			panic(err)
		}
		return err
	})
	return nil
}

func (d *DB) Store(task *Task) error {
	db, err := bolt.Open(d.filepath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Task"))
		if err != nil {
			panic(err)
		}
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		id, _ := b.NextSequence()
		task.ID = int(id)
		buf, err := json.Marshal(task)
		if err != nil {
			panic(err)
		}
		err = b.Put(itob(task.ID), buf)
		if err != nil {
			panic(err)
		}
		return err
	})
	return err
}

func (d *DB) AllTask(statusList []string) ([]*Task, error) {
	db, err := bolt.Open(d.filepath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var taskList []*Task
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Task"))
		if err != nil {
			panic(err)
		}
		return nil
	})
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Task"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var task Task
			err := json.Unmarshal(v, &task)
			if err != nil {
				panic(err)
			}
			for _, status := range statusList {
				if status == task.Status {
					taskList = append(taskList, &task)
				}
			}
		}
		return nil
	})
	return taskList, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func (d *DB) Retrive(key string) (*Task, error) {
	db, err := bolt.Open(d.filepath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var t Task
	db.View(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Task"))
		if err != nil {
			panic(err)
		}
		v := b.Get([]byte("ans"))
		//fmt.Printf("in func retrive %s \n", v)
		err = json.Unmarshal(v, &t)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("struct : %+v\n", t)
		return nil
	})

	return &t, nil
}

func (d *DB) Filter() ([]*Task, error) {
	return nil, nil
}
