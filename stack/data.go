// Code generated by regen.go. DO NOT EDIT.

package stack

import (
	"html/template"
)

const indexHTML = "<!DOCTYPE html>\n{{- /* Join a list */ -}}\n{{- define \"Join\" -}}\n{{- if . -}}\n{{- $l := len . -}}\n{{- $last := minus $l 1 -}}\n{{- range $i, $e := . -}}\n{{- $e -}}\n{{- $isNotLast := ne $i $last -}}\n{{- if $isNotLast}}, {{end -}}\n{{- end -}}\n{{- end -}}\n{{- end -}}\n{{- /* Accepts a Args */ -}}\n{{- define \"RenderArgs\" -}}\n<span class=\"args\"><span>\n{{- $elided := .Elided -}}\n{{- if .Processed -}}\n{{- $l := len .Processed -}}\n{{- $last := minus $l 1 -}}\n{{- range $i, $e := .Processed -}}\n{{- $e -}}\n{{- $isNotLast := ne $i $last -}}\n{{- if or $elided $isNotLast}}, {{end -}}\n{{- end -}}\n{{- else -}}\n{{- $l := len .Values -}}\n{{- $last := minus $l 1 -}}\n{{- range $i, $e := .Values -}}\n{{- $e.String -}}\n{{- $isNotLast := ne $i $last -}}\n{{- if or $elided $isNotLast}}, {{end -}}\n{{- end -}}\n{{- end -}}\n{{- if $elided}}…{{end -}}\n</span></span>\n{{- end -}}\n{{- /* Accepts a Call */ -}}\n{{- define \"RenderCreatedBy\" -}}\n<span class=\"call hastooltip\"><span class=\"tooltip\">\n{{- if and .LocalSrcPath (ne .RemoteSrcPath .LocalSrcPath) -}}\nRemoteSrcPath: {{.RemoteSrcPath}}\n<br>LocalSrcPath: {{.LocalSrcPath}}\n{{- else -}}\nSrcPath: {{.RemoteSrcPath}}\n{{- end -}}\n<br>Func: {{.Func.Complete}}\n<br>Location: {{.Location}}\n</span><a href=\"{{srcURL .}}\">{{.SrcName}}:{{.Line}}</a> <span class=\"{{funcClass .}}\">\n<a href=\"{{pkgURL .}}\">{{.Func.DirName}}.{{.Func.Name}}</a></span>()\n</span>\n{{- end -}}\n{{- /* Accepts a Stack */ -}}\n{{- define \"RenderCalls\" -}}\n<table class=\"stack\">\n{{- range $i, $e := .Calls -}}\n<tr>\n<td>{{$i}}</td>\n<td>\n<a href=\"{{pkgURL $e}}\">{{$e.Func.DirName}}</a>\n</td>\n<td class=\"hastooltip\">\n<span class=\"tooltip\">\n{{- if and $e.LocalSrcPath (ne $e.RemoteSrcPath $e.LocalSrcPath) -}}\nRemoteSrcPath: {{$e.RemoteSrcPath}}\n<br>LocalSrcPath: {{$e.LocalSrcPath}}\n{{- else -}}\nSrcPath: {{$e.RemoteSrcPath}}\n{{- end -}}\n<br>Func: {{$e.Func.Complete}}\n<br>Location: {{$e.Location}}\n</span>\n<a href=\"{{srcURL $e}}\">{{$e.SrcName}}:{{$e.Line}}</a>\n</td>\n<td>\n<span class=\"{{funcClass $e}}\"><a href=\"{{pkgURL $e}}\">{{$e.Func.Name}}</a></span>({{template \"RenderArgs\" $e.Args}})\n</td>\n</tr>\n{{- end -}}\n{{- if .Elided}}<tr><td>(…)</td><tr>{{end -}}\n</table>\n{{- end -}}\n<meta charset=\"UTF-8\">\n<meta name=\"author\" content=\"Marc-Antoine Ruel\" >\n<meta name=\"generator\" content=\"https://github.com/maruel/panicparse\" >\n<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n<title>PanicParse</title>\n<link rel=\"shortcut icon\" type=\"image/gif\" href=\"data:image/gif;base64,{{.Favicon}}\"/>\n<style>\n{{- /* Minimal CSS reset */ -}}\n* {\nfont-family: inherit;\nfont-size: 1em;\nmargin: 0;\npadding: 0;\n}\nhtml {\nbox-sizing: border-box;\nfont-size: 62.5%;\n}\n*, *:before, *:after {\nbox-sizing: inherit;\n}\nh1, h2 {\nmargin-bottom: 0.2em;\nmargin-top: 0.8em;\n}\nh1 {\nfont-size: 1.4em;\n}\nh2 {\nfont-size: 1.2em;\n}\nbody {\nfont-size: 1.6em;\nmargin: 2px;\n}\nli {\nmargin-left: 2.5em;\n}\na {\ncolor: inherit;\ntext-decoration: inherit;\n}\nol, ul {\nmargin-bottom: 0.5em;\nmargin-top: 0.5em;\n}\np {\nmargin-bottom: 2em;\n}\ntable {\nmargin: 0.6em;\n}\ntable tr:nth-child(odd) {\nbackground-color: #F0F0F0;\n}\ntable tr:hover {\nbackground-color: #DDD !important;\n}\ntable td {\nfont-family: monospace;\npadding: 0.2em 0.4em 0.2em;\n}\n.call {\nfont-family: monospace;\n}\n@media screen and (max-width: 500px) {\nh1 {\nfont-size: 1.3em;\n}\n}\n@media screen and (max-width: 500px) and (orientation: portrait) {\n.args span {\ndisplay: none;\n}\n.args::after {\ncontent: '…';\n}\n}\n.created {\nwhite-space: nowrap;\n}\n.race {\nfont-weight: 700;\ncolor: #600;\n}\n#content {\nwidth: 100%;\n}\n.hastooltip:hover .tooltip {\nbackground: #fffAF0;\nborder: 1px solid #DCA;\nborder-radius: 6px;\nbox-shadow: 5px 5px 8px #CCC;\ncolor: #111;\ndisplay: inline;\nposition: absolute;\n}\n.tooltip {\ndisplay: none;\nline-height: 16px;\nmargin-left: 1rem;\nmargin-top: 2.5rem;\npadding: 1rem;\nz-index: 10;\n}\n.bottom-padding {\nmargin-top: 5em;\n}\n{{- /* Highlights based on stack.Location value. */ -}}\n.FuncMain {\ncolor: #880;\n}\n.FuncLocationUnknown {\ncolor: #888;\n}\n.FuncGoMod {\ncolor: #800;\n}\n.FuncGOPATH {\ncolor: #109090;\n}\n.FuncGoPkg {\ncolor: #008;\n}\n.FuncStdlib {\ncolor: #080;\n}\n.Exported {\nfont-weight: 700;\n}\n</style>\n<div id=\"content\">\n{{- if .Aggregated -}}\n{{- range $i, $e := .Aggregated.Buckets -}}\n{{$l := len $e.IDs}}\n<h1>Signature #{{$i}}: {{$l}} routine{{if ne 1 $l}}s{{end}}: <span class=\"state\">{{$e.State}}</span>\n{{- if $e.SleepMax -}}\n{{- if ne $e.SleepMin $e.SleepMax}} <span class=\"sleep\">[{{$e.SleepMin}}~{{$e.SleepMax}} mins]</span>\n{{- else}} <span class=\"sleep\">[{{$e.SleepMax}} mins]</span>\n{{- end -}}\n{{- end -}}\n</h1>\n{{if $e.Locked}} <span class=\"locked\">[locked]</span>\n{{- end -}}\n{{- if $e.CreatedBy.Calls}} <span class=\"created\">Created by: {{template \"RenderCreatedBy\" index $e.CreatedBy.Calls 0}}</span>\n{{- end -}}\n{{template \"RenderCalls\" $e.Signature.Stack}}\n{{- end -}}\n{{- else -}}\n{{- range $i, $e := .Snapshot.Goroutines -}}\n<h1>Routine {{$e.ID}}: <span class=\"state\">{{$e.State}}</span>\n{{- if $e.SleepMax -}}\n{{- if ne $e.SleepMin $e.SleepMax}} <span class=\"sleep\">[{{$e.SleepMin}}~{{$e.SleepMax}} mins]</span>\n{{- else}} <span class=\"sleep\">[{{$e.SleepMax}} mins]</span>\n{{- end -}}\n{{- end -}}\n</h1>\n{{if $e.Locked}} <span class=\"locked\">[locked]</span>\n{{- end -}}\n{{if $e.RaceAddr}} <span class=\"race\">Race {{if $e.RaceWrite}}write{{else}}read{{end}} @ {{$e.RaceAddr}}</span><br>\n{{- end -}}\n{{- if $e.CreatedBy.Calls}} <span class=\"created\">Created by: {{template \"RenderCreatedBy\" index $e.CreatedBy.Calls 0}}</span>\n{{- end -}}\n{{template \"RenderCalls\" $e.Signature.Stack}}\n{{- end -}}\n{{- end -}}\n</div>\n<h2>Metadata</h2>\n<ul>\n<li>Created on {{.Now.String}}</li>\n<li>{{.Version}}</li>\n{{- if and .Snapshot.LocalGOROOT (ne .Snapshot.RemoteGOROOT .Snapshot.LocalGOROOT) -}}\n<li>GOROOT (remote): {{.Snapshot.RemoteGOROOT}}</li>\n<li>GOROOT (local): {{.Snapshot.LocalGOROOT}}</li>\n{{- else -}}\n<li>GOROOT: {{.Snapshot.RemoteGOROOT}}</li>\n{{- end -}}\n<li>GOPATH: {{template \"Join\" .Snapshot.LocalGOPATHs}}</li>\n{{- if .Snapshot.LocalGomods -}}\n<li>go modules (local):\n<ul>\n{{- range $path, $import := .Snapshot.LocalGomods -}}\n<li>{{$path}}: {{$import}}</li>\n{{- end -}}\n</ul>\n</li>\n{{- end -}}\n<li>GOMAXPROCS: {{.GOMAXPROCS}}</li>\n</ul>\n<h2>Legend</h2>\n<table class=\"legend\">\n<thead>\n<th>Type</th>\n<th>Exported</th>\n<th>Private</th>\n</thead>\n<tr class=\"call hastooltip\">\n<td>\nPackage main\n<span class=\"tooltip\">Sources that are in the main package.</span>\n</td>\n<td class=\"FuncMain\">main.Foo()</td>\n<td class=\"FuncMain\">main.foo()</td>\n</tr>\n<tr class=\"call hastooltip\">\n<td>\nGo module\n<span class=\"tooltip\">Sources located inside a directory containing a\n<strong>go.mod</strong> file but outside $GOPATH.</span>\n</td>\n<td class=\"FuncGoMod Exported\">pkg.Foo()</td>\n<td class=\"FuncGoMod\">pkg.foo()</td>\n</tr>\n<tr class=\"call hastooltip\">\n<td>\n$GOPATH/src/...\n<span class=\"tooltip\">Sources located inside the traditional $GOPATH/src\ndirectory.</span>\n</td>\n<td class=\"FuncGOPATH Exported\">pkg.Foo()</td>\n<td class=\"FuncGOPATH\">pkg.foo()</td>\n</tr>\n<tr class=\"call hastooltip\">\n<td>\n$GOPATH/pkg/mod/...\n<span class=\"tooltip\">Sources located inside the go module dependency\ncache under $GOPATH/pkg/mod. These files are unmodified third parties.</span>\n</td>\n<td class=\"FuncGoPkg Exported\">pkg.Foo()</td>\n<td class=\"FuncGoPkg\">pkg.foo()</td>\n</tr>\n<tr class=\"call hastooltip\">\n<td>\nStandard library\n<span class=\"tooltip\">Sources from the Go standard library under\n$GOROOT/src/.</span>\n</td>\n<td class=\"FuncStdlib Exported\">pkg.Foo()</td>\n<td class=\"FuncStdlib\">pkg.foo()</td>\n</tr>\n<tr class=\"call hastooltip\">\n<td>\nUnknown source location\n<span class=\"tooltip\">Sources which location was not successfully\ndetermined.</span>\n</td>\n<td class=\"FuncLocationUnknown Exported\">pkg.Foo()</td>\n<td class=\"FuncLocationUnknown\">pkg.foo()</td>\n</tr>\n</table>\n{{- .Footer -}}\n{{- /* Add unnecessary bottom spacing so the last tooltip from the legend is visible. */ -}}\n<div class=\"bottom-padding\"></div>\n"

