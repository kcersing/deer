import { LinkOutlined,PlusCircleFilled, SearchOutlined,HeartOutlined, SmileOutlined } from '@ant-design/icons';
import type { Settings as LayoutSettings } from '@ant-design/pro-components';

import { SettingDrawer ,ProBreadcrumb} from '@ant-design/pro-components';
import type { RequestConfig, RunTimeLayoutConfig } from '@umijs/max';
import { history, Link } from '@umijs/max';
import React, { useState, useReducer, useEffect, useRef, memo, useCallback } from 'react';
import { Input, Space } from 'antd';
import {
  AvatarDropdown,
  AvatarName,
  Footer,
  Question,
  SelectLang,
} from '@/components';
import defaultSettings from '../config/defaultSettings';
import { errorConfig } from './requestErrorConfig';
import '@ant-design/v5-patch-for-react-19';
import { getMenuList} from "@/services/ant-design-pro/menu";
import {getUser, getRoleMenuAll} from '@/services/ant-design-pro/user';
import type { MenuDataItem } from '@ant-design/pro-components';

import * as allIcons from '@ant-design/icons';
//https://www.iconfont.cn/collections/detail?cid=9402
const isDev = process.env.NODE_ENV === 'development';
const loginPath = '/user/login';

/**
 * @see https://umijs.org/docs/api/runtime-config#getinitialstate
 * */
export async function getInitialState(): Promise<{
  settings?: Partial<LayoutSettings>;
  currentUser?: API.User;
  loading?: boolean;
  fetchUserInfo?: () => Promise<API.User | undefined>;
  menuList?: MenuDataItem;
}> {

  const menuList = async () => {
    try {
      const msg = await getMenuList({});
      console.log(msg)
      // if(msg.code !==0){
      //   history.push(loginPath);
      // }
      return msg.data;
    } catch (_error) {
      // history.push(loginPath);
    }
    return undefined;
  };

  const fetchUserInfo = async () => {
    try {
      const msg = await getUser(
        {id:0}
        // {skipErrorHandler: true,}
      );
      console.log(msg)
      if(msg.code !==0){
        history.push(loginPath);
      }
      return msg.data;
    } catch (_error) {
      history.push(loginPath);
    }
    return undefined;
  };
  // 如果不是登录页面，执行
  const { location } = history;
  if (
    ![loginPath, '/user/register', '/user/register-result'].includes(
      location.pathname,
    )
  ) {
    const currentUser = await fetchUserInfo();
    return {
      fetchUserInfo,
      currentUser,
      settings: defaultSettings as Partial<LayoutSettings>,
    };
  }
  return {
    fetchUserInfo,
    settings: defaultSettings as Partial<LayoutSettings>,
  };
}

const IconMap = (icon:string, iconType = 'Outlined')=>{
  let fixIconName = icon.slice(0, 1).toLocaleUpperCase() + icon.slice(1) + iconType;
  let icont = React.createElement(allIcons[fixIconName] || allIcons[icon]);
  console.log(icont)
   return icont
};

const loopMenuItem = (menus: any[]): MenuDataItem[] =>
  menus.map(({ icon, component,children, ...item }) => ({
    ...item,
    component: component,
    routes: children && loopMenuItem(children),
    icon: icon && IconMap(icon),
    children: children && loopMenuItem(children),
}));

const filterByMenuData = (
  data: MenuDataItem[],
  keyWord: '',
): MenuDataItem[] =>
  data.map((item) => {
      if (item.name?.includes(keyWord)) {
        return { ...item };
      }
      const children = filterByMenuData(item.children || [], keyWord);
      if (children.length > 0) {
        return { ...item, children };
      }
       return undefined;
    })
    .filter((item) => item) as MenuDataItem[];


