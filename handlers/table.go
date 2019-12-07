package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/nftables"
	"github.com/samber/nft-http-api/nft"
)

func getTableOrError(c *gin.Context, tableName string) (table *nftables.Table, ok bool) {
	table, err := nft.GetTable(tableName)
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to fetch tables",
			"message": err.Error(),
		})
		return nil, false
	}

	if table == nil {
		c.JSON(404, gin.H{
			"error":   "Not found",
			"message": fmt.Sprintf("Table %s not found", tableName),
		})
		return nil, false
	}

	return table, true
}

func FindTables(c *gin.Context) {
	tables, err := nft.GetTables()
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Failed to fetch tables",
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, tables)
}

func FindTableByName(c *gin.Context) {
	tableName := c.Param("tableName")

	table, ok := getTableOrError(c, tableName)
	if ok == false {
		return
	}

	c.JSON(200, table)
}
