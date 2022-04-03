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

func TestGetAllCompany(t *testing.T) {
	var res = GetAllCompany()
	t.Log(res)
}

func TestAddPerson(t *testing.T) {
	var personId = AddPerson("NGUYEN VAN A", 52, "http://sth.sthm")
	if personId == -1 {
		t.Fail()
	}
	t.Log(personId)
}

func TestAddOwnerships(t *testing.T) {
	var legalEntityId = addLegalEntity()
	var company_id = AddCompany("ABC")
	var ownershipId = AddOwnerships(legalEntityId, company_id)

	if ownershipId == -1 {
		t.Fail()
	}

	var personId = AddPerson("A", 19, "httplsdf")
	var oId = AddPersonOwnerships(personId, company_id)
	if oId == -1 {
		t.Fail()
	}

	var companyB = AddCompany("ABCD")
	var oId2 = AddCompanyOwnerships(company_id, companyB)
	if oId2 == -1 {
		t.Fail()
	}
}

func TestAddIndustryCategory(t *testing.T) {
	var industryCategoryId = AddIndustryCategory("STEEL")
	t.Log(industryCategoryId)
	if industryCategoryId == -1 {
		t.Fail()
	}
}
