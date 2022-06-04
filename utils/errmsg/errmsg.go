package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	ErrorUsernameUsed            = 1001 // >1000 用户错误
	ErrorPasswordWrong           = 1002
	ErrorUserCreateFail          = 1003
	ErrorUserQueryFail           = 1004
	ErrorUserEditFail            = 1005
	ErrorUserDeleteFail          = 1006
	ErrorUserDntExist            = 1007
	ErrorInsufficientPermissions = 1008
	ErrorTokenCreateFail         = 1009
	ErrorTokenDntExist           = 1010
	ErrorTokenRuntime            = 1011
	ErrorTokenWrong              = 1012
	ErrorTokenFormat             = 1013

	ErrorCategoryNameUsed   = 2001 // >2000 分类错误
	ErrorCategoryCreateFail = 2002
	ErrorCategoryQueryFail  = 2003
	ErrorCategoryEditFail   = 2004
	ErrorCategoryDeleteFail = 2005

	ErrorArticleCreateFail             = 3001 // >3000 文章错误
	ErrorArticleQueryUnderCategoryFail = 3002
	ErrorArticleQueryByIdFail          = 3003
	ErrorArticleQueryFail              = 3004
	ErrorArticleEditFail               = 3005
	ErrorArticleDeleteFail             = 3006
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ErrorUsernameUsed:            "用户名已存在",
	ErrorPasswordWrong:           "密码错误,请检查后重新输入",
	ErrorUserCreateFail:          "用户新增失败",
	ErrorUserQueryFail:           "查询用户列表失败",
	ErrorUserEditFail:            "编辑用户失败",
	ErrorUserDeleteFail:          "删除用户失败",
	ErrorUserDntExist:            "用户不存在",
	ErrorInsufficientPermissions: "用户权限不足",
	ErrorTokenCreateFail:         "TOKEN创建失败",
	ErrorTokenDntExist:           "TOKEN不存在",
	ErrorTokenRuntime:            "TOKEN已过期",
	ErrorTokenWrong:              "TOKEN不正确",
	ErrorTokenFormat:             "TOKEN格式错误",

	ErrorCategoryNameUsed:   "文章分类名已存在",
	ErrorCategoryCreateFail: "文章分类创建失败",
	ErrorCategoryQueryFail:  "查询文章分类列表失败",
	ErrorCategoryEditFail:   "编辑文章分类失败",
	ErrorCategoryDeleteFail: "删除文章分类失败",

	ErrorArticleCreateFail:             "文章创建失败",
	ErrorArticleQueryUnderCategoryFail: "查询分类下所有文章失败",
	ErrorArticleQueryByIdFail:          "根据ID查询文章失败",
	ErrorArticleQueryFail:              "查询文章列表失败",
	ErrorArticleEditFail:               "编辑文章失败",
	ErrorArticleDeleteFail:             "删除文章失败",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
