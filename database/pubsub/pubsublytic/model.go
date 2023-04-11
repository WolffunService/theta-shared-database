package pubsublytic

import "github.com/WolffunService/theta-shared-database/database/pubsub/auditprotobuf"

type Auditlytic[T any] struct {
	Audit    *Audit[T] `json:"audit,omitempty"`
	Analytic *analytic `json:"analytic,omitempty"`
}

type Audit[T any] struct {
	AuditName string `json:"auditName"`
	AuditData T      `json:",inline"`
}

type analytic struct {
	EventName    string                    `json:"eventName"`
	Timestamp    int64                     `json:"timestamp"`
	AnalyticData []*auditprotobuf.KeyPair2 `json:",inline"`
}
