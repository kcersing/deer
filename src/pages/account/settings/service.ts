import { request } from '@umijs/max';
import type { CurrentUser, GeographicItemType } from './data';
import {BaseResp} from "../../../services/typings";
import { headers } from '@/services/ant-design-pro/utils';

export async function queryCurrent(): Promise<{ data: CurrentUser }> {
  return request('/api/accountSettingCurrentUser');
}

export async function queryProvince(): Promise<{ data: GeographicItemType[] }> {
  return request('/api/geographic/province');
}

export async function queryCity(
  province: string,
): Promise<{ data: GeographicItemType[] }> {
  return request(`/api/geographic/city/${province}`);
}

export async function query() {
  return request('/api/users');
}


/**  更新密码  POST /service/user/change-password */
export async function changePassword(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/user/change-password', {
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
