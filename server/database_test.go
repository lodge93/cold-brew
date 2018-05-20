// Copyright © 2018 Mark Spicer
// Made available under the MIT license.

package server

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/nanobox-io/golang-scribble"
)

type TestObject struct {
	Foo string
	Bar string
}

func TestNewDatabase(t *testing.T) {
	db, dir, err := withNewTempDatabase()
	if err != nil {
		t.Fatal("database could not be created:", err)
	}
	defer cleanUpTempDatabase(dir)

	err = ensureEntriesCanBeWritten(db)
	if err != nil {
		t.Fatal("could not access database:", err)
	}
}

func ensureEntriesCanBeWritten(db *scribble.Driver) error {
	testObj := TestObject{
		Foo: "bar",
		Bar: "foo",
	}
	err := db.Write("testobjects", "object1234", testObj)
	if err != nil {
		return err
	}

	testObjOne := TestObject{}
	err = db.Read("testobjects", "object1234", &testObjOne)
	if err != nil {
		return err
	}

	if !reflect.DeepEqual(testObj, testObjOne) {
		log.Println(testObj.Bar)
		log.Println(testObjOne.Bar)
		return errors.New("test object read does not match what was written")
	}

	return nil
}

func cleanUpTempDatabase(dir string) {
	os.RemoveAll(dir)
}

func withNewTempDatabase() (*scribble.Driver, string, error) {
	dir, err := ioutil.TempDir("", "cold-brew-test")
	if err != nil {
		return nil, "", err
	}

	db, err := NewDatabase(dir)
	if err != nil {
		return nil, "", err
	}

	return db, dir, nil
}
