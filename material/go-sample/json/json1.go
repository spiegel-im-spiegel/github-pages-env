package main

import (
	"encoding/json"
	"fmt"
)

//Info is information class for OpenPGP packets
type Info struct {
	Packets []*Item `toml:"packets,omitempty" json:"packets,omitempty"`
}

//Item is information item class
type Item struct {
	Name  string  `toml:"name" json:"name"`
	Value string  `toml:"value,omitempty" json:"value,omitempty"`
	Items []*Item `toml:"items,omitempty" json:"items,omitempty"`
}

//JSON function
func (i *Info) JSON() ([]byte, error) {
	j, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return nil, nil
	}
	return j, nil
}

func main() {
	s := &Info{}
	//j, err := json.MarshalIndent(s, "", "  ")
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(j))
	j, _ := s.JSON()
	fmt.Println(string(j))
}
