# Clean Architecture checker for Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/roblaszczak/go-cleanarch)](https://goreportcard.com/report/github.com/roblaszczak/go-cleanarch)

go-cleanarch was created to keep Clean Architecture rules,
like a _The Dependency Rule_ and _interaction between modules_ in your Go projects.
More about Clean Architecture you can read in [Uncle's Bob article](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html).


Some benefits of using Clean Architecture:

> 1. Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
> 2. Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
> 3. Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
> 4. Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
> 5. Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

_Source: The Clean Architecture_

![Clean Architecture](docs/go-cleanarch.png)

### Project schema requirements

go-cleanarch assumes this files structure:

    [GOPATH]/[PACKAGE_NAME]/[LAYER_NAME]

or

    [GOPATH]/[PACKAGE_NAME]/[MODULE_NAME]/[LAYER_NAME]

For example

* go/src/github.com/roblaszczak/awesome-app
    * auth
        * domain
        * application
        * interfaces
    * content
        * domain
            * submodule1
            * submodule2
            * *etc.*
        * application
        * interfaces
    * frontend
        * domain
        * application
        * interfaces

### Allowed `LAYER_NAME`:

The default layer names are as followed. It is possible to set different names
by command line parameters see -domain/-application/-interfaces/-infrastructure
bellow.

    var LayersAliases = map[string]Layer{
        // Domain
        "domain":   LayerDomain,
        "entities": LayerDomain,

        // Application
        "app":         LayerApplication,
        "application": LayerApplication,
        "usecases":    LayerApplication,
        "usecase":     LayerApplication,
        "use_cases":   LayerApplication,

        // Interfaces
        "interfaces": LayerInterfaces,
        "interface":  LayerInterfaces,
        "adapters":   LayerInterfaces,
        "adapter":    LayerInterfaces,

        // Infrastructure
        "infrastructure": LayerInfrastructure,
        "infra":          LayerInfrastructure,
    }

For examples please go to [examples](examples/) directory,
with contains examples of valid and invalid architectures.

For more informations about Clean Architecture please read [Uncle's Bob article](https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html)


## Installing

    go get -u gopkg.in/roblaszczak/go-cleanarch.v1

_go-cleanarch was only tested on Linux and also should work on OS X.
Probably it doesn't work well on Windows._

## Running

To run in current directory:

    go-cleanarch

To run in provided directory

    go-cleanarch go/src/github.com/roblaszczak/awesome-cms

Process will exit with code `1` if architecture is not valid, otherwise it will exit with `0`.

### -ignore-tests

If you need to ignore `*_test.go` files in `go-cleanarch` check you can pass `-ignore-tests`

    go-cleanarch -ignore-tests

It is useful when you have memory implementation in infrastructure layer
and you need to test application service which depends of it.

### -ignore-package

If for some reason you need to allow to make forbidden import, for example

`github.com/roblaszczak/go-cleanarch/examples/ignore-package/app` to `github.com/roblaszczak/go-cleanarch/examples/ignore-package/domain`.

you can use

    go-cleanarch -ignore-package=github.com/roblaszczak/go-cleanarch/examples/ignore-package/app 

### Layer names

The layer names can be set to a specific value with the following parameters. Each
parameter stands for on layer.

    go-cleanarch -domain dom -application appli -interfaces int -infrastructure outer

This would only allow the domain name to be dom, application hast to be appli,
interafces must be int and infrastructure must be outer.

## Running the tests

    make test

## And coding style tests

    make qa

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## Credits

Made without love by Robert Laszczak </3

## License

This project is licensed under the MIT License.
