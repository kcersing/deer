import {request} from '@umijs/max';
import {ItemListResp,ItemResp,Item} from "./data";
import {TreeResp,Tree,BaseResp} from "@/services/typings";

const headers = {
  'Content-Type': 'application/json', Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};


/** 创建 item  POST /service/item/create*/
export async function createItem(options?: { [key: string]: any }) {
  return request<ItemResp>('/service/item/create', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}

/** 更新item  POST /service/item/update*/
export async function updateItem(options?: { [key: string]: any }) {
  return request<ItemResp>('/service/item/update', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}

/** 删除item  POST /service/item/delete*/
export async function deleteItem(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/item/delete', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
/** 获取item  POST /service/item*/
export async function getItem(options?: { [key: string]: any }) {
console.log(options)
  return request<TreeResp>('/service/item', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
/** item 列表 POST /service/item/list*/
export async function getItemList(params: {
  // query
  /** 当前的页码 */
  current?: number; /** 页面的容量 */
  pageSize?: number; keywords?: string;

}, options?: { [key: string]: any },) {
  return request<ItemListResp>('/service/item/list', {
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




