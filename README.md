[![Build Status](https://travis-ci.org/telyn/centre.svg?branch=master)](https://travis-ci.org/telyn/centre) [![Coverage Status](https://coveralls.io/repos/github/telyn/centre/badge.svg?branch=master)](https://coveralls.io/github/telyn/centre?branch=master)

Tiny go program to generate text centred in = signs.

Each argument is its own piece of text, and for now the width is fixed to 30 characters.

```
» go get github.com/telyn/centre

» centre "Ingredients" "Precautions" "Method"
========================= Ingredients =========================
========================= Precautions =========================
=========================== Method ============================
```
