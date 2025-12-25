
export type Product = {
  id?: number;
  name?: string;
  pic?: string;
  desc?: string;
  status?: number;
  statusName?: string;
  createdId?: number;
  createdName?: string;
  createdAt?: string;
  updatedAt?: string;
  price?: number;
  stock?: number;
  isSales?: number[];
  signSalesAt?: string;
  endSalesAt?: string;
  items?: Item[];
}


export type ProductResp={
  data?: Product;
  baseResp?: BaseResp;
}
export type ProductListResp = {
  data?: Product[];
  baseResp?: BaseResp;
}




