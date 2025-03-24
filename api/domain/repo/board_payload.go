package repo

import (
	"context"
	"encoding/json"
	"github.com/xiaoxlm/monitor-gateway/api/domain/model"
	"gorm.io/gorm"
)

func GetPanelContent(ctx context.Context, db *gorm.DB, id uint, panelID string) (*model.Panel, error) {
	var result struct {
		PanelContent string `gorm:"column:panel_content"`
	}

	err := db.WithContext(ctx).Table("board_payload").
		Select("JSON_EXTRACT(payload, REPLACE(JSON_UNQUOTE(JSON_SEARCH(payload, 'one', ?, NULL, '$.panels[*].id')), '.id', '')) as panel_content",
			panelID).Where("id = ?", id).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	var ret = &model.Panel{}
	if err = json.Unmarshal([]byte(result.PanelContent), ret); err != nil {
		return nil, err
	}

	return ret, nil
}
