package nft

import (
	"github.com/google/nftables"
)

func GetRules(table *nftables.Table, chain *nftables.Chain) ([]*nftables.Rule, error) {
	return conn.GetRule(table, chain)
}

func GetRule(table *nftables.Table, chain *nftables.Chain, id uint64) (*nftables.Rule, error) {
	rules, err := GetRules(table, chain)
	if err != nil {
		return nil, err
	}

	for _, rule := range rules {
		if rule.Position == id {
			return rule, nil
		}
	}

	return nil, nil
}
