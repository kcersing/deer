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
import {Button, Drawer, Input,Dropdown, Space, Tag ,AntDesignOutlined,Avatar, message, type TreeDataNode} from 'antd';

import React, {useCallback, useEffect, useRef, useState} from 'react';

import CreateForm from './components/CreateForm';
import UpdateForm from './components/UpdateForm';

import { Order } from  "./service/data";
import {deleteOrder, getOrderList} from "./service/service";

const TableList: React.FC = () => {
  const actionRef = useRef<ActionType | null>(null);

  const [showDetail, setShowDetail] = useState<boolean>(false);
  const [currentRow, setCurrentRow] = useState<Order>();
  const [selectedRowsState, setSelectedRows] = useState<Order[]>([]);

  const [messageApi, contextHolder] = message.useMessage();

  const { run: delRun, loading } = useRequest(deleteOrder, {
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


  const columns: ProColumns<Order>[] = [
    {
      title: "ID",
      dataIndex: 'id',
      sorter: true,
      hideInForm: true,
      copyable:true,
    },
    {
      title: '订单编号',
      dataIndex: 'sn',
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
      title: "会员",
      dataIndex: 'memberName',
      sorter: true,
      hideInForm: true,
      copyable:true,
    },
    {
      title: "会员手机号",
      dataIndex: 'memberMobile',
      sorter: true,
      hideInForm: true,
      copyable:true,
    },
    {
      title: "订单类型",
      dataIndex: 'nature',
      sorter: true,
      hideInForm: true,

    },
    {
      title: "总金额",
      dataIndex: 'totalAmount',
      sorter: true,
      search: false,
      hideInForm: true,

    },
    {
      title: "实际已付款",
      dataIndex: 'actual',
      search: false,
      sorter: true,
      hideInForm: true,

    },
    {
      title: "员工",
      dataIndex: 'userName',
      sorter: true,
      hideInForm: true,
      copyable:true,
    },
    {
      disable: true,
      title: '标签',
      dataIndex: 'tag',
      search: false,
      renderFormItem: (_, { defaultRender }) => {
        return defaultRender(_);
      },
      render: (_, record) => (
        <Space>
          {/*{record.labels.map(({ name, color }) => (*/}
          {/*  <Tag color={color} key={name}>*/}
          {/*    {name}*/}
          {/*  </Tag>*/}
          {/*))}*/}
        </Space>
      ),
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
    async (selectedRows: Order[]) => {
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
      <ProTable<Order, API.PageParams>
        headerTitle='订单列表'
        actionRef={actionRef}
        rowKey="id"
        pagination={false}

        form={{
          // 由于配置了 transform，提交的参数与定义的不同这里需要转化一下
          syncToUrl: (values, type) => {
            if (type === 'get') {
              return {
                ...values,
                created_at: [values.startTime, values.endTime],
              };
            }
            console.log(values)
            return values;
          },
        }}
        search={{
          labelWidth: 120,
          optionRender: (searchConfig, formProps, dom) => [
            ...dom.reverse(),
            <Button key="export" onClick={() => {
              const values = searchConfig?.form?.getFieldsValue();
              console.log(values);

            }}>导出</Button>,
          ],
        }}

        toolBarRender={() => [
          // <Button type="primary">导出</Button>,
        ]}
        request={getOrderList}
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
          <ProDescriptions<Order>
            column={2}
            title={currentRow?.name}
            request={async () => ({
              data: currentRow || {},
            })}
            params={{
              id: currentRow?.name,
            }}
            columns={columns as ProDescriptionsItemProps<Order>[]}
          />
        )}
      </Drawer>
    </PageContainer>
  );
};

export default TableList;
