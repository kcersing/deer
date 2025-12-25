
export type Item = {
  id?: number;
  name?: string;
  pic?: string;
  desc?: string;
  status?: number;
  duration?: number;
  length?: number;
  count?: number;
  price?: number;
  tagId?: number[];
  tagName?: string[];

  type?: string;
  code?: string;
  createdId?: number;
  createdName?: string;
  createdAt?: string;
  updatedAt?: string;
}

export type ItemResp={
  data?: Item;
  baseResp?: BaseResp;
}
export type ItemListResp = {
  data?: Item[];
  baseResp?: BaseResp;
}




