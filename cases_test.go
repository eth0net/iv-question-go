package pathbuilder

type customTest struct {
	separator  string
	components []string
	expected   string
}

var customTests = []customTest{
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
}
