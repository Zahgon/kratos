package apollo

const contentKey = "content"

type jsonExtParser struct{}

func (parser jsonExtParser) Parse(configContent any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type yamlExtParser struct{}

func (parser yamlExtParser) Parse(configContent any) (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
