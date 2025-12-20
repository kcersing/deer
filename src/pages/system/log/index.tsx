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


import { Log } from  "@/pages/system/log/service/data";
import {deleteLog,getLogList}  from "@/pages/system/log/service/service";

const TableList: React.FC = () => {
  const actionRef = useRef<ActionType | null>(null);

  const [showDetail, setShowDetail] = useState<boolean>(false);
  const [currentRow, setCurrentRow] = useState<Log>();
  const [selectedRowsState, setSelectedRows] = useState<Log[]>([]);

  const [messageApi, contextHolder] = message.useMessage();

  const { run: delRun, loading } = useRequest(deleteLog, {
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

  const columns: ProColumns<Log>[] = [
    {
      title: 'API',
      dataIndex: 'api',
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
      title: "类型",
      dataIndex: 'type',
      sorter: true,
      hideInForm: true,
    },
    {
      title: "请求方法",
      dataIndex: 'method',
      sorter: true,
      hideInForm: true,
    },
    {
      title: "日志请求内容",
      dataIndex: 'req_content',
      sorter: true,
      hideInForm: true,
    },
    {
      title: "日志返回内容",
      dataIndex: 'resp_content',
      sorter: true,
      hideInForm: true,
      valueType: 'textarea',
    },

    {
      title: "IP",
      dataIndex: 'ip',
      sorter: true,
      hideInForm: true,
      valueType: 'textarea',
    },

    {
      title: "操作者ID",
      dataIndex: 'identity',
      sorter: true,
      hideInForm: true,
      valueType: 'textarea',
    },
    {
      title: "日志时间",
      dataIndex: 'time',
      sorter: true,
      hideInForm: true,
      valueType: 'textarea',
    },




  ];

  /**
   *  Delete node
   * @zh-CN 删除节点
   *
   * @param selectedRows
   */
  const handleRemove = useCallback(
    async (selectedRows: Log[]) => {
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
      <ProTable<Log, API.PageParams>
        headerTitle='菜单列表'
        actionRef={actionRef}
        rowKey="id"
        pagination={false}
        search={{
          labelWidth: 120,
        }}
        request={getLogList}
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
          <ProDescriptions<Log>
            column={2}
            title={currentRow?.title}
            request={async () => ({
              data: currentRow || {},
            })}
            params={{
              id: currentRow?.title,
            }}
            columns={columns as ProDescriptionsItemProps<Log>[]}
          />
        )}
      </Drawer>
    </PageContainer>
  );
};



export default TableList;
