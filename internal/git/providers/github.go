package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GitHubProvider GitHub 平台提供者
type GitHubProvider struct{}

// NewGitHubProvider 创建 GitHub 提供者实例
func NewGitHubProvider() *GitHubProvider {
	return &GitHubProvider{}
}

// PlatformName 返回平台名称
func (p *GitHubProvider) PlatformName() string {
	return "github"
}

// GetRepoCreationTime 获取 GitHub 仓库创建时间
func (p *GitHubProvider) GetRepoCreationTime(owner, repo string) (time.Time, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return time.Time{}, fmt.Errorf("GitHub API 请求失败：%s", resp.Status)
	}

	var repoData githubRepo
	if err := json.NewDecoder(resp.Body).Decode(&repoData); err != nil {
		return time.Time{}, err
	}

	return time.Parse(time.RFC3339, repoData.CreatedAt)
}

// GetEarliestReleaseTime 获取 GitHub 最早发布时间
func (p *GitHubProvider) GetEarliestReleaseTime(owner, repo string) (time.Time, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return time.Time{}, fmt.Errorf("GitHub API 请求失败：%s", resp.Status)
	}

	var releases []githubRelease
	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return time.Time{}, err
	}

	if len(releases) == 0 {
		return time.Time{}, fmt.Errorf("没有找到发布记录")
	}

	earliestTime := parseTime(releases[0].CreatedAt)
	for _, release := range releases {
		t := parseTime(release.CreatedAt)
		if t.Before(earliestTime) {
			earliestTime = t
		}
	}

	return earliestTime, nil
}

// githubRepo GitHub API 响应结构
type githubRepo struct {
	CreatedAt string `json:"created_at"`
}

// githubRelease GitHub Release 结构
type githubRelease struct {
	CreatedAt string `json:"created_at"`
}
