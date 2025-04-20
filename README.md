# gqlgen-v0.17.71-issue

With gqlgen v0.17.70, this respository works.
https://github.com/skaji/gqlgen-v0.17.71-issue/tree/gqlgen-v0.17.70

Note that we have the following settings:

* We define "type Status" which has "MarshalJSON" and "UnmarshalJSON" methods
  * https://github.com/skaji/gqlgen-v0.17.71-issue/blob/gqlgen-v0.17.70/enum/enum.go
* We use that model for graphql enum Status
  * https://github.com/skaji/gqlgen-v0.17.71-issue/blob/gqlgen-v0.17.70/gqlgen.yml#L17-L18

---

After upgrading gqlgen from v0.17.70 to v0.17.71, we got the following errors:

```
❯ go run github.com/99designs/gqlgen
validation failed: packages.Load: -: # github.com/skaji/gqlgen-v0.17.71-issue/graph
graph/generated.go:3595:27: cannot use v (variable of interface type any) as []byte value in argument to res.UnmarshalJSON: need type assertion
graph/generated.go:3600:12: ec._Status undefined (type *executionContext has no field or method _Status)
/Users/skaji/src/github.com/skaji/gqlgen-v0.17.71-issue/graph/generated.go:3595:27: cannot use v (variable of interface type any) as []byte value in argument to res.UnmarshalJSON: need type assertion
/Users/skaji/src/github.com/skaji/gqlgen-v0.17.71-issue/graph/generated.go:3600:12: ec._Status undefined (type *executionContext has no field or method _Status)

exit status 1
```
