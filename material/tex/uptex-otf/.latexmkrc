#!/usr/bin/env perl
$latex        = 'uplatex -synctex=1';
$latex_silent = 'uplatex -synctex=1 -interaction=batchmode';
$bibtex       = 'upbibtex';
$biber        = 'biber --bblencoding=utf8 -u -U --output_safechars';
$dvipdf       = 'dvipdfmx -z9 -V 4 -f sourcehanjp.map %O -o %D %S';
#$dvipdf       = 'dvipdfmx -z9 -V 4 %O -o %D %S && ps2pdf14 -dPDFA -dPDFACompatibilityPolicy=1 -sProcessColorModel=DeviceCMYK %D %R-pdfa.pdf';
#$ps2pdf       = 'ps2pdf14 -dPDFA -dPDFACompatibilityPolicy=1 -sProcessColorModel=DeviceCMYK %O %S %D';
$makeindex    = 'mendex %O -o %D %S';
$max_repeat   = 5;
$pdf_mode	  = 3; # generates pdf via dvipdfmx

$pvc_view_file_via_temporary = 0;
