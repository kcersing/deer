// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';
import {headers} from "@/services/ant-design-pro/utils";



/** 此处后端没有提供注释 GET /service/ping */
export async function getNotices(options?: { [key: string]: any }) {
  return request<API.NoticeIconList>('/service/ping', {
    method: 'GET',
    ...(options || {}),
  });
}

