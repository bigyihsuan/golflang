#!/bin/sh

antlr4 -Dlanguage=Go -visitor -package par *.g4 -o .