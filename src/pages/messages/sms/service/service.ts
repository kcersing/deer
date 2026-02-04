import {SmsSendListResp} from "../../message/service/data";

const headers = {
  'Content-Type': 'application/json', Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};

/** 获取 sms  POST /service/message/sms */
export async function getSms(options?: { [key: string]: any }) {
  console.log(options)
  return request<SmsResp>('/service/message/sms', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      ...options,
    },
    ...(options || {}),
  });
}
/** sms  send 列表 POST /service/message/sms/send/list */
export async function getSmsSendList(params: {
  // query
  /** 当前的页码 */
  current?: number; /** 页面的容量 */
  pageSize?: number; keywords?: string;

}, options?: { [key: string]: any },) {
  return request<SmsSendListResp>('/service/message/sms/send/list', {
    method: 'POST', headers: {
      ...headers,
    },
    data: {
      page: params.current, ...params,
    }, ...(options || {}),
  });
}
