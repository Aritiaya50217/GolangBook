package controlscope_test

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

var testDB *sql.DB

func setup() {
	var err error
	testDB, err = sql.Open("mysql", "user:password@/testdb")
	if err != nil {
		log.Fatal(err)
	}
}

func teardown() {
	testDB.Close()
}

func TestInsertData(t *testing.T) {
	setup() // ตั้งค่าก่อนการทดสอบ
	defer teardown()
}

func TestFileProcessing(t *testing.T) {
	file, err := os.Open("testfile.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close() // ทำการปิดไฟล์หลังจากทดสอบเสร็จ

	// การทดสอบกับ file
}
// 33
