// Code generated by qtc from "nav.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/nav.qtpl:1
package views

//line views/nav.qtpl:1
import "net/http"

//line views/nav.qtpl:2
import "strings"

//line views/nav.qtpl:3
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/nav.qtpl:4
import "github.com/bouncepaw/mycorrhiza/user"

//line views/nav.qtpl:5
import "github.com/bouncepaw/mycorrhiza/util"

// This is the <nav> seen on top of many pages.

//line views/nav.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/nav.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/nav.qtpl:9
type navEntry struct {
	path  string
	title string
}

var navEntries = []navEntry{
	{"page", "Hypha"},
	{"edit", "Edit"},
	{"attachment", "Attachment"},
	{"history", "History"},
	{"revision", "NOT REACHED"},
	{"rename-ask", "Rename"},
	{"delete-ask", "Delete"},
	{"text", "Raw text"},
}

//line views/nav.qtpl:25
func StreamNavHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName, navType string, revisionHash ...string) {
//line views/nav.qtpl:25
	qw422016.N().S(`
`)
//line views/nav.qtpl:27
	u := user.FromRequest(rq)

//line views/nav.qtpl:28
	qw422016.N().S(`
	<nav class="hypha-tabs main-width">
		<ul class="hypha-tabs__flex">
`)
//line views/nav.qtpl:31
	for _, entry := range navEntries {
//line views/nav.qtpl:32
		if navType == "revision" && entry.path == "revision" {
//line views/nav.qtpl:32
			qw422016.N().S(`			<li class="hypha-tabs__tab hypha-tabs__tab_active">
				<span class="hypha-tabs__selection">`)
//line views/nav.qtpl:34
			qw422016.E().S(revisionHash[0])
//line views/nav.qtpl:34
			qw422016.N().S(`</span>
			</li>
`)
//line views/nav.qtpl:36
		} else if navType == entry.path {
//line views/nav.qtpl:36
			qw422016.N().S(`			<li class="hypha-tabs__tab hypha-tabs__tab_active">
				<span class="hypha-tabs__selection">`)
//line views/nav.qtpl:38
			qw422016.E().S(entry.title)
//line views/nav.qtpl:38
			qw422016.N().S(`</span>
			</li>
`)
//line views/nav.qtpl:40
		} else if entry.path != "revision" && u.CanProceed(entry.path) {
//line views/nav.qtpl:40
			qw422016.N().S(`			<li class="hypha-tabs__tab">
				<a class="hypha-tabs__link" href="/`)
//line views/nav.qtpl:42
			qw422016.E().S(entry.path)
//line views/nav.qtpl:42
			qw422016.N().S(`/`)
//line views/nav.qtpl:42
			qw422016.E().S(hyphaName)
//line views/nav.qtpl:42
			qw422016.N().S(`">`)
//line views/nav.qtpl:42
			qw422016.E().S(entry.title)
//line views/nav.qtpl:42
			qw422016.N().S(`</a>
			</li>
`)
//line views/nav.qtpl:44
		}
//line views/nav.qtpl:45
	}
//line views/nav.qtpl:45
	qw422016.N().S(`		</ul>
	</nav>
`)
//line views/nav.qtpl:48
}

//line views/nav.qtpl:48
func WriteNavHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName, navType string, revisionHash ...string) {
//line views/nav.qtpl:48
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:48
	StreamNavHTML(qw422016, rq, hyphaName, navType, revisionHash...)
//line views/nav.qtpl:48
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:48
}

//line views/nav.qtpl:48
func NavHTML(rq *http.Request, hyphaName, navType string, revisionHash ...string) string {
//line views/nav.qtpl:48
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:48
	WriteNavHTML(qb422016, rq, hyphaName, navType, revisionHash...)
//line views/nav.qtpl:48
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:48
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:48
	return qs422016
//line views/nav.qtpl:48
}

