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
	null "gopkg.in/guregu/null.v3"
	"log"
	"time"
)

type Deliveryservice struct {
	Id                   int64       `db:"id" json:"id"`
	XmlId                string      `db:"xml_id" json:"xmlId"`
	Active               int64       `db:"active" json:"active"`
	Dscp                 int64       `db:"dscp" json:"dscp"`
	Signed               null.Int    `db:"signed" json:"signed"`
	QstringIgnore        null.Int    `db:"qstring_ignore" json:"qstringIgnore"`
	GeoLimit             null.Int    `db:"geo_limit" json:"geoLimit"`
	HttpBypassFqdn       null.String `db:"http_bypass_fqdn" json:"httpBypassFqdn"`
	DnsBypassIp          null.String `db:"dns_bypass_ip" json:"dnsBypassIp"`
	DnsBypassIp6         null.String `db:"dns_bypass_ip6" json:"dnsBypassIp6"`
	DnsBypassTtl         null.Int    `db:"dns_bypass_ttl" json:"dnsBypassTtl"`
	OrgServerFqdn        null.String `db:"org_server_fqdn" json:"orgServerFqdn"`
	Type                 int64       `db:"type" json:"type"`
	Profile              int64       `db:"profile" json:"profile"`
	CdnId                null.Int    `db:"cdn_id" json:"cdnId"`
	CcrDnsTtl            null.Int    `db:"ccr_dns_ttl" json:"ccrDnsTtl"`
	GlobalMaxMbps        null.Int    `db:"global_max_mbps" json:"globalMaxMbps"`
	GlobalMaxTps         null.Int    `db:"global_max_tps" json:"globalMaxTps"`
	LongDesc             null.String `db:"long_desc" json:"longDesc"`
	LongDesc1            null.String `db:"long_desc_1" json:"longDesc1"`
	LongDesc2            null.String `db:"long_desc_2" json:"longDesc2"`
	MaxDnsAnswers        null.Int    `db:"max_dns_answers" json:"maxDnsAnswers"`
	InfoUrl              null.String `db:"info_url" json:"infoUrl"`
	MissLat              null.Float  `db:"miss_lat" json:"missLat"`
	MissLong             null.Float  `db:"miss_long" json:"missLong"`
	CheckPath            null.String `db:"check_path" json:"checkPath"`
	LastUpdated          time.Time   `db:"last_updated" json:"lastUpdated"`
	Protocol             null.Int    `db:"protocol" json:"protocol"`
	SslKeyVersion        null.Int    `db:"ssl_key_version" json:"sslKeyVersion"`
	Ipv6RoutingEnabled   null.Int    `db:"ipv6_routing_enabled" json:"ipv6RoutingEnabled"`
	RangeRequestHandling null.Int    `db:"range_request_handling" json:"rangeRequestHandling"`
	EdgeHeaderRewrite    null.String `db:"edge_header_rewrite" json:"edgeHeaderRewrite"`
	OriginShield         null.String `db:"origin_shield" json:"originShield"`
	MidHeaderRewrite     null.String `db:"mid_header_rewrite" json:"midHeaderRewrite"`
	RegexRemap           null.String `db:"regex_remap" json:"regexRemap"`
	Cacheurl             null.String `db:"cacheurl" json:"cacheurl"`
	RemapText            null.String `db:"remap_text" json:"remapText"`
	MultiSiteOrigin      null.Int    `db:"multi_site_origin" json:"multiSiteOrigin"`
	DisplayName          string      `db:"display_name" json:"displayName"`
	TrResponseHeaders    null.String `db:"tr_response_headers" json:"trResponseHeaders"`
	InitialDispersion    null.Int    `db:"initial_dispersion" json:"initialDispersion"`
	DnsBypassCname       null.String `db:"dns_bypass_cname" json:"dnsBypassCname"`
}

func handleDeliveryservice(method string, id int, payload []byte) (interface{}, error) {
	if method == "GET" {
		return getDeliveryservice(id)
	} else if method == "POST" {
		return postDeliveryservice(payload)
	} else if method == "PUT" {
		return putDeliveryservice(id, payload)
	} else if method == "DELETE" {
		return delDeliveryservice(id)
	}
	return nil, nil
}

