import std.monkey.bridge.verylong

//println(`cat ~/.mk/lib/std/monkey/bridge/verylong.mk`)

x = verylong.Pyint("990000000000000000")
y = verylong.Pyint("12345678901234567890")

z = x + y

println(z.pynum)