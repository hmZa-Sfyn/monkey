import std.monkey.bridge.verylong

//println(`cat ~/.mk/lib/std/monkey/bridge/verylong.mk`)

xx = 424
yy = 432

x = verylong.Pyint(xx+"**"+xx)
y = verylong.Pyint(yy+"**"+yy)

z = x + y

println(z.pynum)