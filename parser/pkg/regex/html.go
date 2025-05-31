package regex

import "regexp"


var (
	TitleRegex = regexp.MustCompile(`<title>(.*?)<\/title>`)
	DescriptionRegex = regexp.MustCompile(`<meta\s+name=["']description["']\s+content=["']([^"']*)["']`)
	TextRegex = regexp.MustCompile(`<p>(.*?)<\/p>|<h1>(.*?)<\/h1>|<h2>(.*?)<\/h2>|<h3>(.*?)<\/h3>|<h4>(.*?)<\/h4>|<h5>(.*?)<\/h5>|<h6>(.*?)<\/h6>`)

)