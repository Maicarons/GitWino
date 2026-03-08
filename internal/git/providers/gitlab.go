package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GitLabProvider GitLab 平台提供者
type GitLabProvider struct{}

// NewGitLabProvider 创建 GitLab 提供者实例
func NewGitLabProvider() *GitLabProvider {
	return &GitLabProvider{}
}

// PlatformName 返回平台名称
func (p *GitLabProvider) PlatformName() string {
	return "gitlab"
}

// GetRepoCreationTime 获取 GitLab 仓库创建时间
func (p *GitLabProvider) GetRepoCreationTime(owner, repo string) (time.Time, error) {
	url := fmt.Sprintf("https://gitlab.com/api/v4/projects/%s%%2F%s", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return time.Time{}, fmt.Errorf("GitLab API 请求失败：%s", resp.Status)
	}

	var repoData gitlabRepo
	if err := json.NewDecoder(resp.Body).Decode(&repoData); err != nil {
		return time.Time{}, err
	}

	return time.Parse(time.RFC3339, repoData.CreatedAt)
}

// GetEarliestReleaseTime 获取 GitLab 最早发布时间
func (p *GitLabProvider) GetEarliestReleaseTime(owner, repo string) (time.Time, error) {
	url := fmt.Sprintf("https://gitlab.com/api/v4/projects/%s%%2F%s/releases", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return time.Time{}, fmt.Errorf("GitLab API 请求失败：%s", resp.Status)
	}

	var releases []gitlabRelease
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

// gitlabRepo GitLab API 响应结构
type gitlabRepo struct {
	CreatedAt string `json:"created_at"`
}

// gitlabRelease GitLab Release 结构
type gitlabRelease struct {
	CreatedAt string `json:"created_at"`
}
