package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kataras/iris/v12"
)


func main() {
	/* 
	app := iris.New()

	booksAPI := app.Party("/books")
	{
		booksAPI.Use(iris.Compression)

		// GET: http://localhost:8080/books
		booksAPI.Get("/", list)
		// POST: http://localhost:8080/books
		booksAPI.Post("/", create)
	}
	*/

    app := iris.Default()

    // This handler will match /user/john but will not match /user/ or /user
    app.Get("/user/{name}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        ctx.Writef("Hello %s", name)
    })

    // However, this one will match /user/john/ and also /user/john/send
    // If no other routers match /user/john, it will redirect to /user/john/
    app.Get("/user/{name}/{action:path}", func(ctx iris.Context) {
        name := ctx.Params().Get("name")
        action := ctx.Params().Get("action")
        message := name + " is " + action
        ctx.WriteString(message)
    })
	// custom
    app.Get("/welcome", custom_welcome)
	app.Post("/form_post", custom_form_post)
	app.Post("/get_json", custom_get_json)
	

    app.Listen(":8080")
}

// Book example.
type Book struct {
	Title string `json:"title"`
}
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func custom_get_json(ctx iris.Context)  {
	
	var person Person
	rawBodyAsBytes, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil { /* handle the error */ ctx.Writef("%v", err) }
	_ = json.Unmarshal(rawBodyAsBytes, &person)
	
	println(person.Age,person.Name) 
	/*
	err := ctx.ReadJSON(&persons)
	if err != nil {
		ctx.StopWithError(iris.StatusBadRequest, err)
		return
	}

	ctx.Writef("Received: %#+v\n", persons)	
	*/
}

func custom_form_post(ctx iris.Context) {
	message := ctx.PostValue("message")
	nick := ctx.PostValueDefault("nick", "anonymous")

	ctx.JSON(iris.Map{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

func custom_welcome(ctx iris.Context)  {
	
        firstname := ctx.URLParamDefault("firstname", "Guest")
        //lastname := ctx.URLParam("lastname") 
		// shortcut for ctx.Request().URL.Query().Get("lastname")
		lastname := ctx.Request().URL.Query().Get("lastname")
        ctx.Writef("Hello %s %s", firstname, lastname)
    
}

func list(ctx iris.Context) {
	books := []Book{
		{"Mastering Concurrency in Go"},
		{"Go Design Patterns"},
		{"Black Hat Go"},
	}

	ctx.JSON(books)
	// TIP: negotiate the response between server's prioritizes
	// and client's requirements, instead of ctx.JSON:
	// ctx.Negotiation().JSON().MsgPack().Protobuf()
	// ctx.Negotiate(books)
}

func create(ctx iris.Context) {
	var b Book
	err := ctx.ReadJSON(&b)
	// TIP: use ctx.ReadBody(&b) to bind
	// any type of incoming data instead.
	if err != nil {
		ctx.StopWithProblem(iris.StatusBadRequest, iris.NewProblem().
			Title("Book creation failure").DetailErr(err))
		// TIP: use ctx.StopWithError(code, err) when only
		// plain text responses are expected on errors.
		return
	}

	println("Received Book: " + b.Title)

	ctx.StatusCode(iris.StatusCreated)
}
