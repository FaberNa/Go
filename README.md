#for run program go in folder hello and run this command
go run main.go

#for build , this command produce executable file
go build main.go

#for format code
gofmt -w main.go

#to imports missing import reference
goimports -w main.go


#to see how many time
time go main.go

#to run only on one proc
time GOMAXPROCS=1 go run main.go