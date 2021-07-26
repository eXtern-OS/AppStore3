package publisher

type Publisher struct {
}

type ExportedPublisher struct {
}

func (p *Publisher) Export() ExportedPublisher {
	return ExportedPublisher{}
}

func NewPublisher(name string) ExportedPublisher {
	return ExportedPublisher{}
}
