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
	if id == 0 {
		t.Fail()
	}
}

func TestAddCompany(t *testing.T) {
	var companyId = AddCompany("ABC", -1)
	if companyId == 0 {
		t.Fail()
	}
	t.Log(companyId)
}

func TestAddExchange(t *testing.T) {
	var exchangeId = AddExchange("Over the counter", "OTC")
	if exchangeId == 0 {
		t.Fail()
	}
	t.Log(exchangeId)
}

func TestGetAllCompany(t *testing.T) {
	var res = GetAllCompanies()
	t.Log(res)
}

func TestAddPerson(t *testing.T) {
	var personId = AddPerson("NGUYEN VAN A", 52, "http://sth.sthm", 0)
	if personId == 0 {
		t.Fail()
	}
	t.Log(personId)
}

func TestAddOwnerships(t *testing.T) {
	var legalEntityId = addLegalEntity()
	var exchangeId = AddExchange("", "HOSE")
	var stockId = AddStock("ACB", exchangeId)
	var ownershipId = AddOwnerships(legalEntityId, stockId, 100)

	if ownershipId == 0 {
		t.Fail()
	}

	var personId = AddPerson("A", 19, "httplsdf", 0)
	var oId = AddPersonOwnerships(personId, stockId, 100)
	if oId == 0 {
		t.Fail()
	}

	var companyB = AddCompany("ABCD", -1)
	var oId2 = AddCompanyOwnerships(companyB, stockId, 100)
	if oId2 == 0 {
		t.Fail()
	}
}

func TestAddIndustryCategory(t *testing.T) {
	var industryCategoryId = AddIndustryCategory("STEEL")
	t.Log(industryCategoryId)
	if industryCategoryId == 0 {
		t.Fail()
	}
}

func TestAddIndustry(t *testing.T) {
	var industryCategoryId = AddIndustryCategory("PRODUCTION")
	var industryId = AddIndustry("STEEL", industryCategoryId, "")
	t.Log(industryId)
	if industryId == 0 {
		t.Fail()
	}
}
