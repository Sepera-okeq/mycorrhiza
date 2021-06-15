// Code generated by qtc from "mutators.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/mutators.qtpl:1
package views

//line views/mutators.qtpl:1
import "net/http"

//line views/mutators.qtpl:3
import "github.com/bouncepaw/mycorrhiza/cfg"

//line views/mutators.qtpl:4
import "github.com/bouncepaw/mycorrhiza/util"

//line views/mutators.qtpl:5
import "github.com/bouncepaw/mycorrhiza/user"

//line views/mutators.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/mutators.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/mutators.qtpl:7
func StreamToolbar(qw422016 *qt422016.Writer, u *user.User) {
//line views/mutators.qtpl:7
	qw422016.N().S(`
<aside class="edit-toolbar layout-card">
	<h2 class="edit-toolbar__title layout-card__title">Markup</h2>
	<section class="edit-toolbar__buttons">
	`)
//line views/mutators.qtpl:11
	for _, el := range []struct {
		class   string
		onclick string
		display string
	}{
		{"link", "wrapLink()", "[[link]]"},
		{"heading2", "insertHeading2()", "## heading"},
		{"heading3", "insertHeading3()", "### heading"},
		{"bold", "wrapBold()", "<b>**Bold**</b>"},
		{"italic", "wrapItalic()", "<i>//Italic//</i>"},
		{"highlighted", "wrapHighlighted()", "<mark>!!Highlight!!</mark>"},
		{"monospace", "wrapMonospace()", "<code>`Monospace`</code>"},
		{"lifted", "wrapLifted()", "<sup>^^Lifted^^</sup>"},
		{"lowered", "wrapLowered()", "<sub>,,Lowered,,</sub>"},
		{"strikethrough", "wrapStrikethrough()", "<strike>~~Strikethrough~~</strike>"},
		{"rocket", "insertRocket()", "=> rocketlink"},
		{"xcl", "insertXcl()", "<= transclusion"},
		{"img", "insertImgBlock()", "<code>img {}</code>"},
		{"table", "insertTableBlock()", "<code>table {}</code>"},
		{"hr", "insertHorizontalBar()", "Horizontal bar"},
		{"codeblock", "insertCodeblock()", "Code block"},
		{"bulletedlist", "insertBulletedList()", "* bullet list"},
		{"numberedlist", "insertNumberedList()", "*. number list"},
	} {
//line views/mutators.qtpl:34
		qw422016.N().S(`
		<button
			class="btn edit-toolbar__btn edit-toolbar__`)
//line views/mutators.qtpl:36
		qw422016.E().S(el.class)
//line views/mutators.qtpl:36
		qw422016.N().S(`"
			onclick="`)
//line views/mutators.qtpl:37
		qw422016.E().S(el.onclick)
//line views/mutators.qtpl:37
		qw422016.N().S(`">
			`)
//line views/mutators.qtpl:38
		qw422016.N().S(el.display)
//line views/mutators.qtpl:38
		qw422016.N().S(`
		</button>
	`)
//line views/mutators.qtpl:40
	}
//line views/mutators.qtpl:40
	qw422016.N().S(`
	</section>
	<p class="edit-toolbar__ad"><a href="https://mycorrhiza.lesarbr.es/hypha/mycomarkup" target="_blank">Learn more</a> about mycomarkup</p>
	<h2 class="edit-toolbar__title layout-card__title">Actions</h2>
	<section class="edit-toolbar__buttons">
	`)
//line views/mutators.qtpl:45
	for _, el := range []struct {
		class   string
		onclick string
		display string
	}{
		{"date", "insertDate()", "Insert current date"},
		{"time", "insertTimeUTC()", "Insert current time"},
	} {
//line views/mutators.qtpl:52
		qw422016.N().S(`
		<button
			class="btn edit-toolbar__btn edit-toolbar__`)
//line views/mutators.qtpl:54
		qw422016.E().S(el.class)
//line views/mutators.qtpl:54
		qw422016.N().S(`"
			onclick="`)
//line views/mutators.qtpl:55
		qw422016.E().S(el.onclick)
//line views/mutators.qtpl:55
		qw422016.N().S(`">
			`)
//line views/mutators.qtpl:56
		qw422016.N().S(el.display)
//line views/mutators.qtpl:56
		qw422016.N().S(`
		</button>
	`)
//line views/mutators.qtpl:58
	}
//line views/mutators.qtpl:58
	qw422016.N().S(`
	`)
//line views/mutators.qtpl:59
	if u.Group != "anon" {
//line views/mutators.qtpl:59
		qw422016.N().S(`
		<button
			class="btn edit-toolbar__btn edit-toolbar__user-link"
			onclick="insertUserlink()">
			Link yourself
		</button>
	`)
//line views/mutators.qtpl:65
	}
//line views/mutators.qtpl:65
	qw422016.N().S(`
	</section>
</aside>
<script src="/static/toolbar.js"></script>
`)
//line views/mutators.qtpl:69
}

