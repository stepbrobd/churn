#let doc(
  name: none,
  id: none,
  email: none,
  institution: none,
  semester: none,
  course: none,
  instructor: none,
  title: none,
  due: none,
  body,
) = [
  #set document(
    author: name,
    title: name + " - " + institution + " " + semester + " " + course + " - " + title,
  )

  #set par(linebreaks: "optimized")
  #set text(font: "New Computer Modern", lang: "en", size: 10pt)

  #show link: underline

  #set page(
    paper: "us-letter",
    margin: (x: 1in, y: 1in),
    header: locate(loc => if [#loc.page()] == [1] [
      #h(1fr)
      #text(gray)[#due.display("[month repr:long] [day], [year]")]
    ] else [
      #h(1fr)
      #text(gray)[#name]
    ]),
    footer: [
      #h(1fr)
      #text(gray)[#counter(page).display("1/1", both: true)]
    ],
  )

  #align(center)[
    = #title

    #v(0.1in)

    #institution | #semester | #course | #instructor

    #name | #id | #link("mailto:" + email)[#email]

    #v(0.1in)
  ]

  #body
]

#let answer(title: none, body) = [
  = #title

  #block(fill: luma(230), inset: 8pt, width: 100%, [
    #body
  ])
]