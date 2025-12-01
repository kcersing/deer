
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


export type UserResp = {
  data?: User;
  baseResp?: BaseResp;
}
export type UserListResp = {
  data?: User[];
  baseResp?: BaseResp;
}

export type UpdateUserReq= {}

export type CreateUserReq = {}
export type GetUserListReq = {}

export type User = {
  id?: number;
  name?: string;
  username?:string;
  avatar?: string;
  mobile?: string;
  status?: number;
  level?: number;
  gender?: string;
  birthday?: string;
  lastAt?: string;
  lastIp?: string;
  detail?: string;
  role?: Role;
};
