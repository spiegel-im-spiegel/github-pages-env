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

#let key3 = ""

#if "key3" in sys.inputs {
	key3 = sys.inputs.at("key3")
} else {
	assert(false, message: "key3 is not provided. Please provide it using the --input option.")
}

#key3
