package xerr

var (
    ErrNoPermission   = Unauthorized("没有权限进行此操作")
    ErrInvalidRequest = BadRequest("无效的请求，请检查参数是否正确")
    ErrParamEmpty     = BadRequest("请求参数为空")
    ErrParamInvalid   = BadRequest("参数无效，请检查参数是否满足要求")
)
