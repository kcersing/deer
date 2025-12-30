namespace go venue
include "../base/base.thrift"
include "../base/venue.thrift"

struct CreatePlaceReq{
     2: optional string name="" (api.raw = "name")
     3: optional i64 venueId=0 (api.raw = "venueId")
     4: optional string pic="" (api.raw = "pic")
     /**是否开放:0关闭 1开放;*/
     5: optional i64 status=1 (api.raw = "status")
     8: optional i64 number=0 (api.raw = "number")
     9: optional string desc="" (api.raw = "desc")
     /**关联座位*/
     10: optional list<list<base.Seat>> seat=0 (api.raw = "seat")
     11:  optional string startTime="" (api.raw = "startTime")
     12:  optional string endTime="" (api.raw = "endTime")
}

struct UpdatePlaceReq{

}
struct GetPlaceListReq{
     1:optional i64 page=1
     2:optional i64 pageSize=10
}

struct CreateVenueReq{
   2:  optional string name="" (api.raw = "name")
    3:  optional string address="" (api.raw = "address")
    5:  optional string latitude="" (api.raw = "latitude")
    6:  optional string longitude =""(api.raw = "longitude")
    7:  optional string mobile="" (api.raw = "mobile")
    8:  optional string pic =""(api.raw = "pic")
    9:  optional string desc="" (api.raw = "desc")
     /**是否开放:0关闭 1开放;*/
    10:  optional i64 status=1 (api.raw = "status")
    13:  optional string email =""(api.raw = "email")
    14:  optional string startTime="" (api.raw = "startTime")
    15:  optional string endTime="" (api.raw = "endTime")
}
struct UpdateVenueReq{

}
struct GetVenueListReq{
     1:optional i64 page=1
     2:optional i64 pageSize=10
}


struct PlaceResp{
    1:optional venue.Place data= {}
    255:optional base.BaseResp baseResp={}
}
struct PlaceListResp{
    1:optional list<venue.Place> data= []
    255:optional base.BaseResp baseResp={}
}

struct VenueResp{
    1:optional venue.Venue data= {}
    255:optional base.BaseResp baseResp={}
}
struct VenueListResp{
    1:optional list<venue.Venue> data= []
    255:optional base.BaseResp baseResp={}
}
service VenueService {

    PlaceResp CreatePlace(1: CreatePlaceReq req)
    PlaceResp UpdatePlace(1: UpdatePlaceReq req)
    base.NilResponse DeletePlace(1: base.IdReq req)
    PlaceResp GetPlace(1: base.IdReq req)
    PlaceListResp GetPlaceList(1: GetPlaceListReq req)

    VenueResp CreateVenue(1: CreateVenueReq req)
    VenueResp UpdateVenue(1: UpdateVenueReq req)
    base.NilResponse DeleteVenue(1: base.IdReq req)
    VenueResp GetVenue(1: base.IdReq req)
    VenueListResp GetVenueList(1: GetVenueListReq req)
}