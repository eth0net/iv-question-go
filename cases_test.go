package pathbuilder

type customTest struct {
	separator  string
	components []string
	expected   string
}

var customTestCases = []customTest{
	{"/", []string{"abc"}, "abc"},
	{"/", []string{"", "abc"}, "abc"},
	{"/", []string{"abc", "d"}, "abc/d"},
	{"/", []string{"/abc"}, "/abc"},
	{"/", []string{"/abc", "", "d"}, "/abc/d"},
	{"/", []string{"/abc", "/d"}, "/abc/d"},
	{"/", []string{"/abc", "/d", "e/"}, "/abc/d/e/"},
	{"/", []string{"/abc", "/d", "e/", "f"}, "/abc/d/e/f"},
	{"/", []string{"/abc", "/d", "0", "f"}, "/abc/d/0/f"},
	{"/", []string{"/abc", "/d", "", "f"}, "/abc/d/f"},
	{"/", []string{"abc", "/d//e/", "f"}, "abc/d//e/f"},
	{"/", []string{"//cdn.xyz.com", "images"}, "//cdn.xyz.com/images"},
	{"|", []string{"a", "b"}, "a|b"},
	{"|", []string{"a", "|b|", "c"}, "a|b|c"},
	{"~~", []string{"a/", "b"}, "a/~~b"},
}

type test struct {
	components []string
	expected   string
}

var urlTestCases = []test{
	{[]string{""}, ""},
	{[]string{"/", ""}, "/"},
	{[]string{"abc"}, "abc"},
	{[]string{"a", "b"}, "a/b"},
	{[]string{"/", "test"}, "/test"},
	{[]string{"/", "/test"}, "/test"},
	{[]string{"", "/", "/test"}, "/test"},
	{[]string{"/", "", "/test", ""}, "/test"},
	{[]string{"/", "c", "4", "ab"}, "/c/4/ab"},
	{[]string{"/", "c", "0", "ab"}, "/c/0/ab"},
	{[]string{"/", "c", "", "ab"}, "/c/ab"},
	{[]string{"//cdn.domain.tld", "", "/test", ""}, "//cdn.domain.tld/test"},
	{[]string{"/test/", "/subdir/test/"}, "/test/subdir/test/"},
	{[]string{"/test", "/subdir/test", "/"}, "/test/subdir/test/"},
	{[]string{"test", "/subdir/test/"}, "test/subdir/test/"},
	{[]string{"test", "/subdir//test/"}, "test/subdir//test/"},
}

var systemTestCases = []test{
	{[]string{"a", "b"}, "a/b"},
}

var unixTestCases = []test{
	{[]string{"a", "b"}, "a/b"},
}

var windowsTestCases = []test{
	{[]string{"a", "b"}, "a\\b"},
}
