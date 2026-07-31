#show heading: set text(font: "Noto Sans CJK JP")


#let colorText(color: red, it) = {
  text(fill: color)[#it]
};

#colorText[Hello] #colorText(color: blue)[world]
