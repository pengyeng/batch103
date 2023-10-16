package main

type ProcessorType interface {
	Process(data []BatchData) ([]BatchData, error)
	SetParameters(values map[string]string)
}

type BaseProcessor struct {
	parameters map[string]string
}

/* Parameters Setter */
func (b *BaseProcessor) SetParameters(values map[string]string) {
	b.parameters = values
}

/* Parameters Getter */
func (b BaseProcessor) Parameters() map[string]string {
	return b.parameters
}
