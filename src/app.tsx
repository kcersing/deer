import type { Settings as LayoutSettings } from '@ant-design/pro-components';

import { SettingDrawer ,ProBreadcrumb} from '@ant-design/pro-components';
import type { RequestConfig, RunTimeLayoutConfig } from '@umijs/max';
import { history, Link } from '@umijs/max';
import React, { useState, useReducer, useEffect, useRef, memo, useCallback } from 'react';
import {
  AvatarDropdown,
  AvatarName,
  Footer,
  Question,
} from '@/components';
import defaultSettings from '../config/defaultSettings';
import { errorConfig } from './requestErrorConfig';
import '@ant-design/v5-patch-for-react-19';
import { getMenuList} from "@/services/ant-design-pro/menu";
import {getUser, getRoleMenuAll} from '@/services/ant-design-pro/user';
import { matchRoutes } from 'umi';
import { loopMenuItem} from "@/utils";
import Login from '@/pages/login';

const isDev = process.env.NODE_ENV === 'development';
const loginPath = '/login';

let extraRoutes: any[] = [];







export function onRouteChange({ location, clientRoutes, routes, action, basename, isFirst }: any) {
  // const access = useAccess();
  const isLogin = sessionStorage.getItem('token');

  const route = matchRoutes(clientRoutes, location.pathname)?.pop()?.route;
  if (route) {
    document.title = route.title || '';
  }

  // 未登录跳转登录页
  if (!isLogin && isLogin == '' && location.pathname !== loginPath) {
    history.push(loginPath);
  }

  // // 检查页面权限（示例）
  // if (!access.canRead(location.pathname)) {
  //   history.push('/403');
  // }
}


export function patchClientRoutes({ routes }) {
  routes.unshift({
    path: '/login',
    layout: false,
    element: <Login />,
  });
  // routes.unshift({
  //     path: '/',
  //     redirect: '/welcome',
  //   });
}


/**
 * @see https://umijs.org/docs/api/runtime-config#getinitialstate
 * */
export async function getInitialState(): Promise<{
  settings?: Partial<LayoutSettings>;
  currentUser?: API.User;
  loading?: boolean;
  fetchUserInfo?: () => Promise<API.User | undefined>;
}> {
  const fetchUserInfo = async () => {
    try {
      const msg = await getUser({});
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

// ProLayout 支持的api https://procomponents.ant.design/components/layout
export const layout: RunTimeLayoutConfig = ({
  initialState,
  setInitialState,
}) => {

  return {
    actionsRender: () => {
    return   <Question key="doc" />
  },
    headerContentRender:() => {
      return   <ProBreadcrumb />
    },

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
      params: {
        id: initialState?.currentUser?.id,
      },
      request: async (params, defaultMenuData) => {

        if(params.id == undefined || params.id == 0){
          return defaultMenuData
        }
        // initialState.currentUser 中包含了所有用户信息
        const loadMenuData = await getMenuList({});
        let Item = loopMenuItem(loadMenuData.data || [])
        return loopMenuItem(loadMenuData.data || []);
      },

    },
    footerRender: () =>{
      return  <Footer />
    },
    menuHeaderRender: undefined,
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
