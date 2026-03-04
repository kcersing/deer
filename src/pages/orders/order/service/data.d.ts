import {BaseResp} from "@/services/typings";

export type OrderItem = {
  productId?: number;
  quantity?: number;
  price?: number;
  name?: string;
}

export type OrderPay = {
  remission?: number;
  pay?: number;
  reason?: string;
  payAt?: string;
  payWay?: string;
  paySn?: string;
  prepayId?: string;
  payExtra?: string;
  createdId?: number;
}
export type OrderRefund = {
  refundAt?: string;
  refundReason?: string;
  createdId?: number;
  refundAmount?: number;
}

export type Order = {
  id?: number;
  memberId?: number;
  sn?: string;
  totalAmount?: number;
  status?: number;
  orderItem?:OrderItem[];
  nature?: string;
  createdAt?: string;
  completionAt?: string;
  closeAt?: string;
  updatedAt?: string;
  cancelledReason?: string;
  orderPays?: OrderPay[];
  orderRefund?: OrderRefund;
  createdId?: number;
  createdName?: string;
}


export type OrderResp={
  data?: Order;
  baseResp?: BaseResp;
}
export type OrderListResp = {
  data?: Order[];
  baseResp?: BaseResp;
}


