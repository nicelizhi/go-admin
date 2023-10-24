package apis

import (
	"github.com/gin-gonic/gin"
)

const INDEX = `
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>Welcome</title>
<style>
body{
  margin:0; 
  padding:0; 
  overflow-y:hidden
}
</style>
</head>
<body>
<h2>pls use the api document file <a href="/swagger/admin/index.html">swagger/admin/index.html</a></h2>
</body>
</html>
`

func GoAdmin(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(200, INDEX)
}
