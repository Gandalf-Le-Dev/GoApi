package htmlFormatter

import (
	"fmt"
	"strings"

	"github.com/Gandalf-Le-Dev/goapi/models"
)

func postTemplate(post models.Post) string {
	return fmt.Sprintf(`
	<li id="post-%d">
		<div class="view">
			<h3>%s</h3>
			<p>%s</p>
			<button hx-get="/posts/%d/edit" hx-target="#post-%d .edit" hx-swap="innerHTML" hx-headers='{"Accept": "text/html"}'>Edit</button>
			<button hx-delete="/posts/%d" hx-target="#post-%d" hx-swap="outerHTML" hx-headers='{"Accept": "text/html"}' hx-confirm="Are you sure you wish to delete this post?">Delete</button>
		</div>
		<div class="edit" style="display:none;">
			<!-- Edit form will be loaded here -->
		</div>
	</li>`, post.PostID, post.Title, post.Content, post.PostID, post.PostID, post.PostID, post.PostID)
}

func Posts(posts models.PostSlice) string {
	var builder strings.Builder
	for _, post := range posts {
		builder.WriteString(postTemplate(*post))
	}
	return builder.String()
}

func Post(post models.Post) string {
	return postTemplate(post)
}

func EditPost(post models.Post) string {
	return fmt.Sprintf(`
	<form hx-put="/posts/%d" hx-target="#post-%d" hx-swap="outerHTML" hx-headers='{"Accept": "text/html"}'>
		<label for="title">Title:</label>
		<input type="text" id="title" name="title" value="%s">
		<br>
		<label for="content">Content:</label>
		<textarea id="content" name="content">%s</textarea>
		<br>
		<button type="submit">Save</button>
		<button type="button" hx-get="/posts/%d" hx-target="#post-%d" hx-swap="outerHTML" hx-headers='{"Accept": "text/html"}'>Cancel</button>
	</form>`, post.PostID, post.PostID, post.Title, post.Content, post.PostID, post.PostID)
}
