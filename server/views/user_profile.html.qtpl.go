// Code generated by qtc from "user_profile.html.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/user_profile.html.qtpl:1
package views

//line views/user_profile.html.qtpl:1
import "github.com/etherealmachine/rpg.ai/server/models"

//line views/user_profile.html.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/user_profile.html.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/user_profile.html.qtpl:4
type UserProfilePage struct {
	*BasePage
	User       models.User
	UserAssets []models.Asset
}

//line views/user_profile.html.qtpl:11
func StreamUploadAssetComponent(qw422016 *qt422016.Writer) {
//line views/user_profile.html.qtpl:11
	qw422016.N().S(`
<form action="/upload" enctype="multipart/form-data">
  <input type="file" id="fileupload" name="filename[]" multiple>
  <input type="submit" value="Upload">
</form>
`)
//line views/user_profile.html.qtpl:16
}

//line views/user_profile.html.qtpl:16
func WriteUploadAssetComponent(qq422016 qtio422016.Writer) {
//line views/user_profile.html.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/user_profile.html.qtpl:16
	StreamUploadAssetComponent(qw422016)
//line views/user_profile.html.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line views/user_profile.html.qtpl:16
}

//line views/user_profile.html.qtpl:16
func UploadAssetComponent() string {
//line views/user_profile.html.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
//line views/user_profile.html.qtpl:16
	WriteUploadAssetComponent(qb422016)
//line views/user_profile.html.qtpl:16
	qs422016 := string(qb422016.B)
//line views/user_profile.html.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
//line views/user_profile.html.qtpl:16
	return qs422016
//line views/user_profile.html.qtpl:16
}

//line views/user_profile.html.qtpl:18
func (p *UserProfilePage) StreamContent(qw422016 *qt422016.Writer) {
//line views/user_profile.html.qtpl:18
	qw422016.N().S(`
  <div class="container">
    <p class="my-4">`)
//line views/user_profile.html.qtpl:20
	qw422016.E().S(p.User.Email)
//line views/user_profile.html.qtpl:20
	qw422016.N().S(`, a member since `)
//line views/user_profile.html.qtpl:20
	qw422016.E().V(p.User.CreatedAt)
//line views/user_profile.html.qtpl:20
	qw422016.N().S(`</p>
    <h4>Assets</h4>
    <ul>
    `)
//line views/user_profile.html.qtpl:23
	for _, asset := range p.UserAssets {
//line views/user_profile.html.qtpl:23
		qw422016.N().S(`
      <li>`)
//line views/user_profile.html.qtpl:24
		qw422016.E().S(asset.Filename)
//line views/user_profile.html.qtpl:24
		qw422016.N().S(`</li>
    `)
//line views/user_profile.html.qtpl:25
	}
//line views/user_profile.html.qtpl:25
	qw422016.N().S(`
    </ul>
    <div id="asset-uploader"></div>
  </div>
`)
//line views/user_profile.html.qtpl:29
}

//line views/user_profile.html.qtpl:29
func (p *UserProfilePage) WriteContent(qq422016 qtio422016.Writer) {
//line views/user_profile.html.qtpl:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/user_profile.html.qtpl:29
	p.StreamContent(qw422016)
//line views/user_profile.html.qtpl:29
	qt422016.ReleaseWriter(qw422016)
//line views/user_profile.html.qtpl:29
}

//line views/user_profile.html.qtpl:29
func (p *UserProfilePage) Content() string {
//line views/user_profile.html.qtpl:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/user_profile.html.qtpl:29
	p.WriteContent(qb422016)
//line views/user_profile.html.qtpl:29
	qs422016 := string(qb422016.B)
//line views/user_profile.html.qtpl:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/user_profile.html.qtpl:29
	return qs422016
//line views/user_profile.html.qtpl:29
}