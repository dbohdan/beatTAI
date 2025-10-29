# beatTAI

```none
▄                          ▄      ▄ ▄ ▄ ▄ ▄    ▄ ▄    ▄ ▄ ▄
▄ ▄ ▄     ▄ ▄     ▄ ▄ ▄    ▄ ▄ ▄      ▄      ▄     ▄    ▄
▄    ▄  ▄     ▄         ▄  ▄          ▄      ▄ ▄ ▄ ▄    ▄
▄    ▄  ▄ ▀ ▀     ▄ ▀ ▀ ▄  ▄          ▄      ▄     ▄    ▄
▄ ▄ ▄     ▄ ▄ ▄     ▄ ▄ ▄    ▄ ▄      ▄      ▄     ▄  ▄ ▄ ▄

 -:[ a new Internet time for turbonerds & superweirdos ]:-
```

This repository is a fork of [B4UDW3RK5/beatTAI](https://github.com/B4UDW3RK5/beatTAI)&thinsp;<sub>404</sub> by cat K.


## Original abstract

<pre>
   In 1998 Swatch introduced <a href="https://en.wikipedia.org/wiki/Swatch_Internet_Time">Swatch Internet Time</a> (or .beat time), displayed as
@xxx.xx (for example @198.26). It's a cute and fun way to display the time and
I find it aesthetically very pleasing to look at.

   However, Swatch being based in Switzerland they chose to align .beats to
Swiss time (UTC+1) which outside of countries in that timezone means very
little and is impractical.

   I half-jokingly <a href="https://hackers.town/@cat/107951155969547130">posted on the Fediverse</a> that I wanted a new .beat time
aligned to UTC instead of UTC+1, and was met with some small discussion about
aligning it to <a href="https://en.wikipedia.org/wiki/International_Atomic_Time">International Atomic Time</a> (TAI) instead and, well, here we are:
beatTAI or .tai for short.
</pre>


## Format

In line with Swatch .beat time, .tai is a day divided into 1000 and represented as ":xxx.xx".


## Math

beatTAI = (TAIhours &times; 3600 + TAIminutes &times; 60 + TAIseconds) / 86.4


## Etc

- [gbt](gbt)&thinsp;<sub>Go</sub> outputs current time in beatTAI using [brandondube/tai](https://github.com/brandondube/tai)
- [gbtgui](gbtgui)&thinsp;<sub>Go</sub> is a graphical beatTAI clock modified from [peterhellberg/beats](https://github.com/peterhellberg/beats)
- [i9w](i9w)&thinsp;<sub>C</sub> outputs current time in beatTAI using [libtai](https://cr.yp.to/libtai.html)
