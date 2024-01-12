package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length int
}

type Tracks []Track

func (x Tracks) Len() int           { return len(x) }
func (x Tracks) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x Tracks) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

var trackList = template.Must(template.New("trackList").Parse(`
<table>
<tr style='text-align: left'>
  <th>Title</th>
  <th>Artist</th>
  <th>Album</th>
  <th>Year</th>
  <th>Length</th>
</tr>
{{range .}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

var tracks = Tracks{
	{"Go", "Delilah", "From the Roots Up", 2012, 183},
	{"Go", "Moby", "Moby", 1992, 377},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, 276},
	// 更多的Track...
}

func tracksHandler(w http.ResponseWriter, r *http.Request) {
	sort.Sort(tracks)
	trackList.Execute(w, tracks)
}

func main() {
	http.HandleFunc("/tracks", tracksHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}
