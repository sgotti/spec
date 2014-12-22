package types

import (
	"encoding/json"
	"errors"
)

type Dependencies []Dependency

type Dependency struct {
	App     ACName `json:"app"`
	ImageID Hash   `json:"imageID,omitempty"`
	Labels  Labels `json:"labels"`
}

type dependency Dependency

func (d Dependency) assertValid() error {
	if len(d.App) < 1 {
		return errors.New(`App cannot be empty`)
	}
	return nil
}

func (d Dependency) MarshalJSON() ([]byte, error) {
	if err := d.assertValid(); err != nil {
		return nil, err
	}

	if d.ImageID.Empty() {
		return json.Marshal(struct {
			dependency
			// Override ImageID with an empty value so it won't be marshalled
			ImageID bool `json:"imageID,omitempty"`
		}{
			dependency: dependency(d),
		})

	}
	return json.Marshal(dependency(d))
}

func (d *Dependency) UnmarshalJSON(data []byte) error {
	var jd dependency
	if err := json.Unmarshal(data, &jd); err != nil {
		return err
	}
	nd := Dependency(jd)
	if err := nd.assertValid(); err != nil {
		return err
	}
	*d = nd
	return nil
}
