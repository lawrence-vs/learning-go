package config

type IOReplace map[string]string

type SheetType struct {
	Types []string
}

type ModuleType struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ModuleTypes map[string]ModuleType

type IOHeader struct {
	DocumentTitle  string `json:"Document Title"`
	DocumentNumber string `json:"Document Number"`
	RevisionNumber string `json:"Revision Number"`
	RevisionDate   string `json:"Revision Date"`
}

type DocumentFormat struct {
	Header      IOHeader
	ColumnNames []string
}

// Configuration holds the application settings
type Configuration struct {
	IOReplace   IOReplace      `json:"IOReplace"`
	SheetTypes  []string       `json:"SheetTypes"`
	ModuleTypes ModuleTypes    `json:"ModuleTypes"`
	IODocument  DocumentFormat `json:"IODocument"`
}
