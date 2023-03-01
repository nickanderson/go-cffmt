%%%
title = "cffmt 1"
area = "User Commands"
workgroup = "CFEngine"
%%%

cffmt
=====

## Name

cffmt - CFEngine pretty printer/formatter

## Synopsis

cffmt` *[OPTION]*... *[FILE]*

## Description

Cffmt will take a CFengine file out of **FILE** or from standard input and will reformat it.

Options are:

- `-a` print the AST to standard error
- `-p` print the pretty printed document to standard output (defaults to true)
- `-d` show lexer/yacc debug will parsing

## Author

Miek Gieben <miek@miek.nl>.