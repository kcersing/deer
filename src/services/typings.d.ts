// @ts-ignore
/* eslint-disable */

export type BaseResp = {
  code?: number;
  message?: string;
  time?: string;
  total?: number;
}

export type TreeResp = {
  data?: Tree[];
  baseResp?: BaseResp;
}

export type Tree = {
  title?: string;
  value?: string;
  key?: string;
  method?: string;
  children?: Tree[];
}

export type LoginReq = {
  username: string;
  password: string;
  captcha: string;
}
export type LoginResp = {
  expire?: string;
  token?: string;
  status?: number;
  type?: string;
};