func getDeliveryservice(id int) (interface{}, error) {
	if id >= 0 {
		return getDeliveryserviceById(id)
	} else {
		return getDeliveryservices()
	}
}

// @Title getDeliveryserviceById
// @Description retrieves the deliveryservice information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Deliveryservice
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice/{id} [get]
func getDeliveryserviceById(id int) (interface{}, error) {
	ret := []Deliveryservice{}
	arg := Deliveryservice{Id: int64(id)}
	nstmt, err := db.GlobalDB.PrepareNamed(`select * from deliveryservice where id=:id`)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getDeliveryservices
// @Description retrieves the deliveryservice information for a certain id
// @Accept  application/json
// @Success 200 {array}    Deliveryservice
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice [get]
func getDeliveryservices() (interface{}, error) {
	ret := []Deliveryservice{}
	queryStr := "select * from deliveryservice"
	err := db.GlobalDB.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postDeliveryservice
// @Description enter a new deliveryservice
// @Accept  application/json
// @Param                 Body body     Deliveryservice   true "Deliveryservice object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice [post]
func postDeliveryservice(payload []byte) (interface{}, error) {
	var v Deliveryservice
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
	}
	sqlString := "INSERT INTO deliveryservice("
	sqlString += "xml_id"
	sqlString += ",active"
	sqlString += ",dscp"
	sqlString += ",signed"
	sqlString += ",qstring_ignore"
	sqlString += ",geo_limit"
	sqlString += ",http_bypass_fqdn"
	sqlString += ",dns_bypass_ip"
	sqlString += ",dns_bypass_ip6"
	sqlString += ",dns_bypass_ttl"
	sqlString += ",org_server_fqdn"
	sqlString += ",type"
	sqlString += ",profile"
	sqlString += ",cdn_id"
	sqlString += ",ccr_dns_ttl"
	sqlString += ",global_max_mbps"
	sqlString += ",global_max_tps"
	sqlString += ",long_desc"
	sqlString += ",long_desc_1"
	sqlString += ",long_desc_2"
	sqlString += ",max_dns_answers"
	sqlString += ",info_url"
	sqlString += ",miss_lat"
	sqlString += ",miss_long"
	sqlString += ",check_path"
	sqlString += ",protocol"
	sqlString += ",ssl_key_version"
	sqlString += ",ipv6_routing_enabled"
	sqlString += ",range_request_handling"
	sqlString += ",edge_header_rewrite"
	sqlString += ",origin_shield"
	sqlString += ",mid_header_rewrite"
	sqlString += ",regex_remap"
	sqlString += ",cacheurl"
	sqlString += ",remap_text"
	sqlString += ",multi_site_origin"
	sqlString += ",display_name"
	sqlString += ",tr_response_headers"
	sqlString += ",initial_dispersion"
	sqlString += ",dns_bypass_cname"
	sqlString += ") VALUES ("
	sqlString += ":xml_id"
	sqlString += ",:active"
	sqlString += ",:dscp"
	sqlString += ",:signed"
	sqlString += ",:qstring_ignore"
	sqlString += ",:geo_limit"
	sqlString += ",:http_bypass_fqdn"
	sqlString += ",:dns_bypass_ip"
	sqlString += ",:dns_bypass_ip6"
	sqlString += ",:dns_bypass_ttl"
	sqlString += ",:org_server_fqdn"
	sqlString += ",:type"
	sqlString += ",:profile"
	sqlString += ",:cdn_id"
	sqlString += ",:ccr_dns_ttl"
	sqlString += ",:global_max_mbps"
	sqlString += ",:global_max_tps"
	sqlString += ",:long_desc"
	sqlString += ",:long_desc_1"
	sqlString += ",:long_desc_2"
	sqlString += ",:max_dns_answers"
	sqlString += ",:info_url"
	sqlString += ",:miss_lat"
	sqlString += ",:miss_long"
	sqlString += ",:check_path"
	sqlString += ",:protocol"
	sqlString += ",:ssl_key_version"
	sqlString += ",:ipv6_routing_enabled"
	sqlString += ",:range_request_handling"
	sqlString += ",:edge_header_rewrite"
	sqlString += ",:origin_shield"
	sqlString += ",:mid_header_rewrite"
	sqlString += ",:regex_remap"
	sqlString += ",:cacheurl"
	sqlString += ",:remap_text"
	sqlString += ",:multi_site_origin"
	sqlString += ",:display_name"
	sqlString += ",:tr_response_headers"
	sqlString += ",:initial_dispersion"
	sqlString += ",:dns_bypass_cname"
	sqlString += ")"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putDeliveryservice
// @Description modify an existing deliveryserviceentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Deliveryservice   true "Deliveryservice object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice/{id}  [put]
func putDeliveryservice(id int, payload []byte) (interface{}, error) {
	var v Deliveryservice
	err := json.Unmarshal(payload, &v)
	v.Id = int64(id) // overwrite the id in the payload
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v.LastUpdated = time.Now()
	sqlString := "UPDATE deliveryservice SET "
	sqlString += "xml_id = :xml_id"
	sqlString += ",active = :active"
	sqlString += ",dscp = :dscp"
	sqlString += ",signed = :signed"
	sqlString += ",qstring_ignore = :qstring_ignore"
	sqlString += ",geo_limit = :geo_limit"
	sqlString += ",http_bypass_fqdn = :http_bypass_fqdn"
	sqlString += ",dns_bypass_ip = :dns_bypass_ip"
	sqlString += ",dns_bypass_ip6 = :dns_bypass_ip6"
	sqlString += ",dns_bypass_ttl = :dns_bypass_ttl"
	sqlString += ",org_server_fqdn = :org_server_fqdn"
	sqlString += ",type = :type"
	sqlString += ",profile = :profile"
	sqlString += ",cdn_id = :cdn_id"
	sqlString += ",ccr_dns_ttl = :ccr_dns_ttl"
	sqlString += ",global_max_mbps = :global_max_mbps"
	sqlString += ",global_max_tps = :global_max_tps"
	sqlString += ",long_desc = :long_desc"
	sqlString += ",long_desc_1 = :long_desc_1"
	sqlString += ",long_desc_2 = :long_desc_2"
	sqlString += ",max_dns_answers = :max_dns_answers"
	sqlString += ",info_url = :info_url"
	sqlString += ",miss_lat = :miss_lat"
	sqlString += ",miss_long = :miss_long"
	sqlString += ",check_path = :check_path"
	sqlString += ",last_updated = :last_updated"
	sqlString += ",protocol = :protocol"
	sqlString += ",ssl_key_version = :ssl_key_version"
	sqlString += ",ipv6_routing_enabled = :ipv6_routing_enabled"
	sqlString += ",range_request_handling = :range_request_handling"
	sqlString += ",edge_header_rewrite = :edge_header_rewrite"
	sqlString += ",origin_shield = :origin_shield"
	sqlString += ",mid_header_rewrite = :mid_header_rewrite"
	sqlString += ",regex_remap = :regex_remap"
	sqlString += ",cacheurl = :cacheurl"
	sqlString += ",remap_text = :remap_text"
	sqlString += ",multi_site_origin = :multi_site_origin"
	sqlString += ",display_name = :display_name"
	sqlString += ",tr_response_headers = :tr_response_headers"
	sqlString += ",initial_dispersion = :initial_dispersion"
	sqlString += ",dns_bypass_cname = :dns_bypass_cname"
	sqlString += " WHERE id=:id"
	result, err := db.GlobalDB.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delDeliveryserviceById
// @Description deletes deliveryservice information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Deliveryservice
// @Resource /api/2.0
// @Router /api/2.0/deliveryservice/{id} [delete]
func delDeliveryservice(id int) (interface{}, error) {
	arg := Deliveryservice{Id: int64(id)}
	result, err := db.GlobalDB.NamedExec("DELETE FROM deliveryservice WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
