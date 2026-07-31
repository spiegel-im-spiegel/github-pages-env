#let colorText(color: red, it) = {
  text(fill: color)[#it]
}

= Heading 1

== Heading 1.1

#show heading: colorText

= Heading 2

== Heading 2.1

#show heading: it => colorText(color: blue, it)

= Heading 3

== Heading 3.1
