package git

import (
	"fmt"
	"gitwino/internal/git/providers"
	"os"
	"time"
)

// GetRepoHistory 获取仓库历史数据 (主入口函数)
func GetRepoHistory(repoURL string) (*RepoHistory, error) {
	// 确保 URL 包含协议前缀
	if !hasProtocolPrefix(repoURL) {
		repoURL = "https://" + repoURL
	}

	// 确保 URL 以 .git 结尾
	if !hasGitSuffix(repoURL) {
		repoURL = repoURL + ".git"
	}

	// 创建提供者管理器
	pm := providers.NewProviderManager()

	// 解析仓库 URL
	platform, owner, repo, err := pm.ParseRepoURL(repoURL)
	if err != nil {
		return nil, fmt.Errorf("解析仓库 URL 失败：%v", err)
	}

	// 获取对应平台的提供者
	provider, err := pm.GetProvider(platform)
	if err != nil {
		return nil, err
	}

	// 类型断言为 PlatformProvider 接口
	platformProvider, ok := provider.(interface {
		GetRepoCreationTime(owner, repo string) (time.Time, error)
		GetEarliestReleaseTime(owner, repo string) (time.Time, error)
	})
	if !ok {
		return nil, fmt.Errorf("无效的提供者类型")
	}

	// 创建 Git 服务
	gitService := NewGitService()

	// 创建临时目录用于克隆仓库
	tempDir, err := CreateTempDir("gitwino-*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir)

	// 克隆仓库
	repoObj, err := gitService.Clone(tempDir, repoURL)
	if err != nil {
		return nil, err
	}

	// 获取最早提交时间
	earliestCommit, err := gitService.GetEarliestCommit(repoObj)
	if err != nil {
		return nil, fmt.Errorf("获取最早提交失败：%v", err)
	}

	// 获取最早标签时间
	earliestTag, err := gitService.GetEarliestTag(repoObj)
	if err != nil {
		// 标签获取失败不影响整体结果
		earliestTag = time.Time{}
	}

	// 获取仓库创建时间 (通过平台 API)
	repoCreationTime, err := platformProvider.GetRepoCreationTime(owner, repo)
	if err != nil {
		// API 调用失败，使用最早提交时间作为备选
		repoCreationTime = earliestCommit
	}

	// 获取最早发布时间 (通过平台 API)
	earliestReleaseTime, err := platformProvider.GetEarliestReleaseTime(owner, repo)
	if err != nil {
		// API 调用失败，返回空字符串
		earliestReleaseTime = time.Time{}
	}

	// 构建历史数据
	history := &RepoHistory{
		EarliestCommitTime:  formatTime(earliestCommit),
		RepoCreationTime:    formatTime(repoCreationTime),
		EarliestTagTime:     formatTimeOrNil(earliestTag),
		EarliestReleaseTime: formatTimeOrNil(earliestReleaseTime),
		RepoURL:             repoURL,
	}

	return history, nil
}

// hasProtocolPrefix 检查 URL 是否有协议前缀
func hasProtocolPrefix(url string) bool {
	return len(url) > 7 && (url[:7] == "http://" || url[:8] == "https://")
}

// hasGitSuffix 检查 URL 是否以.git 结尾
func hasGitSuffix(url string) bool {
	return len(url) >= 4 && url[len(url)-4:] == ".git"
}

// formatTime 格式化时间为 RFC3339 字符串
func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}

// formatTimeOrNil 格式化时间，零时间返回空字符串
func formatTimeOrNil(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}
