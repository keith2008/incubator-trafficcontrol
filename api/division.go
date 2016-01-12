// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_goto2.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	"github.com/Comcast/traffic_control/traffic_ops/goto2/db"
	_ "github.com/Comcast/traffic_control/traffic_ops/goto2/output_format" // needed for swagger
	"log"
	"time"
)

type Division struct {
	Id          int64     `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	LastUpdated time.Time `db:"last_updated" json:"lastUpdated"`
}

func handleDivision(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getDivision(id)
	} else if method == "POST" {
		return postDivision(payload)
	} else if method == "PUT" {
		return putDivision(id, payload)
	} else if method == "DELETE" {
		return delDivision(id)
	}
	return nil, nil
}

func getDivision(id int) (interface{}, error) {
	if id >= 0 {
		return getDivisionById(id)
	} else {
		return getDivisions()
	}
}

// @Title getDivisionById
// @Description retrieves the division information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Division
// @Resource /api/2.0
// @Router /api/2.0/division/{id} [get]
func getDivisionById(id int) (interface{}, error) {
	ret := []Division{}
	arg := Division{Id: int64(id)}
	nstmt, err := db.GlobalDB.PrepareNamed(`select * from division where id=:id`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getDivisions
// @Description retrieves the division information for a certain id
// @Accept  application/json
// @Success 200 {array}    Division
// @Resource /api/2.0
// @Router /api/2.0/division [get]
func getDivisions() (interface{}, error) {
	ret := []Division{}
	queryStr := "select * from division"
	err := db.GlobalDB.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postDivision
// @Description enter a new division
// @Accept  application/json
// @Param                 Body body     Division   true "Division object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/division [post]
func postDivision(payload []byte) (interface{}, error) {
	var v Division
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO division("
	sqlString += "name"
	sqlString += ") VALUES ("
	sqlString += ":name"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putDivision
// @Description modify an existing divisionentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Division   true "Division object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/division/{id}  [put]
func putDivision(id int, payload []byte) (interface{}, error) {
	var v Division
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE division SET "
	sqlString += "name = :name"
	sqlString += ",last_updated = :last_updated"
	sqlString += " WHERE id=:id"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delDivisionById
// @Description deletes division information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Division
// @Resource /api/2.0
// @Router /api/2.0/division/{id} [delete]
func delDivision(id int) (interface{}, error) {
	arg := Division{Id: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM division WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
