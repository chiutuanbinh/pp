package db

import (
	"context"
	"testing"
)

func TestConnect(t *testing.T) {
	conn, err := pool.Acquire(context.TODO())
	if err != nil {
		t.Fail()
	}
	if conn == nil {
		t.Fail()
	}
	t.Log(conn.Ping(context.Background()))
}

func TestAddLegalEntity(t *testing.T) {
	var id int = addLegalEntity()
	t.Log(id)
	if id == -1 {
		t.Fail()
	}
}

func TestAddCompany(t *testing.T) {
	var companyId = AddCompany("ABC")
	if companyId == -1 {
		t.Fail()
	}
	t.Log(companyId)
}

func TestAddPerson(t *testing.T) {
	var personId = AddPerson("NGUYEN VAN A", 52, "http://sth.sthm")
	if personId == -1 {
		t.Fail()
	}
	t.Log(personId)
}
