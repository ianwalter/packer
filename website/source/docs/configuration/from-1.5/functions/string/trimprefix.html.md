---
layout: "docs"
page_title: "trimprefix - Functions - Configuration Language"
sidebar_current: "configuration-functions-string-trimprefix"
description: |-
  The trimprefix function removes the specified prefix from the start of a
  given string.
---

# `trimprefix` Function


`trimprefix` removes the specified prefix from the start of the given string.

## Examples

```
> trimprefix("helloworld", "hello")
world
```

## Related Functions

* [`trim`](./trim.html) removes characters at the start and end of a string.
* [`trimsuffix`](./trimsuffix.html) removes a word from the end of a string.
* [`trimspace`](./trimspace.html) removes all types of whitespace from
  both the start and the end of a string.
