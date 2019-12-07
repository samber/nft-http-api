package handlers

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/nftables"
	"github.com/samber/nft-http-api/nft"
)

func getRuleOrError(c *gin.Context, tableName string, chainName string, ruleId uint64) (table *nftables.Table, chain *nftables.Chain, rule *nftables.Rule, ok bool) {
	table, chain, ok = getChainOrError(c, tableName, chainName)
	if ok == false {
		return nil, nil, nil, false
	}

	rule, err := nft.GetRule(table, chain, ruleId)
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to fetch rule",
			"message": err.Error(),
		})
		return nil, nil, nil, false
	}

	if rule == nil {
		c.JSON(404, gin.H{
			"error":   "Not found",
			"message": fmt.Sprintf("Rule %d not found", ruleId),
		})
		return nil, nil, nil, false
	}

	return table, chain, rule, true
}

func FindRules(c *gin.Context) {
	tableName := c.Param("tableName")
	chainName := c.Param("chainName")

	table, chain, ok := getChainOrError(c, tableName, chainName)
	if ok == false {
		return
	}

	rules, err := nft.GetRules(table, chain)
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to fetch rules",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, rules)
}

func FindRuleByID(c *gin.Context) {
	tableName := c.Param("tableName")
	chainName := c.Param("chainName")
	_ruleId := c.Param("ruleId")

	ruleId, err := strconv.ParseUint(_ruleId, 10, 64)
	if err != nil || ruleId < 0 {
		c.JSON(422, gin.H{
			"error": "Invalid rule ID",
		})
		return
	}

	_, _, rule, ok := getRuleOrError(c, tableName, chainName, ruleId)
	if ok == false {
		return
	}

	c.JSON(200, rule)
}
