# How to dance with *Tango!*

*Tango!* is a stateless API framework for creating web services with the [Go][2] programming language.

## Getting Go

In order to use *Tango!* you'll need [Go][2].

### Installing Go for OSX

1. Download the apropriate [installer][6] package.
2. Open a terminal and create a `$HOME/go` directory, `cd` to this directory and create three more empty subdirectories: `pkg`, `bin` and `src`.
3. Add `export GOPATH=$HOME/go` to the `$HOME/.profile` (`.zshrc` if you're a zsh user). If the file does not exist, create it.
4. Close the terminal.

### Installing Go for Linux

1. Choose your installation [package][4].
2. Uncompress it to `/usr/local/go`
3. Create a `$HOME/go` directory and create three empty subdirectories inside: `pkg`, `bin` and `src`.
4. Add `$HOME/go/bin` to your `$PATH` (add `PATH=$PATH:$HOME/go/bin` to your `.bashrc`, `.zshrc` or `.profile`).
5. Set `$GOPATH` to `$HOME/go` (add `GOPATH=$HOME/go` to your `.bashrc`, `.zshrc` or `.profile`).
6. Close the terminal.

### Installing Go for Windows

Please refer to the [installation guide][3] if you need more help.

## Getting *Tango!*

Now that you have [Go][2] you can get *Tango!*.

```
go get github.com/astrata/tango
go install github.com/astrata/tango/cmd/tango
```

## The *Tango!* command line

Run `tango` with no arguments to get a list of commands

```
% tango

Tango! (Prelude) - by Astrata

Usage:

  tango [command] [arguments]

The commands are:

  version    Returns current Tango! version.
  workspace  Returns current Tango! app working directory.
  run        Starts a Tango! app.
  help       Shows a short description of each Tango! subcommand.
  test       Runs the standard "go test" on the current Tango! app.
  init       Creates a Tango! basic app in the current directory.
  build      Compiles a Tango! app into a static binary.

Use "tango help [command]" for more information about a command.

```

If you want to see help on a particular topic run `tango help [topic]`:

```
% tango help build

Tango! (Prelude) - by Astrata

Description:

  Compiles a Tango! app into a static binary.

Usage:

  tango build <output>

```

## My first *Tango!* project

Create an empty working directory for your project, let's use `~/projects/tango-example`

```
% mkdir -p ~/projects/tango-example
```

Use `tango init` to initialize a project inside your recently created directory.

```
% cd ~/projects/tango-example
% tango init
Initializing Tango! workspace in /home/xiam/projects/tango-example.
```

A new workspace will be created with the following structure:

* `src` directory that contains your app's source code.
* `static` directory that routes everything that does not have a defined *route*.
* `config` directory that contains the settings.yaml file.
* `main.go` for running your app.

```
% find .
.
./main.go
./src
./src/models
./src/models/static.go
./static
./static/.empty
./.tango
./.tango/.empty
./config
./config/settings.yaml
```

Run your example project with `tango run`

```
% tango run
2012/10/15 15:38:21 Tango! by Astrata

2012/10/15 15:38:21 Initializing server...
2012/10/15 15:38:21 Adding fallback: /

2012/10/15 15:38:21 [::]:8080 is ready to dance.
2012/10/15 15:38:21 Stop server with ^C.

```

This will create an standalone server (by default) listening on any interface on the `8080` port.

Open a browser to http://localhost:8080 to see your first *Tango!* app :-).

## Models

A model represents the smallest part of an application, it usally provides services around a single entity of data (a SQL table or a NoSQL collection).

You're the developer and it is up to you to define the exact extend of an model and its interactions with other models.

### *Tango!* models

It is recommended that you put your models in the `src/models/` subdirectory of your project with a descriptive name, for example, if you want a model to be named `User` then it is reasonable to call the source file `user.go`.

Tango models must implement the `app.Model` interface.

```go
# file app/app.go
type Model interface {
  StartUp()
}
```

Besides having at least the `StartUp()` method, all models source files must have an `init()` function that registers the model and associates it with a particular URL.

```go
func init() {
  app.Register("Hello", &Hello{})
  app.Route("/", app.App("Hello"))
}

```

Finally, models belong to the "models" package.

```go
package models
```

### An example model for *Tango!*

This model will echo a "Hello, world!\n" string when an user visits the "/hello" URL.

```go
# Models belong to the "models" package
package models

import (
  "github.com/astrata/tango/app"
)

# Model structure
type Hello struct {
}

func init() {
  # Registering a model.
  app.Register("Hello", &Hello{})
  # Routing this model to the /hello URL.
  app.Route("/hello", app.App("Hello"))
}

func (self *Hello) StartUp() {
  # Models must implements this function for self initialization.
}

func (self *Hello) Index() string {
  return "Hello, world!\n"
}
```

Try to save it as `src/models/hello.go`, run the `tango run` command and visit http://locahost:8080/hello.

## Routes

A route is a HTTP query path, they are utilized to delegate client queries to specific modules.

For example, consider the string `http://example.org/foo/bar/baz?q=1`, the route is `/foo/bar/baz`.

### Defining a route

Routes are defined from within the `init()` function of each module.

This is how you could setup a route that would delegate every request to `/user` to the `&User{}` module.

```go
// Initialization function on user.go
func init() {
  // Registering the module
  app.Register("User", &User{})

  // Routing /user -> &User{}
  app.Route("/user", app.App("User"))
}
```

### Understanding routes

Routes are split into chunks separated by the `/` character.

The first chunk after the matching route is considered to be the method of `&User{}` that is going to be called,
remaining chunks are threated as arguments for this method, for example `/user/get/1/2/3` would call
`func (*User) Get(int, int, int)`.

Data type conversions (if required) are handled by *Tango!*.

#### Notes

* If there are no chunks after the matching main route, the `Index()` method is tried.
* If no method matches the request, the `CatchAll()` method is tried, if this last method does not exists
a `404` error is sent back to the user.
* A method name is always lowercased and camelized before searching for a match, for example, `/users/my-profile` would
look for `(*User)MyProfile()`.
* Only letters (`a-z`), numbers (`0-9`) and some punctuation symbols (`_-.`) are accepted. The symbols `-`, `.` and `_` are considered
to be word delimitators for camelization.
* As you would expect, *Tango!* can't serve private methods (those that begin with a lower case letter). The `StartUp()` method is an exception.
It cannot be called remotely, no matter what.

### Fallback routes

A fallback route is an special case of route, this route is only probed at the end if no other route has precedence. This is the kind of route
the Static module uses.

```go
init() {
  // Registering the module.
  app.Register("Static", &Static{})

  // Adding a fallback route.
  app.Fallback("/", app.App("Static"))
}
```

## HTTP Parameters

There are two optional structures that you can define in your model if you want to access `GET` and `POST` values.

```go
# import "github.com/astrata/tango"

struct User {
  Get tango.Value
  Post tango.Value
}
```

And there's a third one that will return `POST` or `GET` values (if the former is `nil`).

```go
# import "github.com/astrata/tango"

struct User {
  Params tango.Value
}
```

For example, let's consider the `(*User)Add()`, this is how you could retrieve values from `POST` requests:

```go
func (self *User) Add() string {
  email := self.Post.String("email")
  return fmt.Sprintf("Hello %s!\n", email)
}
```

## Data validation

Given that you can use HTTP `GET` and `POST` variables in your model you may want to add restrictions
on the data the user sends.

*Tango!* provides helper methods for filtering and validating input values.

### Filtering input variables

Use the `(*tango.Value) Filter(args ...string) *tango.Value` method for accepting only
a reduced set of variables as application parameters.

```go
func (self *User) Add() string {

  // the "post" variable will be a tango.Value that only knows "email" and "name".
  post := self.Post.Filter(
    "email",
    "name",
  )

  email := post.String("email")
  name  := post.String("name")

  return fmt.Sprintf("Hello %s <%s>!\n", name, email)
}
```
### Required parameters

If you need the user to send not empty variables then you should use `tango.Value.Require(names ...string) (bool, map[string][]string)`.

This method will check that every given parameter is not empty, otherwise it will return a `false` boolean value and a map of error messages
that you can redirect to the client.

```go
# import "github.com/astrata/tango/validation"
# import "github.com/astrata/tango/body"

func (self *Data) Required() body.Body {

  response := body.Json()

  content := sugar.Tuple{}

  params := self.Params

  var isValid bool
  var messages map[string][]string

  isValid, messages = params.Require("email", "name", "last_name")

  if isValid == false {
    content["error"] = "Missing required data."
    content["data"] = messages
  }

  isValid, messages = self.Rules.Validate(params)

  if isValid == true {
    content["success"] = "Thank you."
  } else {
    content["error"] = "Some errors were found."
    content["data"] = messages
  }

  response.Set(content)

  return response
}
```

### Input validation

You can define rules for input validation, just add a `*validation.Rules` property in your model.

```go
# import "github.com/astrata/tango/validation"

type User struct {
  ...
  Rules *validation.Rules
  ...
}
```

And use the `StartUp()` method to define each rule.

```go
# import "github.com/astrata/tango/validation"

func (self *User) StartUp() {
  self.Rules = validation.New()

  self.Rules.Add("email", validation.Email, "Please enter a valid e-mail address.")
  self.Rules.Add("phone", validation.Numeric, "Please enter some numbers.")
}
```

Now you can validate all the data the client sends by calling the `(*validation.Rules) Validate(params tango.Value)` method.

```go
# import "github.com/astrata/tango/validation"

func (self *User) Validate() body.Body {

  response := body.Json()

  content := sugar.Tuple{}

  params := self.Params

  isValid, messages := self.Rules.Validate(params)

  if isValid == true {
    content["success"] = "Thank you."
  } else {
    content["error"] = "Some errors were found."
    content["data"] = messages
  }

  response.Set(content)

  return response
}
```

## Return values

There is a special `struct` that can define the way *Tango!* replies client requests.

```go
# from github.com/astrata/tango/body
type Body interface {
  Header() http.Header
  Status() int
  Get() []byte
  Set(interface{})
}
```

And there are some predefined structs that can be used as `body.Body` interfaces: `body.Html()`, `body.Json()` and `body.File()`.

This is how you could serve a file from disk.

```go
func (self *User) Icon() body.Body {
  content := body.File()

  // Pass the path to the file.
  content.Set("static/icon.png")

  return content
}
```

Some JSON

```go
func (self *User) Icon() body.Body {
  content := body.Json()

  response := sugar.Tuple{
    "success": "OK",
  }

  // Set the whole JSON object
  content.Set(response)

  return content
}
```

And standard HTML

```go
func (self *User) Icon() body.Body {
  content := body.Html()

  response := "<h1>Hello world!</h1>"

  content.Set(response)

  return content
}
```

Another way of returning HTML is by setting `string` as the return value of your model method.

It is not recommended to output HTML directly from a generated string, it would be better to use a templating package like `html/template`.

Finally, if *Tango!* does not know how to handle a return value, it will try to send it as JSON.

## Serving static files

Use the `static.go` model that comes by default and put your files into the `static` subdirectory, if a route is not taken if will
fallback to the file.

## Testing

One of the best features of [Go][2] is its [testing package][1].

*Tango!* has a wrapper for `go test` that helps on writing tests your apps.

If you want to run tests again your API, write a `main_test.go` file in your app root:

```go
# main_test.go

package main

import (
  # Sugar datatypes
	"github.com/gosexy/sugar"
  # This package handler REST the easy way.
	"github.com/gosexy/rest"
  # The Go testing package
	"testing"
)

var service = rest.New("http://localhost:8080")

func init() {
	rest.Debug = true
}

func TestClientModel(t *testing.T) {

	var result sugar.Tuple

	result, _ = service.Get("/client/add", nil).Post(sugar.Tuple{
		"email":    "test@example.com",
		"password": "password",
	}).Json()

	if result["success"] == nil {
		t.Errorf("Test failed.")
	}

}
```

You can run this tests using `tango test` from any subdirectory inside your *Tango!* app.

Please note that `tango test` does not automatically starta a *Tango!* server (yet), this instance needs
to be started manually.

## Settings

Settings for your app are stored in the `conf/settings.yaml` file.

YAML is a really simple markup language that stores settings in human-readable plain text.

```yaml
## Server configuration
server:
  bind: 0.0.0.0     # Listen on all interfaces.
  port: 8080        # Listen on port 8080.
```

If you're in the need of editing some settings please refer to the `config/settings.yaml` file and look for the
configuration you would like to tweak.

### Note

YAML uses two-spaces for indenting, if you use tabs the YAML parser will complain. Please use two-spaces to indent YAML files.

## Saving your own settings

Just write them down in the `conf/settings.yaml` file, use the `github.com/astrata/tango/config` package to retrieve them.

### Environments

You can use many different settings files to customize special environments, such as testing or production.

Use the `-settings` flag of `tango run` to make the app read settings from another file.

```
tango run -settings=conf/production.yaml
```

## Further reading

* [How to write Go code, testing][5] Covers how to make testing with Go's testing package.

[1]: http://golang.org/pkg/testing/
[2]: http://golang.org
[3]: http://golang.org/doc/install
[4]: http://code.google.com/p/go/downloads/list
[5]: http://golang.org/doc/code.html#Testing
[6]: http://code.google.com/p/go/downloads/list?q=OpSys-OSX+Type-Installer

