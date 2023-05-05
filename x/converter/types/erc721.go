package types

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

type ERC721TokenData struct {
	URI  string
	Data string
}

func NewERC721TokenData(URI string) ERC721TokenData {
	return ERC721TokenData{
		URI: URI,
	}
}
