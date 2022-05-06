package main

/*
This script has to be run before setting Patroni for any given servers.
It checks if the configurations and server details match between the East and Central servers
Also checks if there are existing Clusters with the same name you're trying to setup.
Arguments to be provided from the command line:
c --> Central Hostname
e --> East Hostname
cluster --> Cluster Name
For example:
go run patroniPreChecks.go -c cxxx12345 -e exxx12345 -cluster myDB
*/

import (
	"database/sql"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	_ "github.com/lib/pq"
)

const (
	url  = "https://xxxxinfluxxxxx-postgres.example.com/query?q="
	pg_port = 5432
	pg_user = "userID"
	pg_pass = "REDACTED"
	pg_ssl  = "require"
)

type PullMetrics struct {
	Results []Results
}

type Results struct {
	Series []Series
}

type Series struct {
	Name    string
	Tags    map[string]string
	Columns []string
	Values  [][]interface{}
}

type CurrentObjects struct {
	TableCatalog string
	TableName    string
	TableType    string
}

func queryInflux(qry string) ([]byte, error) {
	resp, err := http.Get(url + url.QueryEscape(qry))
	if err != nil {
		fmt.Printf("ERROR: Unable to run %v. %v\n", qry, err)
		return nil, errors.New("Unable to query from Influx: " + qry)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		return data, nil
	}
}

func getComplianceType(cmpl string) string {
	switch {
	case cmpl == "0":
		return "None"
	case cmpl == "1":
		return "ComplianceType1"
	case cmpl == "2":
		return "ComplianceType2"
	case cmpl == "3":
		return "ComplianceType3"
	case cmpl == "4":
		return "ComplianceType4"
	default:
		return "Unknown"
	}
}
