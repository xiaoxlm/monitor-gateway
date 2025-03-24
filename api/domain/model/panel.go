package model

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
