package main

import (
	"log"
	"os"

	"app/gogen"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	app       = kingpin.New("gogen", "A command-line application to generate code.")
	cGenerate = app.Command("generate", "Generate files from template pairs.")
	//cGenerateProjectFolder  = cGenerate.Flag("project", "Path to the project folder.").Required().String()
	//cGenerateTemplateFolder = cGenerate.Flag("template", "Path to the template folder.").Required().String()
	cGenerateTmpl = cGenerate.Arg("folder/template", "Template pair name. Don't include an extension.").Required().String()
	cGenerateVars = gogen.NewStringList(cGenerate.Arg("key:value", "Key and value required for the template pair."))

	cTemplate = app.Command("template", "Generate templates from source code.")

	cTemplateProject = cTemplate.Arg("folder", "Project folder. Don't include a trailing slash.").Required().String()
	cTemplateName    = cTemplate.Arg("name", "Template name to create.").Required().String()
	cTemplateVars    = gogen.NewStringList(cTemplate.Arg("key:value", "Key and value required for the template pair."))
)

// init sets runtime settings.
func init() {
	// Verbose logging with file name and line number.
	log.SetFlags(log.Lshortfile)
}

func main() {
	argList := os.Args[1:]
	arg := kingpin.MustParse(app.Parse(argList))

	// Get the environment variables.
	projectDir := os.Getenv("GOGEN_PROJECT_DIR")
	if len(projectDir) == 0 {
		app.Fatalf("Missing environment variable: %v", "GOGEN_PROJECT_DIR")
	}
	templateDir := os.Getenv("GOGEN_TEMPLATE_DIR")
	if len(templateDir) == 0 {
		app.Fatalf("Missing environment variable: %v", "GOGEN_TEMPLATE_DIR")
	}

	switch arg {
	case cGenerate.FullCommand():
		// Generate the code.
		err := gogen.Run(argList[1:], projectDir, templateDir)
		if err != nil {
			app.Fatalf("%v", err)
		}
	case cTemplate.FullCommand():
		// Generate the templates.
		err := gogen.CreateTemplate(argList[1:], argList[2], projectDir, templateDir)
		if err != nil {
			app.Fatalf("%v", err)
		}
	}
}
