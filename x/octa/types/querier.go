package types

import "strings"

// Query endpoints supported by the octa querier
const QueryListTransfers = "list"
const QueryGetTransferOfFunds = "get"

type QueryResTransfers []string

// implement fmt.Stringer
func (n QueryResTransfers) String() string {
	return strings.Join(n[:], "\n")
}