// ProLayout 支持的api https://procomponents.ant.design/components/layout
export const layout: RunTimeLayoutConfig = ({
  initialState,
  setInitialState,
}) => {

  const [keyWord, setKeyWord] = useState('');
  const [loadMenuData, setLoadMenuData] = useState<MenuDataItem[]>([]);
  const loadMenu = async () => {
    try {
      const [menuData] = await Promise.all([
        getMenuList({}),
      ]);
      setLoadMenuData(menuData.data || []);
    } catch (error: any) {
      console.error('加载菜单数据失败', error);
    }
  };
  //
  useEffect(() => {
    loadMenu();
  }, [initialState?.currentUser?.id]);
    console.log(loopMenuItem(loadMenuData))
  return {
    actionsRender: () => [
      <Question key="doc" />,
      <SelectLang key="SelectLang" />,
    ],

    headerContentRender:() => [
      <ProBreadcrumb />,
  ],

    avatarProps: {
      src: initialState?.currentUser?.avatar,
      title: <AvatarName />,
      render: (_, avatarChildren) => {
        return <AvatarDropdown>{avatarChildren}</AvatarDropdown>;
      },
    },
    waterMarkProps: {
      content: initialState?.currentUser?.name,
    },
    //https://pro.ant.design/zh-CN/docs/advanced-menu
    menu:{
      locale: false,
      // request:loopMenuItem(loadMenuData),
    },
    menuExtraRender:({ collapsed }) =>
      !collapsed && (
        <Space
          style={{
            marginBlockStart: 16,
          }}
          align="center"
        >
          <Input
            style={{
              borderRadius: 4,
              backgroundColor: 'rgba(0,0,0,0.03)',
            }}
            prefix={
              <SearchOutlined
                style={{
                  color: 'rgba(0, 0, 0, 0.15)',
                }}
              />
            }
            placeholder="搜索方案"
            variant="borderless"
            onPressEnter={(e) => {
              setKeyWord((e.target as HTMLInputElement).value);
            }}
          />
          <PlusCircleFilled
            style={{
              color: 'var(--ant-primary-color)',
              fontSize: 24,
            }}
          />
        </Space>
      ),

     postMenuData:(menus) => filterByMenuData(menus || [], keyWord || ''),

    menuDataRender: () => {
      if (loadMenuData) {
        return loopMenuItem(loadMenuData);
      }
      return [];
    },


    locale:"zh-CN",
    footerRender: () => <Footer />,
    onPageChange: () => {
      const { location } = history;
      // 如果没有登录，重定向到 login
      if (!initialState?.currentUser && location.pathname !== loginPath) {
        history.push(loginPath);
      }
    },

    bgLayoutImgList: [
      {
        src: 'https://mdn.alipayobjects.com/yuyan_qk0oxh/afts/img/D2LWSqNny4sAAAAAAAAAAAAAFl94AQBr',
        left: 85,
        bottom: 100,
        height: '303px',
      },
      {
        src: 'https://mdn.alipayobjects.com/yuyan_qk0oxh/afts/img/C2TWRpJpiC0AAAAAAAAAAAAAFl94AQBr',
        bottom: -68,
        right: -45,
        height: '303px',
      },
      {
        src: 'https://mdn.alipayobjects.com/yuyan_qk0oxh/afts/img/F6vSTbj8KpYAAAAAAAAAAAAAFl94AQBr',
        bottom: 0,
        left: 0,
        width: '331px',
      },
    ],
    links: isDev
      ? [
          // <Link key="openapi" to="/umi/plugin/openapi" target="_blank">
          //   <LinkOutlined />
          //   <span>OpenAPI 文档</span>
          // </Link>,
        ]
      : [],
    menuHeaderRender: undefined,
    // 自定义 403 页面
    // unAccessible: <div>unAccessible</div>,
    // 增加一个 loading 的状态
    childrenRender: (children) => {
      // if (initialState?.loading) return <PageLoading />;
      return (
        <>
          {children}
          {isDev && (
            <SettingDrawer
              disableUrlParams
              enableDarkTheme
              settings={initialState?.settings}
              onSettingChange={(settings) => {
                setInitialState((preInitialState) => ({
                  ...preInitialState,
                  settings,
                }));
              }}
            />
          )}
        </>
      );
    },
    ...initialState?.settings,
  };
};

/**
 * @name request 配置，可以配置错误处理
 * 它基于 axios 和 ahooks 的 useRequest 提供了一套统一的网络请求和错误处理方案。
 * @doc https://umijs.org/docs/max/request#配置
 */
export const request: RequestConfig = {
  // baseURL: 'https://proapi.azurewebsites.net',
  ...errorConfig,
};
