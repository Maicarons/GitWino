package git

import (
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Repository 封装 git.Repository
type Repository struct {
	*git.Repository
}

// GitService Git 服务实现
type GitService struct{}

// NewGitService 创建 Git 服务实例
func NewGitService() *GitService {
	return &GitService{}
}

// Clone 克隆仓库到指定目录
func (s *GitService) Clone(targetDir string, url string) (*Repository, error) {
	repo, err := git.PlainClone(targetDir, false, &git.CloneOptions{
		URL:      url,
		Progress: nil,
	})
	if err != nil {
		return nil, fmt.Errorf("克隆仓库失败：%v (请检查仓库 URL 是否正确且可访问)", err)
	}
	return &Repository{Repository: repo}, nil
}

// GetEarliestCommit 获取最早提交时间
func (s *GitService) GetEarliestCommit(repo *Repository) (time.Time, error) {
	references, err := repo.References()
	if err != nil {
		return time.Time{}, err
	}

	var earliestTime time.Time
	var found bool

	err = references.ForEach(func(ref *plumbing.Reference) error {
		if ref.Name().IsBranch() {
			tree, err := repo.Log(&git.LogOptions{
				From: ref.Hash(),
			})
			if err != nil {
				return err
			}

			err = tree.ForEach(func(c *object.Commit) error {
				if !found || c.Committer.When.Before(earliestTime) {
					earliestTime = c.Committer.When
					found = true
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return time.Time{}, err
	}

	if !found {
		return time.Time{}, fmt.Errorf("未找到提交记录")
	}

	return earliestTime, nil
}

// GetEarliestTag 获取最早标签时间
func (s *GitService) GetEarliestTag(repo *Repository) (time.Time, error) {
	tags, err := repo.Tags()
	if err != nil {
		return time.Time{}, err
	}

	var earliestTime time.Time
	var found bool

	err = tags.ForEach(func(tag *plumbing.Reference) error {
		commit, err := repo.CommitObject(tag.Hash())
		if err != nil {
			return err
		}

		if !found || commit.Committer.When.Before(earliestTime) {
			earliestTime = commit.Committer.When
			found = true
		}

		return nil
	})

	if err != nil {
		return time.Time{}, err
	}

	if !found {
		return time.Time{}, fmt.Errorf("未找到标签")
	}

	return earliestTime, nil
}

// CreateTempDir 创建临时目录
func CreateTempDir(prefix string) (string, error) {
	tempDir, err := os.MkdirTemp("", prefix)
	if err != nil {
		return "", fmt.Errorf("创建临时目录失败：%v", err)
	}
	return tempDir, nil
}