// favicon is the bomb emoji U+1F4A3 in Noto Emoji as a 128x128 base64 encoded
// PNG.
//
// See README.md for license and how to retrieve it.
const favicon template.HTML = "R0lGODlhQABAAMZhACEhISIiIiMjIyQkJCUlJSYmJicnJygoKCkpKSoqKisrKywsLC0tLS4uLi8vLzAwMDExMTIyMjMzMzQ0NDU1NTY2Njc3Nzg4ODk5OTo6Ojs7Ozw8PD09PT4+Pj8/P0BAQEFBQUJCQkNDQ0REREVFRUZGRkdHR4o1MEhISElJSUpKSktLS0xMTE1NTU5OTk9PT1BQUFFRUVJSUv8fQFNTU1RUVKc+N1VVVVZWVldXV/8mPVhYWFlZWf8oPP8pO1paWltbW1xcXF1dXV5eXl9fX2BgYGFhYf80NWJiYv8+MP9ALvRDNv9JKcxYUfRPQ/9PJvRTSP9VI/9aIPRcUf9dHudhV/9lGv91EP91Ef96Dv99C/9+DP+FB/+MA/+PAf+QAP+RAP///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////yH5BAEKAH8ALAAAAABAAEAAAAf+gH+Cg4SFhoeIiYqLjI2Oj5CRjSkjJiqSkl6YfyMXCQCgBRUkm45ahF+NJRIBAAEKEhIKAAIepYVgg0uCVINXjCocAwAEGSWEHwIBI7eoXIRHgklSiyQPrhcohx0AEM2FVkqFM4sfBQALIokqBAKX34PijigVrhYpiwwAzN+7g1NOmqzrYACAgRCNFgAwAa8QlCUnQLwTpGKEhgP2tAEbQGBiMyxPdPypYgPUAAYRJDA450oCv0YfAESAl6tQig4MBIACFYCBhmOKVqh4p8IBAFsNB/kgVBEEiBEeE7loYSLF0BQYJEQt5e8PE0E9MMmAocGAAQgdSgwdmvRPlkH+USThoKFhgQIEBtopwAB1a6kug7ZAAsLjRgYIDxowuFtgQCutDWtCEkJZSI4NFShIiJB4AV7HAVD4bWuoiGkiOWSo6KAhAwYLmiE4WJAgbwAPo0kLKmKkSA4VIZx++OCBwwbXFiZEcMCg9oAJudvyNlKDxIgRJLKPEBH8Q4fjsJc3N4CArW5CRZD8MEGihAkUKFLAN1GChAgQxTNcoLDccwES0X1jBBIwkFCVCiuwoKCCQqFQwgghfMCBfvzNdgAHQp33R28rlCAaCy68AMOIMLzwQgssqIACCRFOmBwEDSSAQF+6mZaChyu08EIMMtBQw480yBDDiSus2CIGFUj+8IBnBogQoCS83YjCCi7AQMMNOOSww5Y54FDDWC0UyaIHG+wHY20EAJgUEUawQMKULsRAAw47/ABEEHgStsMNNMDQggomjABCBxlYIIGFBBwgWkNF3CCCJS1YiQMPQAhBBBFFEDGEEEDsMJefKpQgApkXTAABAwgUEIBWTzKyqQhvsvDCnDwEMcR0RuRaxBCd3iDDC0UKygGSEcRogE4SNRMEES2IoFYLMdywAxBD5GqttbsCkUMNfqbAIqGGOqDAAcMkYN4mnH7w6AovyIDDD0Lgeq2uQhQmgwuAjqpBckuSCwputwjhwgcjoMCClTkAwea88w7xw1zAmiDCB2X+TvAAqgS00sC5kVT6AQixwlDDDssyPC8R2tLwAgsojEBxqacikDEoam3yww0cgJzCwSMHIa/JKOfQJ8sul0lBzDMDsEGrh/zQQs4hj6ywydcGrXKwL5uKcSvdMG3IDyNsQLDBsyZcLdW5OozDrytInPWSMnNtAMeO5MDDBmKv64K7PMSLdhFBFBYDvqJ6sK+SC/i706KRTGqBBh6E8GwMUt9qchFC/LBtt99mkKS4Bgyzk5qR3HBDBRlATbYM0lL7MxGZ73lv2y4Py5+xOu2UbCRfUnBB3s9a2bqtl2oaxA+yr3xjCIaXenHcOwGwOyQ/SmBBBh08BanwdtsJBBD+P/CQg68nAlq754cmLvpOTkoiQw0S+L5B5G/mCMP7WOagPw43fFl+oIPSAMxwFz0A1EwSNDAU6jjwgRC8SQU6uh8NJhikGMDABSliD/M2gKT0HSBpPKGbI2TAAQhM4HoMdGBVchQiE5kIgytIQaA2WCiLjSd3O1GA1wpRIAdEgAIojNwIPGSVFSTIiCqQ4YPwAx5TNWBcBCggACqwQ0JgcAE+BGIG5gcCEWDHPfBBAX2sE6Hv1LA/qcLhTj5QxUHISiVZtAAGNsABD4AgBCLw4nXyCALicEADr1HObNIoRQJYpRQtAEECGODDCVTgAlvkQAc8MJzheMADHfijfir+YKpBhk6KU2wjIYyogARgEQLxs8AFMJABDeDtla3BwAU4qaQGMGl90QuAJZqxAoIgQAENeAAqHWkBVV7gmBewQGYmIAHZLCZVuIweFb8xlJUcwJQNcAAEUjKBblKgm7HgjANsmYADFECN0SMA45pREY6Qx5QMyOYDhAkBxDxgnAygzQHygs7oAawhKthAANx5AAQkQAELWAADFprQUiJgn+0A5U6oKEpEDIUCrhDAAApgloIi4KMPNUABCDAAZUgUFBA4ZFsuupMAaJQjBIgpSUvKtZN2Q6WkGcoFasqTAPiUpzYFAAVwqpuhmCOoQR1AB0SYUxVVAKhIBUUE1FI6UUkMhQQV6KdEV+WkqpZiKCgoYRRBOYAHcMASTNWQIdYSKhBw4K0MBNBa1GpVtrJ1BXTNq173qohAAAA7"
