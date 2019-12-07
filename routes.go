package main

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/nft-http-api/handlers"
)

type Route struct {
	Method  string
	Path    string
	Handler func(*gin.Context)
}

var routes []Route

func init() {
	routes = append(routes, Route{"GET", "/health", handlers.Health})

	routes = append(routes, Route{"GET", "/table", handlers.FindTables})
	routes = append(routes, Route{"GET", "/table/:tableName", handlers.FindTableByName})
	routes = append(routes, Route{"GET", "/table/:tableName/chain", handlers.FindChains})
	routes = append(routes, Route{"GET", "/table/:tableName/chain/:chainName", handlers.FindChainByName})
	routes = append(routes, Route{"GET", "/table/:tableName/chain/:chainName/rule", handlers.FindRules})
	routes = append(routes, Route{"GET", "/table/:tableName/chain/:chainName/rule/:ruleId", handlers.FindRuleByID})
}

func registerRoutes(router *gin.Engine) {
	for _, r := range routes {
		router.Handle(r.Method, r.Path, r.Handler)
	}
}
