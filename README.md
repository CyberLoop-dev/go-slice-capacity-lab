<div align="center">
<img src="https://capsule-render.vercel.app/api?type=waving&color=0:101827,100:0f766e&height=170&section=header&text=Go%20Slice%20Capacity%20Lab&fontSize=32&fontColor=ffffff&animation=fadeIn" width="100%" alt="Waving header"/>
<img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=600&size=18&pause=900&color=14B8A6&center=true&vCenter=true&width=760&lines=A%20dependency-free%20Go%20CLI%20showing%20slice%20length%2C%20capacity%2C%20shared%20backing%20arrays" alt="Typing title"/>

[![License](https://img.shields.io/github/license/CyberLoop-dev/go-slice-capacity-lab)](LICENSE)
![Language](https://img.shields.io/github/languages/top/CyberLoop-dev/go-slice-capacity-lab)
![Stars](https://img.shields.io/github/stars/CyberLoop-dev/go-slice-capacity-lab?style=flat)
![Last commit](https://img.shields.io/github/last-commit/CyberLoop-dev/go-slice-capacity-lab)
<img src="assets/mark.svg" width="112" alt="Go Slice Capacity Lab illustration"/>
</div>

## 🧭 Overview

A dependency-free Go CLI showing slice length, capacity, shared backing arrays, and what changes when append allocates new storage.

## ✨ Features

- 📏 Inspect len and cap after each operation.
- 🧬 See how slices can share a backing array.
- 🌱 Compare append within capacity and reallocation.
- 🧪 Test aliasing behavior with the standard library.

## 🎬 Demo

No hosted demo is provided. The examples run locally and use only the inputs documented below.

## ⚡ Quick Start

~~~bash
go run .
~~~

## 🧰 Tech Stack

- Go 1.22+
- Standard library or the small dependencies listed in the project manifest

## 🗂️ Folder Structure

~~~text
go-slice-capacity-lab/
├── main.go
├── main_test.go
├── go.mod
├── assets/mark.svg
├── .gitignore
├── LICENSE
└── README.md
~~~

## 🗺️ Roadmap

- [x] Core behavior and local examples
- [x] README and custom SVG illustration
- [ ] Add more guided exercises from community feedback

## 🤝 Contributing

Issues and focused pull requests are welcome. Keep examples reproducible, local-first, and within the stated scope.

## 📄 License

MIT License. See [LICENSE](LICENSE).

<div align="center"><img src="https://capsule-render.vercel.app/api?type=waving&color=0:0f766e,100:101827&height=100&section=footer" width="100%" alt="Waving footer"/></div>
