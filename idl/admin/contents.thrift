namespace go contents
include "../base/base.thrift"
include "../base/contents.thrift"


struct CreateArticleReq{
    2:optional string title="",
    3:optional string content="",
    5:optional list<i64> tagId=[],
    9:optional list<string> pic="",
    6:optional i64 status=0,
}
struct UpdateArticleReq{
    1:optional i64 id=0,
    2:optional string title="",
    3:optional string content="",
    5:optional list<i64> tagId=[],
    6:optional i64 status=0,
    9:optional list<string> pic="",
}
struct ArticleListReq{
    1:optional i64 page=1
    2:optional i64 pageSize=10
    3:optional string keyword=""
    5:optional list<i64> tagId=[],
}
struct ArticleListResp{
    1:optional list<contents.Article> data= []
    255:optional base.BaseResp baseResp={}

}
struct ArticleResp{
    1:optional contents.Article data={},
    255:optional base.BaseResp baseResp={}
}
service ContentsService  {

      ArticleResp GetArticle(1:base.IdReq req)(api.post = "/service/article")
      ArticleResp CreateArticle(1: CreateArticleReq req)(api.post = "/service/article/create")
      ArticleResp UpdateArticle(1: UpdateArticleReq req)(api.post = "/service/article/update")
      base.BaseResp DeleteArticle(1:base.IdReq req)(api.post = "/service/article/delete")
      ArticleListResp ArticleList(1: ArticleListReq req)(api.post = "/service/article/list")
}