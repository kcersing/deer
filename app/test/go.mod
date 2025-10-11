module test

go 1.25.0

replace (
	common => ../../common
	gen => ../../gen
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
)

require golang.org/x/exp v0.0.0-20251009144603-d2f985daa21b
