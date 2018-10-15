Iota
----

Go's iota identifier is used in const declarations to simplify definitions of incrementing numbers. Because it can be
used in expressions, it provides a generality beyond that of simple enumerations.

The values of iota is reset to 0 whenever the reserved word const appears in the source (i.e. each const block) and
increments by one after each ConstSpec e.g. each Line. This can be combined with the constant shorthand (leaving out
everything after the constant name) to very concisely define related constants.