# Dragon-cli

## Installation

```
go install github.com/svetozar12/dragon-cli@latest
```

To check if installation was complete

```
dragon-cli --version

Your CLI Tool Version vx.x.x
```

## Usage

To generate project run the following command and follow the steps

```
dragon-cli --projectName=my-example-app --beFramework=nodejs --feFramework=Nextjs --installDeps=true
```

Or
`dragon-cli`

### I wish you happy coding !

<!-- Docs -->

--branch=`<the branch from which you want to get the templates default is master>`
--projectName=`<the name of the project>`
--beFramework=`"nodejs","golang","none"`
--feFramework=`"React(with vite)","Nextjs","Astro","none"`
--installDeps=`"true","false"`
go test -v -coverprofile cover.out ./...
go tool cover -html cover.out -o cover.html
