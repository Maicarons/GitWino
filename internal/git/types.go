package git

import (
	"fmt"
	"time"
)

// RepoHistory 仓库历史数据结构
type RepoHistory struct {
	EarliestCommitTime  string `json:"earliest_commit_time"`  // 最早提交时间
	RepoCreationTime    string `json:"repo_creation_time"`    // 仓库建立时间
	EarliestTagTime     string `json:"earliest_tag_time"`     // 最早标签创建时间
	EarliestReleaseTime string `json:"earliest_release_time"` // 最早发布时间
	RepoURL             string `json:"repo_url"`              // 仓库 URL
}

// GitPlatform 表示 Git 平台类型
type GitPlatform string

const (
	PlatformGitHub  GitPlatform = "github"
	PlatformGitee   GitPlatform = "gitee"
	PlatformGitLab  GitPlatform = "gitlab"
	PlatformUnknown GitPlatform = "unknown"
)

// RepoInfo 仓库基本信息
type RepoInfo struct {
	Platform GitPlatform
	Owner    string
	Repo     string
	RawURL   string
}

// PlatformProvider Git 平台提供者接口
type PlatformProvider interface {
	// GetRepoCreationTime 获取仓库创建时间
	GetRepoCreationTime(owner, repo string) (time.Time, error)
	// GetEarliestReleaseTime 获取最早发布时间
	GetEarliestReleaseTime(owner, repo string) (time.Time, error)
	// PlatformName 返回平台名称
	PlatformName() string
}

// RepositoryService 仓库服务接口
type RepositoryService interface {
	// Clone 克隆仓库到指定目录
	Clone(targetDir string, url string) (*Repository, error)
	// GetEarliestCommit 获取最早提交时间
	GetEarliestCommit(repo *Repository) (time.Time, error)
	// GetEarliestTag 获取最早标签时间
	GetEarliestTag(repo *Repository) (time.Time, error)
}

// JSON 转换为 JSON 字符串
func (h *RepoHistory) JSON() string {
	return fmt.Sprintf(`{
	"earliest_commit_time": "%s",
	"repo_creation_time": "%s",
	"earliest_tag_time": "%s",
	"earliest_release_time": "%s",
	"repo_url": "%s"
}`, h.EarliestCommitTime, h.RepoCreationTime, h.EarliestTagTime, h.EarliestReleaseTime, h.RepoURL)
}
