package actions

import "github.com/gobuffalo/buffalo"

func AboutHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("about/index.html"))
}

func ProjectsHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("projects/index.html"))
}

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("home/index.html"))
}