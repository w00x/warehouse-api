package v1

import (
	"os"
	"testing"
	"warehouse/tests"
)

func TestMain(m *testing.M) {
	tests.SetDbConn()
	ret := m.Run()
	tests.CleanStock()
	tests.CleanPrice()
	tests.CleanInventory()
	tests.CleanItem()
	tests.CleanMarket()
	tests.CleanRack()
	os.Exit(ret)
}
