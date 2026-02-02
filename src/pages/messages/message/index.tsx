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
import {Button, Drawer, Dropdown, Space, Tag , Input, message, type TreeDataNode} from 'antd';

import React, {useCallback, useEffect, useRef, useState} from 'react';

import CreateForm from './components/CreateForm';

import { Messages } from  "./service/data";
import {deleteMessages, getMessagesList,getMessagesTypes} from "./service/service";
import UpdateForm from './components/UpdateForm';

const TableList: React.FC = () => {
  const actionRef = useRef<ActionType | null>(null);

  const [showDetail, setShowDetail] = useState<boolean>(false);
  const [currentRow, setCurrentRow] = useState<Messages>();
  const [selectedRowsState, setSelectedRows] = useState<Messages[]>([]);

  const [types, setTypes] = useState([]);



  useEffect(() => {
    const getTypes = async () => {
      const [res] = await Promise.all([
        getMessagesTypes(),
      ]);
      setTypes(res.data)
    }
    getTypes();
  }, []);



  const [messageApi, contextHolder] = message.useMessage();

  const { run: delRun, loading } = useRequest(deleteMessages, {
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


  const columns: ProColumns<Messages>[] = [
    {
      title: "id",
      dataIndex: 'id',
      sorter: true,
      hideInForm: true,
      hideInTable: true,
    },
    {
      title: '消息主题',
      dataIndex: 'title',
      render: (dom, entity) => {
        console.log(entity)
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
      title: '消息类型',
      dataIndex: 'type',
      hideInForm: true,
      render: (dom, entity) => {

        return (
          <>
            {types.map(type => (
              type.value == dom && <span key={type.title}>{type.title}</span>
            ))}
          </>
        );
      },
    },


    {
      title: '消息状态',
      dataIndex: 'status',
      hideInForm: true,
      valueEnum: {
        0:{
          text: '草稿',
          status: 'Error',
        },
        1: {
          text: '已发布/发送完成',
          status: 'Success',
        },
        2: {
          text: '定时发布中',
          status: 'Success',
        },
        3: {
          text: '已撤销',
          status: 'Success',
        },
        4: {
          text: '已归档',
          status: 'Success',
        },
        5: {
          text: '已删除',
          status: 'Success',
        },

      },
    },

    {
      title: "发送者",
      dataIndex: 'createdName',
      sorter: false,
      hideInForm:false,
    },

    {
      title: "创建时间",
      dataIndex: 'createdAt',
      sorter: false,
      hideInForm:false,
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
    async (selectedRows: Messages[]) => {
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
      <ProTable<Messages, API.PageParams>
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
        request={getMessagesList}
        columns={columns}
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

        {currentRow?.title && (
          <ProDescriptions<Messages>
            column={2}
            title={currentRow?.title}
            request={async () => ({
              data: currentRow || {},
            })}
            params={{
              id: currentRow?.title,
            }}
            columns={columns as ProDescriptionsItemProps<Messages>[]}
          />
        )}
      </Drawer>
    </PageContainer>
  );
};

export default TableList;
