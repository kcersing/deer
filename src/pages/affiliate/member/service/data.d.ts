
export type Member = {
  id?: number;
  name?: string;
  avatar?: string;
  mobile?: string;
  status?: number;
  level?: number;
  gender?: number;
  birthday?: string;
  lastAt?: string;
  lastIp?: string;
  intention?: number;
  createdAt?: string;
  updatedAt?: string;
  createdId?: number;
  createdName?: string;

  source?: number;
  note?: string[];
}

export type MemberResp={
  data?: Member;
  baseResp?: BaseResp;
}
export type MemberListResp = {
  data?: Member[];
  baseResp?: BaseResp;
}
