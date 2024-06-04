#!/bin/sh

JAR="antlr-4.13.1-complete.jar"

if [ ! -f "$JAR" ]; then
    curl -O https://www.antlr.org/download/"$JAR"
fi
rm -f *.interp *.tokens

alias antlr4='java -Xmx500M -cp "./${JAR}:$CLASSPATH" org.antlr.v4.Tool'
antlr4 -Dlanguage=Go -package parser *.g4
