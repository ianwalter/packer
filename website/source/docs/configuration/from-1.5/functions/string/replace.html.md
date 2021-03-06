---
layout: "docs"
page_title: "replace - Functions - Configuration Language"
sidebar_current: "configuration-functions-string-replace"
description: |-
  The replace function searches a given string for another given substring,
  and replaces all occurrences with a given replacement string.
---

# `replace` Function


`replace` searches a given string for another given substring, and replaces
each occurrence with a given replacement string.

```hcl
replace(string, substring, replacement)
```

## Examples

```
> replace("1 + 2 + 3", "+", "-")
1 - 2 - 3

> replace("hello world", "world", "everybody")
hello everybody
```

## Related Functions

- [`regexreplace`](./regexreplace.html) searches a given string for another given substring, 
  and replaces each occurrence with a given replacement string.
