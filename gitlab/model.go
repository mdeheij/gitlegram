package gitlab

import (
	"errors"
	"time"

	"github.com/mdeheij/gitlegram/interfaces"
)

type Project struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	WebURL            string `json:"web_url"`
	AvatarURL         string `json:"avatar_url"`
	GitSSHURL         string `json:"git_ssh_url"`
	GitHTTPURL        string `json:"git_http_url"`
	Namespace         string `json:"namespace"`
	VisibilityLevel   int    `json:"visibility_level"`
	PathWithNamespace string `json:"path_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
	Homepage          string `json:"homepage"`
	URL               string `json:"url"`
	SSHURL            string `json:"ssh_url"`
	HTTPURL           string `json:"http_url"`
}
type Repository struct {
	Name            string `json:"name"`
	URL             string `json:"url"`
	Description     string `json:"description"`
	Homepage        string `json:"homepage"`
	GitHTTPURL      string `json:"git_http_url"`
	GitSSHURL       string `json:"git_ssh_url"`
	VisibilityLevel int    `json:"visibility_level"`
}

type Request struct {
	ObjectKind        string     `json:"object_kind"`
	Before            string     `json:"before"`
	After             string     `json:"after"`
	Ref               string     `json:"ref"`
	CheckoutSha       string     `json:"checkout_sha"`
	UserID            int        `json:"user_id"`
	UserName          string     `json:"user_name"`
	UserUsername      string     `json:"user_username"`
	UserEmail         string     `json:"user_email"`
	UserAvatar        string     `json:"user_avatar"`
	ProjectID         int        `json:"project_id"`
	Project           Project    `json:"project"`
	Repository        Repository `json:"repository"`
	Commits           []Commit   `json:"commits"`
	TotalCommitsCount int        `json:"total_commits_count"`
}

//currently supported request types (object kinds)
var supportedObjectKinds = []string{"push", "tag_push"}

//IsValid returns true if it seems like a valid request
func (r *Request) IsValid() bool {
	for _, b := range supportedObjectKinds {
		if b == r.ObjectKind {
			return true
		}
	}
	return false
}

func (r Request) GetRepository() interfaces.Repository {
	return r.Repository
}
func (r Request) GetUser() (interfaces.User, error) {
	if r.ObjectKind == "push" {
		return User{
			Username:  r.UserUsername,
			Name:      r.UserName,
			AvatarURL: r.UserAvatar,
		}, nil
	}

	return User{}, errors.New("Could not fetch user data from request")
}
func (r Repository) GetName() string {
	return r.Name
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Commit struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	URL       string    `json:"url"`
	Author    Author    `json:"author"`
	Added     []string  `json:"added"`
	Modified  []string  `json:"modified"`
	Removed   []string  `json:"removed"`
}

type IssueHook struct {
	ObjectKind       string           `json:"object_kind"`
	User             User             `json:"user"`
	Project          Project          `json:"project"`
	Repository       Repository       `json:"repository"`
	ObjectAttributes ObjectAttributes `json:"object_attributes"`
	Assignees        []Assignee       `json:"assignees"`
	Assignee         Assignee         `json:"assignee"`
	Labels           []Label          `json:"labels"`
}
type ObjectAttributes struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	AssigneeIds []int       `json:"assignee_ids"`
	AssigneeID  int         `json:"assignee_id"`
	AuthorID    int         `json:"author_id"`
	ProjectID   int         `json:"project_id"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Position    int         `json:"position"`
	BranchName  interface{} `json:"branch_name"`
	Description string      `json:"description"`
	MilestoneID interface{} `json:"milestone_id"`
	State       string      `json:"state"`
	Iid         int         `json:"iid"`
	URL         string      `json:"url"`
	Action      string      `json:"action"`
}

type User struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

func (u User) GetUsername() string {
	return u.Username
}
func (u User) GetName() string {
	return u.Name
}
func (u User) GetAvatarURL() string {
	return u.AvatarURL
}

type Assignee struct {
	User
}

type Label struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Color       string    `json:"color"`
	ProjectID   int       `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Template    bool      `json:"template"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	GroupID     int       `json:"group_id"`
}
