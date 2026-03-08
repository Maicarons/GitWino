package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GiteeProvider Gitee 平台提供者
type GiteeProvider struct{}

// NewGiteeProvider 创建 Gitee 提供者实例
func NewGiteeProvider() *GiteeProvider {
	return &GiteeProvider{}
}

// PlatformName 返回平台名称
func (p *GiteeProvider) PlatformName() string {
	return "gitee"
}

// GetRepoCreationTime 获取 Gitee 仓库创建时间
func (p *GiteeProvider) GetRepoCreationTime(owner, repo string) (time.Time, error) {
	url := fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return time.Time{}, fmt.Errorf("Gitee API 请求失败：%s", resp.Status)
	}

	var repoData giteeRepo
	if err := json.NewDecoder(resp.Body).Decode(&repoData); err != nil {
		return time.Time{}, err
	}

	return time.Parse(time.RFC3339, repoData.CreatedAt)
}

// GetEarliestReleaseTime 获取 Gitee 最早发布时间
func (p *GiteeProvider) GetEarliestReleaseTime(owner, repo string) (time.Time, error) {
	url := fmt.Sprintf("https://gitee.com/api/v5/repos/%s/%s/releases", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return time.Time{}, fmt.Errorf("Gitee API 请求失败：%s", resp.Status)
	}

	var releases []giteeRelease
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

// giteeRepo Gitee API 响应结构
type giteeRepo struct {
	CreatedAt string `json:"created_at"`
}

// giteeRelease Gitee Release 结构
type giteeRelease struct {
	CreatedAt string `json:"created_at"`
}
