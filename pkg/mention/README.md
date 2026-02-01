# mention [![Build Status](https://travis-ci.org/gernest/mention.svg)](https://travis-ci.org/gernest/mention) [![Coverage Status](https://coveralls.io/repos/gernest/mention/badge.svg?branch=master&service=github)](https://coveralls.io/github/gernest/mention?branch=master) [![GoDoc](https://godoc.org/github.com/gernest/mention?status.svg)](https://godoc.org/github.com/gernest/mention)

`mention` parses twitter like mentions and hashtags like @gernest and #Tanzania from text input.

# Installation

	go get github.com/gernest/mention
	

# Usage

`mention` is flexible, meaning it is not only limited to `@` and `#` tags. You can choose whatever tag you like and mention will take it from there.

## twitter like mentions

For instance you have the following message

```
hello @gernesti I would like to follow you on twitter
```

And you want to know who was mentioned in the text.

```go
package main

import (
	"fmt"
	"strings"

	"github.com/gernest/mention"
)

func main() {
	message := "hello @gernest I would like to follow you on twitter"

	tags := mention.GetTags('@', message)
	tagStrings := mention.GetTagsAsUniqueStrings('@', message)

	fmt.Println(tags)
	fmt.Println(tagStrings)
}
```

If you run the above example it will print `[gernest]` is the stdout.

## twitter like hashtags

For instance you have the following message

```
how does it feel to be rejected? #loner
```

And you want to know the hashtags

```go
package main

import (
	"fmt"
	"strings"

	"github.com/gernest/mention"
)

func main() {
	message := "how does it feel to be rejected? #loner"

	tags := mention.GetTags('#', message)

	fmt.Println(tags)
}
```

If you run the above example it will print `[loner]` in the stdout.

# The API
mention exposes only one function `GetTags(char rune, src string) []string`

The first argument `char` is the prefix for your tag, this can be `@` or `#` or whatever unicode character you prefer. Don't be worried by its type `rune` it is just your normal characters but in single quotes. See the examples for more information.

The second argument is the source of the input which can be from texts.

# Contributing

Start with clicking the star button to make the author and his neighbors happy. Then fork the repository and submit a pull request for whatever change you want to be added to this project.

If you have any questions, just open an issue.

# Author
Geofrey Ernest
Twitter  : [@gernesti](https://twitter.com/gernesti)

Chad Barraford
Github  : [@cbarraford](https://github.com/cbarraford)


# Licence

This project is released under the MIT licence. See [LICENCE](LICENCE) for more details.
