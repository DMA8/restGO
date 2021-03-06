package domain

import (
	"time"
)

const (
	timeFormat = "02-01-2006"
)

type Props struct {
	Props []*Prop `json"props"`
}

type Prop struct {
	ID            uint    `json:"id"`
	RawInstalDate string `json:"date,omitempty"`
	InstalDate    time.Time
}

func (p *Props) ConverTime() error {
	var err error
	for _, prop := range p.Props {
		if prop.RawInstalDate != "" {
			prop.InstalDate, err = time.Parse("2006-02-01", prop.RawInstalDate)
			if err != nil {
				return err
			}
		}
	}
	return err
}
