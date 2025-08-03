package pkg

import "fmt"

var pkgName string = getPkgName()

const pkgNameStr string = "pkg"

func init() {
	fmt.Println("pkg init method invoke")
}

func getPkgName() string {
	fmt.Println("create var pkg pkg1Name")
	return pkgNameStr
}
