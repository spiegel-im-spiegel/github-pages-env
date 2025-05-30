// for main text
#set text(font: (
  (
    name: "New Computer Modern",
    covers: "latin-in-cjk",
  ),
  "BIZ UDMincho"
))

// for headings
#show heading: it => {
    set text(font: (
      (
        name: "Liberation Sans",
        covers: "latin-in-cjk",
      ),
      "BIZ UDGothic"
    ))
	it.body
}

= Albert Einsteinについて

Albert Einsteinは1879年3月14日，ドイツ生まれの理論物理学者です。
彼による革命的な3つの論文「光電効果の理論」「ブラウン運動の理論」「特殊相対性理論」が発表された1905年は「奇跡の年」と呼ばれています。
