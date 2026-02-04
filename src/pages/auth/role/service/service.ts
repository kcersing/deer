import {request} from '@umijs/max';
import {RoleListResp,RoleResp,Role} from "./data";
import {TreeResp,Tree,BaseResp} from "@/services/typings";


/** 获取role 列表 POST /service/role/list*/
export async function getRoleList(params: {
  // query
  /** 当前的页码 */
  current?: number; /** 页面的容量 */
  pageSize?: number; keywords?: string;

}, options?: { [key: string]: any },) {
  return request<RoleListResp>('/service/role/list', {
    method: 'POST', headers: {
      ...headers,
    }, data: {
      page: params.current, ...params,
    }, ...(options || {}),
  });
}

/** 获取role  POST /service/role/tree*/
export async function getRoleTree() {
  return request<TreeResp>('/service/role/tree', {
    method: 'POST', headers: {
      ...headers,
    },
  });
}

/** 删除role  POST /service/role/delete*/
export async function deleteRole(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/role/delete', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}

const headers = {
  'Content-Type': 'application/json', Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};

/** 更新role  POST /service/role/update*/
export async function updateRole(options?: { [key: string]: any }) {
  return request<RoleResp>('/service/role/update', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
/** 创建 role  POST /service/role/create*/
export async function createRole(options?: { [key: string]: any }) {
  return request<RoleResp>('/service/role/create', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
/** 获取 role  POST /service/role */
export async function getRole(options?: { [key: string]: any }) {
  return request<RoleResp>('/service/role', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}


/** 创建 role 菜单权限  POST /service/role/create/menu */
export async function CreateRoleMenu(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/role/create/menu', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}

/** 创建 role API权限  POST /service/role/create/api */
export async function CreateRoleApi(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/role/create/api', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}


/** 获取 role API列表  POST /service/role/api */
export async function GetRoleApi(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/role/api', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}

/** 获取 role 菜单列表  POST /service/role/menu */
export async function GetRoleMenu(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/role/menu', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
