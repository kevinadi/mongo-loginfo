package main

import (
	"sync"
	"testing"
)

func TestPatternInitandlisten(t *testing.T) {
	var wg_main sync.WaitGroup
	r := res_initandlisten
	ch := make(chan string)

	go Matcher(func_array_initandlisten, ch, output, &wg_main)
	wg_main.Add(1)

	ch <- "2017-06-25T22:41:52.991-0400 I CONTROL  [initandlisten] db version v3.4.3"
	ch <- `2017-08-04T11:28:47.780-0500 I CONTROL  [initandlisten] options: { auditLog: { destination: "syslog" }, config: "/data/rs0_8/automation-mongod.conf", net: { compression: { compressors: "snappy" }, port: 27003 }, processManagement: { fork: true }, replication: { replSetName: "rs0" }, security: { enableEncryption: true, encryptionKeyFile: "/etc/mongodb-keyfile"}, authorization: "enabled", keyFile: "/opt/app/mongo/mongodb-mms-automation/keyfile", storage: { dbPath: "/data/rs0_8", engine: "wiredTiger" }, systemLog: { destination: "file", path: "/data/rs0_8/mongodb.log" } }`
	close(ch)
	wg_main.Wait()

	if r.db_version != "3.4.3" {
		t.Error("db version is", r.db_version, "expecting 3.4.3")
	}
	if r.storage_engine != "wiredTiger" {
		t.Error("storage engine is", r.storage_engine, "expecting wiredTiger")
	}
	if r.enterprise != false {
		t.Error("enterprise is", r.enterprise, "expecting false")
	}
	if r.audit != "syslog" {
		t.Error("audit is", r.audit, "expecting syslog")
	}
	if r.keyfile != "/opt/app/mongo/mongodb-mms-automation/keyfile" {
		t.Error("keyfile is", r.keyfile, "expecting /opt/app/mongo/mongodb-mms-automation/keyfile")
	}
}
