import {request} from '@umijs/max';
import {MemberListResp,MemberResp,Member} from "./data";
import {TreeResp,Tree,BaseResp} from "@/services/typings";

//     base.NilResponse GetOrder(1:GetOrderReq req)(api.post = "/service/order")
//     base.NilResponse GetOrderList(1:GetOrderListReq req)(api.post = "/service/order/list")
//     base.NilResponse DeleteOrder(1:base.IdReq req)(api.post = "/service/order/delete")
//     base.NilResponse CreateOrder(1:CreateOrderReq req)(api.post = "/service/order/create")
//     base.NilResponse Payment(1:PaymentReq req)(api.post = "/service/order/payment")
//     base.NilResponse CancelledOrder(1:CancelledOrderReq req)(api.post = "/service/order/cancelled")
//     base.NilResponse RefundOrder(1:RefundOrderReq req)(api.post = "/service/order/refund")

/** 获取 order  POST /service/order */
export async function getOrder(options?: { [key: string]: any }) {
  return request<OrderResp>('/service/order', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
/** 获取 Order 列表 POST /service/order/list */
export async function getOrderList(params: {
  // query
  /** 当前的页码 */
  current?: number; /** 页面的容量 */
  pageSize?: number; keywords?: string;

}, options?: { [key: string]: any },) {
  return request<MemberOrderResp>('/service/order/list', {
    method: 'POST', headers: {
      ...headers,
    }, data: {
      page: params.current, ...params,
    }, ...(options || {}),
  });
}
/** 创建 Member  POST /service/member/create */
export async function createMember(options?: { [key: string]: any }) {
  return request<MemberResp>('/service/member/create', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
/** 删除 Order  POST /service/order/delete*/
export async function deleteOrder(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/order/delete', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
/** 更新 Member  POST /service/member/update*/
export async function updateMember(options?: { [key: string]: any }) {
  return request<MemberResp>('/service/member/update', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}




/** ChangePassword  POST /service/member/change-password */
export async function changePassword() {
  return request<TreeResp>('/service/member/change-password', {
    method: 'POST', headers: {
      ...headers,
    },
  });
}

const headers = {
  'Content-Type': 'application/json', Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};
