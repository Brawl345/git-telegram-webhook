package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	TelegramAPIURL = "https://api.telegram.org/bot%s/sendMessage"

	BranchPrefix = "refs/heads/"
)

type WebhookPayload struct {
	Ref        string `json:"ref"`
	Before     string `json:"before"`
	After      string `json:"after"`
	Repository struct {
		Id       int    `json:"id"`
		NodeId   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
			Name              string `json:"name"`
			Email             string `json:"email"`
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
		CreatedAt        int         `json:"created_at"`
		UpdatedAt        time.Time   `json:"updated_at"`
		PushedAt         int         `json:"pushed_at"`
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
		Stargazers               int           `json:"stargazers"`
		MasterBranch             string        `json:"master_branch"`
	} `json:"repository"`
	Pusher struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher"`
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
	Created bool        `json:"created"`
	Deleted bool        `json:"deleted"`
	Forced  bool        `json:"forced"`
	BaseRef interface{} `json:"base_ref"`
	Compare string      `json:"compare"`
	Commits []struct {
		Id        string    `json:"id"`
		TreeId    string    `json:"tree_id"`
		Distinct  bool      `json:"distinct"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		Url       string    `json:"url"`
		Author    struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"committer"`
		Added    []interface{} `json:"added"`
		Removed  []interface{} `json:"removed"`
		Modified []string      `json:"modified"`
	} `json:"commits"`
	HeadCommit struct {
		Id        string    `json:"id"`
		TreeId    string    `json:"tree_id"`
		Distinct  bool      `json:"distinct"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		Url       string    `json:"url"`
		Author    struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"author"`
		Committer struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"committer"`
		Added    []interface{} `json:"added"`
		Removed  []interface{} `json:"removed"`
		Modified []string      `json:"modified"`
	} `json:"head_commit"`
}

func (p *WebhookPayload) TotalCommits() int {
	return len(p.Commits)
}

func SendTelegramMessage(token, chatID, message string) error {
	apiURL := fmt.Sprintf(TelegramAPIURL, token)
	resp, err := http.PostForm(apiURL, url.Values{
		"chat_id":                  {chatID},
		"text":                     {message},
		"parse_mode":               {"HTML"},
		"disable_web_page_preview": {"true"},
	})
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram API error: %s", body)
	}
	return nil
}

func VerifySignature(payload []byte, secret, signature string) bool {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	expectedMAC := mac.Sum(nil)
	expectedSignature := "sha256=" + hex.EncodeToString(expectedMAC)
	return hmac.Equal([]byte(expectedSignature), []byte(signature))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	chatID := r.URL.Query().Get("chat_id")
	if chatID == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println("Missing chat_id in query parameters")
		return
	}

	telegramToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if telegramToken == "" {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Missing TELEGRAM_BOT_TOKEN")
		return
	}

	secret := os.Getenv("GITHUB_WEBHOOK_SECRET")
	if secret == "" {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Println("Missing GITHUB_WEBHOOK_SECRET")
		return
	}

	signature := r.Header.Get("X-Hub-Signature-256")
	if signature == "" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		log.Println("Missing X-Hub-Signature-256")
		return
	}

	event := r.Header.Get("X-GitHub-Event")
	if event == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println("Missing X-GitHub-Event")
		return
	}

	if event != "push" {
		http.Error(w, "Not Implemented", http.StatusNotImplemented)
		log.Println("Unsupported X-GitHub-Event:", event)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		log.Println("Error reading request body:", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(r.Body)

	if os.Getenv("DANGEROUS_SKIP_GITHUB_WEBHOOK_SECRET_CHECK") == "" {
		if !VerifySignature(body, secret, signature) {
			http.Error(w, "Forbidden", http.StatusForbidden)
			log.Println("Invalid signature")
			return
		}
	} else {
		log.Println("Skipping signature check, this is DANGEROUS!")
	}

	var p WebhookPayload
	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	var branchName = strings.TrimPrefix(p.Ref, BranchPrefix)
	var commitDesc string

	var titleLink string
	if p.TotalCommits() == 1 {
		commitDesc = "1 new commit"
		titleLink = p.Commits[0].Url
	} else {
		commitDesc = fmt.Sprintf("%d new commits", p.TotalCommits())
		titleLink = p.Compare
	}
	if titleLink == "" {
		titleLink = p.Repository.HtmlUrl + "/src/" + branchName
	}
	title := fmt.Sprintf(`[<a href="%s">%s</a>:<a href="%s">%s</a>] %s`, p.Repository.HtmlUrl, p.Repository.FullName, titleLink, branchName, commitDesc)

	var text string
	// for each commit, generate attachment text
	for i, commit := range p.Commits {
		var authorName string
		if commit.Author.Name != "" {
			authorName = " - " + commit.Author.Name
		}
		text += fmt.Sprintf(`[<a href="%s">%s</a>] %s`, commit.Url, commit.Id[:7],
			strings.TrimRight(commit.Message, "\r\n")) + authorName
		// add linebreak to each commit but the last
		if i < len(p.Commits)-1 {
			text += "\n"
		}
	}

	if err := SendTelegramMessage(telegramToken, chatID, title+"\n"+text); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("Error sending message to Telegram:", err)
		return
	}

	_, err = fmt.Fprintf(w, "Webhook received successfully")
	if err != nil {
		log.Println("Error writing response:", err)
		return // ignore error
	}
}
