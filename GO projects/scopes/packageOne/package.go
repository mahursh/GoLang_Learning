package packageOne

// lowercase (first letter of variables name)
// private (NOT exported)
// only available inside this package.
var privateVar = " I am private"

// uppercase (first letter of variables name)
// public (exported)
//availabale outside this packaage.
var PublicVar = "I am public (or exported)"

func notExported() {}

func Exported() {}
