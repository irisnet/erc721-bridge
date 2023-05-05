package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// ERC721Data is the struct that holds the data of an ERC721 token
type ERC721Data struct {
	Name   string
	Symbol string
	URI    string
	Data   string
}

// ERC721StringResponse defines the string value from the call response
type ERC721StringResponse struct {
	Value string
}

// NewERC721Data creates a new ERC20Data instance
func NewERC721Data(name, symbol string) ERC721Data {
	return ERC721Data{
		Name:   name,
		Symbol: symbol,
	}
}

type ERC721Contract struct {
	Contract common.Address
	URI      string
	Data     string
}

func (e ERC721Contract) GetID() string {
	return e.Contract.Hex()
}

func (e ERC721Contract) GetURI() string {
	return e.URI
}

func (e ERC721Contract) GetData() string {
	return e.Data
}

type ERC721Token struct {
	Contract common.Address
	ID       *big.Int
	URI      string
	Data     string
}

func (e ERC721Token) GetClassID() string {
	return e.Contract.Hex()
}

func (e ERC721Token) GetID() string {
	return e.ID.String()
}

func (e ERC721Token) GetURI() string {
	return e.URI
}

func (e ERC721Token) GetData() string {
	return e.Data
}

type ERC721TokenData struct {
	URI  string
	Data string
}

func NewERC721TokenData(URI string) ERC721TokenData {
	return ERC721TokenData{
		URI: URI,
	}
}
