import {Role} from "@/pages/auth/role/service/data";

export type MenuResp={
  data?: Menu;
  baseResp?: BaseResp;
}
export type MenuListResp = {
  data?: Menu[];
  baseResp?: BaseResp;
}

export type Menu = {
  id?: number;
  name?: string;
  parentId?: number;
  level?: number;
  path?: string;
  redirect?: string;
  component?: string;
  menuType?: number;
  hidden?: number;
  sort?: number;
  status?: number;
  url?: string;
  children?: Menu[];
  title?: string;
  type?: string;
  icon?: string;
}
