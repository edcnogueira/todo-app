package handlers

import (
	"github.com/a-h/templ"
	"github.com/edcnogueira/todo-app/views"
	"github.com/edcnogueira/todo-app/views/auth"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
)

func HandleViewHome(ctx fiber.Ctx) error {
	hindex := views.HomeIndex(fromProtected)
	home := views.Home("", fromProtected, hindex)

	handler := adaptor.HTTPHandler(templ.Handler(home))

	return handler(ctx)
}

func HandleViewLogin(c fiber.Ctx) error {
	lindex := auth.LoginIndex(fromProtected)
	login := auth.Login(" | Login", fromProtected, lindex)

	handler := adaptor.HTTPHandler(templ.Handler(login))

	// if c.Method() == "POST" {
	// 	var (
	// 		user models.User
	// 		err  error
	// 	)
	// 	fm := fiber.Map{
	// 		"type": "error",
	// 	}

	// 	// notice: in production you should not inform the user
	// 	// with detailed messages about login failures
	// 	if user, err = models.CheckEmail(c.FormValue("email")); err != nil {
	// 		fm["message"] = "There is no user with that email"

	// 		return c.Redirect("/login")
	// 	}

	// 	err = bcrypt.CompareHashAndPassword(
	// 		[]byte(user.Password),
	// 		[]byte(c.FormValue("password")),
	// 	)
	// 	if err != nil {
	// 		fm["message"] = "Incorrect password"

	// 		return flash.WithError(c, fm).Redirect("/login")
	// 	}

	// 	session, err := store.Get(c)
	// 	if err != nil {
	// 		fm["message"] = fmt.Sprintf("something went wrong: %s", err)

	// 		return flash.WithError(c, fm).Redirect("/login")
	// 	}

	// 	session.Set(AUTH_KEY, true)
	// 	session.Set(USER_ID, user.ID)

	// 	err = session.Save()
	// 	if err != nil {
	// 		fm["message"] = fmt.Sprintf("something went wrong: %s", err)

	// 		return flash.WithError(c, fm).Redirect("/login")
	// 	}

	// 	fm = fiber.Map{
	// 		"type":    "success",
	// 		"message": "You have successfully logged in!!",
	// 	}

	// 	return flash.WithSuccess(c, fm).Redirect("/todo/list")
	// }

	return handler(c)
}
