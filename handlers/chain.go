package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/nftables"
	"github.com/samber/nft-http-api/nft"
)

func getChainOrError(c *gin.Context, tableName string, chainName string) (table *nftables.Table, chain *nftables.Chain, ok bool) {
	table, ok = getTableOrError(c, tableName)
	if ok == false {
		return nil, nil, false
	}

	chain, err := nft.GetChain(table, chainName)
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to fetch chain",
			"message": err.Error(),
		})
		return nil, nil, false
	}

	if chain == nil {
		c.JSON(404, gin.H{
			"error":   "Not found",
			"message": fmt.Sprintf("Chain %s not found", chainName),
		})
		return nil, nil, false
	}

	return table, chain, true
}

func FindChains(c *gin.Context) {
	tableName := c.Param("tableName")

	table, ok := getTableOrError(c, tableName)
	if ok == false {
		return
	}

	chains, err := nft.GetChains(table)
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to fetch chains",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, chains)
}

func FindChainByName(c *gin.Context) {
	tableName := c.Param("tableName")
	chainName := c.Param("chainName")

	_, chain, ok := getChainOrError(c, tableName, chainName)
	if ok == false {
		return
	}

	c.JSON(200, chain)
}
