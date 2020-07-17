// Code generated by qtc from "campaign.html.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/campaign.html.qtpl:1
package views

//line views/campaign.html.qtpl:1
import "github.com/etherealmachine/rpg.ai/server/models"

//line views/campaign.html.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/campaign.html.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/campaign.html.qtpl:4
type CampaignPage struct {
	*BasePage
	Campaign *models.Campaign
}

//line views/campaign.html.qtpl:10
func (p *CampaignPage) StreamContent(qw422016 *qt422016.Writer) {
//line views/campaign.html.qtpl:10
	qw422016.N().S(`
  <div class="container">
  </div>
`)
//line views/campaign.html.qtpl:13
}

//line views/campaign.html.qtpl:13
func (p *CampaignPage) WriteContent(qq422016 qtio422016.Writer) {
//line views/campaign.html.qtpl:13
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/campaign.html.qtpl:13
	p.StreamContent(qw422016)
//line views/campaign.html.qtpl:13
	qt422016.ReleaseWriter(qw422016)
//line views/campaign.html.qtpl:13
}

//line views/campaign.html.qtpl:13
func (p *CampaignPage) Content() string {
//line views/campaign.html.qtpl:13
	qb422016 := qt422016.AcquireByteBuffer()
//line views/campaign.html.qtpl:13
	p.WriteContent(qb422016)
//line views/campaign.html.qtpl:13
	qs422016 := string(qb422016.B)
//line views/campaign.html.qtpl:13
	qt422016.ReleaseByteBuffer(qb422016)
//line views/campaign.html.qtpl:13
	return qs422016
//line views/campaign.html.qtpl:13
}
