package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const assetsDir = "_assets"

var httpClient = http.Client{
	Timeout: 2 * time.Second,
}

var menuMapping = map[string]topLevelMenu{
	"dcs/storage":                        {"Decentralized Cloud Storage", 10},
	"dcs/downloads":                      {"Downloads", 20},
	"dcs/getting-started":                {"Getting Started", 30},
	"dcs/api-reference":                  {"SDK & Reference", 40},
	"dcs/how-tos":                        {"How To's", 50},
	"dcs/solution-architectures":         {"Solution Architectures", 60},
	"dcs/concepts":                       {"Concepts", 70},
	"dcs/support":                        {"Support", 80},
	"dcs/billing-payment-and-accounts-1": {"Billing, Payment & Accounts", 90},

	"node/before-you-begin":       {"Before You Begin", 10},
	"node/dependencies":           {"Dependencies", 20},
	"node/setup":                  {"Setup", 30},
	"node/sno-applications":       {"SNO Applications", 40},
	"node/resources":              {"Resources", 50},
	"node/solution-architectures": {"Solution Architectures", 60},
}

var urlToTitle = map[string]string{
	"https://docs.microsoft.com/en-us/windows-server/administration/openssh/openssh_install_firstuse":                                           "Get started with OpenSSH",
	"https://docs.microsoft.com/en-us/windows-server/administration/openssh/openssh_install_firstuse#installing-openssh-with-powershell":        "Install OpenSSH using Windows Settings",
	"https://docs.microsoft.com/en-us/windows-server/administration/openssh/openssh_server_configuration#windows-configurations-in-sshd_config": "Windows Configurations in sshd_config",
	"https://docs.microsoft.com/en-us/windows/wsl/install-win10":                                                                                "Install WSL",
	"https://osxdaily.com/2016/08/16/enable-ssh-mac-command-line/":                                                                              "How to Enable SSH on a Mac from the Command Line",
	"https://osxdaily.com/2016/08/16/enable-ssh-mac-command-line":                                                                               "How to Enable SSH on a Mac from the Command Line",
	"https://superuser.com/questions/364304/how-do-i-configure-ssh-on-os-x":                                                                     "How do I configure SSH on OS X?",
}

