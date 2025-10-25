import { request } from '@umijs/max';

/** 获取menu 列表 POST /service/menu/list*/
export async function menuList(
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

const headers = {
  'Content-Type': 'application/json',
  Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};