//line views/mutators.qtpl:69
func WriteToolbar(qq422016 qtio422016.Writer, u *user.User) {
//line views/mutators.qtpl:69
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/mutators.qtpl:69
	StreamToolbar(qw422016, u)
//line views/mutators.qtpl:69
	qt422016.ReleaseWriter(qw422016)
//line views/mutators.qtpl:69
}

//line views/mutators.qtpl:69
func Toolbar(u *user.User) string {
//line views/mutators.qtpl:69
	qb422016 := qt422016.AcquireByteBuffer()
//line views/mutators.qtpl:69
	WriteToolbar(qb422016, u)
//line views/mutators.qtpl:69
	qs422016 := string(qb422016.B)
//line views/mutators.qtpl:69
	qt422016.ReleaseByteBuffer(qb422016)
//line views/mutators.qtpl:69
	return qs422016
//line views/mutators.qtpl:69
}

//line views/mutators.qtpl:71
func StreamEditHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName, textAreaFill, warning string) {
//line views/mutators.qtpl:71
	qw422016.N().S(`
`)
//line views/mutators.qtpl:72
	qw422016.N().S(NavHTML(rq, hyphaName, "edit"))
//line views/mutators.qtpl:72
	qw422016.N().S(`
<div class="layout">
<main class="main-width edit edit_no-preview">
	<h1 class="edit__title">Edit `)
//line views/mutators.qtpl:75
	qw422016.E().S(util.BeautifulName(hyphaName))
//line views/mutators.qtpl:75
	qw422016.N().S(`</h1>
	`)
//line views/mutators.qtpl:76
	qw422016.N().S(warning)
//line views/mutators.qtpl:76
	qw422016.N().S(`
	<form method="post" class="edit-form"
			action="/upload-text/`)
//line views/mutators.qtpl:78
	qw422016.E().S(hyphaName)
//line views/mutators.qtpl:78
	qw422016.N().S(`">
		<textarea name="text" class="edit-form__textarea" autofocus>`)
//line views/mutators.qtpl:79
	qw422016.E().S(textAreaFill)
//line views/mutators.qtpl:79
	qw422016.N().S(`</textarea>
		<br/><br/>
		<label for="text">Describe your changes:</label><br>
		<input id="text" type="text" name="message" class="edit-form__message" placeholder="(optional)">
		<br/><br/>
		<input type="submit" name="action" class="btn" value="Save" class="edit-form__save"/>
		<input type="submit" name="action" class="btn" value="Preview" class="edit-form__preview">
		<a href="/hypha/`)
//line views/mutators.qtpl:86
	qw422016.E().S(hyphaName)
//line views/mutators.qtpl:86
	qw422016.N().S(`" class="btn btn_weak">Cancel</a>
	</form>
</main>
`)
//line views/mutators.qtpl:89
	qw422016.N().S(Toolbar(user.FromRequest(rq)))
//line views/mutators.qtpl:89
	qw422016.N().S(`
</div>
`)
//line views/mutators.qtpl:91
	streameditScripts(qw422016)
//line views/mutators.qtpl:91
	qw422016.N().S(`
`)
//line views/mutators.qtpl:92
}

//line views/mutators.qtpl:92
func WriteEditHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName, textAreaFill, warning string) {
//line views/mutators.qtpl:92
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/mutators.qtpl:92
	StreamEditHTML(qw422016, rq, hyphaName, textAreaFill, warning)
//line views/mutators.qtpl:92
	qt422016.ReleaseWriter(qw422016)
//line views/mutators.qtpl:92
}

//line views/mutators.qtpl:92
func EditHTML(rq *http.Request, hyphaName, textAreaFill, warning string) string {
//line views/mutators.qtpl:92
	qb422016 := qt422016.AcquireByteBuffer()
//line views/mutators.qtpl:92
	WriteEditHTML(qb422016, rq, hyphaName, textAreaFill, warning)
//line views/mutators.qtpl:92
	qs422016 := string(qb422016.B)
//line views/mutators.qtpl:92
	qt422016.ReleaseByteBuffer(qb422016)
//line views/mutators.qtpl:92
	return qs422016
//line views/mutators.qtpl:92
}

