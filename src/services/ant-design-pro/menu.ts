import { request } from '@umijs/max';
import type { TreeDataNode } from 'antd';
/** 获取menu 列表 POST /service/menu/list*/
export async function getMenuList(
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
  return request<API.MenuListResp>('/service/menu/list', {
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

/** 获取menu 列表 POST /service/menu/tree*/
export async function getMenuTree() {
  return request<API.TreeResp>('/service/menu/tree', {
    method: 'POST',
    headers: {
      ...headers,
    },
  });
}

/** 删除menu 列表 POST /service/menu/delete*/
export async function deleteMenu(options?: { [key: string]: any }) {
  return request<API.TreeResp>('/service/menu/delete', {
    method: 'POST',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
const headers = {
  'Content-Type': 'application/json',
  Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};


export async function updateMenu(options?: { [key: string]: any }) {

  console.log(options)
  return
  return request<API.TreeResp>('/service/menu/update', {
    method: 'POST',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}

export async function createMenu(options?: { [key: string]: any }) {
  return request<API.TreeResp>('/service/menu/create', {
    method: 'POST',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
export async function getMenu(options?: { [key: string]: any }) {
  return request<API.TreeResp>('/service/menu', {
    method: 'POST',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}

