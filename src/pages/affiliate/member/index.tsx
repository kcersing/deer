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
import {Button, Drawer, Input,AntDesignOutlined,Avatar, message, type TreeDataNode} from 'antd';
import React, {useCallback, useEffect, useRef, useState} from 'react';

import CreateForm from './components/CreateForm';
import UpdateForm from './components/UpdateForm';

import { Member } from  "@/pages/affiliate/member/service/data";
import {deleteMember, getMemberList} from "@/pages/affiliate/member/service/service";

const TableList: React.FC = () => {
  const actionRef = useRef<ActionType | null>(null);

  const [showDetail, setShowDetail] = useState<boolean>(false);
  const [currentRow, setCurrentRow] = useState<Member>();
  const [selectedRowsState, setSelectedRows] = useState<Member[]>([]);

  const [messageApi, contextHolder] = message.useMessage();

  const { run: delRun, loading } = useRequest(deleteMember, {
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


  const columns: ProColumns<Member>[] = [
    {
      title: "ID",
      dataIndex: 'id',
      sorter: true,
      hideInForm: true,
      copyable:true,
    },
    {
      title: "图片",
      dataIndex: 'avatar',
      sorter: true,
      hideInForm: true,
      render: (dom, entity) => {
        return (
          <Avatar
            size={{ xs: 24, sm: 32, md: 40, lg: 64, xl: 80, xxl: 100 }}
            src={entity}
          />
        );
      },

    },

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
      copyable:true,
    },
    {
      title: "等级",
      dataIndex: 'level',
      sorter: false,
      hideInForm:false,
    },
    {
      title: "性别",
      dataIndex: 'gender',
      sorter: true,
      hideInForm: true,

    },
    {
      title: "出生日期",
      dataIndex: 'birthday',
      sorter: true,
      hideInForm: true,

    },
    {
      title: "意向",
      dataIndex: 'intention',
      sorter: true,
      hideInForm: true,

    },
    {
      title: "来源",
      dataIndex: 'source',
      sorter: true,
      hideInForm: true,

    },
    {
      title: "备注",
      dataIndex: 'note',
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
    async (selectedRows: Member[]) => {
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
      <ProTable<Member, API.PageParams>
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
        request={getMemberList}
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
          <ProDescriptions<Member>
            column={2}
            title={currentRow?.name}
            request={async () => ({
              data: currentRow || {},
            })}
            params={{
              id: currentRow?.name,
            }}
            columns={columns as ProDescriptionsItemProps<Member>[]}
          />
        )}
      </Drawer>
    </PageContainer>
  );
};

export default TableList;
