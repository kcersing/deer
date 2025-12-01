
export type ApiListResp = {
  data?: Api[];
  baseResp?: BaseResp;
}
export type ApiResp = {
  data?: Api;
  baseResp?: BaseResp;
}
export type Api = {
  id?: number;
  path?: string;
  title?: string;
  group?: string;
  desc?: string;
  method?: string;
  status?: number;
}
