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
	Type            string                 `json:"type"`
	ID              string                 `json:"id"`
	Layout          Layout                 `json:"layout"`
	Version         string                 `json:"version"`
	DatasourceCate  string                 `json:"datasourceCate"`
	DatasourceValue string                 `json:"datasourceValue"`
	Targets         []Target               `json:"targets"`
	Transformations []Transformation       `json:"transformations"`
	Name            string                 `json:"name"`
	MaxPerRow       int                    `json:"maxPerRow,omitempty"`
	Custom          map[string]interface{} `json:"custom,omitempty"`
	Options         map[string]interface{} `json:"options,omitempty"`
	Links           []interface{}          `json:"links,omitempty"`
	Description     string                 `json:"description,omitempty"`
	Collapsed       bool                   `json:"collapsed,omitempty"`
	Panels          []Panel                `json:"panels,omitempty"`
	GraphTooltip    string                 `json:"graphTooltip,omitempty"`
	Overrides       []Override             `json:"overrides,omitempty"`
}

type Layout struct {
	H           int    `json:"h"`
	W           int    `json:"w"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	I           string `json:"i"`
	IsResizable bool   `json:"isResizable"`
}

type Target struct {
	RefId         string `json:"refId,omitempty"`
	Expr          string `json:"expr"`
	MaxDataPoints int    `json:"maxDataPoints,omitempty"`
	Legend        string `json:"legend,omitempty"`
	Instant       bool   `json:"instant,omitempty"`
	Mode          string `json:"__mode__,omitempty"`
}

type Transformation struct {
	ID      string                 `json:"id"`
	Options map[string]interface{} `json:"options"`
}

type Override struct {
	Matcher    Matcher     `json:"matcher"`
	Properties interface{} `json:"properties"`
}

type Matcher struct {
	ID string `json:"id"`
}
