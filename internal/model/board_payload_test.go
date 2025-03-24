package model

import (
	"fmt"
	"github.com/xiaoxlm/monitor-gateway/config"
	"testing"
)

func TestBoardPayload(t *testing.T) {
	db := config.Config.Mysql.GetDB()

	list, err := FetchByIDs(db, []uint{2})
	if err != nil {
		t.Fatal(err)
	}

	panel, err := GetPanelByBoardIDAndPanelID(list, 2, "b93d912c-b3bf-4c6e-a77b-77ea6123ffe9")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Print(panel)
}
