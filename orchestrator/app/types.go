package app

// 0 => Success, -1 => Fail
type Response struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type DashboardResponse struct {
	ID           any    `json:"id"`
	UID          string `json:"uid"`
	Title        string `json:"title"`
	Tags         []any  `json:"tags"`
	Timezone     string `json:"timezone"`
	Editable     bool   `json:"editable"`
	GraphTooltip int    `json:"graphTooltip"`
	Panels       []any  `json:"panels"`
	Time         struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"time"`
	Timepicker struct {
		RefreshIntervals []any `json:"refresh_intervals"`
	} `json:"timepicker"`
	Templating struct {
		List []any `json:"list"`
	} `json:"templating"`
	Annotations struct {
		List []any `json:"list"`
	} `json:"annotations"`
	Refresh       string `json:"refresh"`
	SchemaVersion int    `json:"schemaVersion"`
	Version       int    `json:"version"`
	Links         []any  `json:"links"`
}
