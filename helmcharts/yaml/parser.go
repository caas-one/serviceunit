package yaml

import (
	"io/ioutil"

	mlog "github.com/maxwell92/log"
)

var log = mlog.Log

// Parser parse the profile.yaml to Profile Object
type Parser struct {
	Profile Profile
}

// NewParser function give a new parser instance
func NewParser() *Parser {
	return &Parser{
		Profile: Profile{
			ServiceUnits: make([]ServiceUnit, 0),
		},
	}
}

// Parse function parse the profile.yaml
func (p *Parser) Parse(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf("Read profile.yaml error: err=%s", err)
	}

	p.Profile.UnmarshalFromYaml(data)
	if err != nil {
		log.Errorf("yaml.Unmarshal error: err=%s", err)
	}
	return err
}
