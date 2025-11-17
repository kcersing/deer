import type {
  ActionType,
  ProColumns,
  ProDescriptionsItemProps,
} from '@ant-design/pro-components';
import {
  FooterToolbar,
  PageContainer,
  ProDescriptions,
  ProTable,
} from '@ant-design/pro-components';
import {  useRequest } from '@umijs/max';
import {Button, Drawer, Input, message, type TreeDataNode} from 'antd';
import React, {useCallback, useEffect, useRef, useState} from 'react';

import CreateForm from './components/CreateForm';
import UpdateForm from './components/UpdateForm';
import {getUserList} from "@/services/ant-design-pro/user";

const TableList: React.FC = () => {
  const actionRef = useRef<ActionType | null>(null);

  const [showDetail, setShowDetail] = useState<boolean>(false);
  const [currentRow, setCurrentRow] = useState<API.User>();
  const [selectedRowsState, setSelectedRows] = useState<API.User[]>([]);

  const [messageApi, contextHolder] = message.useMessage();


  const columns: ProColumns<API.User>[] = [
    {
      title: '名称',
      dataIndex: 'name',
      render: (dom, entity) => {
        return (
          <a
            onClick={() => {
              setCurrentRow(entity);
              setShowDetail(true);
            }}
          >
            {dom}
          </a>
        );
      },
    },

    {
      title: "手机号",
      dataIndex: 'mobile',
      sorter: true,
      hideInForm: true,
    },

    {
      title: "账号",
      dataIndex: 'username',
      sorter: true,
      hideInForm: true,
      hideInTable: true,
    },
    {
      title: "性别",
      dataIndex: 'gender',
      sorter: true,
      hideInForm: true,
    },
    {
      title: "部门",
      dataIndex: '',
      sorter: true,
      hideInForm: true,
    },
    {
      title: "职位",
      dataIndex: '',
      sorter: true,
      hideInForm: true,
    },

    {
      title: "角色",
      dataIndex: '',
      sorter: true,
      hideInForm: true,
    },
    {
      title: '状态',
      dataIndex: 'status',
      hideInForm: true,
      valueEnum: {
        0:{
          text: '禁用',
          status: 'Error',
        },
        1: {
          text: '正常',
          status: 'Success',
        },

      },
    },
    {
      title: "详情",
      dataIndex: 'detail',
      sorter: true,
      hideInForm: true,
      hideInTable: true,
    },

    {
      title: '操作',
      dataIndex: 'option',
      valueType: 'option',
      render: (_, record) => [
        <UpdateForm
          trigger={
            <a>更新</a>
          }
          key="config"
          onOk={actionRef.current?.reload}
          values={record}
        />,
      ],
    },
  ];

  return (
    <PageContainer>
      {contextHolder}
      <ProTable<API.User, API.PageParams>
        headerTitle='人员列表'
        actionRef={actionRef}
        rowKey="id"
        pagination={false}
        search={{
          labelWidth: 120,
        }}
        toolBarRender={() => [
          <CreateForm key="create" reload={actionRef.current?.reload} />,
        ]}
        request={getUserList}
        columns={columns}
        rowSelection={{
          onChange: (_, selectedRows) => {
            setSelectedRows(selectedRows);
          },
        }}
      />


      <Drawer
        width={600}
        open={showDetail}
        onClose={() => {
          setCurrentRow(undefined);
          setShowDetail(false);
        }}
        closable={false}
      >
        {currentRow?.name && (
          <ProDescriptions<API.User>
            column={2}
            title={currentRow?.name}
            request={async () => ({
              data: currentRow || {},
            })}
            params={{
              id: currentRow?.name,
            }}
            columns={columns as ProDescriptionsItemProps<API.Menu>[]}
          />
        )}
      </Drawer>
    </PageContainer>
  );
};


// const [menuData, setMenuData] = useState<TreeDataNode[]>([]);
//
// const loadData = async () => {
//   try {
//     const [menuData] = await Promise.all([
//       getMenuTree(),
//     ]);
//     setMenuData(menuData.data)
//   } catch (error: any) {
//     console.error('加载问卷数据失败', error);
//     message.error(error.message || '加载问卷数据失败');
//   } finally {
//     // dispatch({ type: 'LOADING', payload: false });
//   }
// }
//
// useEffect(() => {
//   loadData();
// }, []);
//
// console.log(menuData);

export default TableList;
