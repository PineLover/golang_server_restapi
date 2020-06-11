package main_test

import (
	"fmt"
	"os"
	"testing"
)

var a main.App

var ip, dbName, collectionName = "192.168.99.100", "godb", "movies"

func TestMain(m *testing.M) {
	a = main.App{}
	a.Initialize(
		ip,
		dbName,
		collectionName,
	)

	prevTesting()

	code := m.Run()

	postTesting()

	os.Exit(code)
}

func prevTesting() {
	fmt.Println("prevTesting....")
}

func postTesting() {
	fmt.Println("postTesting....")
}

func TestExam(t *testing.T) {

	fmt.Println("TestExam inside")
}
