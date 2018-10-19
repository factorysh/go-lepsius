package input

import (
	"bufio"
	"os"

	"github.com/factorysh/go-lepsius/model"
	"github.com/factorysh/go-lepsius/parser"
	"github.com/mitchellh/mapstructure"
)

func init() {
	register("stdin", &Stdin{})
}

type Stdin struct {
	config *StdinConf
	parser model.Parser
}

type StdinConf struct {
	parser string
}

func (s *Stdin) Lines() chan *model.Line {
	lines := make(chan *model.Line)
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for scanner.Scan() {
			line := scanner.Bytes()
			l, err := s.parser.Parse(line)
			if err != nil {
				panic(err)
			}
			lines <- &model.Line{
				Values: map[string]interface{}{
					"message": l,
				},
			}
		}
	}()
	return lines
}

func (s *Stdin) Configure(conf map[string]interface{}) error {
	var cfg StdinConf
	err := mapstructure.Decode(conf, &cfg)
	if err != nil {
		return err
	}
	if cfg.parser == "" {
		cfg.parser = "raw"
	}
	s.config = &cfg
	s.parser = parser.Parser[cfg.parser]
	return nil
}
