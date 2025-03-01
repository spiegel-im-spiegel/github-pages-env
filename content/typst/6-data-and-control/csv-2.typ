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
#let header = holidays.first().keys()
#let data = holidays.map(holiday => holiday.values())

#header

#data
