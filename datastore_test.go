package main

import (
        "testing"
        "os"
        "reflect"
)

var dbName = "./testDB"

func TestDataStore(t *testing.T) {
        ds, err := createDataStore(dbName)
        if err != nil {
                t.Fatal(err)
        }
        defer func() {
                ds.Close()
                os.RemoveAll(dbName)
        } ()

        var recs []Record
        recs = append(recs, Record{0, "woah", 21})
        recs = append(recs, Record{1, "test", 211})
        recs = append(recs, Record{1, "woah2", 222})
        for _, rec := range recs {
                err = ds.InsertRecord(rec)
                if err != nil {
                        t.Fatal(err)
                }
        }
        records, err := ds.GetLatestRecords()
        if err != nil {
                t.Fatal(err)
        }

        if !reflect.DeepEqual(records, recs) {
                t.Fatal("Records not equal")
        }
}
