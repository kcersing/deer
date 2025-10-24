namespace go dict
include "../base/base.thrift"
// 字典信息
struct Dict {
    1:  i64 id=0 (api.raw = "id" )
    2:  string title="" (api.raw = "title" )
    3:  string name="" (api.raw = "name" )
    5:  i64 status=1 (api.raw = "status" )
    6:  string description="" (api.raw = "description" )
    7:  string createdAt="" (api.raw = "createdAt" )
    8:  string updatedAt="" (api.raw = "updatedAt" )
}

// 字典键值信息
struct Dictht {
    1:  i64 id=0 (api.raw = "id" )
    2:  string title="" (api.raw = "title" )
    3:  string key="" (api.raw = "key" )
    4:  string value="" (api.raw = "value" )
    5:  i64 status=1 (api.raw = "status" )
    6:  string createdAt="" (api.raw = "createdAt" )
    7:  string updatedAt="" (api.raw = "updatedAt" )
    8:  i64 parentID (api.raw = "parentID" )
}
// 字典列表请求数据
struct DictListReq {
    1:  optional string title  = ""(api.raw = "title" )
    2:  optional string name = "" (api.raw = "name" )
    3:  optional i64 page=1 (api.raw = "page")
    4:  optional i64 pageSize=100 (api.raw = "pageSize")
}

//字典名获取字典键值请求数据
struct DicthtListReq{
    1:  optional string name = "" (api.raw = "name" )
    2:  optional i64 dictionaryId= 0 (api.raw = "dictionaryId" )
}



service DictService {
    // 创建字典信息
      base.NilResponse CreateDict(1: Dict req) (api.post = "/service/dict/create")
      // 更新字典信息
      base.NilResponse UpdateDict(1: Dict req) (api.post = "/service/dict/update")
      // 删除字典信息
      base.NilResponse DeleteDict(1: base.IdReq req) (api.post =  "/service/dict")
      // 获取字典列表
      base.NilResponse DictList(1: DictListReq req) (api.get = "/service/dict/list")
      // 创建字典键值信息
      base.NilResponse CreateDictht(1: Dictht req) (api.post = "/service/dict/dictht/create")
      // 更新字典键值信息
      base.NilResponse UpdateDictht(1: Dictht req) (api.post = "/service/dict/dictht/update")
      // 删除字典键值信息
      base.NilResponse DeleteDictht(1: base.IdReq req) (api.get = "/service/dict/dictht")
      // 根据字典名获取字典键值列表
      base.NilResponse DicthtList(1: DicthtListReq req) (api.post = "/service/dict/dictht/list")

}