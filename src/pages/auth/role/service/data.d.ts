
export type Role = {
  id?: number;
  name?: string;
  code?: string;
  desc?: number;
  order_no?: number;
  apis?: number[];
  menus?: number[];
}
export type RoleResp={
  data?: Role;
  baseResp?: BaseResp;
}
export type RoleListResp = {
  data?: Role[];
  baseResp?: BaseResp;
}