func run(exe string, args ...string) {
	cmd := exec.Command(exe, args...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	_ = cmd.Run()
}

func mustRun(exe string, args ...string) {
	cmd := exec.Command(exe, args...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var skipDownload = flag.Bool("skip-download", false, "don't try to cache new images")
var skipWorktree = flag.Bool("skip-worktree", false, "skips worktree logic")

func main() {
	flag.Parse()

	if !*skipWorktree {
		fmt.Println("# Processing worktree")
		// reset gitbook
		run("git", "worktree", "remove", "gitbook/dcs")
		mustRun("git", "worktree", "add", "gitbook/dcs", "origin/gitbook-sync")
		run("git", "worktree", "remove", "gitbook/node")
		mustRun("git", "worktree", "add", "gitbook/node", "origin/gitbook-node-sync")
	}

	// cleanup previous run
	os.RemoveAll("content")
	os.Mkdir("content", 0755)

	// start conversion
	failures := []error{}
	warnings := []error{}
	convs := []Convert{
		{
			SourceDir:  "gitbook/dcs/docs",
			ContentDir: "content",
			ExtraDir:   "content-extra",
			TargetDir:  "dcs",
		},
		{
			SourceDir:  "gitbook/node",
			ContentDir: "content",
			ExtraDir:   "content-extra",
			TargetDir:  "node",
		},
	}

	for _, conv := range convs {
		fmt.Println("# Converting", conv.SourceDir)
		conv.Run()
		failures = append(failures, conv.Failures...)
		warnings = append(warnings, conv.Warnings...)
	}
	if len(warnings) > 0 {
		fmt.Println("# WARNINGS")
		for _, warn := range warnings {
			fmt.Println(warn)
		}
	}
	if len(failures) > 0 {
		fmt.Println("# ERRORS")
		for _, fail := range failures {
			fmt.Println(fail)
		}
		os.Exit(1)
	}
}

type Convert struct {
	SourceDir  string
	ContentDir string
	ExtraDir   string
	TargetDir  string

	OrderByFolder map[string][]SummaryItem
	SummaryByItem map[string]SummaryItem

	Failures []error
	Warnings []error
}

type SummaryItem struct {
	Title       string
	ContentPath string
}

func (conv *Convert) Run() {
	conv.CreateOrder()
	conv.Files()
	conv.AddSectionIndices()
	conv.CopyExtra()
}

func (conv *Convert) Files() {
	err := filepath.WalkDir(conv.SourceDir,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if filepath.Base(path) == ".git" || filepath.Base(path) == ".github" {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
			if d.IsDir() {
				return nil
			}

			if err := conv.Convert(filepath.ToSlash(path)); err != nil {
				conv.Failures = append(conv.Failures,
					fmt.Errorf("failed to convert %s: %w", path, err))
			}
			return nil
		})
	if err != nil {
		conv.Failures = append(conv.Failures, err)
	}
}

func (conv *Convert) CopyExtra() {
	sourceDir := path.Join(conv.ExtraDir, conv.TargetDir)
	err := filepath.WalkDir(sourceDir,
		func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if filepath.Base(p) == ".git" {
				return filepath.SkipDir
			}
			if d.IsDir() {
				return nil
			}

			fullPath := filepath.ToSlash(p)
			fmt.Println("  - ", fullPath)
			contentPath := trimPrefix(fullPath, sourceDir)

			targetPath := path.Join(conv.ContentDir, conv.TargetDir, contentPath)
			err = copyFile(fullPath, targetPath)
			if err != nil {
				conv.Failures = append(conv.Failures, fmt.Errorf("failed to copy %q: %w", fullPath, err))
			}
			return nil
		})
	if err != nil {
		conv.Failures = append(conv.Failures, err)
	}
}

func (conv *Convert) Convert(fullPath string) error {
	fmt.Println("  - ", fullPath)
	contentPath := trimPrefix(fullPath, conv.SourceDir)

	switch path.Ext(contentPath) {
	case ".png", ".jpg", ".jpeg", ".svg", ".gif":
		if !strings.HasPrefix(contentPath, ".gitbook/assets") {
			return fmt.Errorf("don't know where to move")
		}
		noPrefix := trimPrefix(contentPath, ".gitbook/assets")
		targetPath := path.Join(conv.ContentDir, conv.TargetDir, assetsDir, noPrefix)
		err := copyFile(fullPath, targetPath)
		if err != nil {
			return fmt.Errorf("failed to copy: %w", err)
		}
		return nil

	case ".md":
	default:
		switch contentPath {
		case ".gitbook/assets/0", ".gitbook/assets/1", ".gitbook/assets/2", ".gitbook/assets/3":
			noPrefix := contentPath[len(contentPath)-1:] + "-fix.png"
			targetPath := path.Join(conv.ContentDir, conv.TargetDir, assetsDir, noPrefix)
			err := copyFile(fullPath, targetPath)
			if err != nil {
				return fmt.Errorf("failed to copy: %w", err)
			}
			return nil
		}

		return fmt.Errorf("don't know how to handle %q", contentPath)
	}

	// markdown handling
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return fmt.Errorf("failed to load: %w", err)
	}

	page := ParsePage(contentPath, string(data))
	conv.AddWeight(&page)
	conv.AddAlias(&page)
	conv.LiftTitle(&page)
	conv.FixJoiners(&page)
	conv.ReplaceContentRefs(&page)
	conv.ReplaceTags(&page)
	conv.FixTrailingSpace(&page)
	conv.FixLinksToReadme(&page)
	conv.FixImageLinks(&page)
	conv.FixRegularLinks(&page)
	conv.ReplaceMath(&page)
	conv.ReplaceStarryNight(&page)
	conv.ReplaceUnderlines(&page)
	conv.ReplaceBoldPeriod(&page)
	conv.ReplaceZeroWidth(&page)

	// patch a filename containing a .
	if strings.Contains(contentPath, "config.yaml.md") {
		contentPath = strings.ReplaceAll(contentPath, "config.yaml.md", "config-yaml.md")
	}

	targetPath := path.Join(conv.ContentDir, conv.TargetDir, contentPath)
	if strings.EqualFold(path.Base(targetPath), "README.md") {
		targetPath = targetPath[:len(targetPath)-len("README.md")] + "_index.md"
	}

	return page.WriteToFile(targetPath)
}

func copyFile(from, to string) error {
	data, err := os.ReadFile(from)
	if err != nil {
		return err
	}
	return writeFile(to, data)
}

func writeFile(to string, data []byte) error {
	if err := ensureFileDir(to); err != nil {
		return err
	}
	return os.WriteFile(to, data, 0644)
}

func trimPrefix(path, prefix string) string {
	return strings.TrimLeft(strings.TrimPrefix(path, prefix), "\\/")
}

func ensureFileDir(path string) error {
	return os.MkdirAll(filepath.Dir(path), 0755)
}

type Page struct {
	ContentPath string
	FrontMatter string
	Content     string
}

func ParsePage(contentPath, content string) Page {
	tokens := strings.SplitN(content, "---\n", 3)
	if len(tokens) == 1 {
		return Page{
			ContentPath: contentPath,
			Content:     content,
		}
	}

	return Page{
		ContentPath: contentPath,
		FrontMatter: tokens[1],
		Content:     tokens[2],
	}
}

func (page *Page) WriteToFile(path string) error {
	return writeFile(path, []byte(strings.Join([]string{
		"",
		page.FrontMatter,
		page.Content,
	}, "---\n")))
}

func (conv *Convert) CreateOrder() {
	conv.OrderByFolder = map[string][]SummaryItem{}
	conv.SummaryByItem = map[string]SummaryItem{}

	data, err := os.ReadFile(path.Join(conv.SourceDir, "SUMMARY.md"))
	if err != nil {
		conv.Failures = append(conv.Failures, fmt.Errorf("failed to read summary: %w", err))
		return
	}

	rx := mustCompile(`\[([^\]]*)\]\(([^)]*)\)`)
	for _, match := range rx.FindAllStringSubmatch(string(data), -1) {
		title := match[1]
		contentPath := match[2]

		dir := path.Dir(contentPath)
		if filepath.Base(contentPath) == "README.md" {
			dir = path.Dir(dir)
		}

		sum := SummaryItem{
			Title:       title,
			ContentPath: contentPath,
		}
		conv.SummaryByItem[contentPath] = sum
		conv.OrderByFolder[dir] = append(conv.OrderByFolder[dir], sum)
	}
}

func (conv *Convert) AddWeight(page *Page) {
	if page.ContentPath == "SUMMARY.md" {
		page.FrontMatter = "draft: true\n" + page.FrontMatter
		return
	}

	dir := path.Dir(page.ContentPath)
	if filepath.Base(page.ContentPath) == "README.md" {
		dir = path.Dir(dir)
	}

	for i, item := range conv.OrderByFolder[dir] {
		if item.ContentPath == page.ContentPath {
			page.FrontMatter = "weight: " + strconv.Itoa(-100+i*10) + "\n" + page.FrontMatter
			return
		}
	}
	conv.Failures = append(conv.Failures, fmt.Errorf("order missing for %s", page.ContentPath))
}

// LiftTitle moves `# XYZ` to front matter `title: `
func (conv *Convert) LiftTitle(page *Page) {
	if match(`title\s*:`, page.FrontMatter) {
		return
	}

	if page.ContentPath == "_index.md" || page.ContentPath == "README.md" {
		switch conv.TargetDir {
		case "dcs":
			page.FrontMatter = "title: \"DCS\"\n" + page.FrontMatter
		case "node":
			page.FrontMatter = "title: \"Node Operator\"\n" + page.FrontMatter
		default:
			panic(conv.TargetDir)
		}
		return
	}

	const rxTitle = `#\s*([^\n]+)\n`

	var title string
	ok := match(rxTitle, page.Content, nil, &title)
	if !ok {
		return
	}

	summary, ok := conv.SummaryByItem[page.ContentPath]
	if ok && summary.Title != title {
		conv.Warnings = append(conv.Warnings, fmt.Errorf("%q title differs in summary", page.ContentPath))
		page.FrontMatter = "title: \"" + summary.Title + "\"\n" + page.FrontMatter
		page.Content = mustReplaceFirst("\n?"+rxTitle, page.Content, "\n# "+summary.Title+"\n")
		return
	}

	page.FrontMatter = "title: \"" + title + "\"\n" + page.FrontMatter
	// hugo-book does not add the title automatically
	// page.Content = mustReplaceFirst("\n?"+rxTitle, page.Content, "")
}

func (conv *Convert) AddAlias(page *Page) {
	if conv.TargetDir == "node" && page.ContentPath == "resources/faq/where-can-i-find-a-config.yaml.md" {
		page.FrontMatter = "aliases: [\"/node/resources/faq/where-can-i-find-a-config.yaml\"]\n" + page.FrontMatter
		return
	}

	if conv.TargetDir != "dcs" {
		return
	}
	if page.ContentPath == "_index.md" || page.ContentPath == "README.md" {
		return
	}

	alias := ""
	if path.Base(page.ContentPath) == "_index.md" || path.Base(page.ContentPath) == "README.md" {
		alias = "/" + path.Dir(page.ContentPath)
	} else {
		alias = "/" + strings.TrimSuffix(page.ContentPath, ".md")
	}

	page.FrontMatter = "aliases: [\"" + alias + "\"]\n" + page.FrontMatter
}

func (conv *Convert) FixJoiners(page *Page) {
	page.Content = strings.ReplaceAll(page.Content, `\
‌\
`, "\n\n")
	page.Content = strings.ReplaceAll(page.Content, "‌", "")
}

// ReplaceContentRefs implements replacing multi-line content-ref tags:
//
//   {% content-ref url="before-you-begin/auth-token.md" %}
//   [auth-token.md](before-you-begin/auth-token.md)
//   {% endcontent-ref %}
func (conv *Convert) ReplaceContentRefs(page *Page) {
	rxContentRef := mustCompile(
		`{%\s+content-ref url="([^"]+)"\s+%}\n` +
			`\[([^\]]+)\]\(([^)]+)\)\n` +
			`{%\s+endcontent-ref\s+%}`)

	page.Content = rxContentRef.ReplaceAllStringFunc(page.Content, func(match string) string {
		matches := rxContentRef.FindStringSubmatch(match)
		url := matches[1]
		title := matches[2]
		link := matches[3]
		ref := strings.TrimSpace(url)
		expectedTitle := path.Base(link)

		if url == "broken-reference" {
			return `{{< biglink >}}Broken Reference{{< /biglink >}}`
		}
		if strings.HasSuffix(url, "/") {
			expectedTitle = path.Base(path.Dir(link))
			ref += "_index.md"
		}

		if url != link {
			panic(fmt.Sprintf("content-ref link mismatch: %s\nurl:%q\nlink:%q", match, url, link))
		}
		if title != expectedTitle {
			panic(fmt.Sprintf("content-ref title mismatch: %s\ntitle:%q\nexpected:%q", match, title, expectedTitle))
		}

		ref = conv.NearRef(page, ref)
		return `{{< biglink relref="` + ref + `" />}}`
	})
}

const youtubeLink = `{% embed url="https://www.youtube.com/watch?v=H6bRljVjR48" %}
Video Tutorial for the Setup Process
{% endembed %}`

// ReplaceTags implements replacing tags of `{% *** %}`
func (conv *Convert) ReplaceTags(page *Page) {
	tabIndex := 0

	if page.ContentPath == "sno-applications/qnap-storage-node-app.md" {
		page.Content = strings.ReplaceAll(page.Content, youtubeLink, "{{< youtube H6bRljVjR48 >}}")
	}

	rxTag := mustCompile(`{%\s*([a-zA-Z0-9-]+)\s(.*)\s*%}`)
	page.Content = rxTag.ReplaceAllStringFunc(page.Content, func(tag string) string {
		tok := rxTag.FindStringSubmatch(tag)
		switch tok[1] {
		case "code", "endcode":
			return ""
		case "tabs":
			tabIndex++
			return fmt.Sprintf(`{{< tabs id%d >}}`, tabIndex)
		case "endtabs":
			return `{{< /tabs >}}`
		case "tab":
			var title string
			if match(`^title="(.*)"$`, strings.TrimSpace(tok[2]), nil, &title) {
				title = strings.TrimSpace(title)
				if strings.EqualFold(title, "macOS") { // fix some misnamed tabs
					title = "macOS"
				}
				return `{{< tab "` + title + `" >}}`
			}
		case "endtab":
			return `{{< /tab >}}`
		case "hint":
			switch strings.TrimSpace(tok[2]) {
			case `style="info"`:
				return `{{< hint info >}}`
			case `style="warning"`:
				return `{{< hint warning >}}`
			case `style="danger"`:
				return `{{< hint danger >}}`
			case `style="success"`:
				return `{{< hint success >}}`
			}
		case "endhint":
			return `{{< /hint >}}`
		case "embed":
			var url string
			if match(`^url="(.*)"$`, strings.TrimSpace(tok[2]), nil, &url) {
				url = strings.TrimSpace(url)
				title, ok := urlToTitle[url]
				if ok {
					return `{{< biglink href="` + strings.TrimSpace(url) + `" >}}` + title + `{{< /biglink >}}`
				}
			}
		}

		panic("unhandled: " + tag)
	})
}

func (conv *Convert) NearRef(page *Page, rel string) string {
	if strings.HasPrefix(rel, "http") || strings.HasPrefix(rel, "/") {
		return rel
	}
	if strings.HasPrefix(rel, "../") || strings.HasPrefix(rel, "/") {
		return conv.AbsRef(page, rel)
	}
	return rel
}

func (conv *Convert) AbsRef(page *Page, rel string) string {
	if strings.HasPrefix(rel, "http") || strings.HasPrefix(rel, "/") {
		return rel
	}
	if strings.HasPrefix(rel, "/") {
		return "/" + path.Join(conv.TargetDir, rel)
	}
	return "/" + path.Clean(path.Join(conv.TargetDir, path.Dir(page.ContentPath), rel))
}

// FixTrailingSpace fixes some weird content issues in the markdown files.
func (conv *Convert) FixTrailingSpace(page *Page) {
	page.Content = replaceAll(`&#x20;`, page.Content, "")
	page.Content = replaceAll(` *$`, page.Content, "")
}

// FixLinksToReadme fixes links to README.md -> _index.md.
func (conv *Convert) FixLinksToReadme(page *Page) {
	page.Content = replaceAll(`where-can-i-find-a-config\.yaml\.md`, page.Content, "where-can-i-find-a-config-yaml.md)")
	page.Content = replaceAll(`README\.md\)`, page.Content, "_index.md)")
}

// FixImageLinks fixes links to "![xyz](<a/b/c>)" --> "![xyz](a/b/c)".
func (conv *Convert) FixImageLinks(page *Page) {
	rx := mustCompile(`([^\\])!\[([^\]]*)\]\((<[^>]*>|[^\)]*)\)`)
	page.Content = rx.ReplaceAllStringFunc(page.Content, func(m string) string {
		match := rx.FindStringSubmatch(m)
		prefix, title, url := match[1], match[2], match[3]
		url = strings.ReplaceAll(url, "\\_", "_")

		hasAngle := url[0] == '<'
		if hasAngle {
			url = url[1 : len(url)-1]
		}

		p := strings.Index(url, ".gitbook/assets")
		if p >= 0 {
			noPrefix := url[p+8+7:]
			// special case fix for images that are missing file extension
			if noPrefix == "/0" || noPrefix == "/1" || noPrefix == "/2" || noPrefix == "/3" {
				noPrefix += "-fix.png"
			}
			url = "/" + conv.TargetDir + "/" + assetsDir + noPrefix
		}

		if strings.HasPrefix(url, "http") {
			var asset string
			var err error
			asset, err = conv.DownloadAndCacheImage(page, url)
			if err != nil {
				conv.Warnings = append(conv.Warnings, fmt.Errorf("failed to download %s in %s: %w", url, page.ContentPath, err))
			}
			url = asset
		}

		if hasAngle {
			url = "<" + url + ">"
		}
		return prefix + "![" + title + "](" + url + ")"
	})
}

func fileExists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}