//line views/mutators.qtpl:94
func StreamPreviewHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName, textAreaFill, message, warning string, renderedPage string) {
//line views/mutators.qtpl:94
	qw422016.N().S(`
`)
//line views/mutators.qtpl:95
	qw422016.N().S(NavHTML(rq, hyphaName, "edit"))
//line views/mutators.qtpl:95
	qw422016.N().S(`
<div class="layout">
<main class="main-width edit edit_with-preview">
	<h1>Edit `)
//line views/mutators.qtpl:98
	qw422016.E().S(util.BeautifulName(hyphaName))
//line views/mutators.qtpl:98
	qw422016.N().S(` (preview)</h1>
	`)
//line views/mutators.qtpl:99
	qw422016.N().S(warning)
//line views/mutators.qtpl:99
	qw422016.N().S(`
	<form method="post" class="edit-form"
			action="/upload-text/`)
//line views/mutators.qtpl:101
	qw422016.E().S(hyphaName)
//line views/mutators.qtpl:101
	qw422016.N().S(`">
		<textarea class="edit-form__textarea" name="text">`)
//line views/mutators.qtpl:102
	qw422016.E().S(textAreaFill)
//line views/mutators.qtpl:102
	qw422016.N().S(`</textarea>
		<br/><br/>
		<label for="text">Edit description:</label>
		<input id="text" type="text" name="message" class="edit-form__message" value="`)
//line views/mutators.qtpl:105
	qw422016.E().S(message)
//line views/mutators.qtpl:105
	qw422016.N().S(`">
		<br/><br/>
		<input type="submit" name="action" value="Save" class="edit-form__save"/>
		<input type="submit" name="action" value="Preview" class="edit-form__preview">
		<a href="/hypha/`)
//line views/mutators.qtpl:109
	qw422016.E().S(hyphaName)
//line views/mutators.qtpl:109
	qw422016.N().S(`" class="edit-form__cancel">Cancel</a>
	</form>
	<p class="warning">Note that the hypha is not saved yet. You can preview the changes ↓</p>
	<article class="edit__preview">`)
//line views/mutators.qtpl:112
	qw422016.N().S(renderedPage)
//line views/mutators.qtpl:112
	qw422016.N().S(`</article>
</main>
`)
//line views/mutators.qtpl:114
	qw422016.N().S(Toolbar(user.FromRequest(rq)))
//line views/mutators.qtpl:114
	qw422016.N().S(`
</div>
`)
//line views/mutators.qtpl:116
	streameditScripts(qw422016)
//line views/mutators.qtpl:116
	qw422016.N().S(`
`)
//line views/mutators.qtpl:117
}

//line views/mutators.qtpl:117
func WritePreviewHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName, textAreaFill, message, warning string, renderedPage string) {
//line views/mutators.qtpl:117
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/mutators.qtpl:117
	StreamPreviewHTML(qw422016, rq, hyphaName, textAreaFill, message, warning, renderedPage)
//line views/mutators.qtpl:117
	qt422016.ReleaseWriter(qw422016)
//line views/mutators.qtpl:117
}

//line views/mutators.qtpl:117
func PreviewHTML(rq *http.Request, hyphaName, textAreaFill, message, warning string, renderedPage string) string {
//line views/mutators.qtpl:117
	qb422016 := qt422016.AcquireByteBuffer()
//line views/mutators.qtpl:117
	WritePreviewHTML(qb422016, rq, hyphaName, textAreaFill, message, warning, renderedPage)
//line views/mutators.qtpl:117
	qs422016 := string(qb422016.B)
//line views/mutators.qtpl:117
	qt422016.ReleaseByteBuffer(qb422016)
//line views/mutators.qtpl:117
	return qs422016
//line views/mutators.qtpl:117
}

//line views/mutators.qtpl:119
func streameditScripts(qw422016 *qt422016.Writer) {
//line views/mutators.qtpl:119
	qw422016.N().S(`
`)
//line views/mutators.qtpl:120
	for _, scriptPath := range cfg.EditScripts {
//line views/mutators.qtpl:120
		qw422016.N().S(`
<script src="`)
//line views/mutators.qtpl:121
		qw422016.E().S(scriptPath)
//line views/mutators.qtpl:121
		qw422016.N().S(`"></script>
`)
//line views/mutators.qtpl:122
	}
//line views/mutators.qtpl:122
	qw422016.N().S(`
`)
//line views/mutators.qtpl:123
}

//line views/mutators.qtpl:123
func writeeditScripts(qq422016 qtio422016.Writer) {
//line views/mutators.qtpl:123
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/mutators.qtpl:123
	streameditScripts(qw422016)
//line views/mutators.qtpl:123
	qt422016.ReleaseWriter(qw422016)
//line views/mutators.qtpl:123
}

//line views/mutators.qtpl:123
func editScripts() string {
//line views/mutators.qtpl:123
	qb422016 := qt422016.AcquireByteBuffer()
//line views/mutators.qtpl:123
	writeeditScripts(qb422016)
//line views/mutators.qtpl:123
	qs422016 := string(qb422016.B)
//line views/mutators.qtpl:123
	qt422016.ReleaseByteBuffer(qb422016)
//line views/mutators.qtpl:123
	return qs422016
//line views/mutators.qtpl:123
}
