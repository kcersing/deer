import { request } from '@umijs/max';
import multer from "@umijs/preset-umi/compiled/multer";




/** 登录接口 POST /service/user/login */
export async function login(body: API.LoginReq, options?: { [key: string]: any }) {
  return request<API.LoginResp>('/service/user/login', {
    method: 'POST',
    headers: {

      ...headers,
    },
    data: body,
    ...(options || {}),
  });
}
/** 退出登录接口 POST /service/user/logout */
export async function logout(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/service/user/logout', {
    method: 'POST',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}

/** 创建接口 POST /service/user/create */
export async function createUser(body: API.CreateUserReq, options?: { [key: string]: any }) {
  return request<API.UserResp>('/service/user/create', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: body,
    ...(options || {}),
  });
}
/** 获取用户信息 POST /service/user */
export async function getUser(body: { id:number }, options?: { [key: string]: any }) {
  return request<API.LoginResp>('/service/user', {
    method: 'POST',
    headers: {
    ...headers,
    },
    data: body,
    ...(options || {}),
  });
}


/** 获取用户信息 POST /service/user */
export async function getRoleMenuAll(body: { roleId:number }, options?: { [key: string]: any }) {
  return request<API.LoginResp>('/service/user', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: body,
    ...(options || {}),
  });
}




/** 获取user列表 POST /service/user/list*/
export async function getUserList(
  params: {
    // query
    /** 当前的页码 */
    current?: number;
    /** 页面的容量 */
    pageSize?: number;
    keywords?: string;

  },
  options?: { [key: string]: any },
) {
  return request<API.UserListResp>('/service/user/list', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: {
      page: params.current,
      ...params,
    },
    ...(options || {}),
  });
}

const headers = {
  'Content-Type': 'application/json',
  Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};

