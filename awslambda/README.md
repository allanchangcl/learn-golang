// add above struct
//go:generate easytags $GOFILE

// install easytags
go get -u github.com/betacraft/easytags/...

// generate the struct tags
go generate
