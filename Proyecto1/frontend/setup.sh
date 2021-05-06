#!/bin/sh -
URL=${1:-'https://proyecto.grupo16-api.codes'}
URL=$(echo $URL | awk 'gsub("/","\\/",$0)')
sed -i 's/const\s\+url\s\+=\s*'"'"'.\+'"'"';/const url = '"'$URL';/g" *.html
