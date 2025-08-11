package pkg1

import "fmt"

var pkgName string = getPkgName()

const pkgNameStr string = "pkg1"

func init() {
	fmt.Println("pkg1 init method invoke")
}

func getPkgName() string {
	fmt.Println("create var pkg1 pkg1Name")
	return pkgNameStr
}
