import {BaseResp} from "@/services/typings";

export type Article = {
  id?: number;
  title?: string;
  pic?: string[];
  content?: string;
  tagId?: number[];
  status?: number;
  createdId?: number;
  createdName?: string;
  createdAt?: string;
  updatedAt?: string;
}


export type ArticleResp={
  data?: Article;
  baseResp?: BaseResp;
}
export type ArticleListResp = {
  data?: Article[];
  baseResp?: BaseResp;
}




