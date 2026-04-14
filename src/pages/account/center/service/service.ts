import {request} from '@umijs/max';
import {ItemListResp, ItemResp, Item, MessagesSendListResp, MessagesListResp, SmsSendListResp} from "./data";
import {TreeResp,Tree,BaseResp} from "@/services/typings";

const headers = {
  'Content-Type': 'application/json', Authorization: 'Bearer ' + sessionStorage.getItem('token') || '',
};




  /** 发送记录 POST /service/message/send/list */
  export async function getMessagesSendList(params: {
    // query
    /** 当前的页码 */
    current?: number; /** 页面的容量 */
    pageSize?: number; keywords?: string;

  }, options?: { [key: string]: any },) {
    return request<MessagesSendListResp>('/service/message/send/list', {
      method: 'POST', headers: {
        ...headers,
      },
      data: {
        page: params.current, ...params,
      }, ...(options || {}),
    });
  }




