# goldmark-plantuml

[![master_test](https://github.com/OhYee/goldmark-plantuml/workflows/master_test/badge.svg)](https://github.com/OhYee/goldmark-plantuml/actions?workflow=master_test)

goldmark-plantuml is an extension for [goldmark](https://github.com/yuin/goldmark).  

You can use plantuml to build svg image in your markdown like [mume](https://github.com/shd101wyy/mume)

## screenshot

There are two demo(using `'` instead of &#8242; in the code block)

1. default config

```markdown
'''go
package main

import ()

func main(){}
'''

'''uml
@startuml
Alice -> Bob: test
@enduml
'''
```

2. using `plantuml-svg` and [goldmark-highlighting extension](https://github.com/yuin/goldmark-highlighting)

```markdown
'''go
package main

import ()

func main(){}
'''

'''uml-svg
@startuml
Alice -> Bob: test
@enduml
'''
```

## Installation

```bash
go get -u github.com/OhYee/goldmark-plantuml
```

## Usage

See [uml_test.go](./uml_test.go)

## License

[MIT](LICENSE)
