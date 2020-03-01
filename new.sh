#!/bin/bash
y=$(date '+%Y')
m=$(date '+%m')
d=$(date '+%d')

section=$1
fname=$2
p=""
if [ ${section:-0} == 0 ]; then
	echo "入力ファイルを指定してください。"
	exit 1
elif [ $section == "remark" ]; then
	if [ ${fname:-0} == 0 ]; then
		p="$section/$y/$m/$d-stories.md"
	else
		p="$section/$y/$m/$fname"
	fi
elif [ $section == "bookmarks" ]; then
	if [ ${fname:-0} == 0 ]; then
		p="$section/$y/$m/$d-bookmarks.md"
	else
		p="$section/$y/$m/$fname"
	fi
elif [ $section == "release" ]; then
	if [ ${fname:-0} == 0 ]; then
		echo "入力ファイルを指定してください。"
		exit 1
	else
		p="$section/$y/$m/$fname"
	fi
elif [ $section == "cc-licenses" ]; then
	if [ ${fname:-0} == 0 ]; then
		echo "入力ファイルを指定してください。"
		exit 1
	else
		p="$section/$fname"
	fi
elif [ $section == "golang" ]; then
	if [ ${fname:-0} == 0 ]; then
		echo "入力ファイルを指定してください。"
		exit 1
	else
		p="$section/$fname"
	fi
elif [ $section == "rust-lang" ]; then
	if [ ${fname:-0} == 0 ]; then
		echo "入力ファイルを指定してください。"
		exit 1
	else
		p="$section/$fname"
	fi
elif [ $section == "hugo" ]; then
	if [ ${fname:-0} == 0 ]; then
		echo "入力ファイルを指定してください。"
		exit 1
	else
		p="$section/$fname"
	fi
elif [ $section == "openpgp" ]; then
	if [ ${fname:-0} == 0 ]; then
		echo "入力ファイルを指定してください。"
		exit 1
	else
		p="$section/$fname"
	fi
else
	fname=$section
	section=""
	p=$fname
fi

echo "Create file content/$p"
hugo new $p
