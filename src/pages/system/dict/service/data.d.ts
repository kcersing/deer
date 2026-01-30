import {BaseResp} from "@/services/typings";
export type Dictht = {
  id?: number;
  title?: string;
  value?: string;
  dictId?: number;
  status?: number;
  createdAt?: string;
  updatedAt?: string;
  createdId?: number;
  createdName?: string;
}


export type DicthtList = {
  data?: Dictht[];
  /** 列表的内容总数 */
  baseResp?: BaseResp;
}
export type Dict = {
  id?: number;
  title?: string;
  code?: string;
  status?: number;
  desc?: string;
  createdAt?: string;
  updatedAt?: string;
  createdId?: number;
  createdName?: string;
}

export type DictList = {
  data?: Dict[];
  /** 列表的内容总数 */
  baseResp?: BaseResp;
}



