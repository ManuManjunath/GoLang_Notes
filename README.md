# GoLang_Notes

To create a workspace in - /users/you/go - This is the default install location

To view this - $ go env GOPATH


Create folders - 

Go to /users/you/go

Create a folder called "src" for your source code.

Create a folder called "pkg" for any packages you import

Create a folder called "bin" for the programs you compile

mkdir "your project name" inside "src"


// To run, go back to the terminal and -
// $ go run filename.go

// To compile without running -
// $ go build
// This creates an executable in the project folders.
// To run the executale -
// $ ./the_build_file

// Another way -
// $ go install
// The executable is placed in the bin folder

// To initialize your project, go to your workspace direcvtory and -
// $ go mod init <your github repo>
// After this, you can run your code also using -
// $ go run <your github repo>
