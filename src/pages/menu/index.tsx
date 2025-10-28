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
import { FormattedMessage, useIntl, useRequest } from '@umijs/max';
import {Button, Drawer, Input, message, type TreeDataNode} from 'antd';
import React, {useCallback, useEffect, useRef, useState} from 'react';

import CreateForm from './components/CreateForm';
import UpdateForm from './components/UpdateForm';
import {deleteMenu, getMenuList} from "@/services/ant-design-pro/menu";

const TableList: React.FC = () => {
  const actionRef = useRef<ActionType | null>(null);

  const [showDetail, setShowDetail] = useState<boolean>(false);
  const [currentRow, setCurrentRow] = useState<API.Menu>();
  const [selectedRowsState, setSelectedRows] = useState<API.Menu[]>([]);

  const [messageApi, contextHolder] = message.useMessage();

  const { run: delRun, loading } = useRequest(deleteMenu, {
    manual: true,
    onSuccess: () => {
      setSelectedRows([]);
      actionRef.current?.reloadAndRest?.();

      messageApi.success('删除成功，即将刷新');
    },
    onError: () => {
      messageApi.error('删除失败，请重试');
    },
  });

  const columns: ProColumns<API.RuleListItem>[] = [
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
      title: "标题",
      dataIndex: 'title',
      sorter: true,
      hideInForm: true,
    },
    {
      title: "图标",
      dataIndex: 'icon',
      sorter: true,
      hideInForm: true,
    },
    {
      title: "路由路径",
      dataIndex: 'path',
      sorter: true,
      hideInForm: true,
    },
    {
      title: "组件路径",
      dataIndex: 'component',
      sorter: true,
      hideInForm: true,
      valueType: 'textarea',
    },

    {
      title: "跳转路径",
      dataIndex: 'redirect',
      sorter: true,
      hideInForm: true,
      valueType: 'textarea',
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

  /**
   *  Delete node
   * @zh-CN 删除节点
   *
   * @param selectedRows
   */
  const handleRemove = useCallback(
    async (selectedRows: API.RuleListItem[]) => {
      if (!selectedRows?.length) {
        messageApi.warning('请选择删除项');

        return;
      }

      // await delRun({
      //   data: {
      //     key: selectedRows.map((row) => row.key),
      //   },
      // });
    },
    [delRun, messageApi.warning],
  );

  return (
    <PageContainer>
      {contextHolder}
      <ProTable<API.RuleListItem, API.PageParams>
        headerTitle='菜单列表'
        actionRef={actionRef}
        rowKey="id"
        pagination={false}
        search={{
          labelWidth: 120,
        }}
        toolBarRender={() => [
          <CreateForm key="create" reload={actionRef.current?.reload} />,
        ]}
        request={getMenuList}
        columns={columns}
        rowSelection={{
          onChange: (_, selectedRows) => {
            setSelectedRows(selectedRows);
          },
        }}
      />
      {selectedRowsState?.length > 0 && (
        <FooterToolbar
          extra={
            <div>
              被选中{' '}
              <a style={{ fontWeight: 600 }}>{selectedRowsState.length}</a>{' '}
              项
            </div>
          }
        >
          <Button
            loading={loading}
            onClick={() => {
              handleRemove(selectedRowsState);
            }}
          >
            批量删除
          </Button>
          {/*<Button type="primary">*/}
          {/*  批量更新*/}
          {/*</Button>*/}
        </FooterToolbar>
      )}

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
          <ProDescriptions<API.Menu>
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
