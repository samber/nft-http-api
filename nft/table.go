package nft

import (
	"github.com/google/nftables"
)

func GetTables() ([]*nftables.Table, error) {
	return conn.ListTables()
}

func GetTable(name string) (*nftables.Table, error) {
	tables, err := GetTables()
	if err != nil {
		return nil, err
	}

	for _, t := range tables {
		if t.Name == name {
			return t, nil
		}
	}

	return nil, nil
}
