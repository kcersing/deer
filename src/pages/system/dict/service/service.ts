// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';
import { headers } from '@/services/ant-design-pro/utils';

import {Dict,DictList,Dictht,DicthtList} from "./data";
import {BaseResp} from "@/services/typings";

/** 创建字典信息 POST /service/dict/create */
export async function createDict(options?: { [key: string]: any }) {
  return request<Dict>('/service/dict/create', {
    method: 'POST',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 更新字典信息 POST /service/dict/update */
export async function updateDict(options?: { [key: string]: any }) {
  return request<Dict>('/service/dict/update', {
    method: 'POST',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 删除字典信息 POST /service/dict/delete */
export async function deleteDict(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/dict/delete', {
    method: 'POST',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 获取字典列表 POST /service/dict/list */
export async function getDictList(options?: { [key: string]: any }) {
  return request<DictList>('/service/dict/list', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data:{
      key:options.key,
    },
    ...(options || {}),
  });
}
/** 创建字典键值信息 POST /service/dict/dictht/create */
export async function createDictht(options?: { [key: string]: any }) {
  return request<Dictht>('/service/dict/dictht/create', {
    method: 'POST',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 更新字典键值信息 POST /service/dict/dictht/update */
export async function updateDictht(options?: { [key: string]: any }) {
  return request<Dictht>('/service/dict/dictht/update', {
    method: 'POST',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 删除字典键值信息 POST /service/dict/dictht/delete */
export async function deleteDictht(options?: { [key: string]: any }) {
  return request<BaseResp>('/service/dict/dictht/delete', {
    method: 'POST',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 根据字典名获取字典键值列表 POST /service/dict/dictht/list */
export async function getDicthtList(options?: { [key: string]: any }) {
  console.log(options)
  return request<DicthtList>('/service/dict/dictht/list', {
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
