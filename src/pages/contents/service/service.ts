import {request} from '@umijs/max';

import {TreeResp,Tree,BaseResp} from "@/services/typings";

const headers = {
  'Content-Type': 'application/json', Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};


/** 创建 product  POST /service/article/create*/
export async function createArticle(options?: { [key: string]: any }) {
  return request<ArticleResp>('/service/article/create', {
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

/** 更新article  POST /service/article/update*/
export async function updateArticle(options?: { [key: string]: any }) {
  return request<ArticleResp>('/service/article/update', {
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

/** 删除article  POST /service/article/delete*/
export async function deleteArticle(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/article/delete', {
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
/** 获取article  POST /service/article*/
export async function getArticle(options?: { [key: string]: any }) {
  return request<TreeResp>('/service/article', {
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

/** article 列表 POST /service/article/list*/
export async function getArticleList(params: {
  // query
  /** 当前的页码 */
  current?: number; /** 页面的容量 */
  pageSize?: number; keywords?: string;

}, options?: { [key: string]: any },) {
  return request<ArticleListResp>('/service/article/list', {
    method: 'POST', headers: {
      ...headers,
    }, data: {
      page: params.current, ...params,
    }, ...(options || {}),
  });
}
