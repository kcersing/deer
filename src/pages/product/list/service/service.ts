import {request} from '@umijs/max';
import {ItemListResp,ItemResp,Item} from "./data";
import {TreeResp,Tree,BaseResp} from "@/services/typings";

const headers = {
  'Content-Type': 'application/json', Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};


/** 创建 product  POST /service/product/create*/
export async function createProduct(options?: { [key: string]: any }) {
  return request<ProductResp>('/service/product/create', {
    method: 'POST', headers: {
      ...headers,
    }, ...(options || {}),
  });
}

/** 更新product  POST /service/product/update*/
export async function updateProduct(options?: { [key: string]: any }) {
  return request<ProductResp>('/service/product/update', {
    method: 'POST', headers: {
      ...headers,
    }, ...(options || {}),
  });
}

/** 删除product  POST /service/product/delete*/
export async function deleteProduct(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/product/delete', {
    method: 'POST', headers: {
      ...headers,
    }, ...(options || {}),
  });
}
/** 获取product  POST /service/product*/
export async function getProduct() {
  return request<TreeResp>('/service/product', {
    method: 'POST', headers: {
      ...headers,
    },
  });
}

/** online product  POST /service/product/online */
export async function onlineProduct() {
  return request<TreeResp>('/service/product/online', {
    method: 'POST', headers: {
      ...headers,
    },
  });
}
/**  offline product  POST /service/product/offline */
export async function offlineProduct() {
  return request<TreeResp>('/service/product/offline', {
    method: 'POST', headers: {
      ...headers,
    },
  });
}
/**  search product  POST /service/product/search */
export async function searchProduct() {
  return request<TreeResp>('/service/product/search', {
    method: 'POST', headers: {
      ...headers,
    },
  });
}


/** product 列表 POST /service/product/list*/
export async function getProductList(params: {
  // query
  /** 当前的页码 */
  current?: number; /** 页面的容量 */
  pageSize?: number; keywords?: string;

}, options?: { [key: string]: any },) {
  return request<ProductListResp>('/service/product/list', {
    method: 'POST', headers: {
      ...headers,
    }, data: {
      page: params.current, ...params,
    }, ...(options || {}),
  });
}
