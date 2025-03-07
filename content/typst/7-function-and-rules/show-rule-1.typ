#let colorText(color: red, it) = {
  text(fill: color)[#it]
}

#show heading: it => colorText(color: blue, it)
// #show heading: colorText.with(color: blue)

= Heading 1

== Heading 1.1
