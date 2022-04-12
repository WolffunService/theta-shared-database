package pubsublytic

type Auditlytic struct {
	Audit    *Audit    `json:"audit,omitempty"`
	Analytic *Analytic `json:"analytic,omitempty"`
}

type Audit struct {
	AuditName string      `json:"auditName"`
	AuditData interface{} `json:",inline"`
}

type Analytic struct {
	EventName    string      `json:"eventName"`
	Timestamp    int64       `json:"timestamp"`
	AnalyticData interface{} `json:",inline"`
}
