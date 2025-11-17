// @ts-ignore
/* eslint-disable */
declare namespace API {
  type User = {
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
  type LoginReq = {
    username: string;
    password: string;
    captcha: string;
  }


  type CreateUserReq = {}
  type GetUserListReq = {}

  type BaseResp = {
    code?: number;
    message?: string;
    time?: string;
    total?: number;
  }

  export type LoginResp = {
    expire?: string;
    token?: string;
    status?: number;
    type?: string;
  };
  type UserResp = {
    data?: User;
    baseResp?: BaseResp;
  }
  type UserListResp = {
    data?: User[];
    baseResp?: BaseResp;
  }

  type Role = {
    id?: number;
    name?: string;
    code?: string;
    desc?: number;
    order_no?: number;
    apis?: number[];
    menus?: number[];
  }

  type RoleListResp = {
    data?: Role[];
    baseResp?: BaseResp;
  }
  type MenuListResp = {
    data?: Menu[];
    baseResp?: BaseResp;
  }
  type TreeResp = {
    data?: Tree[];
    baseResp?: BaseResp;
  }

  type Tree = {
    title?: string;
    value?: string;
    key?: string;
    method?: string;
    children?: Tree[];

  }

  type Menu = {
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


  type Api = {
    id?: number;
    path?: string;
    title?: string;
    group?: string;
    desc?: string;
    method?: string;
    status?: number;
  }

  type Dictht = {
    id?: number;
    dictId?: number;
    key?: string;
    value?: string;
    status?: number;
    title?: string;
    createdAt?: string;
    updatedAt?: string;
  }

  type DicthtList = {
    data?: Dictht[];
    /** 列表的内容总数 */
    total?: number;
  }
  type Dict = {
    id?: number;
    desc?: string;
    name?: string;
    status?: number;
    title?: string;
    createdAt?: string;
    updatedAt?: string;
  }
  type DictList = {
    data?: Dict[];
    /** 列表的内容总数 */
    total?: number;
  }


  type LoginResult = {
    status?: string;
    type?: string;
    currentAuthority?: string;
  };

  type PageParams = {
    current?: number;
    pageSize?: number;
  };

  type RuleListItem = {
    key?: number;
    disabled?: boolean;
    href?: string;
    avatar?: string;
    name?: string;
    owner?: string;
    desc?: string;
    callNo?: number;
    status?: number;
    updatedAt?: string;
    createdAt?: string;
    progress?: number;
  };

  type RuleList = {
    data?: RuleListItem[];
    /** 列表的内容总数 */
    total?: number;
    success?: boolean;
  };

  type FakeCaptcha = {
    code?: number;
    status?: string;
  };

  type LoginParams = {
    username?: string;
    password?: string;
    autoLogin?: boolean;
    type?: string;
  };

  type ErrorResponse = {
    /** 业务约定的错误码 */
    errorCode: string;
    /** 业务上的错误信息 */
    errorMessage?: string;
    /** 业务上的请求是否成功 */
    success?: boolean;
  };

  type NoticeIconList = {
    data?: NoticeIconItem[];
    /** 列表的内容总数 */
    total?: number;
    success?: boolean;
  };

  type NoticeIconItemType = 'notification' | 'message' | 'event';

  type NoticeIconItem = {
    id?: string;
    extra?: string;
    key?: string;
    read?: boolean;
    avatar?: string;
    title?: string;
    status?: string;
    datetime?: string;
    description?: string;
    type?: NoticeIconItemType;
  };
}
