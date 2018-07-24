# gogen

These instructions are for testing code generation.

```bash
# Set the environment variables.
export GOGEN_PROJECT_DIR=$GOPATH/src/app/webapi
export GOGEN_TEMPLATE_DIR=$GOPATH/src/app/webapi/template

# CD to the correct folder.
cd src/app/gogen/cmd/gogen

# Generate templates from the project.
go run main.go template component/user component itemUpper:User itemLower:user allUpper:USER

# Generate code from the new templates.
go run main.go generate component/default itemUpper:Monkey itemLower:monkey allUpper:MONKEY
```