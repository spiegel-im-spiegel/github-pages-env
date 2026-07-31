#let colorText(color: red, it) = {
  text(fill: color)[#it]
}

#show: colorText.with(color: blue)
// #show: it => colorText(color: blue)[#it]

#lorem(40)
