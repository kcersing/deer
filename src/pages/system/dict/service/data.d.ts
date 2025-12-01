export type Dictht = {
  id?: number;
  dictId?: number;
  key?: string;
  value?: string;
  status?: number;
  title?: string;
  createdAt?: string;
  updatedAt?: string;
}

export type DicthtList = {
  data?: Dictht[];
  /** 列表的内容总数 */
  baseResp?: BaseResp;
}
export type Dict = {
  id?: number;
  desc?: string;
  name?: string;
  status?: number;
  title?: string;
  createdAt?: string;
  updatedAt?: string;
}

export type DictList = {
  data?: Dict[];
  /** 列表的内容总数 */
  baseResp?: BaseResp;
}



