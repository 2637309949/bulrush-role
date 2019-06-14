# bulrush-role
## Example
Basic permissions and role control.
- EXAMPLE:   
```go
app.Use(&plugins.Role{
	FailureHandler: func(c *gin.Context, action string) {
		c.JSON(http.StatusForbidden, gin.H{
			"message": 	"No permission to access this api",
		})
	},
	RoleHandler:	func(c *gin.Context, action string) bool {
		fmt.Println("action :", action)
		return false
	},
	RoleHandler1:	func(c *gin.Context, action string) bool {
		fmt.Println("action :", action)
		return false
	},
})

app.Use(bulrush.PNQuick(func(testInject string, router *gin.RouterGroup, role *plugins.Role) {
	router.GET("/bulrushApp", role.Can("super@testPermts"), role.Can1("super@testPermts"), func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": 	testInject,
		})
	})
}))
```
## API

### RoleHandle defined api for role rule check

	func RoleHandler(c *gin.Context, action string) bool

	func RoleHandler1(c *gin.Context, action string) bool

	func RoleHandler2(c *gin.Context, action string) bool

	func RoleHandler3(c *gin.Context, action string) bool

	func RoleHandler4(c *gin.Context, action string) bool

	func RoleHandler5(c *gin.Context, action string) bool

	func RoleHandler6(c *gin.Context, action string) bool

### Can defined api for role rule use

	func Can(action string) gin.HandlerFunc

	func Can1(action string) gin.HandlerFunc

	func Can2(action string) gin.HandlerFunc

	func Can3(action string) gin.HandlerFunc

	func Can4(action string) gin.HandlerFunc

	func Can5(action string) gin.HandlerFunc

	func Can6(action string) gin.HandlerFunc


## MIT License

Copyright (c) 2018-2020 Double

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.