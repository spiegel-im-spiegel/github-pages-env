unset key
set style line 1 pointsize 0.1 pointtype 7 linecolor rgb "black"
f(x)=a*x+b
fit f(x) "qq100k.dat" via a,b
plot "qq100k.dat" linestyle 1, f(x)
