
// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';
import { headers } from '@/services/ant-design-pro/utils';


/** 创建字典信息 PSOT /service/dict/create */
export async function CreateDict(options?: { [key: string]: any }) {
  return request<API.NoticeIconList>('/service/dict/create', {
    method: 'PSOT',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 更新字典信息 PSOT /service/dict/update */
export async function UpdateDict(options?: { [key: string]: any }) {
  return request<API.Dict>('/service/dict/update', {
    method: 'PSOT',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 删除字典信息 PSOT /service/dict/delete */
export async function deleteDict(options?: { [key: string]: any }) {
  return request<API.Dict>('/service/dict/delete', {
    method: 'PSOT',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 获取字典列表 POST /service/dict/list */
export async function getDictList(options?: { [key: string]: any }) {
  return request<API.DictList>('/service/dict/list', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data:{
      name:options.name,
    },
    ...(options || {}),
  });
}
/** 创建字典键值信息 PSOT /service/dict/dictht/create */
export async function CreateDictht(options?: { [key: string]: any }) {
  return request<API.Dictht>('/service/dict/dictht/create', {
    method: 'PSOT',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 更新字典键值信息 PSOT /service/dict/dictht/update */
export async function UpdateDictht(options?: { [key: string]: any }) {
  return request<API.Dictht>('/service/dict/dictht/update', {
    method: 'PSOT',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 删除字典键值信息 PSOT /service/dict/dictht/delete */
export async function DeleteDictht(options?: { [key: string]: any }) {
  return request<API.Dictht>('/service/dict/dictht/delete', {
    method: 'PSOT',
    headers: {
      ...headers,
    },
    ...(options || {}),
  });
}
/** 根据字典名获取字典键值列表 POST /service/dict/dictht/list */
export async function getDicthtList(options?: { [key: string]: any }) {
  console.log(options)
  return request<API.DicthtList>('/service/dict/dictht/list', {
    method: 'POST',
    headers: {
      ...headers,
    },
    data:{
      name:options.name,
      dictId:options.dictId,
    },
    ...(options || {}),
  });
}
