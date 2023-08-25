package main

import (
	"regexp"
)

var exp = regexp.MustCompile(`^(?P<remote>[^ ]*) (?P<host>[^ ]*) (?P<user>[^ ]*) \[(?P<time>[^\]]*)\] \"(?P<method>\w+)(?:\s+(?P<path>[^\"]*?)(?:\s+\S*)?)?\" (?P<status_code>[^ ]*) (?P<size>[^ ]*)(?:\s"(?P<referer>[^\"]*)") "(?P<agent>[^\"]*)" (?P<urt>[^ ]*)$`)
