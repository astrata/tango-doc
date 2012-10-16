package models

import (
	"fmt"
	"github.com/astrata/tango"
	"github.com/astrata/tango/app"
	"github.com/astrata/tango/body"
	"github.com/gosexy/sugar"
	"strings"
	"time"
)

type Phrase struct {
	Params tango.Value
}

func init() {
	app.Register("Phrase", &Phrase{})
	app.Fallback("/phrase", app.App("Phrase"))
}

func (self *Phrase) StartUp() {

}

/* Counts the number of A's in a phrase. */
func (self *Phrase) Count() body.Body {
	content := body.Json()

	input := self.Params.GetString("text")

	counter := strings.Count(strings.ToLower(input), "a")

	response := sugar.Tuple{
		"data": sugar.Tuple{
			"input":  input,
			"output": counter,
		},
		"time": time.Now(),
	}

	if counter > 0 {
		response["success"] = fmt.Sprintf("Has %d \"A\"s.", counter)
	} else {
		response["error"] = "Does not contain any \"A\"."
	}

	content.Set(response)

	return content
}

/* Separes every character in a phrase and makes an array with them. */
func (self *Phrase) Split() body.Body {
	content := body.Json()

	input := self.Params.GetString("text")

	content.Set(
		sugar.Tuple{
			"success": "OK",
			"data": sugar.Tuple{
				"input":  input,
				"output": strings.Split(input, ""),
			},
			"time": time.Now(),
		},
	)

	return content
}

/* Changes every character in a phrase to lowercase. */
func (self *Phrase) Lower() body.Body {
	content := body.Json()

	input := self.Params.GetString("text")

	content.Set(
		sugar.Tuple{
			"success": "OK",
			"data": sugar.Tuple{
				"input":  input,
				"output": strings.ToLower(input),
			},
			"time": time.Now(),
		},
	)

	return content
}

/* Changes every character in a phrase to uppercase. */
func (self *Phrase) Upper() body.Body {
	content := body.Json()

	input := self.Params.GetString("text")

	content.Set(
		sugar.Tuple{
			"success": "OK",
			"data": sugar.Tuple{
				"input":  input,
				"output": strings.ToUpper(input),
			},
			"time": time.Now(),
		},
	)

	return content
}
