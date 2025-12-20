import {request} from '@umijs/max';
import {MemberListResp,MemberResp,Member} from "./data";
import {TreeResp,Tree,BaseResp} from "@/services/typings";

/** 创建 Member  POST /service/member/create */
export async function createMember(options?: { [key: string]: any }) {
  return request<MemberResp>('/service/member/create', {
    method: 'POST', headers: {
      ...headers,
    }, ...(options || {}),
  });
}
/** 删除 Member  POST /service/member/delete*/
export async function deleteMember(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/member/delete', {
    method: 'POST', headers: {
      ...headers,
    }, ...(options || {}),
  });
}
/** 更新 Member  POST /service/member/update*/
export async function updateMember(options?: { [key: string]: any }) {
  return request<MemberResp>('/service/member/update', {
    method: 'POST', headers: {
      ...headers,
    }, ...(options || {}),
  });
}

/** 获取 Member  POST /service/member */
export async function getMember(options?: { [key: string]: any }) {
  return request<MemberResp>('/service/member', {
    method: 'POST', headers: {
      ...headers,
    }, ...(options || {}),
  });
}
/** 获取 Member 列表 POST /service/member/list */
export async function getMemberList(params: {
  // query
  /** 当前的页码 */
  current?: number; /** 页面的容量 */
  pageSize?: number; keywords?: string;

}, options?: { [key: string]: any },) {
  return request<MemberListResp>('/service/member/list', {
    method: 'POST', headers: {
      ...headers,
    }, data: {
      page: params.current, ...params,
    }, ...(options || {}),
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
