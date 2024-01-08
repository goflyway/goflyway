package history

type SchemaData struct {
	InstalledRank int    `json:"installedRank" column:"installed_rank"`
	Version       string `json:"version" column:"version"`
	Description   string `json:"description" column:"description"`
	Type          string `json:"type" column:"type"`
	Script        string `json:"script" column:"script"`
	Checksum      int    `json:"checksum" column:"checksum"`
	InstalledBy   string `json:"installedBy" column:"installed_by"`
	InstalledOn   string `json:"installedOn" column:"installed_on"`
	ExecutionTime int    `json:"executionTime" column:"execution_time"`
	Success       bool   `json:"success" column:"success"`
}
