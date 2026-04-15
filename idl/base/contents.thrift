namespace go base

struct Article {
    1:optional i64 id=0,
    2:optional string title="",
    3:optional string content="",
    5:optional list<i64> tagId=[],
    6:optional i64 createdId=0,
    7:optional string createdAt="",
    8:optional string updatedAt=""
    9:optional list<string> pic="",
    10:optional i64 status=0,
}