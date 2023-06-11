// `package view` is declaring that this file belongs to the `view` package.
// This allows other files in the same package to access the functions and variables defined in this file.
package view

// The `import` statement is used to import packages that are necessary for the program to run.
// In this case, the program is importing the following packages:
import "text/template"

type View struct {
	Template *template.Template
}

// The function creates a new view by parsing a list of files and returning a pointer to a View struct.
func NewView(files ...string) *View {
	files = append(files, "view/front-end/header.html", "view/front-end/menu.html")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}
