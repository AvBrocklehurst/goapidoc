package main

import (
	"html/template"
	"os"
)

const documentationHTMLTemplate = `
<!doctype html>

<html lang="en">

<head>
	<meta charset="utf-8">
	<title>Go API Docs</title>
	<meta name="description" content="Documentation produced by goapidoc">

	<script src="https://cdn.rawgit.com/google/code-prettify/master/loader/run_prettify.js"></script>

</head>

<body>
	<div id="sidebar">
		<h1>{{.Title}}</h1>

		<h3 class="underline">Services</h3>
		<ul>
			{{range $key, $value := .Endpoints}} 
				<a href="#{{$key}}" ><li>{{ $key }}</li></a>
			{{end}}
		</ul>
	</div>
	
	<div id="main">
		{{range $key, $value := .Endpoints}} 
			<div class="service" id="{{$key}}">
				<h1>{{ $key }}</h1>
				{{range $key, $value := $value}} 
					<div class="endpoint">
						<h3>{{$value.Name}}</h3>
						<code>{{$value.Method}} {{$value.Route}}</code>
						<p>{{$value.Description}}</p>
						
						{{ if $value.Returns }} 
							<h4>Returns:</h4>
							<pre class="prettyprint">{{$value.Returns}}</pre>
						{{ end }}

						{{ if $value.Params }} 
							<h4>Params:</h4>
							<table>
								<thead>
									<tr>
										<th>Name</th>
										<th>Type</th>
										<th>Location</th>
										<th>Description</th>
									</tr>
								</thead>
								<tbody>
									{{ range $param := $value.Params }} 
									<tr>
										<td>{{ $param.Name }}</td>
										<td>{{ $param.Type }}</td>
										<td>{{ $param.Location }}</td>
										<td>{{ $param.Description }}</td>
									</tr>
									{{ end }}
								</tbody>
							</table>
						{{ end }}
					</div>
				{{end}}
			</div>
		{{end}}
	</div>
</body>


<style>
	html, body {
		font-family: Helvetica,Arial,sans-serif;
		top: 0;
		left: 0;
		margin: 0;
		padding: 0;
		position: absolute;
	}

	pre.prettyprint {
		margin-left: 28px;
		max-width: 50%;
	}

	#main h1, #main h2, #main h3, #main h4, #main h5, #main h6, #main p, #main table, #main ul, #main ol, #main aside, #main dl {
		// margin-right: 50%;
		padding: 0 28px;
		// box-sizing: border-box;
		display: block;
	}

	#sidebar {
		position: fixed;
		height: 100vh;
		width: 20vw;
		padding: 20px;
		box-sizing: border-box;
		background: #2F3336;
		color: white;
	}

	h3.underline {
		border-bottom: 1px solid white;
    	padding-bottom: 2px;
	}

	#sidebar h1 {
		text-align: center;
	}

	a {
		color: white;
		text-decoration: none;
	}


	#sidebar ul {
		padding: 0;
    	list-style: none;
	}

	#sidebar ul li {
		padding-bottom:10px;
	}

	#main {
		margin-left: 20vw;
		width: 80vw;
	}

	#main code {
		background-color: rgba(0,0,0,0.05);
		padding: 3px;
		margin-left: 28px;
		border-radius: 3px;
	}

	#main .service:first-of-type h1 {
		border-top-width: 0;
		margin-top: 0;
	}

	#main h1 {
		font-size: 25px;
		padding-top: 0.5em;
		padding-bottom: 0.5em;
		margin-bottom: 0;
		margin-top: 0;
		border-top: 1px solid #ccc;
		border-bottom: 1px solid #ccc;
		background-color: #fdfdfd;
		color: #333;
	}

	#main table tr:last-child {
		border-bottom: 1px solid #ccc;
	}

	#main table th {
		padding: 5px 10px;
		border-bottom: 1px solid #ccc;
		vertical-align: bottom;
	}

	#main table tr:nth-child(even)>td {
		background-color: #fbfcfd;
	}
	
	#main table td {
		padding: 10px;
	}
	#main table th, #main table td {
		text-align: left;
		vertical-align: top;
		line-height: 1.6;
	}

	h1 {
		font-size: 24px;
		line-height: 32px;
	}

	#main .endpoint {
		border-top: 1px solid #ccc;
		background: #F4F7F9;
		padding-bottom: 20px;
	}

	#main .service>.endpoint:first-of-type() {
		border-top: 0px;
	}

	#main .endpoint h3 {
		margin: 0;
		padding-top: 18px;
		padding-bottom: 18px;

	}

	#main .endpoint p {
		margin-bottom: 0;
	}

	// .endpoint:nth-child(odd) {
		
	// }


</style>

</html>
`

type templateStruct struct {
	Title     string
	Endpoints map[string][]endpoint
}

func (fp *fileParser) createHTMLDocumentation(title string) (err error) {
	file, err := os.Create("documentation.html")
	if err != nil {
		return
	}
	defer file.Close()
	tmpl, err := template.New("documentation template").Parse(documentationHTMLTemplate)
	if err != nil {
		return
	}
	if len(title) < 1 {
		title = "API Documentation created by Go API Doc"
	}
	ts := templateStruct{
		Title:     title,
		Endpoints: fp.endpoints,
	}
	err = tmpl.Execute(file, ts)
	return
}
