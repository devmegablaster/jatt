---
layout: default
---

# Title Header (H1 header)

### Introduction (H3 header)

This is some placeholder text to show examples of Markdown formatting.
We have [a full article template](https://github.com/do-community/do-article-templates) you can use when writing a DigitalOcean article.
Please refer to our style and formatting guidelines for more detailed explanations: <https://do.co/style>

## Prerequisites (H2 header)

Before you begin this guide you'll need the following:

- Familiarity with [Markdown](https://daringfireball.net/projects/markdown/)

## Step 1 — Basic Markdown

This is _italics_, this is **bold**, this is **underline**, and this is ~~strikethrough~~.

- This is a list item.
- This list is unordered.

1. This is a list item.
2. This list is ordered.

> This is a quote.
>
> > This is a quote inside a quote.
>
> - This is a list in a quote.
> - Another item in the quote list.

Here's how to include an image with alt text and a title:

![Alt text for screen readers](https://assets.digitalocean.com/logos/DO_Logo_horizontal_blue.png)
Caption for the image

_We also support some extra syntax for setting the width, height and alignment of images. You can provide pixels (`200`/`200px`), or a percentage (`50%`), for the width/height. The alignment can be either `left` or `right`, with images being centered by default. These settings are all optional._

![](https://assets.digitalocean.com/public/mascot.png)

Use horizontal rules to break up long sections:

---

Rich transformations are also applied:

- On ellipsis: ...
- On quote pairs: "sammy", 'test'
- On dangling single quotes: it's
- On en/em dashes: a -- b, a --- b

<!-- Comments will be removed from the output -->

| Tables | are   | also  | supported | and   | will  | overflow | cleanly | if    | needed |
| ------ | ----- | ----- | --------- | ----- | ----- | -------- | ------- | ----- | ------ |
| col 1  | col 2 | col 3 | col 4     | col 5 | col 6 | col 7    | col 8   | col 9 | col 10 |
| col 1  | col 2 | col 3 | col 4     | col 5 | col 6 | col 7    | col 8   | col 9 | col 10 |
| col 1  | col 2 | col 3 | col 4     | col 5 | col 6 | col 7    | col 8   | col 9 | col 10 |
| col 1  | col 2 | col 3 | col 4     | col 5 | col 6 | col 7    | col 8   | col 9 | col 10 |
| col 1  | col 2 | col 3 | col 4     | col 5 | col 6 | col 7    | col 8   | col 9 | col 10 |

## Step 2 — Code

This is `inline code`. This is a <^>variable<^>. This is an `in-line code <^>variable<^>`. You can also have [`code` in links](https://www.digitalocean.com).

Here's a configuration file with a label:

```nginx
[label /etc/nginx/sites-available/default]
server {
    listen 80 <^>default_server<^>;
    . . .
}
```

Examples can have line numbers, and every code block has a 'Copy' button to copy just the code:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

Here's output from a command with a secondary label:

This is a non-root user command example:

```bash
adduser sammy
shutdown
```

This is a custom prefix command example:

```sql
FLUSH PRIVILEGES;
SELECT * FROM articles;
```

A custom prefix can contain a space by using `\s`:

```sql
FLUSH PRIVILEGES;
SELECT * FROM articles;
```

Indicate where commands are being run with environments:

```bash
[environment local]
ssh root@server_ip
```

```bash
[environment second]
echo "Secondary server"
```

```bash
[environment third]
echo "Tertiary server"
```

```bash
[environment fourth]
echo "Quaternary server"
```

```bash
[environment fifth]
echo "Quinary server"
```

And all of these can be combined together, with a language for syntax highlighting as well as a line prefix (line numbers, command, custom prefix, etc.), and even an environment and label:

```html
[environment second]
[label index.html]
<html>
<body>
<head>
  <title><^>My Title<^></title>
</head>
<body>
  . . .
</body>
</html>
```

## Step 3 — Callouts

Here is a note, a warning, some info and a draft note:

<$>[note]
**Note:** Use this for notes on a publication.
<$>

<$>[warning]
**Warning:** Use this to warn users.
<$>

<$>[info]
**Info:** Use this for product information.
<$>

<$>[draft]
**Draft:** Use this for notes in a draft publication.
<$>

A callout can also be given a label, which supports inline markdown as well:

<$>[note]
[label Labels support _inline_ **markdown**]
**Note:** Use this for notes on a publication.
<$>

You can also mention users by username:

@MattIPv4

## Step 4 — Layout

Columns allow you to customise the layout of your Markdown:

[column
Content inside a column is regular Markdown block content.

> Any block or inline syntax can be used, including quotes.
> ]

[column
Two or more columns adjacent to each other are needed to create a column layout.

On desktop the columns will be evenly distributed in a single row, on tablets they will wrap naturally, and on mobile they will be in a single stack.
]

[details Content can be hidden using `details`.
Inside the details block you can use any block or inline syntax.

You could hide the solution to a problem:

```js
// Write a message to console
console.log("Hello, world!");
```
