package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Redirect(200, "/items")
}

func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		//if uid := c.Session().Get("current_user_id"); uid == nil {
		//	c.Flash().Add("danger", "You must be authorized to see that page")
		//	return c.Redirect(302, "/items")
		//}
		c.Flash().Add("danger", "You must be authorized to see that page")
		return c.Redirect(302, "/items")
		//return next(c)
	}
}
