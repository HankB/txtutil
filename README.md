# txtutil
Go (golang) text utility code not specific to a particular project.

This began as the IsPrint() and Dump() functions in the s20 project
but they have nothing to do with the Orvibo S20 - Dump is just
useful for debugging to see messages as sent and received. I figured
I'd create a separate 'util' repo and put stuff like this in there.
And I did. Then I read at https://blog.golang.org/package-names
"Avoid meaningless package names. Packages named util, common, or
misc provide clients with no sense of what the package contains."

I renamed to 'txtutil' hoping that was a good name (and not really
thinking so.)

## testing
`go test [-v] example_test.go`
