# goldmark-plantuml

[![Sync to Gitee](https://github.com/OhYee/goldmark-plantuml/workflows/Sync%20to%20Gitee/badge.svg)](https://gitee.com/OhYee/goldmark-plantuml) [![w
orkflow state](https://github.com/OhYee/goldmark-plantuml/workflows/test/badge.svg)](https://github.com/OhYee/goldmark-plantuml/actions) [![codecov](https://codecov.io/gh/OhYee/goldmark-plantuml/branch/master/graph/badge.svg)](https://codecov.io/gh/OhYee/goldmark-plantuml) [![version](https://img.shields.io/github/v/tag/OhYee/goldmark-plantuml)](https://github.com/OhYee/goldmark-plantuml/tags)

goldmark-plantuml is an extension for [goldmark](https://github.com/yuin/goldmark).  

You can plantuml to build svg image in your markdown like [mume](https://github.com/shd101wyy/mume)

## screenshot

There are two demo(using `'` instead of &#8242; in the code block)

1. default config

[Demo1](demo/demo1/main.go)
[Output1](demo/demo1/output.html)

```markdown
'''go
package main

import ()

func main(){}
'''

'''plantuml
@startuml
Alice -> Bob: test
@enduml
'''
```

![](img/default.png)

2. using `plantuml-svg` and [goldmark-highlighting extension](https://github.com/yuin/goldmark-highlighting)

[Demo2](demo/demo1/main.go)
[Output2](demo/demo1/output.html)

```markdown
'''go
package main

import ()

func main(){}
'''

'''plantuml-svg
@startuml
Alice -> Bob: test
@enduml
'''
```

![](img/highlighting.png)

## Installation

```bash
go get -u github.com/OhYee/goldmark-plantuml
```

## License

[MIT](LICENSE)
