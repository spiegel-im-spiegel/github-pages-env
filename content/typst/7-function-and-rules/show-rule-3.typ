#let colorText(color: red, it) = {
  text(fill: color)[#it]
}

#show heading.where(level: 1).or(heading.where(level: 2)): colorText

= Heading 1

== Heading 1.1

=== Heading 1.1.1
