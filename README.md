install gvm, atom1.6 (go-plus, tree-view-git-status, pretty json)

re-session

```
gvm install go1.4
gvm install go1.6
gvm use go1.6

export GOPATH=`pwd`

cd src/echo
go build
go install

cd src/hello
go build
go install

bin/hello

or

go build hello
```
