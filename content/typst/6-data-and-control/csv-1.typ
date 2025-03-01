#show raw: body => {
    set text(font: (
      (
        name: "Inconsolata",
        covers: "latin-in-cjk",
      ),
      "Noto Sans CJK JP"
    ))
    body
}

#let holidays = csv(
  "./holidays.csv",
  delimiter: ",",
  row-type: dictionary,
)

#holidays
