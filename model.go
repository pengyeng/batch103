package batch103

const (
	StgRead    = "READ"
	StgProcess = "PROCESS"
	StgWrite   = "WRITE"
	StatActive = "A"
	StatReject = "R"
)

type BatchData struct {
	GenericData []string
	status      string
	stage       string
}

func (b *BatchData) Reject(currStage string) *BatchData {
	b.status = StatReject
	b.stage = currStage
	return b
}

func (b *BatchData) IsActive() bool {
	if b.status == StatActive {
		return true
	} else {
		return false
	}
}

/* Status Getter */
func (b BatchData) Status() string {
	return b.status
}

/* Status Getter */
func (b BatchData) Stage() string {
	return b.stage
}

/* Constructor BEGIN */
func (b *BatchData) Create(data []string) *BatchData {
	return &BatchData{
		GenericData: data,
		status:      StatActive,
		stage:       StgRead,
	}
}
