Code Kata02: Karate Chop
===

Implementation of Code Kata02 exercice. Link to the challenge [here](http://codekata.com/kata/kata02-karate-chop/).

## Table of Contents

* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
  * [Start tests](#start-tests)
  * [Start each implementation](#start-each-implementation)
* [TODO list](#todo-list)
* [License](#license)
* [Contact](#contact)
* [Acknowledgements](#acknowledgements)


## Getting Started

### Prerequisites

In order to compile and run the different implementation, you will need :
 - Docker installed. You can find instructions [here](https://docs.docker.com/get-docker/)
 - Git
 
### Installation
1. Clone this repository in your GOPATH : 
```bash
mkdir -p $GOPATH/src/github.com/clnbs
cd $GOPATH/src/github.com/clnbs
git clone https://github.com/clnbs/autorace/kararteChop
cd kararteChop
```

1. Compiling implementations
```
make all
```
This command should create a binary for each implementation 
 
## Usage

### Start tests
Tests are trigger in a Docker environment in order to be compliant with the compilation and development environment. Tests can be easily done using the Makefile : 
```bash
make testing
```

This command will start all statics tests and create an HTML file named `cover.html` in the root directory of this project. You can open it with your favorite web browser.

### Start each implementation
For each implementation, you can start a specific binary in the form `<implementation_name>.bin` at the root directory of this project. To start a binary, you can simply run it, e.g :
```bash
./recursive.bin
```

## TODO list
 - missing 2 out of 5 implementations
 - better testing code

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact

Colin Bois - <colin.bois@rocketmail.com>

Project Link: [https://github.com/clnbs/karateChop](https://github.com/clnbs/karateChop)

## Acknowledgements

* [Code Kata](http://codekata.com/)
* [Cycloid.io](https://www.cycloid.io/)
* [Golang testify from stretchr](https://github.com/stretchr/testify)