//line views/nav.qtpl:50
func StreamUserMenuHTML(qw422016 *qt422016.Writer, u *user.User) {
//line views/nav.qtpl:50
	qw422016.N().S(`
`)
//line views/nav.qtpl:51
	if user.AuthUsed {
//line views/nav.qtpl:51
		qw422016.N().S(`
	<li class="header-links__entry header-links__entry_user">
		`)
//line views/nav.qtpl:53
		if u.Group == "anon" {
//line views/nav.qtpl:53
			qw422016.N().S(`
		<a href="/login" class="header-links__link">Login</a>
		`)
//line views/nav.qtpl:55
		} else {
//line views/nav.qtpl:55
			qw422016.N().S(`
		<a href="/hypha/`)
//line views/nav.qtpl:56
			qw422016.E().S(cfg.UserHypha)
//line views/nav.qtpl:56
			qw422016.N().S(`/`)
//line views/nav.qtpl:56
			qw422016.E().S(u.Name)
//line views/nav.qtpl:56
			qw422016.N().S(`" class="header-links__link">`)
//line views/nav.qtpl:56
			qw422016.E().S(util.BeautifulName(u.Name))
//line views/nav.qtpl:56
			qw422016.N().S(`</a>
		`)
//line views/nav.qtpl:57
		}
//line views/nav.qtpl:57
		qw422016.N().S(`
	</li>
`)
//line views/nav.qtpl:59
	}
//line views/nav.qtpl:59
	qw422016.N().S(`
`)
//line views/nav.qtpl:60
	if user.AuthUsed && cfg.UseRegistration && u.Group == "anon" {
//line views/nav.qtpl:60
		qw422016.N().S(`
	<li class="header-links__entry header-links__entry_register">
		<a href="/register" class="header-links__link">Register</a>
	</li>
`)
//line views/nav.qtpl:64
	}
//line views/nav.qtpl:64
	qw422016.N().S(`
`)
//line views/nav.qtpl:65
}

//line views/nav.qtpl:65
func WriteUserMenuHTML(qq422016 qtio422016.Writer, u *user.User) {
//line views/nav.qtpl:65
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:65
	StreamUserMenuHTML(qw422016, u)
//line views/nav.qtpl:65
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:65
}

//line views/nav.qtpl:65
func UserMenuHTML(u *user.User) string {
//line views/nav.qtpl:65
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:65
	WriteUserMenuHTML(qb422016, u)
//line views/nav.qtpl:65
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:65
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:65
	return qs422016
//line views/nav.qtpl:65
}

//line views/nav.qtpl:67
func StreamRelativeHyphaeHTML(qw422016 *qt422016.Writer, relatives string) {
//line views/nav.qtpl:67
	qw422016.N().S(`
<aside class="relative-hyphae layout-card">
	<h2 class="relative-hyphae__title layout-card__title">Relative hyphae</h2>
	`)
//line views/nav.qtpl:70
	qw422016.N().S(relatives)
//line views/nav.qtpl:70
	qw422016.N().S(`
</aside>
`)
//line views/nav.qtpl:72
}

//line views/nav.qtpl:72
func WriteRelativeHyphaeHTML(qq422016 qtio422016.Writer, relatives string) {
//line views/nav.qtpl:72
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:72
	StreamRelativeHyphaeHTML(qw422016, relatives)
//line views/nav.qtpl:72
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:72
}

//line views/nav.qtpl:72
func RelativeHyphaeHTML(relatives string) string {
//line views/nav.qtpl:72
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:72
	WriteRelativeHyphaeHTML(qb422016, relatives)
//line views/nav.qtpl:72
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:72
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:72
	return qs422016
//line views/nav.qtpl:72
}

//line views/nav.qtpl:74
func StreamSubhyphaeHTML(qw422016 *qt422016.Writer, subhyphae string) {
//line views/nav.qtpl:74
	qw422016.N().S(`
`)
//line views/nav.qtpl:75
	if strings.TrimSpace(subhyphae) != "" {
//line views/nav.qtpl:75
		qw422016.N().S(`
<section class="subhyphae">
	<h2 class="subhyphae__title">Subhyphae</h2>
	<nav class="subhyphae__nav">
		<ul class="subhyphae__list">
		`)
//line views/nav.qtpl:80
		qw422016.N().S(subhyphae)
//line views/nav.qtpl:80
		qw422016.N().S(`
		</ul>
	</nav>
</section>
`)
//line views/nav.qtpl:84
	}
//line views/nav.qtpl:84
	qw422016.N().S(`
`)
//line views/nav.qtpl:85
}

//line views/nav.qtpl:85
func WriteSubhyphaeHTML(qq422016 qtio422016.Writer, subhyphae string) {
//line views/nav.qtpl:85
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/nav.qtpl:85
	StreamSubhyphaeHTML(qw422016, subhyphae)
//line views/nav.qtpl:85
	qt422016.ReleaseWriter(qw422016)
//line views/nav.qtpl:85
}

//line views/nav.qtpl:85
func SubhyphaeHTML(subhyphae string) string {
//line views/nav.qtpl:85
	qb422016 := qt422016.AcquireByteBuffer()
//line views/nav.qtpl:85
	WriteSubhyphaeHTML(qb422016, subhyphae)
//line views/nav.qtpl:85
	qs422016 := string(qb422016.B)
//line views/nav.qtpl:85
	qt422016.ReleaseByteBuffer(qb422016)
//line views/nav.qtpl:85
	return qs422016
//line views/nav.qtpl:85
}
