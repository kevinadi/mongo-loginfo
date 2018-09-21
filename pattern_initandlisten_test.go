package main

import (
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_initandlisten_storage_engine_1(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_storage_engine_1(&res, ch, &wg)
	wg.Add(1)

	ch <- `2017-08-04T11:28:47.780-0500 I CONTROL  [initandlisten] options: { auditLog: { destination: "syslog" }, config: "/data/rs0_8/automation-mongod.conf", net: { compression: { compressors: "snappy" }, port: 27003 }, processManagement: { fork: true }, replication: { replSetName: "rs0" }, security: { enableEncryption: true, encryptionKeyFile: "/etc/mongodb-keyfile"}, authorization: "enabled", keyFile: "/opt/app/mongo/mongodb-mms-automation/keyfile", storage: { dbPath: "/data/rs0_8", engine: "wiredTiger" }, systemLog: { destination: "file", path: "/data/rs0_8/mongodb.log" } }`
	close(ch)
	wg.Wait()

	if res != "wiredTiger" {
		t.Error("storage engine is", res, "expecting wiredTiger")
	}
}

func Test_initandlisten_storage_engine_2(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_storage_engine_2(&res, ch, &wg)
	wg.Add(1)

	ch <- `2018-09-04T19:46:40.120+1000 I CONTROL  [initandlisten] options: { storage: { dbPath: "." }, systemLog: { destination: "file", path: "./ent.log" } }`
	ch <- `2018-09-04T19:46:40.127+1000 I STORAGE  [initandlisten] wiredtiger_open config: create,cache_size=7680M,session_max=20000,eviction=(threads_min=4,threads_max=4),config_base=false,statistics=(fast),cache_cursors=false,log=(enabled=true,archive=true,path=journal,compressor=snappy),file_manager=(close_idle_time=100000),statistics_log=(wait=0),verbose=(recovery_progress),`
	ch <- `2018-09-04T19:46:40.777+1000 I STORAGE  [initandlisten] WiredTiger message [1536054400:777902][27894:0x7fffa8e63380], txn-recover: Set global recovery timestamp: 0`

	close(ch)
	wg.Wait()

	if res != "wiredtiger" {
		t.Error("storage engine is", res, "expecting wiredtiger")
	}
}

func Test_initandlisten_storage_engine_3(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_storage_engine_1(&res, ch, &wg)
	wg.Add(1)

	ch <- `2017-08-04T11:28:47.780-0500 I CONTROL  [initandlisten] options: { auditLog: { destination: "syslog" }, config: "/data/rs0_8/automation-mongod.conf", net: { compression: { compressors: "snappy" }, port: 27003 }, processManagement: { fork: true }, replication: { replSetName: "rs0" }, security: { enableEncryption: true, encryptionKeyFile: "/etc/mongodb-keyfile"}, authorization: "enabled", keyFile: "/opt/app/mongo/mongodb-mms-automation/keyfile", storage: { dbPath: "/data/rs0_8", engine: "wiredTiger" }, systemLog: { destination: "file", path: "/data/rs0_8/mongodb.log" } }`

	close(ch)
	wg.Wait()

	if strings.ToLower(res) != "wiredtiger" {
		t.Error("storage engine is", res, "expecting wiredtiger")
	}
}

func Test_initandlisten_host(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_host(&res, ch, &wg)
	wg.Add(1)

	ch <- `2018-09-04T19:46:40.120+1000 I CONTROL  [initandlisten] MongoDB starting : pid=27894 port=27017 dbpath=. 64-bit host=Triptykon.local
`

	close(ch)
	wg.Wait()

	if res != "Triptykon.local" {
		t.Error("Host is", res, "expecting Triptykon.local")
	}
}

func Test_initandlisten_port(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_port(&res, ch, &wg)
	wg.Add(1)

	ch <- `2018-09-04T19:46:40.120+1000 I CONTROL  [initandlisten] MongoDB starting : pid=27894 port=27017 dbpath=. 64-bit host=Triptykon.local
`

	close(ch)
	wg.Wait()

	if res != "27017" {
		t.Error("Port is", res, "expecting 27017")
	}
}

func Test_initandlisten_db_version(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_db_version(&res, ch, &wg)
	wg.Add(1)

	ch <- `2018-09-04T19:46:40.120+1000 I CONTROL  [initandlisten] db version v3.6.5`

	close(ch)
	wg.Wait()

	if res != "3.6.5" {
		t.Error("db version is", res, "expecting 3.6.5")
	}
}

func Test_initandlisten_keyfile(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_keyfile(&res, ch, &wg)
	wg.Add(1)

	ch <- `2018-02-08T13:21:09.966Z I CONTROL  [initandlisten] options: { auditLog: { destination: "file", format: "BSON", path: "/audit/auditLog.bson" }, config: "/mongodb1.conf", net: { bindIp: true, http: { JSONPEnabled: false, RESTInterfaceEnabled: false, enabled: false }, ipv6: false, port: 27033, ssl: { allowConnectionsWithoutCertificates: false, allowInvalidCertificates: false, allowInvalidHostnames: false } }, replication: { oplogSizeMB: 2048, replSetName: "REPENH033" }, security: { authorization: "enabled", enableEncryption: true, encryptionCipherMode: "AES256-CBC", encryptionKeyFile: "/encryption.key", keyFile: "/replica.key" }, setParameter: { authenticationMechanisms: "PLAIN,MONGODB-CR", saslauthdPath: "/run/saslauthd/mux" }, storage: { wiredTiger: { engineConfig: { cacheSizeGB: 4.0 } } }, systemLog: { destination: "file", logAppend: false, logRotate: "rename", path: "/var/log/mongodb/mongodb.log", timeStampFormat: "iso8601-utc", verbosity: 0 } }`

	close(ch)
	wg.Wait()

	if res != "/replica.key" {
		t.Error("keyfile is", res, "expecting /replica.key")
	}
}

func Test_initandlisten_audit(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_audit(&res, ch, &wg)
	wg.Add(1)

	ch <- `2018-02-08T13:21:09.966Z I CONTROL  [initandlisten] options: { auditLog: { destination: "file", format: "BSON", path: "/audit/auditLog.bson" }, config: "/mongodb1.conf", net: { bindIp: true, http: { JSONPEnabled: false, RESTInterfaceEnabled: false, enabled: false }, ipv6: false, port: 27033, ssl: { allowConnectionsWithoutCertificates: false, allowInvalidCertificates: false, allowInvalidHostnames: false } }, replication: { oplogSizeMB: 2048, replSetName: "REPENH033" }, security: { authorization: "enabled", enableEncryption: true, encryptionCipherMode: "AES256-CBC", encryptionKeyFile: "/encryption.key", keyFile: "/replica.key" }, setParameter: { authenticationMechanisms: "PLAIN,MONGODB-CR", saslauthdPath: "/run/saslauthd/mux" }, storage: { wiredTiger: { engineConfig: { cacheSizeGB: 4.0 } } }, systemLog: { destination: "file", logAppend: false, logRotate: "rename", path: "/var/log/mongodb/mongodb.log", timeStampFormat: "iso8601-utc", verbosity: 0 } }`

	close(ch)
	wg.Wait()

	if res != "file" {
		t.Error("audit is", res, "expecting file")
	}
}

func Test_initandlisten_enterprise(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_enterprise(&res, ch, &wg)
	wg.Add(1)

	ch <- `2018-09-04T19:46:40.120+1000 I CONTROL  [initandlisten] modules: enterprise `

	close(ch)
	wg.Wait()

	if res != "true" {
		t.Error("enterprise is", res, "expecting true")
	}
}

func Test_initandlisten_encryption(t *testing.T) {
	var wg sync.WaitGroup
	var res string

	ch := make(chan string)
	go Matcher_encryption(&res, ch, &wg)
	wg.Add(1)

	ch <- `2018-02-08T13:21:09.966Z I CONTROL  [initandlisten] options: { auditLog: { destination: "file", format: "BSON", path: "/audit/auditLog.bson" }, config: "/mongodb1.conf", net: { bindIp: true, http: { JSONPEnabled: false, RESTInterfaceEnabled: false, enabled: false }, ipv6: false, port: 27033, ssl: { allowConnectionsWithoutCertificates: false, allowInvalidCertificates: false, allowInvalidHostnames: false } }, replication: { oplogSizeMB: 2048, replSetName: "REPENH033" }, security: { authorization: "enabled", enableEncryption: true, encryptionCipherMode: "AES256-CBC", encryptionKeyFile: "/encryption.key", keyFile: "/replica.key" }, setParameter: { authenticationMechanisms: "PLAIN,MONGODB-CR", saslauthdPath: "/run/saslauthd/mux" }, storage: { wiredTiger: { engineConfig: { cacheSizeGB: 4.0 } } }, systemLog: { destination: "file", logAppend: false, logRotate: "rename", path: "/var/log/mongodb/mongodb.log", timeStampFormat: "iso8601-utc", verbosity: 0 } }`

	close(ch)
	wg.Wait()

	if res != "true" {
		t.Error("encryption is", res, "expecting true")
	}
}

func Test_initandlisten_group_1(t *testing.T) {
	var wg sync.WaitGroup

	ch := make(chan string)
	outch := make(chan *Res_InitAndListen)
	go MatcherGroup_initandlisten(ch, outch, &wg)
	wg.Add(1)

	ch <- `2017-06-25T22:41:52.991-0400 I CONTROL  [initandlisten] db version v3.4.3`
	ch <- `2018-03-15T01:27:13.858-0700 I CONTROL  [initandlisten] MongoDB starting : pid=18315 port=27017 dbpath=/base/data/mongo 64-bit host=fancyhostname`
	ch <- `2018-02-08T13:21:09.966Z I CONTROL  [initandlisten] options: { auditLog: { destination: "file", format: "BSON", path: "/audit/auditLog.bson" }, config: "/mongodb1.conf", net: { bindIp: true, http: { JSONPEnabled: false, RESTInterfaceEnabled: false, enabled: false }, ipv6: false, port: 27033, ssl: { allowConnectionsWithoutCertificates: false, allowInvalidCertificates: false, allowInvalidHostnames: false } }, replication: { oplogSizeMB: 2048, replSetName: "REPENH033" }, security: { authorization: "enabled", enableEncryption: true, encryptionCipherMode: "AES256-CBC", encryptionKeyFile: "/encryption.key", keyFile: "/replica.key" }, setParameter: { authenticationMechanisms: "PLAIN,MONGODB-CR", saslauthdPath: "/run/saslauthd/mux" }, storage: { wiredTiger: { engineConfig: { cacheSizeGB: 4.0 } } }, systemLog: { destination: "file", logAppend: false, logRotate: "rename", path: "/var/log/mongodb/mongodb.log", timeStampFormat: "iso8601-utc", verbosity: 0 } }`
	ch <- `2018-02-08T13:21:09.966Z I CONTROL  [initandlisten] modules: enterprise`
	ch <- `2018-04-06T15:44:27.119-0500 I ACCESS   [conn2598] Successfully authenticated as principal mms-automation on admin`
	close(ch)
	res := <-outch
	wg.Wait()

	assert.Equal(t, res.host, "fancyhostname", "host")
	assert.Equal(t, res.port, "27017", "port")
	assert.Equal(t, res.db_version, "3.4.3", "db_version")
	assert.Equal(t, res.storage_engine, "wiredTiger", "storage_engine")
	assert.Equal(t, res.auth, "true", "auth")
	assert.Equal(t, res.auth_type, "keyfile", "auth_type")
	assert.Equal(t, res.keyfile, "/replica.key", "keyfile")
	assert.Equal(t, res.encrypted, "true", "encrypted")
	assert.Equal(t, res.enterprise, "true", "enterprise")
	assert.Equal(t, res.audit, "file", "audit")
}
