package parser

type Raw struct {
}

func init() {
	register("raw", &Raw{})
}

func (r *Raw) Configure(conf map[string]interface{}) error {
	return nil
}

func (r *Raw) Parse(input []byte) (map[string]interface{}, error) {
	return map[string]interface{}{
		"message": string(input),
	}, nil
}