func (conv *Convert) DownloadAndCacheImage(page *Page, url string) (string, error) {
	h := sha256.Sum256([]byte(url))
	hash := hex.EncodeToString(h[:6])
	suffix := "-" + path.Base(url)
	ext := path.Ext(url)
	if len(ext) != 4 { // e.g. ".png"
		ext = "" // probably some garbage
	}
	if len(suffix) > 16 {
		suffix = ext
	}
	cachedFile := "cache-" + hash + suffix
	target := path.Join(conv.ExtraDir, conv.TargetDir, assetsDir, cachedFile)

	if ext == "" {
		for _, guess := range []string{".png", ".gif", ".jpg", ".svg"} {
			if fileExists(target + guess) {
				return "/" + path.Join(conv.TargetDir, assetsDir, cachedFile) + guess, nil
			}
		}
	} else {
		if fileExists(target) {
			return "/" + path.Join(conv.TargetDir, assetsDir, cachedFile), nil
		}
	}

	if *skipDownload {
		return url, nil
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return url, fmt.Errorf("get failed: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return url, fmt.Errorf("got status %v %v: %w", resp.Status, resp.StatusCode, err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return url, fmt.Errorf("read-all failed: %w", err)
	}

	if ext == "" {
		switch {
		case bytes.HasPrefix(data, []byte("GIF8")):
			ext = ".gif"
		case bytes.HasPrefix(data, []byte("\x89PNG")):
			ext = ".png"
		case bytes.HasPrefix(data, []byte("\xFF\xD8\xFF")):
			ext = ".jpg"
		default:
			return url, fmt.Errorf("unknown file type %q", data[:8])
		}
		cachedFile += ext
		target += ext
	}

	return "/" + path.Join(conv.TargetDir, assetsDir, cachedFile), writeFile(target, data)
}

// FixRegularLinks fixes links to "[xyz](<../b/c>)" --> "[xyz](/a/b/c)".
func (conv *Convert) FixRegularLinks(page *Page) {
	rx := mustCompile(`([^!])\[([^\]]*)\]\((<[^>]*>|[^\)]*)\)`)
	page.Content = rx.ReplaceAllStringFunc(page.Content, func(m string) string {
		match := rx.FindStringSubmatch(m)
		nonMatch, title, url := match[1], match[2], match[3]

		if strings.HasPrefix(url, "http") {
			// replace escaped underscores
			url = strings.ReplaceAll(url, `\_`, `_`)
			return nonMatch + "[" + title + "](" + url + ")"
		}

		return nonMatch + "[" + title + "](" + conv.NearRef(page, url) + ")"
	})
}

// ReplaceMath replaces multiline $$\sqrt{}$$ with {{< katex >}}\sqrt{}{{< /katex >}}.
func (conv *Convert) ReplaceMath(page *Page) {
	page.Content = replaceAll(
		`\$\$\n(.*)\n\$\$`,
		page.Content,
		"{{< katex display >}}\n$1\n{{< /katex >}}",
	)
}

// ReplaceStarryNight replaces **** in a row, which seems a weird gitbook artifact.
func (conv *Convert) ReplaceStarryNight(page *Page) {
	page.Content = replaceAll(`( *\*\*\*\* +|( +\*\*\*\* *))`, page.Content, " ")
	page.Content = replaceAll(`\*\*\*\*`, page.Content, "")
}

// ReplaceUnderline replaces __ in a row, which seems a weird gitbook artifact.
func (conv *Convert) ReplaceUnderlines(page *Page) {
	page.Content = replaceAll(`( *__ +|( +__ *))`, page.Content, " ")
	page.Content = replaceAll(`\b__\b`, page.Content, "")
}

// ReplaceBoldPeriod replaces **.**
func (conv *Convert) ReplaceBoldPeriod(page *Page) {
	page.Content = replaceAll(`\*\*\.\*\*`, page.Content, ".")
}

// ReplaceZeroWidth replaces \u200b and \u200c
func (conv *Convert) ReplaceZeroWidth(page *Page) {
	page.Content = replaceAll("\\*\\*\u200b\\*\\*", page.Content, "")
	page.Content = replaceAll("\u200b", page.Content, "")
	page.Content = replaceAll("\u200c", page.Content, "")
}

func (conv *Convert) AddSectionIndices() {
	entries, err := os.ReadDir(filepath.Join(conv.ContentDir, conv.TargetDir))
	if err != nil {
		conv.Failures = append(conv.Failures, err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if entry.Name() == assetsDir {
			continue
		}

		dir := path.Join(conv.TargetDir, entry.Name())
		info, ok := menuMapping[dir]
		if !ok {
			conv.Failures = append(conv.Failures, fmt.Errorf("menu mapping missing for %s", dir))
		}

		content := "---\n"
		if info.title != "" {
			content += "title: \"" + info.title + "\"\n"
			content += "weight: " + strconv.Itoa(info.weight) + "\n"
		}
		content += "bookFlatSection: true\n"
		content += "---\n"

		if err := writeFile(path.Join(conv.ContentDir, dir, "_index.md"), []byte(content)); err != nil {
			conv.Failures = append(conv.Failures, fmt.Errorf("menu failed for %s", dir))
		}
	}
}

var rxCache = map[string]*regexp.Regexp{}

func match(regex, content string, submatch ...*string) bool {
	rx := mustCompile(regex)
	matches := rx.FindStringSubmatch(content)
	if len(matches) == 0 {
		return false
	}
	if len(submatch) == 0 { // ignore when we don't want submatches
		return true
	}

	if len(submatch) != len(matches) {
		panic("match count mismatch")
	}

	for i, v := range matches {
		p := submatch[i]
		if p == nil {
			continue
		}
		*p = v
	}

	return true
}

func replaceAll(regex, content, newContent string) string {
	rx := mustCompile(regex)
	return rx.ReplaceAllString(content, newContent)
}

func mustReplaceFirst(regex, content, newContent string) string {
	rx := mustCompile(regex)
	loc := rx.FindStringIndex(content)
	if len(loc) == 0 {
		panic("did not match")
	}

	return content[:loc[0]] + newContent + content[loc[1]:]
}

func mustCompile(s string) *regexp.Regexp {
	rx, ok := rxCache[s]
	if !ok {
		rx = regexp.MustCompile(s)
		rxCache[s] = rx
	}
	return rx
}

type topLevelMenu struct {
	title  string
	weight int
}

type SummaryItem2 struct {
	Top        bool
	Title      string
	Dir        string
	Slug       string
	Link       string
	LineNumber int
	Indent     int
	Children   []*SummaryItem2
}

// ParseSummary parses SUMMARY.md files.
func ParseSummary(content string) []*SummaryItem2 {
	rxEntry := mustCompile(`\[([^\]]*)\]\(([^)]*)\)`)

	root := &SummaryItem2{}
	root.Indent = -2
	stack := []*SummaryItem2{root}

	last := func() *SummaryItem2 {
		return stack[len(stack)-1]
	}
	push := func(item *SummaryItem2) {
		if item.Slug == "" {
			if item.Link == "" || strings.HasPrefix(item.Link, "http") {
				item.Slug = slugify(item.Title)
			} else {
				if strings.HasSuffix(item.Link, "README.md") {
					item.Slug = path.Base(strings.TrimSuffix(item.Link, "/README.md"))
				} else {
					item.Slug = path.Base(strings.TrimSuffix(item.Link, ".md"))
				}
			}
		}

		// pop to the right level
		for len(stack) > 0 {
			if last().Indent >= item.Indent {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}

		// calculate dir
		for _, parent := range stack[1:] {
			item.Dir += parent.Slug + "/"
		}
		if item.Dir != "" {
			item.Dir = item.Dir[:len(item.Dir)-1]
		}

		// add to the parent
		parent := last()
		parent.Children = append(parent.Children, item)
		// add itself to the stack
		stack = append(stack, item)
	}

	for lineNumber, line := range strings.Split(content, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			title, slug := parseSummaryTitle(line)
			push(&SummaryItem2{
				Top:        true,
				Title:      title,
				Slug:       slug,
				Indent:     -1,
				LineNumber: lineNumber + 1,
			})
		} else {
			matches := rxEntry.FindStringSubmatch(line)
			push(&SummaryItem2{
				Title:      matches[1],
				Link:       matches[2],
				Indent:     countIndent(line),
				LineNumber: lineNumber,
			})
		}
	}

	return root.Children
}

func countIndent(line string) int {
	for i, p := range line {
		if p != ' ' {
			return i
		}
	}
	return len(line)
}

func parseSummaryTitle(title string) (_, slug string) {
	if p := strings.Index(title, "<"); p >= 0 {
		slug = strings.TrimSpace(title[p:])
		title = title[:p]
	}
	if slug != "" {
		rxID := mustCompile(`^<a\s+href="[^"]*"\s+id="([^"]+)"><\/a>$`)
		matches := rxID.FindStringSubmatch(slug)
		slug = matches[1]
	}
	title = strings.Trim(title, " #")
	return title, slug
}

func slugify(s string) string {
	cutdash := true
	emitdash := false

	slug := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsNumber(r) || unicode.IsLetter(r) {
			if emitdash && !cutdash {
				slug = append(slug, '-')
			}
			slug = append(slug, unicode.ToLower(r))

			emitdash = false
			cutdash = false
			continue
		}
		switch r {
		case '/', '=':
			if len(slug) == 0 || slug[len(slug)-1] != r {
				slug = append(slug, r)
			}
			emitdash = false
			cutdash = true
		default:
			emitdash = true
		}
	}

	if len(slug) == 0 {
		return "-"
	}

	return string(slug)
}

func testSummaryTree(in, out string) {
	data, err := ioutil.ReadFile(in)
	if err != nil {
		panic(err)
	}

	toplevel := ParseSummary(string(data))
	var b bytes.Buffer

	var recurse func(indent int, item *SummaryItem2)
	recurse = func(indent int, item *SummaryItem2) {
		if indent == -2 {
			fmt.Fprintf(&b, "## %s <%s>\n", item.Title, item.Slug)
		} else {
			link := path.Join(item.Dir, item.Slug)
			if len(item.Children) > 0 {
				link += "/_index.md"
			} else {
				link += ".md"
			}
			fmt.Fprintf(&b, "%s* [%s](%s)\n", strings.Repeat(" ", indent), item.Title, link)
		}

		for _, c := range item.Children {
			recurse(indent+2, c)
		}
	}

	for _, t := range toplevel {
		recurse(-2, t)
	}

	os.WriteFile(out, b.Bytes(), 0755)
}
