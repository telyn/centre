[![Build Status](https://travis-ci.org/telyn/centre.svg?branch=master)](https://travis-ci.org/telyn/centre) [![Coverage Status](https://coveralls.io/repos/github/telyn/centre/badge.svg?branch=master)](https://coveralls.io/github/telyn/centre?branch=master)

Tiny go program to generate text centred in = signs.

Each argument is its own piece of text, and specify width with `--width <n>`

```
» go get github.com/telyn/centre

» centre "Ingredients" "Precautions" "Method"
========================= Ingredients =========================
========================= Precautions =========================
=========================== Method ============================


» centre -width 32 "Ingredients" "Precautions" "Method"
========================== Ingredients ==========================
========================== Precautions ==========================
============================ Method =============================
```
