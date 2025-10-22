import { request } from '@umijs/max';
import multer from "@umijs/preset-umi/compiled/multer";




/** 登录接口 POST /service/user/login */
export async function login(body: LoginReq, options?: { [key: string]: any }) {
  return request<LoginResp>('/service/user/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
/** 创建接口 POST /service/user/create */
export async function create(body: CreateUserReq, options?: { [key: string]: any }) {
  return request<API.LoginResult>('/service/user/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
/** 获取用户信息 POST /service/user */
export async function user(body: { id:number }, options?: { [key: string]: any }) {
  return request<API.LoginResult>('/service/user', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    ...headers,
    },
    data: body,
    ...(options || {}),
  });
}

/** 获取user列表 POST /service/user/list*/
export async function userList(
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
  return request<API.RuleList>('/service/user/list', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: {
      page: params.current,
      ...params,
    },
    ...(options || {}),
  });
}

const headers = {
  Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};

type LoginReq = {
  username: string;
  password: string;
  captcha: string;
}


type CreateUserReq = {

}
type GetUserListReq = {

}



type BaseResp ={
  code?: number;
  message?: string;
  time?: string;
  total?: number;
}

export type LoginResp= {
  expire?: string;
  token?: string;
};
type UserResp ={
  data?:User;
  baseResp?: BaseResp;
}
type UserListResp ={
  data?:User[];
  baseResp?: BaseResp;
}


export type User = {
  id?: number;
  name?: string;
  avatar?: string;
  mobile?: string;
  status?: number;
  level?: number;
  gender?: string;
  birthday?: string;
  lastAt?: string;
  lastIp?: string;
  detail?: string;
  role?: Role;

};
type Role = {
  id?: number;
  name?: string;
  value?: string;
  defaultRouter?: string;
  remark?: number;
  apis?: number[];
}
