package types

import "fmt"

// NewGenesisState creates a new GenesisState object
func NewGenesisState(classCollections []ClassCollection) *GenesisState {
	return &GenesisState{
		ClassCollections: classCollections,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	seenErc721 := make(map[string]bool)
	seenClass := make(map[string]bool)

	for _, b := range gs.ClassCollections {
		if seenErc721[b.ClassPair.ContractAddress] {
			return fmt.Errorf("token ERC721 contract duplicated on genesis '%s'", b.ClassPair.ContractAddress)
		}
		if seenClass[b.ClassPair.ClassId] {
			return fmt.Errorf("class duplicated on genesis: '%s'", b.ClassPair.ClassId)
		}

		if err := b.ClassPair.Validate(); err != nil {
			return err
		}
		seenErc721[b.ClassPair.ContractAddress] = true
		seenClass[b.ClassPair.ClassId] = true
	}

	return nil
}
