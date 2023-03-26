#!/bin/sh

#alias antlr='java -Xmx500M -cp "./antlr-4.12.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
antlr -Dlanguage=Go -package parser *.g4
