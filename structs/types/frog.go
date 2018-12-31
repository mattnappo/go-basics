package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/fatih/color"
)

// ErrNilParameter - Error for invalid frog constructor parameters
var ErrNilParameter = errors.New("invalid parameters to construct frog")

// ErrCannotRee - When the frog cant ree
var ErrCannotRee = errors.New("cannot ree")

// Frog - a frog that may be able to ree
type Frog struct {
	Name       string  `json:"name"`
	Weight     float64 `json:"weight"`
	ReeAbility bool    `json:"ReeAbility"`
}

// NewFrog - initialize new frog instance
func NewFrog(name string, weight float64, reeAbility bool) (*Frog, error) {
	if name == "" || weight == 0 {
		return nil, ErrNilParameter
	}

	return &Frog{
		Name:       name,
		Weight:     weight,
		ReeAbility: reeAbility,
	}, nil
}

// Ree - reeee!
func (frog *Frog) Ree() error {
	if frog.ReeAbility {
		color.Red("frog: reeeeeeeeeeee!!!!!")
		color.Green("frog: reeeeeeeeeeee!!!!!")
		color.Blue("frog: reeeeeeeeeeee!!!!!")
		return nil
	}
	return ErrCannotRee
}

// String - encode frog to string
func (frog *Frog) String() string {
	json, _ := json.MarshalIndent(*frog, "", "  ") // Marshal JSON

	return string(json)
}

// Bytes - encode frog to byte array
func (frog *Frog) Bytes() []byte {
	json, _ := json.MarshalIndent(*frog, "", "  ") // Marhsal JSON

	return json
}

// FrogFromBytes - decode byte array into Frog
func FrogFromBytes(b []byte) (*Frog, error) {
	buffer := &Frog{}
	err := json.Unmarshal(b, buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil // Return buffer
}

// WriteToMemory - Write frog to persistnfe memory
func (frog *Frog) WriteToMemory() error {
	json, err := json.MarshalIndent(*frog, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("froggy_%s.json", frog.Name), json, 0644)
	if err != nil {
		return err
	}
	return nil
}

// ReadFrogFromMemory - read a frog from json
func ReadFrogFromMemory(name string) (*Frog, error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("froggy_%s.json", name))
	if err != nil {
		return nil, err
	}
	buffer := &Frog{}
	err = json.Unmarshal(data, buffer)
	if err != nil {
		return nil, err
	}
	return buffer, nil
}
