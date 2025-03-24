package model

import (
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type BoardPayload struct {
	ID
	Payload string `json:"-" gorm:"type:text"`
}

func (BoardPayload) TableName() string {
	return "board_payload"
}

// ConvertToDashboard converts a BoardPayload to a Dashboard
func (bp BoardPayload) ConvertToDashboard() (Dashboard, error) {
	var dashboard Dashboard

	if bp.Payload == "" {
		return dashboard, nil
	}

	err := json.Unmarshal([]byte(bp.Payload), &dashboard)
	if err != nil {
		return Dashboard{}, fmt.Errorf("failed to unmarshal board payload: %w", err)
	}

	return dashboard, nil
}

func GetPanelByBoardIDAndPanelID(boardPayloads []BoardPayload, boardPayloadID uint, panelID string) (*Panel, error) {

	for _, bp := range boardPayloads {
		if bp.ID.ID != boardPayloadID {
			continue
		}

		dashboard, err := bp.ConvertToDashboard()
		if err != nil {
			return nil, fmt.Errorf("failed to convert board payload with ID %d: %w", bp.ID.ID, err)
		}

		for _, panel := range dashboard.Panels {
			if panel.ID == panelID {
				return &panel, nil
			}
		}
	}

	return nil, fmt.Errorf("failed to find panel with boardPayloadID %d, panelID: %s", boardPayloadID, panelID)
}

// FetchByIDs retrieves board payloads by their IDs
func FetchByIDs(db *gorm.DB, ids []uint) ([]BoardPayload, error) {
	var boardPayloads []BoardPayload

	if len(ids) == 0 {
		return boardPayloads, nil
	}

	if err := db.Where("id IN ?", ids).Find(&boardPayloads).Error; err != nil {
		return nil, err
	}

	return boardPayloads, nil
}

type Dashboard struct {
	Version string        `json:"version"`
	Links   []interface{} `json:"links"`
	Var     []Variable    `json:"var"`
	Panels  []Panel       `json:"panels"`
}

type Variable struct {
	Type       string      `json:"type,omitempty"`
	Name       string      `json:"name"`
	Label      string      `json:"label,omitempty"`
	Hide       bool        `json:"hide,omitempty"`
	Datasource *Datasource `json:"datasource,omitempty"`
	Definition string      `json:"definition,omitempty"`
	AllValue   interface{} `json:"allValue,omitempty"`
	AllOption  bool        `json:"allOption,omitempty"`
	Multi      bool        `json:"multi,omitempty"`
	Reg        string      `json:"reg,omitempty"`
}

type Datasource struct {
	Cate  string `json:"cate"`
	Value string `json:"value"`
}

type Panel struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Type            string           `json:"type"`
	Links           []Link           `json:"links"`
	Custom          Custom           `json:"custom"`
	Layout          Layout           `json:"layout"`
	Options         Options          `json:"options"`
	Targets         []Tgt            `json:"targets"`
	Version         string           `json:"version"`
	MaxPerRow       int              `json:"maxPerRow"`
	Description     string           `json:"description"`
	DatasourceCate  string           `json:"datasourceCate"`
	DatasourceValue string           `json:"datasourceValue"`
	Transformations []Transformation `json:"transformations"`
}

type Link struct {
	// Add fields as needed
}

type Custom struct {
	Calc       string `json:"calc"`
	TextMode   string `json:"textMode"`
	ValueField string `json:"valueField"`
}

type Layout struct {
	H           int    `json:"h"`
	I           string `json:"i"`
	W           int    `json:"w"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	IsResizable bool   `json:"isResizable"`
}

type Options struct {
	Thresholds      Thresholds      `json:"thresholds"`
	StandardOptions StandardOptions `json:"standardOptions"`
}

type Thresholds struct {
	Steps []Step `json:"steps"`
}

type Step struct {
	Type  string   `json:"type"`
	Color string   `json:"color"`
	Value *float64 `json:"value"` // Using pointer to handle null values
}

type StandardOptions struct {
	Max      float64 `json:"max"`
	Min      float64 `json:"min"`
	Util     string  `json:"util"`
	Decimals int     `json:"decimals"`
}

type Tgt struct {
	Expr          string `json:"expr"`
	RefID         string `json:"refId"`
	MaxDataPoints int    `json:"maxDataPoints"`
}

type Transformation struct {
	ID      string                 `json:"id"`
	Options map[string]interface{} `json:"options"`
}
