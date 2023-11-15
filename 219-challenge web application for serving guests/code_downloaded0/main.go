package main
import (
	"net/http"
)

func main() {
	
}

// indexHandler serves the main page
func indexHandler(w http.ResponseWriter, req *http.Request) {
	
}

// addHandler add a name to the names list
func addHandler(w http.ResponseWriter, req *http.Request) {
	
}

var indexHTML = `
<!DOCTYPE html>
<html>
    <head>
		<title>Guest Book ::Web GUI</title>
    </head>
    <body>
		<h1>Guest Book :: Web GUI</h1>
		<form action="/add" method="post">
		Name: <input name="name" /><submit value="Sign Guest Book">
		</form>
		<hr />
		<h4>Previous Guests</h4>
		<ul>
			{{range .}}
			<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
</html>
`