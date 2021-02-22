// Code generated by qtc from "unattach.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/unattach.qtpl:1
package views

//line views/unattach.qtpl:1
import "net/http"

//line views/unattach.qtpl:2
import "github.com/bouncepaw/mycorrhiza/util"

//line views/unattach.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/unattach.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/unattach.qtpl:4
func StreamUnattachAskHTML(qw422016 *qt422016.Writer, rq *http.Request, hyphaName string, isOld bool) {
//line views/unattach.qtpl:4
	qw422016.N().S(`
`)
//line views/unattach.qtpl:5
	StreamNavHTML(qw422016, rq, hyphaName, "unattach-ask")
//line views/unattach.qtpl:5
	qw422016.N().S(`
<div class="layout">
<main class="main-width">
	<form class="modal" action="/unattach-confirm/`)
//line views/unattach.qtpl:8
	qw422016.E().S(hyphaName)
//line views/unattach.qtpl:8
	qw422016.N().S(`">
		<fieldset class="modal__fieldset">
			<legend class="modal__title">Unattach `)
//line views/unattach.qtpl:10
	qw422016.E().S(util.BeautifulName(hyphaName))
//line views/unattach.qtpl:10
	qw422016.N().S(`?</legend>
			<p class="modal__confirmation-msg">Do you really want to unattach hypha <em>`)
//line views/unattach.qtpl:11
	qw422016.E().S(hyphaName)
//line views/unattach.qtpl:11
	qw422016.N().S(`</em>?</p>
			<input type="submit" value="Confirm" class="modal__action modal__submit" autofocus>
			<a href="/hypha/`)
//line views/unattach.qtpl:13
	qw422016.E().S(hyphaName)
//line views/unattach.qtpl:13
	qw422016.N().S(`" class="modal__action modal__cancel">Cancel</a>
		</fieldset>
	</form>
</main>
</div>
`)
//line views/unattach.qtpl:18
}

//line views/unattach.qtpl:18
func WriteUnattachAskHTML(qq422016 qtio422016.Writer, rq *http.Request, hyphaName string, isOld bool) {
//line views/unattach.qtpl:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/unattach.qtpl:18
	StreamUnattachAskHTML(qw422016, rq, hyphaName, isOld)
//line views/unattach.qtpl:18
	qt422016.ReleaseWriter(qw422016)
//line views/unattach.qtpl:18
}

//line views/unattach.qtpl:18
func UnattachAskHTML(rq *http.Request, hyphaName string, isOld bool) string {
//line views/unattach.qtpl:18
	qb422016 := qt422016.AcquireByteBuffer()
//line views/unattach.qtpl:18
	WriteUnattachAskHTML(qb422016, rq, hyphaName, isOld)
//line views/unattach.qtpl:18
	qs422016 := string(qb422016.B)
//line views/unattach.qtpl:18
	qt422016.ReleaseByteBuffer(qb422016)
//line views/unattach.qtpl:18
	return qs422016
//line views/unattach.qtpl:18
}
