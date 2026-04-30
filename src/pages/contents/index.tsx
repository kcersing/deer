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

import { Article } from  "./service/data";
import {deleteArticle, getArticleList} from "./service/service";
import {getDicthtList} from "../system/dict/service/service";
import {getMessagesSendList} from "../account/center/service/service";

const TableList: React.FC = () => {
  const actionRef = useRef<ActionType | null>(null);

  const [showDetail, setShowDetail] = useState<boolean>(false);
  const [currentRow, setCurrentRow] = useState<Article>();
  const [selectedRowsState, setSelectedRows] = useState<Article[]>([]);

  const [messageApi, contextHolder] = message.useMessage();

  const { run: delRun, loading } = useRequest(deleteArticle, {
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



  const { data: tagEnum } = useRequest(() => {
    return  getDicthtList({'dictId':2})
  });


  const columns: ProColumns<Article>[] = [
    {
      title: "id",
      dataIndex: 'id',
      sorter: true,
      hideInForm: true,
      filters: true,
      onFilter: true,
      ellipsis: true,
      width: 60,
      hideInTable: true,
      search: false,
    },
    {
      title: '标题',
      dataIndex: 'title',
      ellipsis: true,
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
      title: "内容",
      dataIndex: 'content',
      sorter: true,
      search: false,
      hideInForm: false,
      hideInTable: true,
    },
    {
      title: "标签",
      dataIndex: 'tagId',
      sorter: true,
      hideInForm: true,
      valueType: 'select',
      fieldProps: {
        options: tagEnum,
        showSearch: true,
        multiple: true,
        mode:"multiple",
        fieldNames:{ label: 'title', value: 'id' }
      },
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
      title: "创建时间",
      dataIndex: 'createdAt',
      sorter: false,
      hideInForm:false,
      valueType: 'date',
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
    async (selectedRows: Article[]) => {
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
      <ProTable<Article, API.PageParams>
        headerTitle='菜单列表'
        actionRef={actionRef}
        rowKey="id"
        pagination={false}

        toolBarRender={() => [
          <CreateForm key="create" reload={actionRef.current?.reload} />,
        ]}



        search={{

          defaultCollapsed: false,
          labelWidth: 120,
          optionRender: (searchConfig, formProps, dom) => [
            ...dom.reverse(),
            <Button key="export" onClick={() => {
              const values = searchConfig?.form?.getFieldsValue();
              console.log(values);
            }}>导出</Button>,
          ],
        }}






        request={getArticleList}
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
          <ProDescriptions<Article>
            column={2}
            title={currentRow?.title}
            request={async () => ({
              data: currentRow || {},
            })}
            params={{
              id: currentRow?.title,
            }}
            columns={columns as ProDescriptionsItemProps<Article>[]}
          />
        )}
      </Drawer>
    </PageContainer>
  );
};

export default TableList;
