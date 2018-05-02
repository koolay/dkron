package flow

// EnumDB  enum of db
type EnumDB string

// Flow can contains nodes
type Flow struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DependFlowID string `json:"depend_flow_id"`
	BuildIn      int    `json:"build_in"`
	RunStatus    EnumDB `json:"run_status"`
	Schedule     string `json:"schedule"`
	Status       EnumDB `json:"status"`
	Type         EnumDB `json:"type"`
}

// Node node of flow
type Node struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	Description string `json:"description"`
	FlowID      string `json:"flow_id"`
	IsEnd       int8   `json:"is_end"`
	IsStart     int8   `json:"is_start"`
	Name        string `json:"name"`
	Position    string `json:"position"`
	Type        EnumDB `json:"type"`
}

// Line line of flow
type Line struct {
	ID           string `json:"id"`
	FlowID       string `json:"flow_id"`
	AheadNodeID  string `json:"ahead_node_id"`
	BehindNodeID string `json:"behind_node_id"`
}
