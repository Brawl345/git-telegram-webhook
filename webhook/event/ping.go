package event

import (
	"fmt"
	"strings"
	"time"
)

const (
	PingEventName = "ping"
)

type PingPayload struct {
	Zen    string `json:"zen"`
	HookId int    `json:"hook_id"`
	Hook   struct {
		Type   string   `json:"type"`
		Id     int      `json:"id"`
		Name   string   `json:"name"`
		Active bool     `json:"active"`
		Events []string `json:"events"`
		Config struct {
			ContentType string `json:"content_type"`
			InsecureSsl string `json:"insecure_ssl"`
			Secret      string `json:"secret"`
			Url         string `json:"url"`
		} `json:"config"`
		UpdatedAt     time.Time `json:"updated_at"`
		CreatedAt     time.Time `json:"created_at"`
		Url           string    `json:"url"`
		TestUrl       string    `json:"test_url"`
		PingUrl       string    `json:"ping_url"`
		DeliveriesUrl string    `json:"deliveries_url"`
		LastResponse  struct {
			Code    interface{} `json:"code"`
			Status  string      `json:"status"`
			Message interface{} `json:"message"`
		} `json:"last_response"`
	} `json:"hook"`
	Repository struct {
		Id       int    `json:"id"`
		NodeId   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Login             string `json:"login"`
			Id                int    `json:"id"`
			NodeId            string `json:"node_id"`
			AvatarUrl         string `json:"avatar_url"`
			GravatarId        string `json:"gravatar_id"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"owner"`
		HtmlUrl          string      `json:"html_url"`
		Description      string      `json:"description"`
		Fork             bool        `json:"fork"`
		Url              string      `json:"url"`
		ForksUrl         string      `json:"forks_url"`
		KeysUrl          string      `json:"keys_url"`
		CollaboratorsUrl string      `json:"collaborators_url"`
		TeamsUrl         string      `json:"teams_url"`
		HooksUrl         string      `json:"hooks_url"`
		IssueEventsUrl   string      `json:"issue_events_url"`
		EventsUrl        string      `json:"events_url"`
		AssigneesUrl     string      `json:"assignees_url"`
		BranchesUrl      string      `json:"branches_url"`
		TagsUrl          string      `json:"tags_url"`
		BlobsUrl         string      `json:"blobs_url"`
		GitTagsUrl       string      `json:"git_tags_url"`
		GitRefsUrl       string      `json:"git_refs_url"`
		TreesUrl         string      `json:"trees_url"`
		StatusesUrl      string      `json:"statuses_url"`
		LanguagesUrl     string      `json:"languages_url"`
		StargazersUrl    string      `json:"stargazers_url"`
		ContributorsUrl  string      `json:"contributors_url"`
		SubscribersUrl   string      `json:"subscribers_url"`
		SubscriptionUrl  string      `json:"subscription_url"`
		CommitsUrl       string      `json:"commits_url"`
		GitCommitsUrl    string      `json:"git_commits_url"`
		CommentsUrl      string      `json:"comments_url"`
		IssueCommentUrl  string      `json:"issue_comment_url"`
		ContentsUrl      string      `json:"contents_url"`
		CompareUrl       string      `json:"compare_url"`
		MergesUrl        string      `json:"merges_url"`
		ArchiveUrl       string      `json:"archive_url"`
		DownloadsUrl     string      `json:"downloads_url"`
		IssuesUrl        string      `json:"issues_url"`
		PullsUrl         string      `json:"pulls_url"`
		MilestonesUrl    string      `json:"milestones_url"`
		NotificationsUrl string      `json:"notifications_url"`
		LabelsUrl        string      `json:"labels_url"`
		ReleasesUrl      string      `json:"releases_url"`
		DeploymentsUrl   string      `json:"deployments_url"`
		CreatedAt        time.Time   `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
		PushedAt         time.Time   `json:"pushed_at"`
		GitUrl           string      `json:"git_url"`
		SshUrl           string      `json:"ssh_url"`
		CloneUrl         string      `json:"clone_url"`
		SvnUrl           string      `json:"svn_url"`
		Homepage         string      `json:"homepage"`
		Size             int         `json:"size"`
		StargazersCount  int         `json:"stargazers_count"`
		WatchersCount    int         `json:"watchers_count"`
		Language         string      `json:"language"`
		HasIssues        bool        `json:"has_issues"`
		HasProjects      bool        `json:"has_projects"`
		HasDownloads     bool        `json:"has_downloads"`
		HasWiki          bool        `json:"has_wiki"`
		HasPages         bool        `json:"has_pages"`
		HasDiscussions   bool        `json:"has_discussions"`
		ForksCount       int         `json:"forks_count"`
		MirrorUrl        interface{} `json:"mirror_url"`
		Archived         bool        `json:"archived"`
		Disabled         bool        `json:"disabled"`
		OpenIssuesCount  int         `json:"open_issues_count"`
		License          struct {
			Key    string `json:"key"`
			Name   string `json:"name"`
			SpdxId string `json:"spdx_id"`
			Url    string `json:"url"`
			NodeId string `json:"node_id"`
		} `json:"license"`
		AllowForking             bool          `json:"allow_forking"`
		IsTemplate               bool          `json:"is_template"`
		WebCommitSignoffRequired bool          `json:"web_commit_signoff_required"`
		Topics                   []interface{} `json:"topics"`
		Visibility               string        `json:"visibility"`
		Forks                    int           `json:"forks"`
		OpenIssues               int           `json:"open_issues"`
		Watchers                 int           `json:"watchers"`
		DefaultBranch            string        `json:"default_branch"`
	} `json:"repository"`
	Sender struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"sender"`
}

func (p *PingPayload) Text() string {
	var sb strings.Builder
	sb.WriteString("<b>PONG! Webhook is working!</b>\n")

	sb.WriteString(
		fmt.Sprintf("Received ping event from <a href=\"%s\">%s</a>.\n",
			p.Repository.HtmlUrl,
			p.Repository.FullName,
		),
	)

	sb.WriteString(fmt.Sprintf("Now for some Zen: <code>%s</code>", p.Zen))

	return sb.String()
}
