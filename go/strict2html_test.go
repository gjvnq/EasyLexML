package easyLexML

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStrict2HTML_EmptyDocument(t *testing.T) {
	input := bytes.NewBufferString(``)
	output := bytes.NewBuffer(nil)
	err := Strict2HTML(input, output)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "no <EasyLexML> found")
}

func TestStrict2HTML_MinimalDocument(t *testing.T) {
	input := bytes.NewBufferString(`<?xml?><EasyLexML><toc id="toc"><label href="#toc">Table of Contents</label><ul/></toc><corpus some-attr="my-value" id="corpus"/></EasyLexML>`)
	output := bytes.NewBuffer(nil)
	err := Strict2HTML(input, output)
	require.Nil(t, err)

	expected := `<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title></title>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/mathjax/2.7.5/latest.js?config=MML_SVG" async="1"></script>
		<style>
			@import url('https://fonts.googleapis.com/css?family=Libre+Baskerville:400,400i,700&subset=latin-ext');
			body {
				font-size: 18px;
				margin-top: 2rem;
				font-family: "Libre Baskerville", sans-serif;
				font-variant-numeric: tabular-nums;
				text-align: left;
				text-rendering: optimizeLegibility;
				hyphens: auto;
				line-height: 200%;
			}
			@media screen and (min-width: 36rem) {
				body {
					width: 36rem;
					margin-left: calc(50% - 36rem / 2);
				}
			}
			@media screen and (max-width: 36rem) {
				body {
					font-size: 14px;
				}
			}
			section#toc ul {
				list-style-type: none;
			}
			section#toc > ul {
				padding: 0;
			}
			h1, h2, h3, h4, h5, h6 {
				text-align: center;
			}
			a.label, a.label:visited {
				color: #000;
				text-decoration: none;
			}
			section#toc > a.label {
				padding-bottom: 0;
			}
			section > a {
				margin-top: 1rem;
				text-align: center;
				display: block;
				font-weight: bold;
				padding-top: 2rem;
				padding-bottom: 1rem;
			}
			section > a > span {
				display: block;
				margin-top: -0.5rem;
				margin-bottom: -0.5rem;
			}
			section > a > span:nth-child(1) {
				font-style: italic;
				font-weight: normal;
			}
			section > a > span:last-child {
				font-weight: bold;
				font-style: normal;
			}
			p {
				margin: 0;
			}
			section.sub > section.sub {
				padding-left: 1.5rem;
			}
			section.note {
		    border: 1px solid #aaa;
		    border-radius: 0.5rem;
		    padding: 0px 10px;
			}
			section.note, section.note a.label {
		    color: #aaa;
			}
			section.note a {
		    color: #00a;
			}
			hr {
				margin-top: 1rem;
				margin-bottom: 1rem;
			}
			:target {
				animation-name: highlight;
				animation-duration: 2s;
			}
			@keyframes highlight {
				0% {background-color: #ffa;}
				75% {background-color: #ffa;}
				100% {background-color: transparent;}
			}
			a, a:visited {
				overflow-wrap: break-word;
				color: blue;
			}
			a:hover, a:active {
				text-decoration: underline;
			}
		</style>
	</head>
	<body>
		<header id="header">
			<h1></h1>
		</header>
		<section id="metadata">
		</section>
		<section id="abstract">
			<h1></h1>
			
		</section>
		<section id="toc" data-tag="toc">
	<a href="#toc" class="label" data-tag="label">Table of Contents</a><ul/></section>
		<section data-some-attr="my-value" id="corpus" data-tag="corpus"/>
	</body>
</html>
`
	assert.Equal(t, expected, output.String())
}
