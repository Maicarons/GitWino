package providers

import (
	"fmt"
	"strings"
)

// ProviderManager 提供者管理器
type ProviderManager struct {
	providers map[string]interface{}
}

// NewProviderManager 创建提供者管理器
func NewProviderManager() *ProviderManager {
	pm := &ProviderManager{
		providers: make(map[string]interface{}),
	}

	// 注册所有支持的提供者
	pm.Register("github", NewGitHubProvider())
	pm.Register("gitee", NewGiteeProvider())
	pm.Register("gitlab", NewGitLabProvider())

	return pm
}

// Register 注册平台提供者
func (pm *ProviderManager) Register(platform string, provider interface{}) {
	pm.providers[platform] = provider
}

// GetProvider 获取指定平台的提供者
func (pm *ProviderManager) GetProvider(platform string) (interface{}, error) {
	provider, exists := pm.providers[platform]
	if !exists {
		return nil, fmt.Errorf("不支持的 Git 平台：%s", platform)
	}
	return provider, nil
}

// ParseRepoURL 解析仓库 URL
func (pm *ProviderManager) ParseRepoURL(repoURL string) (string, string, string, error) {
	var platform string
	var owner string
	var repo string

	if strings.Contains(repoURL, "github.com") {
		platform = "github"
		parts := strings.Split(repoURL, "/")
		if len(parts) >= 5 {
			owner = parts[3]
			repo = strings.TrimSuffix(parts[4], ".git")
		}
	} else if strings.Contains(repoURL, "gitee.com") {
		platform = "gitee"
		parts := strings.Split(repoURL, "/")
		if len(parts) >= 5 {
			owner = parts[3]
			repo = strings.TrimSuffix(parts[4], ".git")
		}
	} else if strings.Contains(repoURL, "gitlab.com") {
		platform = "gitlab"
		parts := strings.Split(repoURL, "/")
		if len(parts) >= 5 {
			owner = parts[3]
			repo = strings.TrimSuffix(parts[4], ".git")
		}
	} else {
		return "", "", "", fmt.Errorf("无法识别的 Git 平台")
	}

	if owner == "" || repo == "" {
		return "", "", "", fmt.Errorf("无法解析仓库 URL")
	}

	return platform, owner, repo, nil
}
