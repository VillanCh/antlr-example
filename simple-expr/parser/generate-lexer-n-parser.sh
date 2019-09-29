#!/usr/bin/env sh
antlr -Dlanguage=Go -o gen -package gen -visitor SimpleExpr.g4