package repo

import (
	"context"
	"github.com/xiaoxlm/monitor-gateway/internal/model"
	"gorm.io/gorm"
)

func ListBoardPayload(ctx context.Context, db *gorm.DB, boardPayloadIDList []uint) ([]model.BoardPayload, error) {
	return model.FetchByIDs(db, boardPayloadIDList)
}
