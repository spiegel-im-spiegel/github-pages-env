#!/usr/bin/env perl

# LaTeX commands
$pdflualatex         = 'lualatex %O -synctex=1 %S';
$latex               = 'uplatex %O -synctex=1 %S';
$latex_silent_switch = '-interaction=batchmode -c-style-errors';

# bibTeX commands
$bibtex    = 'upbibtex %O %B';
$biber     = 'biber %O --bblencoding=utf8 -u -U --output_safechars %B';
$makeindex = 'mendex %O -o %D %S';

# Device Driver
$dvipdf = 'dvipdfmx %O -z9 -V 7 -o %D %S';
$dvips  = 'dvips %O -z -f %S | convbkmk -u > %D';
$ps2pdf = 'ps2pdf14 -dPDFA -dPDFACompatibilityPolicy=1 -sProcessColorModel=DeviceCMYK %O %S %D';

# Typeset mode (generate a PDF)
$pdf_mode = 4; # 0: do not generate a pdf , 1: using $pdflatex , 2: using $ps2pdf , 3: using $dvipdf , 4: using $pdflualatex

# Other configuration
$pvc_view_file_via_temporary = 0;
$max_repeat = 5;
$clean_ext = "xmpdata";
