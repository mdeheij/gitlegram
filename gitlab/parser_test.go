package gitlab

import (
	"reflect"
	"testing"
)

var gitlabPushRequest = `{
  "object_kind": "push",
  "before": "95790bf891e76fee5e1747ab589903a6a1f80f22",
  "after": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
  "ref": "refs/heads/master",
  "checkout_sha": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
  "user_id": 4,
  "user_name": "John Smith",
  "user_username": "jsmith",
  "user_email": "john@example.com",
  "user_avatar": "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80",
  "project_id": 15,
  "project":{
    "name":"Diaspora",
    "description":"",
    "web_url":"http://example.com/mike/diaspora",
    "avatar_url":null,
    "git_ssh_url":"git@example.com:mike/diaspora.git",
    "git_http_url":"http://example.com/mike/diaspora.git",
    "namespace":"Mike",
    "visibility_level":0,
    "path_with_namespace":"mike/diaspora",
    "default_branch":"master",
    "homepage":"http://example.com/mike/diaspora",
    "url":"git@example.com:mike/diaspora.git",
    "ssh_url":"git@example.com:mike/diaspora.git",
    "http_url":"http://example.com/mike/diaspora.git"
  },
  "repository":{
    "name": "Diaspora",
    "url": "git@example.com:mike/diaspora.git",
    "description": "",
    "homepage": "http://example.com/mike/diaspora",
    "git_http_url":"http://example.com/mike/diaspora.git",
    "git_ssh_url":"git@example.com:mike/diaspora.git",
    "visibility_level":0
  },
  "commits": [
    {
      "id": "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
      "message": "Update Catalan translation to e38cb41.",
      "timestamp": "2011-12-12T14:27:31+02:00",
      "url": "http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
      "author": {
        "name": "Jordi Mallach",
        "email": "jordi@softcatala.org"
      },
      "added": ["CHANGELOG"],
      "modified": ["app/controller/application.rb"],
      "removed": []
    },
    {
      "id": "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
      "message": "fixed readme",
      "timestamp": "2012-01-03T23:36:29+02:00",
      "url": "http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
      "author": {
        "name": "GitLab dev user",
        "email": "gitlabdev@dv6700.(none)"
      },
      "added": ["CHANGELOG"],
      "modified": ["app/controller/application.rb"],
      "removed": []
    }
  ],
  "total_commits_count": 4
}`

func GitlabSamplePushRequest() Request {
	return Request{
		ObjectKind:   "push",
		Before:       "95790bf891e76fee5e1747ab589903a6a1f80f22",
		After:        "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
		Ref:          "refs/heads/master",
		CheckoutSha:  "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
		UserID:       4,
		UserName:     "John Smith",
		UserUsername: "jsmith",
		UserEmail:    "john@example.com",
		UserAvatar:   "https://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=8://s.gravatar.com/avatar/d4c74594d841139328695756648b6bd6?s=80",
		ProjectID:    15,
		Project: Project{
			Name:              "Diaspora",
			Description:       "",
			WebURL:            "http://example.com/mike/diaspora",
			GitSSHURL:         "git@example.com:mike/diaspora.git",
			GitHTTPURL:        "http://example.com/mike/diaspora.git",
			Namespace:         "Mike",
			VisibilityLevel:   0,
			PathWithNamespace: "mike/diaspora",
			DefaultBranch:     "master",
			Homepage:          "http://example.com/mike/diaspora",
			URL:               "git@example.com:mike/diaspora.git",
			SSHURL:            "git@example.com:mike/diaspora.git",
			HTTPURL:           "http://example.com/mike/diaspora.git",
		},
		Repository: Repository{
			Name:            "Diaspora",
			URL:             "git@example.com:mike/diaspora.git",
			Description:     "",
			Homepage:        "http://example.com/mike/diaspora",
			GitHTTPURL:      "http://example.com/mike/diaspora.git",
			GitSSHURL:       "git@example.com:mike/diaspora.git",
			VisibilityLevel: 0,
		},
		Commits: []Commit{
			Commit{
				ID:        "b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
				Message:   "Update Catalan translation to e38cb41.",
				Timestamp: "2011-12-12T14:27:31+02:00",
				URL:       "http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327",
				Author: Author{
					Name:  "Jordi Mallach",
					Email: "jordi@softcatala.org",
				},
				Added:    []string{"CHANGELOG"},
				Modified: []string{"app/controller/application.rb"},
				Removed:  []string{},
			},
			Commit{
				ID:        "da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				Message:   "fixed readme",
				Timestamp: "2012-01-03T23:36:29+02:00",
				URL:       "http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7",
				Author: Author{
					Name:  "GitLab dev user",
					Email: "gitlabdev@dv6700.(none)",
				},
				Added:    []string{"CHANGELOG"},
				Modified: []string{"app/controller/application.rb"},
				Removed:  []string{},
			},
		},
		TotalCommitsCount: 4,
	}

}
func TestParse(t *testing.T) {
	type args struct {
		jsonBody string
	}
	tests := []struct {
		name  string
		args  args
		wantR Request
	}{
		{"push", args{jsonBody: gitlabPushRequest}, GitlabSamplePushRequest()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := Parse(tt.args.jsonBody); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("Parse() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
