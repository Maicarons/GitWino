package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitwino/internal/git"
	"log"
	"net/http"
)

// Handler 处理Git仓库历史数据查询请求
func Handler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	repoURL := r.URL.Query().Get("repo")
	if repoURL == "" {
		http.Error(w, "Missing repository URL parameter", http.StatusBadRequest)
		return
	}

	// Get repository history data
	history, err := git.GetRepoHistory(repoURL)
	if err != nil {
		log.Printf("Failed to get repository history: %v", err)
		http.Error(w, fmt.Sprintf("Failed to get repository history: %v", err), http.StatusInternalServerError)
		return
	}

	// 返回JSON响应
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(history.JSON()))
}

// GinHandler 用于 Gin 框架的处理函数
func GinHandler(c *gin.Context) {
	repoURL := c.Query("repo")
	if repoURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing repository URL parameter"})
		return
	}

	// Print debug information
	log.Printf("Received request, repository URL: %s", repoURL)

	history, err := git.GetRepoHistory(repoURL)
	if err != nil {
		// Print detailed error
		log.Printf("Failed to process request: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to get repository history: %v", err)})
		return
	}

	log.Printf("Successfully retrieved repository data: %+v", history)
	c.JSON(http.StatusOK, history)
}
