package pubsublytic

type Auditlytic struct {
	Audit    *Audit    `json:"audit,omitempty"`
	Analytic *analytic `json:"analytic,omitempty"`
}

type Audit struct {
	AuditName string      `json:"auditName"`
	AuditData interface{} `json:",inline"`
}

type analytic struct {
	EventName    string      `json:"eventName"`
	Timestamp    int64       `json:"timestamp"`
	AnalyticData interface{} `json:",inline"`
}
