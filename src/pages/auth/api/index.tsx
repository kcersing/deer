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
import { Api } from  "@/pages/auth/api/service/data";
import {deleteApi, getApiList} from "@/pages/auth/api/service/service";

const TableList: React.FC = () => {
  const actionRef = useRef<ActionType | null>(null);

  const [showDetail, setShowDetail] = useState<boolean>(false);
  const [currentRow, setCurrentRow] = useState<Api>();
  const [selectedRowsState, setSelectedRows] = useState<Api[]>([]);

  const [messageApi, contextHolder] = message.useMessage();

  const { run: delRun, loading } = useRequest(deleteApi, {
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

  const columns: ProColumns<Api>[] = [
    {
      title: "id",
      dataIndex: 'id',
      sorter: true,
      hideInForm: true,
      hideInTable: true,
    },
    {
      title: '名称',
      dataIndex: 'title',
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
      title: "路径",
      dataIndex: 'path',
      sorter: true,
      hideInForm: true,
    },
    {
      title: "描述",
      dataIndex: 'desc',
      sorter: true,
      hideInForm: true,
    },

    {
      title: "分组",
      dataIndex: 'group',
      sorter: true,
      hideInForm: true,
      valueType: 'textarea',
    },

    {
      title: "请求类型",
      dataIndex: 'method',
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
    async (selectedRows: Api[]) => {
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
      <ProTable<Api, API.PageParams>
        headerTitle='API列表'
        actionRef={actionRef}
        rowKey="id"
        pagination={false}
        search={{
          labelWidth: 120,
        }}
        toolBarRender={() => [
          <CreateForm key="create" reload={actionRef.current?.reload} />,
        ]}
        request={getApiList}
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
        {currentRow?.title && (
          <ProDescriptions<Api>
            column={2}
            title={currentRow?.title}
            request={async () => ({
              data: currentRow || {},
            })}
            params={{
              id: currentRow?.title,
            }}
            columns={columns as ProDescriptionsItemProps<Api>[]}
          />
        )}
      </Drawer>
    </PageContainer>
  );
};



export default TableList;
