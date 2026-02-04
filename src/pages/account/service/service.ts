import { request } from '@umijs/max';
import {UserListResp,LoginReq,LoginResp,CreateUserReq,UserResp,UpdateUserReq} from "./data";

/** 登录接口 POST /service/user/login */
export async function login(body: LoginReq, options?: { [key: string]: any }) {
  return request<LoginResp>('/service/user/login', {
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
    data:{
      key:options.key,
      dictId:options.dictId,
    },
    ...(options || {}),
  });
}

/** 创建接口 POST /service/user/create */
export async function createUser(body: CreateUserReq, options?: { [key: string]: any }) {
  return request<UserResp>('/service/user/create', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: body,
    ...(options || {}),
  });
}

/** 更新接口 POST /service/user/update */
export async function updateUser(body: UpdateUserReq, options?: { [key: string]: any }) {
  return request<UserResp>('/service/user/update', {
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
  return request<UserResp>('/service/user', {
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
  return request<LoginResp>('/service/user', {
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
  return request<UserListResp>('/service/user/list', {
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





// base.NilResponse CreateDepartments(1: CreateDepartmentsReq req)(api.post = "/service/departments/create")
// base.NilResponse DeleteDepartments(1: base.IdReq req)(api.post = "/service/departments/delete")
// base.NilResponse UpdateDepartments(1: UpdateDepartmentsReq req)(api.post = "/service/departments/update")
// base.NilResponse GetDepartments(1: base.IdReq req)(api.post = "/service/departments")
// base.NilResponse GetDepartmentsList(1: GetDepartmentsListReq req)(api.post = "/service/departments/list")
