package easyLexML

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDraft2Strict_EmptyDocument(t *testing.T) {
	input := bytes.NewBufferString(``)
	output := bytes.NewBuffer(nil)
	err := Draft2Strict(input, output)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "no <EasyLexML> found")
}

func TestDraft2Strict_MinimalDocument(t *testing.T) {
	input := bytes.NewBufferString(`<EasyLexML><corpus></corpus></EasyLexML>`)
	output := bytes.NewBuffer(nil)
	err := Draft2Strict(input, output)
	require.Nil(t, err)

	expected := `<?xml?>
<EasyLexML>
	<toc id="toc">
		<label href="#toc">
			Table of Contents
		</label>
		<ul/>
	</toc>
	<corpus id="corpus"/>
</EasyLexML>`
	assert.Equal(t, expected, output.String())
}

func TestDraft2Strict_CustomTOCTitle(t *testing.T) {
	input := bytes.NewBufferString(`<EasyLexML><set-meta TocTitle="𝕿𝖆𝖇𝖑𝖊 𝖔𝖋 𝕮𝖔𝖓𝖙𝖊𝖓𝖙𝖘"/><corpus></corpus></EasyLexML>`)
	output := bytes.NewBuffer(nil)
	err := Draft2Strict(input, output)
	require.Nil(t, err)

	expected := `<?xml?>
<EasyLexML>
	<toc id="toc">
		<label href="#toc">
			𝕿𝖆𝖇𝖑𝖊 𝖔𝖋 𝕮𝖔𝖓𝖙𝖊𝖓𝖙𝖘
		</label>
		<ul/>
	</toc>
	<corpus id="corpus"/>
</EasyLexML>`
	assert.Equal(t, expected, output.String())
}

func TestDraft2Strict_CustomAbstract1(t *testing.T) {
	input := bytes.NewBufferString(`
<EasyLexML>
	<set-meta AbstractTitle="𝕬𝖇𝖘𝖙𝖗𝖆𝖈𝖙"/>
	<abstract>
		Just a simple example of EasyLexML.
	</abstract>
	<corpus>
	</corpus>
</EasyLexML>`)
	output := bytes.NewBuffer(nil)
	err := Draft2Strict(input, output)
	require.Nil(t, err)

	expected := `<?xml?>
<EasyLexML>
	<abstract>
		<label href="#abstract">
			𝕬𝖇𝖘𝖙𝖗𝖆𝖈𝖙
		</label>
		<p>
			Just a simple example of EasyLexML.
		</p>
	</abstract>
	<toc id="toc">
		<label href="#toc">
			Table of Contents
		</label>
		<ul/>
	</toc>
	<corpus id="corpus">
	</corpus>
</EasyLexML>`
	assert.Equal(t, expected, output.String())
}

func TestDraft2Strict_CustomAbstract2(t *testing.T) {
	input := bytes.NewBufferString(`
<EasyLexML>
	<set-meta AbstractTitle="𝕬𝖇𝖘𝖙𝖗𝖆𝖈𝖙"/>
	<abstract label="𝕸𝖞 𝕬𝖇𝖘𝖙𝖗𝖆𝖈𝖙">
		Just a simple example of EasyLexML.
	</abstract>
	<corpus>
	</corpus>
</EasyLexML>`)
	output := bytes.NewBuffer(nil)
	err := Draft2Strict(input, output)
	require.Nil(t, err)

	expected := `<?xml?>
<EasyLexML>
	<abstract>
		<label href="#abstract">
			𝕸𝖞 𝕬𝖇𝖘𝖙𝖗𝖆𝖈𝖙
		</label>
		<p>
			Just a simple example of EasyLexML.
		</p>
	</abstract>
	<toc id="toc">
		<label href="#toc">
			Table of Contents
		</label>
		<ul/>
	</toc>
	<corpus id="corpus">
	</corpus>
</EasyLexML>`
	assert.Equal(t, expected, output.String())
}

func TestDraft2Strict_Doc1(t *testing.T) {
	input := bytes.NewBufferString(`
<EasyLexML>
	<corpus>
		<cls>Lorem <a href="https://en.wikipedia.org/wiki/Lorem_ipsum">Ipsum</a>.</cls>
	</corpus>
</EasyLexML>`)
	output := bytes.NewBuffer(nil)
	err := Draft2Strict(input, output)
	require.Nil(t, err)

	expected := `<?xml?>
<EasyLexML>
	<toc id="toc">
		<label href="#toc">
			Table of Contents
		</label>
		<ul>
			<li>
				<a href="#cls1">
					Cls. 1
				</a>
			</li>
		</ul>
	</toc>
	<corpus id="corpus">
		<cls id="cls1" num="1">
			<p id="cls1_p1">
				<label href="#cls1">
					Cls. 1
				</label>Lorem
				<a href="https://en.wikipedia.org/wiki/Lorem_ipsum">Ipsum</a>.</p>
		</cls>
	</corpus>
</EasyLexML>`
	assert.Equal(t, expected, output.String())
}

func TestDraft2Strict_Doc2(t *testing.T) {
	input := bytes.NewBufferString(`
<EasyLexML>
	<set-meta TocTitle="𝕿𝖆𝖇𝖑𝖊 𝖔𝖋 𝕮𝖔𝖓𝖙𝖊𝖓𝖙𝖘"/>
	<set-meta AbstractTitle="𝕬𝖇𝖘𝖙𝖗𝖆𝖈𝖙"/>
	<corpus>
		<sec title="My Title">
			<cls>Lorem <a href="https://en.wikipedia.org/wiki/Lorem_ipsum">Ipsum</a>.</cls>
		</sec>
		<cls>Final topic! </cls>
	</corpus>
</EasyLexML>`)
	output := bytes.NewBuffer(nil)
	err := Draft2Strict(input, output)
	require.Nil(t, err)

	expected := `<?xml?>
<EasyLexML>
	<toc id="toc">
		<label href="#toc">
			𝕿𝖆𝖇𝖑𝖊 𝖔𝖋 𝕮𝖔𝖓𝖙𝖊𝖓𝖙𝖘
		</label>
		<ul>
			<li>
				<a href="#sec1">
					Section 1
					-
					My Title
				</a>
				<ul>
					<li>
						<a href="#sec1_cls1">
							Cls. 1
						</a>
					</li>
				</ul>
			</li>
			<li>
				<a href="#cls1">
					Cls. 2
				</a>
			</li>
		</ul>
	</toc>
	<corpus id="corpus">
		<sec id="sec1" num="1">
			<label href="#sec1">
				<span>
					Section 1
				</span>
				<span>
					My Title
				</span>
			</label>
			<cls id="sec1_cls1" num="1">
				<p id="sec1_cls1_p1">
					<label href="#sec1_cls1">
						Cls. 1
					</label>Lorem
					<a href="https://en.wikipedia.org/wiki/Lorem_ipsum">Ipsum</a>.</p>
			</cls>
		</sec>
		<cls id="cls1" num="2">
			<p id="cls1_p1">
				<label href="#cls1">
					Cls. 2
				</label>Final topic!
			</p>
		</cls>
	</corpus>
</EasyLexML>`
	assert.Equal(t, expected, output.String())
}
