/*
 * Matter replacement sample PoC API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConfigsDelete - Delete configurations
func ConfigsDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// ConfigsGet - Get configurations
func ConfigsGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}