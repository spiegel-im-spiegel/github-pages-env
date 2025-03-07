#let colorText(color: red, it) = {
  text(fill: color)[#it]
}

= Heading 1

== Heading 1.1

#show heading: colorText

= Heading 2

== Heading 2.1
