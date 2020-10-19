package pkg

import (
	"testing"
	//"fmt"
	"reflect"
	//"github.com/dragtor/gopherism/htmllink/pkg"
)

func TestGetLinks(t *testing.T) {
	htmltxt := `
    <a href="/dog">
    <span>Something in a span</span>
    Text not in a span
    <b>Bold text!</b>
    </a>

    `
	expectedOutput := []Link{Link{Href: "/dog",
		Text: "Something in a span Text not in a span Bold text!",
	}}
	out, _ := GetLinks([]byte(htmltxt))
	if !reflect.DeepEqual(out, expectedOutput) {
		t.Errorf("Expected result : %+v , Actual Result : %+v", expectedOutput, out)
		return
	}
	//fmt.Printf("Correct output")
}
