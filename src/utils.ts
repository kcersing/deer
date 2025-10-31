import React from "react";
import * as allIcons from "@ant-design/icons";
import type {MenuDataItem} from "@ant-design/pro-components";

export const loopMenuItem = (menus: any[]): MenuDataItem[] =>
  menus.map(({path, icon, component,children, ...item }) => ({
    ...item,
    icon: icon && React.createElement(allIcons[icon]),
    component:component,
    path:path,
    routes: children && loopMenuItem(children),
  }));

// export const loopMenu = (menus: any[]): MenuDataItem[] =>
//   menus.map(({ path,icon, component,children, ...item }) => ({
//     ...item,
//     path:path,
//     icon: icon && React.createElement(allIcons[icon]),
//      component: require(`@/pages/${component}`).default,
//     routes: children && loopMenuItem(children),
//   }));


export const filterByMenuData = (
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
  }).filter((item) => item) as MenuDataItem[];


