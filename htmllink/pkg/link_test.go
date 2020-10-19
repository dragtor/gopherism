package pkg

import (
	"reflect"
	"testing"
)

var testCases TestCase

type TestCase struct {
	InputText      string
	ExpectedResult []Link
}

var TestCases = []TestCase{
	TestCase{
		InputText: `
       <a href="/dog">
       <span>Something in a span</span>
       Text not in a span
       <b>Bold text!</b>
       </a>
       `,
		ExpectedResult: []Link{Link{Href: "/dog",
			Text: "Something in a span Text not in a span Bold text!",
		}},
	},
	TestCase{
		InputText: `
       <html>
       <body>
       <h1>Hello!</h1>
       <a href="/other-page">A link to another page</a>
       </body>
       </html>
       `,
		ExpectedResult: []Link{Link{Href: "/other-page",
			Text: "A link to another page",
		}},
	},
	TestCase{
		InputText: `
        <html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>
       `,
		ExpectedResult: []Link{Link{Href: "https://www.twitter.com/joncalhoun",
			Text: "Check me out on twitter",
		},
			Link{Href: "https://github.com/gophercises",
				Text: "Gophercises is on Github !",
			},
		},
	},
}

func TestGetLinks(t *testing.T) {
	for _, tc := range TestCases {
		out, _ := GetLinks([]byte(tc.InputText))
		if !reflect.DeepEqual(out, tc.ExpectedResult) {
			t.Errorf("Expected result : %+v , Actual Result : %+v", tc.ExpectedResult, out)
			return
		}
	}
}
